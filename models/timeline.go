package models

type Timeline struct {
	Id      interface{} `json:"-"`
	Uid     int32         `json:"uid"`
	AidList []int64     `json:"aid_list"`
}
