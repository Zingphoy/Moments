package models

import (
	"Moments/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
)

type Album struct {
	Uid     int32  `json:"uid"`
	AidList bson.A `json:"aid_list"` // different
}

func makeAlbumObj() {

}

func AddAlbum() error {
	return nil
}

// AppendAlbum append aid to user's specific article album
func AppendAlbum(filter bson.M, aid int64) error {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	var aids bson.M
	collection := db.Collection("album")
	err := collection.FindOne(ctx, filter).Decode(&aids)
	if err != nil {
		return err
	}

	tempList, ok := aids["aid_list"].(bson.A)
	if !ok {
		log.Error("aid_list is not slice")
	}
	tempList = append(tempList, aid)
	update := bson.D{{"$set",
		bson.D{
			{"aid_list", tempList},
		},
	}}
	_, err = collection.UpdateOne(ctx, filter, update)
	return err
}
