package models

import (
	"Moments/pkg/log"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Article struct {
	Aid        int64  `json:"aid"`
	Uid        int32  `json:"uid"`
	Post_time  int64  `json:"post_time"`
	Content    string `json:"content"`
	Photo_list bson.A `json:"photo_list"`
	Privacy    int32  `json:"privacy"`
}

func makeArticleObj(a bson.M) *Article {
	article := Article{
		Aid:        a["aid"].(int64),
		Uid:        a["uid"].(int32),
		Post_time:  a["post_time"].(int64),
		Content:    a["content"].(string),
		Photo_list: a["photo_list"].(bson.A),
		Privacy:    a["privacy"].(int32),
	}
	return &article
}

// IsAidExist check if specific aid is already existed
func IsAidExist(dbname string, aid int64) bool {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	cur, err := collection.Find(ctx, bson.M{"aid": aid})
	defer cur.Close(ctx)
	if err != nil || cur.Next(ctx) {
		return false
	}
	return true
}

// GetDetail get detail of an article with specific filter
func GetDetail(dbname string, filter map[string]interface{}) (*Article, error) {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	var row bson.M
	collection := db.Collection(dbname)
	err := collection.FindOne(ctx, bson.M(filter)).Decode(&row)
	if err != nil {
		log.Error(fmt.Sprintf("Find method in databese \"%s\" error: %s", dbname, err.Error()))
		return nil, err
	}

	log.Info("get data success: ", row)
	article := makeArticleObj(row)
	return article, nil
}

// AddArticle add article to database
func AddArticle(dbname string, data map[string]interface{}) error {
	article := Article{
		Aid:        data["aid"].(int64),
		Uid:        data["uid"].(int32),
		Post_time:  data["post_time"].(int64),
		Content:    data["content"].(string),
		Photo_list: data["photo_list"].(bson.A),
		Privacy:    data["privacy"].(int32),
	}

	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	_, err := collection.InsertOne(ctx, article)
	return err
}

// DeleteArticle softly delete an aticle with specific filter
func DeleteArticle(dbname string, filter map[string]interface{}) error {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
