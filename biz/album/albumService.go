package album

import (
	"Moments/pkg/hint"
	"Moments/pkg/log"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AlbumService struct {
	Data *Album
	Impl AlbumModel
}

func NewAlbumService(data *Album, impl *AlbumModelImpl) *AlbumService {
	return &AlbumService{
		Data: data,
		Impl: impl,
	}
}

func (srv *AlbumService) CreateAlbum(c *gin.Context) error {
	err := srv.Impl.CreateAlbumByUid(srv.Data.Uid)
	return err
}

// AppendAlbum append to the article_id (aid) into data row
func (srv *AlbumService) AppendAlbum(c *gin.Context) error {
	uid := srv.Data.Uid
	aid := srv.Data.AidList[0]
	_, err := srv.Impl.GetAlbumDetailByUid(uid)

	// if album is empty, create one and then append aid
	if err != nil && err.(hint.CustomError).Code == hint.ALBUM_EMPTY {
		err = srv.Impl.CreateAlbumByUid(uid)
		if err != nil {
			log.Error(c, fmt.Sprintf("create new album failed, uid=%d", srv.Data.Uid))
			return err
		}
	}
	err = srv.Impl.AppendAlbumByUidAid(uid, aid)
	if err != nil {
		log.Error(c, fmt.Sprintf("append article into album failed, uid=%d", srv.Data.Uid))
		return err
	}
	return err
}

// DeleteArticleInAlbum delete an article from album
func (srv *AlbumService) DeleteArticleInAlbum(c *gin.Context) error {
	aid := srv.Data.AidList[0]
	err := srv.Impl.RemoveArticleInAlbumByUidAid(srv.Data.Uid, aid)
	return err
}

func (srv *AlbumService) DetailAlbum(c *gin.Context) error {
	album, err := srv.Impl.GetAlbumDetailByUid(srv.Data.Uid)
	if err != nil {
		return err
	}
	srv.Data.AidList = album.AidList
	return nil
}

func (srv *AlbumService) DeleteWholeAlum(c *gin.Context) error {
	err := srv.Impl.DeleteAlbumByUid(srv.Data.Uid)
	if err != nil {
		return err
	}
	return nil
}
