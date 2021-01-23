package model

import (
	"Moments/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
)

type Album struct {
	Uid     int32  `json:"uid"`
	AidList bson.A `json:"aid_list"` // this will use as one single value if needed
}

// NewAlbum add new album for a new user
func NewAlbum(uid int32) error {
	err := insert("album", bson.M{"uid": uid, "aid_list": bson.A{}})
	return err
}

// AppendAlbum append aid to user's specific Article album
func AppendAlbum(filter map[string]interface{}, aid int64) error {
	aids, err := queryOne("album", filter)
	if err != nil {
		return err
	}

	// todo 追加aid有问题
	tempList, ok := aids["aid_list"].(bson.A)
	if !ok {
		log.Error("aid_list is not slice")
	}
	tempList = append(tempList, aid)
	data := map[string]interface{}{"aid_list": tempList}
	err = update("album", filter, data)
	return err
}

// RemoveAlbum delete Article from album permanently
func RemoveAlbum(filter map[string]interface{}, aid int64) error {
	aids, err := queryOne("album", filter)
	if err != nil {
		return err
	}

	tempList, ok := aids["aid_list"].(bson.A)
	if !ok {
		log.Error("aid_list is not slice")
	}

	var list bson.A
	for _, v := range tempList {
		if v == aid {
			continue
		}
		list = append(list, v)
	}

	data := map[string]interface{}{"aid_list": tempList}
	err = update("album", bson.M{"aid": filter["aid"]}, data)
	return err
}

func DetailAlbum(filter map[string]interface{}) (map[string]interface{}, error) {
	album, err := queryOne("album", filter)
	if err != nil {
		log.Error("get album detail failed,", err.Error())
		return nil, err
	}
	return album, err
}
