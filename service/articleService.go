package service

import (
	"Moments/model"
	"Moments/pkg/log"
	"fmt"

	"github.com/pkg/errors"
)

type ArticleHandler struct {
	*model.Article
}

func (handler *ArticleHandler) DetailArticle() error {
	am := model.NewArticleModel(model.ArticleModelImpl{})
	detail, err := am.GetArticleDetailByAid(handler.Aid)
	if err != nil {
		log.Warn("get article detail by aid failed, aid:", handler.Aid)
		return err
	}
	handler.Aid = detail.Aid
	handler.Uid = detail.Uid
	handler.Content = detail.Content
	handler.PostTime = detail.PostTime
	handler.PhotoList = detail.PhotoList
	handler.Privacy = detail.Privacy
	handler.IsDeleted = detail.IsDeleted
	return nil
}

func (handler *ArticleHandler) AddArticle() error {
	am := model.NewArticleModel(model.ArticleModelImpl{})
	aid, err := am.GenerateAid(handler.Uid)
	if err != nil {
		return err
	}
	handler.Aid = aid
	err = am.AddArticle(handler.Article)
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
func (handler *ArticleHandler) DeleteArticle() error {
	err := handler.Delete(true)
	if err != nil {
		log.Warn("delete article failed")
		return errors.Wrap(err, "delete article failed")
	}
	return nil
}

func (handler *ArticleHandler) Delete(isSoftDelete bool) error {
	var err error
	uid, aid := handler.Uid, handler.Aid
	am := model.NewArticleModel(model.ArticleModelImpl{})
	if isSoftDelete {
		err = am.DeleteArticleSoftByUidAid(uid, aid)
	} else {
		err = am.DeleteArticleByUidAid(uid, aid)
		if err != nil {
			log.Error("delete article failed")
		}

		err = model.RemoveAlbum(map[string]interface{}{"uid": uid}, aid)
		if err != nil {
			log.Error("remove album failed")
		}
	}

	if err != nil {
		log.Error(fmt.Sprintf("delete article failed, aid=%d, error: %s", aid, err.Error()))
		return err
	}
	return nil
}

func (handler *ArticleHandler) Comment() error {
	return nil
}

func (handler *ArticleHandler) Like() error {
	return nil
}
