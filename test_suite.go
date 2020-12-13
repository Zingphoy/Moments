package main

import (
	"Moments/models"
	"Moments/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
)

/*
 This is the utility file for test.
 To make test more convenient, here will offer some database testing functions for inserting/deleting/updating data.
*/

func init() {
	log.InitLogger(true)
}

func insertData(dbname string, data interface{}) {
	db, client, ctx, _ := models.ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal("insert data failed, ", err.Error())
	}
}

func deleteData(dbname string, data interface{}) {
	db, client, ctx, _ := models.ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	_, err := collection.DeleteOne(ctx, data)
	if err != nil {
		log.Fatal("insert data failed, ", err.Error())
	}
}

func updateData(dbname string, filter interface{}, update interface{}) {
	db, client, ctx, _ := models.ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	collection := db.Collection(dbname)

	//filter := bson.M{"uid": 90001}
	//update := bson.D{{"$set", bson.D{{"aid_list", tempList}}}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("update data failed, ", err.Error())
	}
}

// example for copy&pasted
func updateAlbum() {
	filter := bson.M{"uid": 90001}
	update := bson.D{{"$set",
		bson.D{
			{"aid_list", bson.A{900011604900530, 900011604900529}},
			{"uid", 90001},
		},
	}}
	updateData("album", filter, update)
}

// write test code here
func main() {
	filter := bson.M{"uid": 90001}
	update := bson.D{{"$set",
		bson.D{
			{"privacy", int32(0)},
		},
	}}
	updateData("article_1", filter, update)
	updateData("article_2", filter, update)
	updateData("article_3", filter, update)

}
