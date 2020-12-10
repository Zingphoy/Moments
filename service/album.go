package service

import (
	"Moments/models"

	"go.mongodb.org/mongo-driver/bson"
)

type Album struct {
	Uid int32 `json:"uid"`
	Aid int64 `json:"aid"`
}

// Append append to the article_id (aid) into data row
func (a *Album) Append() error {
	err := models.AppendAlbum(bson.M{"uid": a.Uid}, a.Aid)
	return err
}

// Delete soft-delete a album, do nothing
func (a *Album) Delete() error {
	return nil
}
