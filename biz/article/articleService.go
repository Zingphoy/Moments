package article

import (
	"Moments/biz/album"
	"Moments/pkg/log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

/*
 todo Model层不再做过度抽象，直接一层interface加一个实现的结构体，service层的结构体接受model层结构体为参数，在具体函数中条用model层结构体即可
*/

type ArticleHandler struct {
	Data *Article
	Impl ArticleModel
}

// todo 改成依赖注入
func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{
		Data: &Article{},
		Impl: &ArticleModelImpl{},
	}
}

func (handler *ArticleHandler) DetailArticle(c *gin.Context) error {
	detail, err := handler.Impl.GetArticleDetailByAid(handler.Data.Aid)
	if err != nil {
		log.Warn(c, "get article detail by aid failed, aid:", handler.Data.Aid)
		return err
	}
	handler.Data.Aid = detail.Aid
	handler.Data.Uid = detail.Uid
	handler.Data.Content = detail.Content
	handler.Data.PostTime = detail.PostTime
	handler.Data.PhotoList = detail.PhotoList
	handler.Data.Privacy = detail.Privacy
	handler.Data.IsDeleted = detail.IsDeleted
	return nil
}

func (handler *ArticleHandler) AddArticle(c *gin.Context) error {
	aid, err := handler.Impl.GenerateAid(handler.Data.Uid)
	if err != nil {
		return err
	}
	handler.Data.Aid = aid
	err = handler.Impl.AddArticle(handler.Data)
	if err != nil {
		return errors.Wrap(err, "add article failed")
	}

	//album := Album{
	//	Uid:     a.Uid,
	//	AidList: bson.A{a.Aid},
	//}
	//err = album.Append()
	//if err != nil {
	//	return err
	//}

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
func (handler *ArticleHandler) DeleteArticle(c *gin.Context) error {
	err := handler.Delete(c, true)
	if err != nil {
		log.Warn(c, "delete article failed")
		return errors.Wrap(err, "delete article failed")
	}
	return nil
}

func (handler *ArticleHandler) Delete(c *gin.Context, isSoftDelete bool) error {
	var err error
	uid, aid := handler.Data.Uid, handler.Data.Aid
	if isSoftDelete {
		err = handler.Impl.DeleteArticleSoftByUidAid(uid, aid)
	} else {
		err = handler.Impl.DeleteArticleByUidAid(uid, aid)
		if err != nil {
			log.Error(c, "delete article failed")
		}

		albumHandler := album.NewAlbumHandler()
		albumHandler.Data.Uid = uid
		albumHandler.Data.AidList = append(albumHandler.Data.AidList, aid)
		err = albumHandler.DeleteAlbum(c)
		if err != nil {
			log.Error(c, "remove album failed")
		}
	}

	if err != nil {
		log.Error(c, fmt.Sprintf("delete article failed, aid=%d, error: %s", aid, err.Error()))
		return err
	}
	return nil
}

func (handler *ArticleHandler) CommentArticle(c *gin.Context) error {
	return nil
}

func (handler *ArticleHandler) LikeArticle(c *gin.Context) error {
	return nil
}
