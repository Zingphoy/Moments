package album

import (
	"Moments/biz/database"
	"Moments/pkg/hint"
	"Moments/pkg/log"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type Album struct {
	Uid     int32   `bson:"uid" json:"uid"`
	AidList []int64 `bson:"aid_list" json:"aid_list"`
}

type AlbumModel interface {
	GetAlbumDetailByUid(uid int32) (*Album, error)
	CreateAlbumByUid(uid int32) error
	AppendAlbumByUidAid(uid int32, aid int64) error
	RemoveArticleInAlbumByUidAid(uid int32, aid int64) error
}

type AlbumModelImpl struct {
}

var (
	dbname = "album"
)

func (a *AlbumModelImpl) GetAlbumDetailByUid(uid int32) (*Album, error) {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect()

	album, err := client.Query(dbname, database.Map{"uid": uid})
	if err != nil {
		log.Error(nil, "get album detail failed,", err.Error())
		return nil, err
	}

	if len(album) == 0 {
		return nil, hint.CustomError{
			Code: hint.ALBUM_EMPTY,
			Err:  errors.New("album is empty"),
		}
	}

	ret := Album{
		Uid:     album[0]["uid"].(int32),
		AidList: database.BsonToSliceInt64(album[0]["aid_list"].(bson.A)),
	}
	return &ret, err
}

// CreateAlbumByUid add new album for a new user
func (a *AlbumModelImpl) CreateAlbumByUid(uid int32) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	err = client.Insert(dbname, []interface{}{database.Map{"uid": uid, "aid_list": bson.A{}}})
	return err
}

// AppendAlbumByUidAid append aid to user's specific Article album
func (a *AlbumModelImpl) AppendAlbumByUidAid(uid int32, aid int64) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	filter := database.Map{"uid": uid}
	albums, err := client.Query(dbname, filter)
	if err != nil {
		return err
	}

	// todo 追加aid有问题
	log.Info(nil, "88 ", albums)
	log.Info(nil, "89 ", albums[0])
	tempList, ok := albums[0]["aid_list"].(bson.A)
	if !ok {
		log.Error(nil, "aid_list is not slice")
	}
	tempList = append(tempList, aid)
	return client.Update("album", filter, bson.M{"aid_list": tempList})
}

// RemoveArticleInAlbumByUidAid delete Article from album permanently
func (a *AlbumModelImpl) RemoveArticleInAlbumByUidAid(uid int32, aid int64) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	filter := database.Map{"uid": uid}
	aids, err := client.Query(dbname, filter)
	if err != nil {
		return err
	}

	tempList, ok := aids[0]["aid_list"].(bson.A)
	if !ok {
		log.Error(nil, "aid_list is not slice")
	}

	var head bson.A
	var tail bson.A
	for i, v := range tempList {
		if v == aid {
			tail = tempList[i+1:]
			break
		}
		head = append(head, v)
	}
	head = append(head, tail)
	return client.Update(dbname, bson.M{"aid": filter["aid"]}, bson.M{"aid_list": head})
}

// DeleteAlbumByUid totally delete a whole album
func (a *AlbumModelImpl) DeleteAlbumByUid(uid int32) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	filter := database.Map{"uid": uid}
	err = client.Remove(dbname, filter)
	if err != nil {
		return err
	}
	return nil
}
