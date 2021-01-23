package model

/*
 timeline is sorted in order depends on its timing of insertion.
 When getting timeline from database, it will offer the newest 10 (if needs) since the latest at local cache.
*/

import (
	"Moments/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	AUTHORIZED  = 0
	DENIED      = 1
	timelineClt = "timeline"
)

type TimelineObject interface {
	GetUserTimelineRefresh(aid int64) ([]int64, error)
	GetUserTimelineLoadMore(aid int64) ([]int64, error)
	AppendUserTimeline(aid int64) error
	RemoveUserTimeline(aid int64) error
}

type Timeline struct {
	Uid      int32      `json:"uid"`
	Articles []*Article `json:"articles"`
}

// todo 补充鉴权中间件
// params: uid int32
func hasPermission() error {
	return nil
}

func NewTimelineObject(data map[string]interface{}) TimelineObject {
	tl := Timeline{}
	if _, ok := data["uid"]; ok {
		tl.Uid = data["uid"].(int32)
	}
	if _, ok := data["articles"]; ok {
		tl.Articles = data["articles"].([]*Article)
	}
	return &tl
}

// GetUserTimelineArticleDetail get details of all articles from timeline
func (tl *Timeline) GetUserTimelineArticleDetail() ([]*Article, error) {
	err := Client.Connect()
	if err != nil {
		return nil, err
	}
	defer Client.Disconnect()

	filter := map[string]interface{}{"uid": tl.Uid}
	data, err := Client.Query(timelineClt, filter)
	if err != nil {
		return nil, err
	}

	var ret = make([]*Article, len(data))
	for _, aid := range data {
		// todo
		// 通过aid获取database
		// 从database中获取文章信息
		log.Info(aid)
		ret = append(ret, nil) // todo article要能拿到detail再塞进来
	}
	return ret, nil
}

// GetUserTimelineRefresh when client refresh timeline, it will fetch 10 newest Article and check privacy,
// finally offer the matched Article regardless of amount. Don't deal with the case when less than 1 will be returned.
func (tl *Timeline) GetUserTimelineRefresh(aid int64) ([]int64, error) {
	err := Client.Connect()
	if err != nil {
		return nil, err
	}
	defer Client.Disconnect()

	filter := map[string]interface{}{"uid": tl.Uid}
	data, err := Client.Query(timelineClt, filter)
	if err != nil {
		return nil, err
	}

	list := (data[0]["aid_list"]).(bson.A)
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

// GetUserTimelineLoadMore almost the same as GetTimelineRefreshByUid, for the opposite operation
func (tl *Timeline) GetUserTimelineLoadMore(aid int64) ([]int64, error) {
	err := Client.Connect()
	if err != nil {
		return nil, err
	}
	defer Client.Disconnect()

	filter := map[string]interface{}{"uid": tl.Uid}
	data, err := Client.Query(timelineClt, filter)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var index int
	list := (data[0]["aid_list"]).(bson.A)
	for i, v := range list {
		if v == aid {
			index = i
		}
	}

	aids := make([]int64, 0, len(list))
	count := 10
	for i := 1; index+i < len(list) && i < count+1; i++ {
		// todo checkPrivacy here
		aids = append(aids, list[index+i].(int64))
	}
	// todo 边界条件aid会导致bug
	return aids, nil
}

// AppendUserTimeline append timeline to existing user
func (tl *Timeline) AppendUserTimeline(aid int64) error {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database:", err.Error())
		}
	}()

	var row bson.M
	collection := db.Collection("timeline")
	err := collection.FindOne(ctx, bson.M{"uid": tl.Uid}).Decode(&row)
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
	_, err = collection.UpdateOne(ctx, bson.M{"uid": tl.Uid}, data)
	if err != nil {
		log.Error("append timeline failed,", err.Error())
		return err
	}
	return nil
}

// RemoveTimeline remove one Article from timeline
func (tl *Timeline) RemoveUserTimeline(aid int64) error {
	filter := map[string]interface{}{"uid": tl.Uid}
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
	filter = map[string]interface{}{"uid": tl.Uid}
	err = update("timeline", filter, data)
	if err != nil {
		log.Error("append timeline failed,", err.Error())
		return err
	}
	return nil
}

// InsertNewTimeline insert a new timeline for a new user
func (tl *Timeline) InsertNewTimeline(uid int32, aids []int64) error {
	data := map[string]interface{}{"uid": uid, "aid_list": aids}
	err := insert("timeline", data)
	if err != nil {
		log.Error("insert timeline data failed,", err.Error())
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
