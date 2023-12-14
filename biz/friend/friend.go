package friend

import (
	"Moments/biz/database"
	"Moments/pkg/hint"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type Friend struct {
	Uid        int32   `json:"uid"`
	FriendList []int32 `json:"friend_list"`
}

var collectionName = "friend"

func (f *Friend) GetFriend(uid int32) (*Friend, error) {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect()

	row, err := client.Query(collectionName, bson.M{"uid": uid})
	if err != nil {
		return nil, err
	}
	friend := Friend{Uid: uid}
	friend.FriendList = row[0]["friend_list"].([]int32)
	// todo 可能是查不到这个user，或者user没有好友
	if len(friend.FriendList) == 0 {
		return nil, hint.CustomError{
			Code: hint.USER_NOT_FOUND,
			Err:  errors.New("empty friend list"),
		}
	}
	return &friend, nil
}

func (f *Friend) AddNewFriend(uid int32, fuid int32) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	row, err := client.Query(collectionName, bson.M{"uid": uid})
	if err != nil {
		return err
	}
	friendList := row[0]["friend_list"].([]int32)
	friendList = append(friendList, fuid)
	return client.Update(collectionName, bson.M{"uid": uid}, bson.M{"friend_list": friendList})
}

func (f *Friend) RemoveFriend() error {
	// todo not implement
	return nil
}

func (f *Friend) BlockFriend() error {
	// todo not implement
	return nil
}
