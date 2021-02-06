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

func NewAlbumService(data *Album, impl *AlbumModelImpl) *AlbumHandler {
	return &AlbumHandler{
		Data: data,
		Impl: impl,
	}
}

func (srv *AlbumHandler) CreateAlbum(c *gin.Context) error {
	err := srv.Impl.CreateAlbumByUid(srv.Data.Uid)
	return err
}

// AppendAlbum append to the article_id (aid) into data row
func (srv *AlbumHandler) AppendAlbum(c *gin.Context) error {
	uid := srv.Data.Uid
	aid := srv.Data.AidList[0]
	_, err := srv.Impl.GetAlbumDetailByUid(uid)

	// if album is empty, create one and then append aid
	if err != nil && err.(hint.CustomError).Code == hint.ALBUM_EMPTY {
		err = srv.Impl.CreateAlbumByUid(uid)
		if err != nil {
			log.Error(c, "add album failed,", err.Error())
			return err
		}
		err = srv.Impl.AppendAlbumByUidAid(uid, aid)
		if err != nil {
			log.Error(c, "add album failed,", err.Error())
			return err
		}
	}
	return err
}

// DeleteArticleInAlbum delete an article from album
func (srv *AlbumHandler) DeleteArticleInAlbum(c *gin.Context) error {
	aid := srv.Data.AidList[0]
	err := srv.Impl.RemoveArticleInAlbumByUidAid(srv.Data.Uid, aid)
	return err
}

func (srv *AlbumHandler) DetailAlbum(c *gin.Context) error {
	album, err := srv.Impl.GetAlbumDetailByUid(srv.Data.Uid)
	if err != nil {
		log.Error(c, "get album detail error")
		return err
	}
	srv.Data.AidList = album.AidList
	return err
}
