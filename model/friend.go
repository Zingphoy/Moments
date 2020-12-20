package model

type Friend struct {
	Id         interface{} `json:"-"`
	Uid        int32         `json:"uid"`
	FriendList []int       `json:"friend_list"`
}
