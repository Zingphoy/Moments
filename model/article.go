package model

import (
	"Moments/pkg/log"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Article struct {
	Aid       int64  `bson:"aid"`
	Uid       int32  `bson:"uid"`
	PostTime  int64  `bson:"post_time"`
	Content   string `bson:"content"`
	PhotoList bson.A `bson:"photo_list"`
	Privacy   int32  `bson:"privacy"`
	IsDeleted int32  `bson:"is_deleted"`
}

func makeModelArticleObj(a bson.M) *Article {
	article := Article{
		Aid:       a["aid"].(int64),
		Uid:       a["uid"].(int32),
		PostTime:  a["post_time"].(int64),
		Content:   a["content"].(string),
		PhotoList: a["photo_list"].(bson.A),
		Privacy:   a["privacy"].(int32),
		IsDeleted: 0,
	}
	return &article
}

// todo 有bug，为什么aid会被判断存在
// IsAidExist check if specific aid is already existed
func IsAidExist(dbname string, aid int64) bool {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	cur, err := collection.Find(ctx, bson.M{"aid": aid})
	if err != nil {
		log.Info("error: ", err.Error())
		return false
	}
	defer cur.Close(ctx)
	if !cur.Next(ctx) {
		return false
	}
	return true
}

// GetDetail get detail of an article with specific filter
func GetDetail(dbname string, filter map[string]interface{}) (*Article, error) {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	var row bson.M
	collection := db.Collection(dbname)
	err := collection.FindOne(ctx, bson.M(filter)).Decode(&row)
	if err != nil {
		log.Info(fmt.Sprintf("Find() method in databese \"%s\" error: %s", dbname, err.Error()))
		return nil, err
	}

	article := makeModelArticleObj(row)
	return article, nil
}

// AddArticle add article to database
func AddArticle(dbname string, data map[string]interface{}) (err error) {
	article := Article{
		Aid:       data["aid"].(int64),
		Uid:       data["uid"].(int32),
		PostTime:  data["post_time"].(int64),
		Content:   data["content"].(string),
		PhotoList: data["photo_list"].(bson.A),
		Privacy:   data["privacy"].(int32),
	}

	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	if _, err = collection.InsertOne(ctx, article); err != nil {
		log.Error("add article failed")
	}

	// todo 这里引入扩散写：调用mq服务，mq服务将文章写到好友的timeline数据库
	// 应该调用某个东西，让他来完成具体的expand消息发送，同时还要有地方统一定义消息
	return err
}

// DeleteArticle delete article permanently
func DeleteArticle(dbname string, filter map[string]interface{}) (err error) {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	if _, err = collection.DeleteOne(ctx, filter); err != nil {
		log.Error("permanently delete article failed, error:", err.Error())
	}
	log.Info("delete article success")
	return err
}

// DeleteArticleSoft delete article softly
func DeleteArticleSoft(dbname string, filter map[string]interface{}) (err error) {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	update := bson.D{{"$set",
		bson.D{
			{"is_deleted", 1},
		}},
	}
	if _, err = collection.UpdateOne(ctx, filter, update); err != nil {
		log.Error("softly delete article failed, error:", err.Error())
	}
	log.Info("softly delete article success")
	return err
}
