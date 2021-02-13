package timeline

/*
 timeline is sorted in order  descending order depends on its timing of insertion.
 When getting timeline from database, it will offer the newest 10 (if needs) since the latest at local cache.
*/

import (
	"Moments/biz/database"
	"Moments/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	AUTHORIZED  = 0
	DENIED      = 1
	timelineClt = "timeline"
)

var (
	dbname = "timeline"
)

type Timeline struct {
	Uid     int32   `bson:"uid" json:"uid"`
	AidList []int64 `bson:"aid_list" json:"aid_list"`
}

type TimelineModel interface {
	GetTimelineRefreshByUidAid(uid int32, aid int64) ([]int64, error)
	GetTimelineLoadMoreByUidAid(uid int32, aid int64) ([]int64, error)
	AppendArticleIntoTimelineByUid(uid int32, aid int64) error
	RemoveArticleFromTimelineByUid(uid int32, aid int64) error
	CreateTimelineByUid(uid int32) error
	DeleteTimelineByUid(uid int32) error
}

type TimelineModelImpl struct {
}

// todo 补充鉴权中间件
// params: uid int32
func hasPermission() error {
	return nil
}

// GetTimelineRefreshByUidAid when client refresh timeline, it will fetch 10 newest Article and check privacy,
// finally offer the matched Article regardless of amount. Don't deal with the case when less than 1 will be returned.
func (t *TimelineModelImpl) GetTimelineRefreshByUidAid(uid int32, aid int64) ([]int64, error) {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect()

	data, err := client.Query(dbname, database.Map{"uid": uid})
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

// GetTimelineLoadMoreByUidAid almost the same as GetTimelineRefreshByUid, for the opposite operation
func (t *TimelineModelImpl) GetTimelineLoadMoreByUidAid(uid int32, aid int64) ([]int64, error) {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect()

	data, err := client.Query(dbname, database.Map{"uid": uid})
	if err != nil {
		log.Error(nil, err.Error())
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

// AppendArticleIntoTimelineByUid append timeline to existing user
func (t *TimelineModelImpl) AppendArticleIntoTimelineByUid(uid int32, aid int64) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	data, err := client.Query(dbname, database.Map{"uid": uid})
	if err != nil {
		log.Error(nil, err.Error())
		return err
	}

	aids := []int64{aid}
	list := data[0]["aid_list"].(bson.A)
	for _, v := range list {
		aids = append(aids, v.(int64))
	}
	err = client.Update(dbname, database.Map{"uid": uid}, database.Map{"aid_list": aids})
	if err != nil {
		log.Error(nil, "append timeline failed,", err.Error())
		return err
	}
	return nil
}

// RemoveArticleFromTimelineByUid remove one Article from timeline
func (t *TimelineModelImpl) RemoveArticleFromTimelineByUid(uid int32, aid int64) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	data, err := client.Query(dbname, database.Map{"uid": uid})
	if err != nil {
		log.Error(nil, err.Error())
		return err
	}

	var aids bson.A
	var index int
	list := data[0]["aid_list"].(bson.A)
	if len(list) > 1 {
		for i, v := range list {
			if v == aid {
				index = i
			}
		}
		tmpA := list[0:index]
		tmpB := list[index+1:]
		aids = append(tmpA, tmpB...)
	}
	err = client.Update("timeline", database.Map{"uid": uid}, database.Map{"aid_list": aids})
	if err != nil {
		log.Error(nil, "append timeline failed,", err.Error())
		return err
	}
	return nil
}

// CreateTimelineByUid insert a new timeline for a new user
func (t *TimelineModelImpl) CreateTimelineByUid(uid int32) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	err = client.Insert(dbname, []interface{}{database.Map{"uid": uid, "aid_list": bson.A{}}})
	if err != nil {
		log.Error(nil, "insert timeline data failed,", err.Error())
		return err
	}
	return nil
}

// DeleteTimelineByUid just for test, to delete a timeline permanently
func (t *TimelineModelImpl) DeleteTimelineByUid(uid int32) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	err = client.Remove(dbname, database.Map{"uid": uid})
	if err != nil {
		log.Error(nil, "delete whole timeline of a user failed,", err.Error())
		return err
	}
	return nil
}
