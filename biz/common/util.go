package common

import (
	"Moments/biz/database"
	"Moments/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
)

/*
 Common utility functions for package model
*/

func queryOne(dbname string, filter map[string]interface{}) (map[string]interface{}, error) {
	db, client, ctx, _ := database.ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	var row bson.M
	collection := db.Collection(dbname)
	err := collection.FindOne(ctx, filter).Decode(&row)
	if err != nil {
		log.Error("query data failed", err.Error())
		return nil, err
	}
	return row, nil
}

// useless function
func query(dbname string, filter map[string]interface{}) ([]interface{}, error) {
	db, client, ctx, _ := database.ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error("query data failed", err.Error())
		return nil, err
	}
	defer cur.Close(ctx)

	var rows []bson.M
	if err = cur.All(ctx, &rows); err != nil {
		log.Error(err.Error())
		return nil, err
	}
	ret := make([]interface{}, 0, len(rows))
	for _, result := range rows {
		ret = append(ret, result)
	}
	return ret, nil
}

func update(dbname string, filter map[string]interface{}, data map[string]interface{}) error {
	db, client, ctx, _ := database.ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	var up bson.D
	for k, v := range data {
		up = bson.D{{"$set",
			bson.D{
				{k, v},
			},
		}}
	}
	collection := db.Collection(dbname)
	_, err := collection.UpdateOne(ctx, filter, up)
	if err != nil {
		log.Error("update database failed,", err.Error())
		return err
	}
	return nil
}

func insert(dbname string, data interface{}) error {
	db, client, ctx, _ := database.ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Error("insert data failed,", err.Error())
		return err
	}
	return nil
}

func remove(dbname string, filter map[string]interface{}) error {
	db, client, ctx, _ := database.ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error("remove data failed,", err.Error())
		return err
	}
	return nil
}

/* ---------------------------------- Pubilc ---------------------------------- */

func QueryOne(dbname string, filter map[string]interface{}) (map[string]interface{}, error) {
	db, client, ctx, _ := database.ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	var row bson.M
	collection := db.Collection(dbname)
	err := collection.FindOne(ctx, filter).Decode(&row)
	if err != nil {
		log.Error("query data failed", err.Error())
		return nil, err
	}
	return row, nil
}

func Insert(dbname string, data interface{}) error {
	db, client, ctx, _ := database.ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection(dbname)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Error("insert data failed,", err.Error())
		return err
	}
	return nil
}
