package model

type User struct {
	Id     interface{} `json:"-"`
	Uid    int32         `json:"uid"`
	Name   string      `json:"name"`
	Passwd string      `json:"passwd"`
}
