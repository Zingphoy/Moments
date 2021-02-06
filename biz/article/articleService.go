package article

import (
	"Moments/biz/album"
	"Moments/pkg/log"
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
 todo Model层不再做过度抽象，直接一层interface加一个实现的结构体，service层的结构体接受model层结构体为参数，在具体函数中条用model层结构体即可
*/

type ArticleService struct {
	Data *Article
	Impl ArticleModel
}

func NewArticleService(data *Article, impl *ArticleModelImpl) *ArticleService {
	return &ArticleService{
		Data: data,
		Impl: impl,
	}
}

func (srv *ArticleService) DetailArticle(c *gin.Context) error {
	detail, err := srv.Impl.GetArticleDetailByAid(srv.Data.Aid)
	if err != nil {
		log.Info(c, "get article detail by aid failed, aid=", srv.Data.Aid)
		return err
	}
	srv.Data.Aid = detail.Aid
	srv.Data.Uid = detail.Uid
	srv.Data.Content = detail.Content
	srv.Data.PostTime = detail.PostTime
	srv.Data.PhotoList = detail.PhotoList
	srv.Data.Privacy = detail.Privacy
	srv.Data.IsDeleted = detail.IsDeleted
	return nil
}

func (srv *ArticleService) AddArticle(c *gin.Context) error {
	aid, err := srv.Impl.GenerateAid(srv.Data.Uid)
	if err != nil {
		return err
	}
	log.Info(c, "generated aid=", aid)
	srv.Data.Aid = aid
	err = srv.Impl.AddArticle(srv.Data)
	if err != nil {
		log.Info(c, fmt.Sprintf("add article failed, uid=%d", srv.Data.Uid))
		return err
	}

	albumSrv := album.NewAlbumService(&album.Album{}, &album.AlbumModelImpl{})
	albumSrv.Data.Uid = srv.Data.Uid
	albumSrv.Data.AidList = []int64{srv.Data.Aid}
	err = albumSrv.AppendAlbum(c)
	if err != nil {
		return err
	}

	// send a message to MQ, going to insert this article into users' friends' timeline
	//msg := mq.Message{
	//	MsgType:  mq.EXPAND_TIMELINE_ADD,
	//	Aid:      aid,
	//	Uid:      a.Uid,
	//	Desc:     "",
	//	NeedSafe: true,
	//}
	//err = msg.ExpandTimeline()
	//if err != nil {
	//	log.Error("Expand article into friends' timeline failed,", err.Error())
	//	return err
	//}
	return nil
}

// DeleteArticle delete an article by aid softly, a wrapper function of Delete
func (srv *ArticleService) DeleteArticle(c *gin.Context) error {
	err := srv.Delete(c, true)
	if err != nil {
		return err
	}
	return nil
}

func (srv *ArticleService) Delete(c *gin.Context, isSoftDelete bool) error {
	var err error
	uid, aid := srv.Data.Uid, srv.Data.Aid
	if isSoftDelete {
		err = srv.Impl.DeleteArticleSoftByUidAid(uid, aid)
		if err != nil {
			log.Error(c, fmt.Sprintf("delete article failed, uid=%d, aid:=%d", uid, aid))
			return err
		}
	} else {
		// won't work online
		err = srv.Impl.DeleteArticleByUidAid(uid, aid)
		if err != nil {
			log.Error(c, fmt.Sprintf("delete article failed, uid=%d, aid:=%d", uid, aid))
			return err
		}

		albumHandler := album.NewAlbumService(&album.Album{}, &album.AlbumModelImpl{})
		albumHandler.Data.Uid = uid
		albumHandler.Data.AidList = append(albumHandler.Data.AidList, aid)
		err = albumHandler.DeleteArticleInAlbum(c)
		if err != nil {
			log.Error(c, fmt.Sprintf("remove article from album failed, uid=%d, aid:=%d", uid, aid))
			return err
		}
	}
	return nil
}

func (srv *ArticleService) CommentArticle(c *gin.Context) error {
	return nil
}

func (srv *ArticleService) LikeArticle(c *gin.Context) error {
	return nil
}
