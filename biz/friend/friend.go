package model

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Friend struct {
	Uid        int32   `json:"uid"`
	FriendList []int32 `json:"friend_list"`
}

var collectionName = "friend"

func GetFriend(uid int32) ([]int32, error) {
	row, err := queryOne(collectionName, bson.M{"uid": uid})
	if err != nil {
		return make([]int32, 0, 0), err
	}
	return row["friend_list"].([]int32), nil
}

func AddNewFriend(uid int32, fuid int32) error {
	row, err := queryOne(collectionName, bson.M{"uid": uid})
	if err != nil {
		return err
	}

	friendList := row["friend_list"].([]int32)
	friendList = append(friendList, fuid)
	err = update(collectionName, bson.M{"uid": uid}, bson.M{"friend_list": friendList})
	return err
}

func RemoveFriend() error {
	// todo not implement
	return nil
}

func BlockFriend() error {
	// todo not implement
	return nil
}
