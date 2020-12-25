package model

/*
 timeline is sorted in order depends on its timing of insertion. When getting timeline from database, it will
 offer the newest 10 (if needs) since the latest at local cache.
*/

import (
	"Moments/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
)

type Timeline struct {
	Uid     int32   `json:"uid"`
	AidList []int64 `json:"aid_list"`
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

// GetTimelineRefreshByUid when client refresh timeline, it will fetch 10 newest article and check privacy,
// finally offer the matched article regardless of amount. Don't deal with the case when less than 1 will be returned.
func GetTimelineRefreshByUid(uid int32, aid int64) ([]int64, error) {
	filter := map[string]interface{}{"uid": uid}
	row, err := queryOne("timeline", filter)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	list := (row["aid_list"]).(bson.A)
	aids := make([]int64, 0, len(list))
	for _, v := range list {
		n := v.(int64)
		if n <= aid {
			break
		}
		// todo checkPrivacy here
		aids = append(aids, v.(int64))
	}
	return aids, nil
}

// GetTimelineLoadMoreByUid almost the same as GetTimelineRefreshByUid, for the opposite operation
func GetTimelineLoadMoreByUid(uid int32, aid int64) ([]int64, error) {
	filter := map[string]interface{}{"uid": uid}
	data, err := queryOne("timeline", filter)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var index int
	list := (data["aid_list"]).(bson.A)
	for i, v := range list {
		if v == aid {
			index = i
		}
	}

	//var aids []int64
	aids := make([]int64, 0, len(list))
	count := 10
	for i := 1; index+i < len(list) && i < count+1; i++ {
		// todo checkPrivacy here
		aids = append(aids, list[index+i].(int64))
	}
	// todo 边界条件aid会导致bug
	return aids, nil
}

// AppendTimeline append timeline to existing user
func AppendTimeline(uid int32, aid int64) error {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	var row bson.M
	collection := db.Collection("timeline")
	err := collection.FindOne(ctx, bson.M{"uid": uid}).Decode(&row)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	aids := []int64{aid}
	list := row["aid_list"].(bson.A)
	for _, v := range list {
		aids = append(aids, v.(int64))
	}
	data := bson.D{{"$set",
		bson.D{
			{"aid_list", aids},
		},
	}}
	_, err = collection.UpdateOne(ctx, bson.M{"uid": uid}, data)
	if err != nil {
		log.Error("append timeline failed,", err.Error())
		return err
	}
	return nil
}

// InsertNewTimeline insert a new timeline for a new user
func InsertNewTimeline(uid int32, aids []int64) error {
	data := map[string]interface{}{"uid": uid, "aid_list": aids}
	err := insert("timeline", data)
	if err != nil {
		log.Error("insert timeline data failed,", err.Error())
		return err
	}
	return nil
}

// RemoveTimeline remove one article from timeline
func RemoveTimeline(uid int32, aid int64) error {
	filter := map[string]interface{}{"uid": uid}
	row, err := queryOne("timeline", filter)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	var index int
	list := row["aid_list"].(bson.A)
	for i, v := range list {
		if v == aid {
			index = i
		}
	}
	tmpA := list[0 : index-1]
	tmpB := list[index:len(list)]
	aids := append(tmpA, tmpB...)
	data := map[string]interface{}{"aid_list": aids}
	filter = map[string]interface{}{"uid": uid}
	err = update("timeline", filter, data)
	if err != nil {
		log.Error("append timeline failed,", err.Error())
		return err
	}
	return nil
}

// DeleteRowTimeline todo 这里没封装好，但只用来测试函数
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
