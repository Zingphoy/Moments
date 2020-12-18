package service

import (
	"Moments/models"
	"Moments/pkg/log"
)

type Timeline struct {
	Uid      int32     `json:"uid"`
	Articles []Article `json:"articles"`
}

func (tl *Timeline) RefreshTimeline(uid int32, latestAid int64, schema string) error {
	var err error
	var aids []int64

	switch schema {
	case "refresh":
		aids, err = models.GetTimelineRefreshByUid(uid, latestAid)
	case "loadmore":
		aids, err = models.GetTimelineLoadMoreByUid(uid, latestAid)
	}
	if err != nil {
		log.Error("get timeline failed")
		return err
	}

	tl.Articles = []Article{}
	for _, aid := range aids {
		articleSrv := Article{Aid: aid}
		err = articleSrv.GetDetailByAid()
		if err != nil {
			log.Error("get article detail failed")
			return err
		}

		tl.Articles = append(tl.Articles, Article{
			Aid:        aid,
			Uid:        articleSrv.Uid,
			Post_time:  articleSrv.Post_time,
			Content:    articleSrv.Content,
			Photo_list: articleSrv.Photo_list,
			Privacy:    articleSrv.Privacy,
		})
	}

	log.Info("get timeline success, aid list:", aids)
	return nil
}
