package service

import (
	"Moments/models"
	"Moments/pkg/log"
	"Moments/pkg/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type Article struct {
	Aid        int64  `json:"aid"`
	Uid        int32  `json:"uid"`
	Post_time  int64  `json:"post_time"`
	Content    string `json:"content"`
	Photo_list bson.A `json:"photo_list"`
	Privacy    int32  `json:"privacy"`
}

// getDatabaseName articles has been split into 4 collections, find the correct collection
func getDatabaseName(aid int64) string {
	return "article_" + strconv.Itoa(int(aid%4))
}

// generateAid generate global unique aid，rule as uid + timestamp
// also restrict article sending frequency of one user to 1 time per second
func generateAid(uid int32) (int64, error) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	tmp := strconv.FormatInt(int64(uid), 10) + ts
	aid := utils.Str(tmp).MustInt64()
	yes := models.IsAidExist(getDatabaseName(aid), aid)
	if yes {
		log.Info("aid already existed")
		return 0, errors.New("aid already existed")
	}
	log.Info("generated aid:", aid)
	return aid, nil
}

func makeArticleObj(a *models.Article) *Article {
	article := Article{
		Aid:        a.Aid,
		Uid:        a.Uid,
		Post_time:  a.Post_time,
		Content:    a.Content,
		Photo_list: a.Photo_list,
		Privacy:    a.Privacy,
	}
	return &article
}

func (a *Article) GetDetailByAid() (*Article, error) {
	dbname := getDatabaseName(a.Aid)
	modelArticle, err := models.GetDetail(dbname, bson.M{"aid": a.Aid})
	if err != nil {
		log.Info("get article detail by aid failed, aid:", a.Aid)
		return nil, err
	}
	return makeArticleObj(modelArticle), nil
}

func (a *Article) RefreshTimeline() error {
	return nil
}

func (a *Article) Add() error {
	aid, err := generateAid(a.Uid)
	if err != nil {
		return err
	}
	a.Aid = aid

	article := map[string]interface{}{
		"aid":        aid,
		"uid":        a.Uid,
		"post_time":  a.Post_time,
		"content":    a.Content,
		"photo_list": a.Photo_list,
		"privacy":    a.Privacy,
	}

	dbname := getDatabaseName(aid)
	log.Info("find database name:", dbname)
	err = models.AddArticle(dbname, article)
	return err
}

func (a *Article) DeleteByAid(isSoftDelete bool) error {
	var err error
	dbname := getDatabaseName(a.Aid)
	if isSoftDelete {
		err = models.DeleteArticleSoft(dbname, bson.M{"aid": a.Aid})
	} else {
		err = models.DeleteArticle(dbname, bson.M{"aid": a.Aid})
	}

	if err != nil {
		log.Error(fmt.Sprintf("delete article failed, aid=%d, error: %s", a.Aid, err.Error()))
		return err
	}
	return nil
}

// DeleteByAidSoft delete an article by aid softly
func (a *Article) DeleteByAidSoft() error {
	return a.DeleteByAid(true)
}

func (a *Article) Comment() error {
	return nil
}

func (a *Article) Like() error {
	return nil
}
