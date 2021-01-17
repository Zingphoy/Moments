package service

import (
	"Moments/model"
	"Moments/pkg/log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

type Album struct {
	Uid     int32  `json:"uid"`
	AidList bson.A `json:"aid_list"` // this will use as one single value if needed
}

func (a *Album) Add() error {
	err := model.NewAlbum(a.Uid)
	return err
}

// Append append to the article_id (aid) into data row
func (a *Album) Append() error {
	err := model.AppendAlbum(bson.M{"uid": a.Uid}, a.AidList[0].(int64))
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			err = a.Add()
			if err != nil {
				log.Error("add album failed,", err.Error())
				return err
			}
			err = model.AppendAlbum(bson.M{"uid": a.Uid}, a.AidList[0].(int64))
			if err != nil {
				log.Error("add album failed,", err.Error())
				return err
			}
		} else {
			log.Error("add album failed,", err.Error())
		}
	}
	return err
}

// Delete delete an article from album
func (a *Album) Delete() error {
	err := model.DeleteAlbum(map[string]interface{}{"uid": a.Uid, "aid": a.AidList[0].(int64)})
	return err
}

func (a *Album) Detail() error {
	album, err := model.DetailAlbum(map[string]interface{}{"uid": a.Uid})
	if err != nil {
		log.Error("get album detail error")
		return err
	}
	a.AidList = album["aid_list"].(bson.A)
	return err
}
