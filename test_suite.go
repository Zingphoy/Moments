package main

import (
	"Moments/model"
	"Moments/pkg/log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

/*
 This is the utility file for test.
 To make test more convenient, here will offer some database testing functions for inserting/deleting/updating data.
*/

func init() {
	log.InitLogger(true)
	log.RedirectLogStd()
}

func insertData(dbname string, data interface{}) {
	db, client, ctx, _ := model.ConnectDatabase()
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
	db, client, ctx, _ := model.ConnectDatabase()
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
	db, client, ctx, _ := model.ConnectDatabase()
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

// when running article_test.go, it may create some test articles and we need to clear them
func clearTestArticle() {
	db, client, ctx, _ := model.ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	for i := 0; i < 4; i++ {
		dbname := "article_" + strconv.Itoa(i)
		collection := db.Collection(dbname)
		_, err := collection.DeleteOne(ctx, bson.M{"privacy": 999})
		if err != nil {
			log.Error("delete test article failed")
		}
	}
}

func print(done chan int) {
	log.Info("print sth")
	log.Error("print sth")
	if done != nil {
		done <- 0
	}
}

// write test code heret
func main() {
	log.RedirectLogStd()
	done := make(chan int)
	go print(done)
	<-done
}
