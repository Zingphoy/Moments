package models

/*
 timeline is sorted in order depends on its timing of insertion. When getting timeline from database, it will
 offer the newest 10 (if needs) since the latest at local cache.
*/

import (
	"Moments/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
)

type Timeline struct {
	Uid      int32   `json:"uid"`
	Aid_list []int64 `json:"aid_list"`
}

const (
	AUTHORIZED = 0
	DENIED     = 1
)

// todo 补充鉴权中间件
// params: uid int32
func hasPermission() error {
	return nil
}

// todo Unit Test
// GetTimelineRefreshByUid when client refresh timeline, it will fetch 10 newest atticle and check privacy,
// finally offer the matched artile regardless of amount. Don't deal with the case when less than 1 will be returned.
func GetTimelineRefreshByUid(uid int32, aid int64) ([]int64, error) {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection("timeline")
	//opts := options.Find().SetLimit(10)
	var row bson.M
	err := collection.FindOne(ctx, bson.M{"uid": uid}).Decode(&row)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var aids []int64
	for _, v := range (row["aid_list"]).(bson.A) {
		if v == aid {
			break
		}
		aids = append(aids, v.(int64))
	}
	return aids, nil
}

// todo
// GetTimelineLoadmoreByUid
func GetTimelineLoadmoreByUid(uid int32, aid int64) ([]int64, error) {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection("timeline")
	collection.Find(ctx, bson.M{})

	return nil, nil
}

// AppendTimeline
func AppendTimeline(filter interface{}, update interface{}) error {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection("timeline")
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		log.Error("")
	}

	return nil
}

// InsertNewTimeline
func InsertNewTimeline(data interface{}) error {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection("timeline")
	if _, err := collection.InsertOne(ctx, data); err != nil {
		log.Error("insert timeline data failed:", err.Error())
		return err
	}
	return nil
}

// RemoveTimeline
func RemoveTimeline() error {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection("timeline")
	if _, err := collection.UpdateOne(ctx, bson.M{}, bson.D{}); err != nil {

	}
	return nil
}

// DeleteRowTimeline
func DeleteRowTimeline(filter interface{}) error {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	collection := db.Collection("timeline")
	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		log.Error("delete whole timeline of a user failed,", err.Error())
		return err
	}
	return nil
}
