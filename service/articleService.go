package service

import (
	"Moments/model"
	"Moments/pkg/log"
	"fmt"
	"github.com/pkg/errors"
)

type ArticleHandler struct {
	model.Article
}

func (handler *ArticleHandler) DetailArticle() error {
	err := handler.Article.GetArticleDetailByAid()
	if err != nil {
		log.Warn("get article detail by aid failed, aid:", handler.Aid)
		return err
	}
	return nil
}

func (handler *ArticleHandler) AddArticle(data Map) error {
	aid, err := generateAid(data["uid"].(int32))
	if err != nil {
		return err
	}
	handler.Aid = aid
	err = handler.Article.AddArticle()
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

// DeleteArticle delete an article by aid softly
func (handler *ArticleHandler) DeleteArticle() error {
	err := handler.Article.DeleteArticleSoftByUidAid()
	if err != nil {
		log.Warn("delete article failed")
		return errors.Wrap(err, "delete article failed")
	}
	return nil
}

func (handler *ArticleHandler) Delete(uid int32, aid int64, isSoftDelete bool) error {
	var err error
	if isSoftDelete {
		err = handler.DeleteArticleSoftByUidAid()
	} else {
		err = handler.DeleteArticleByUidAid()
		if err != nil {
			log.Error("delete article failed")
		}

		err = model.RemoveAlbum(Map{"uid": uid}, aid)
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
