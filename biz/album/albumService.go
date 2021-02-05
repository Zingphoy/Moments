package album

import (
	"Moments/pkg/hint"
	"Moments/pkg/log"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	Data *Album
	Impl AlbumModel
}

func NewAlbumService() *AlbumHandler {
	return &AlbumHandler{
		Data: &Album{},
		Impl: &AlbumModelImpl{},
	}
}

func (handler *AlbumHandler) CreateAlbum(c *gin.Context) error {
	err := handler.Impl.CreateAlbumByUid(handler.Data.Uid)
	return err
}

// AppendAlbum append to the article_id (aid) into data row
func (handler *AlbumHandler) AppendAlbum(c *gin.Context) error {
	uid := handler.Data.Uid
	aid := handler.Data.AidList[0]
	_, err := handler.Impl.GetAlbumDetailByUid(uid)

	// if album is empty, create one and then append aid
	if err != nil && err.(hint.CustomError).Code == hint.ALBUM_EMPTY {
		err = handler.Impl.CreateAlbumByUid(uid)
		if err != nil {
			log.Error(c, "add album failed,", err.Error())
			return err
		}
		err = handler.Impl.AppendAlbumByUidAid(uid, aid)
		if err != nil {
			log.Error(c, "add album failed,", err.Error())
			return err
		}
	}
	return err
}

// DeleteArticleInAlbum delete an article from album
func (handler *AlbumHandler) DeleteArticleInAlbum(c *gin.Context) error {
	aid := handler.Data.AidList[0]
	err := handler.Impl.RemoveArticleInAlbumByUidAid(handler.Data.Uid, aid)
	return err
}

func (handler *AlbumHandler) DetailAlbum(c *gin.Context) error {
	album, err := handler.Impl.GetAlbumDetailByUid(handler.Data.Uid)
	if err != nil {
		log.Error(c, "get album detail error")
		return err
	}
	handler.Data.AidList = album.AidList
	return err
}
