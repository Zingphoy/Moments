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

// IsAidExist check if specific aid is already existed
func IsAidExist(dbname string, aid int64) bool {
	filter := map[string]interface{}{"aid": aid}
	rows, err := query(dbname, filter)
	if err != nil {
		log.Info("error: ", err.Error())
		return false
	}
	if len(rows) <= 0 {
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

// AddArticle add article to database, and expand this article into friends' timeline
func AddArticle(dbname string, data map[string]interface{}) error {
	article := Article{
		Aid:       data["aid"].(int64),
		Uid:       data["uid"].(int32),
		PostTime:  data["post_time"].(int64),
		Content:   data["content"].(string),
		PhotoList: data["photo_list"].(bson.A),
		Privacy:   data["privacy"].(int32),
	}

	err := insert(dbname, article)
	if err != nil {
		log.Error("add article failed")
	}
	return err
}

// DeleteArticle delete article permanently
func DeleteArticle(dbname string, filter map[string]interface{}) error {
	err := remove(dbname, filter)
	if err != nil {
		log.Error("permanently delete article failed, error:", err.Error())
	}
	log.Info("delete article success")
	return err
}

// DeleteArticleSoft delete article softly
func DeleteArticleSoft(dbname string, filter map[string]interface{}) error {
	data := map[string]interface{}{"is_deleted": 1}
	err := update(dbname, filter, data)
	if err != nil {
		log.Error("softly delete article failed, error:", err.Error())
	}
	log.Info("softly delete article success")
	return err
}
