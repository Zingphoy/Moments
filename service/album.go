package service

import (
	"Moments/models"

	"go.mongodb.org/mongo-driver/bson"
)

type Album struct {
	Uid int32 `json:"uid"`
	Aid int64 `json:"aid"`
}

func (a *Album) Append() error {
	err := models.AppendAlbum(bson.M{"uid": a.Uid}, a.Aid)
	return err
}

func (a *Album) Delete() error {
	return nil
}
