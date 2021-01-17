package service

import (
	"Moments/model"
	"Moments/pkg/log"
)

//type Timeline struct {
//	Uid      int32     `json:"uid"`
//	Articles []Article `json:"articles"`
//}

type TimelineService model.Timeline

func (tl *TimelineService) RefreshTimeline(uid int32, latestAid int64, schema string) error {
	var err error
	var aids []int64

	switch schema {
	case "refresh":
		aids, err = model.GetTimelineRefreshByUid(uid, latestAid)
	case "loadmore":
		aids, err = model.GetTimelineLoadMoreByUid(uid, latestAid)
	}
	if err != nil {
		log.Error("get timeline failed")
		return err
	}

	tl.Articles = []model.Article{}
	for _, aid := range aids {
		article := new(ArticleService)
		article.Aid = aid
		err = article.GetDetailByAid()
		if err != nil {
			log.Error("get article detail failed")
			return err
		}

		tl.Articles = append(tl.Articles, model.Article{
			Aid:       aid,
			Uid:       article.Uid,
			PostTime:  article.PostTime,
			Content:   article.Content,
			PhotoList: article.PhotoList,
			Privacy:   article.Privacy,
		})
	}

	log.Info("get timeline success, aid list:", aids)
	return nil
}
