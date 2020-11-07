package service

import (
	"Moments/models"
	"Moments/pkg/log"
	"Moments/pkg/utils"
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

// generateAid generate global unique aid，rule as 9 + uid(four digit) + timestamp
// also restrict article sending frequency of one user to 1 time per second
func generateAid(uid int32) (int64, error) {
	ts := string(time.Now().Unix())
	tmp := "9" + string(uid) + ts
	aid := utils.Str(tmp).MustInt64()
	dbname := getDatabaseName(aid)
	yes := models.IsAidExist(dbname, aid)
	if !yes {
		return 0, errors.New("aid already existed")
	}
	return aid, nil
}

func (a *Article) GetDetailByAid() (*models.Article, error) {
	dbname := getDatabaseName(a.Aid)
	article, err := models.GetDetail(dbname, bson.M{"aid": a.Aid})
	if err != nil {
		log.Error("get article detail by aid failed, aid: ", a.Aid)
		return nil, err
	}
	return article, nil
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
	log.Info(dbname, article)
	err = models.AddArticle(dbname, article)
	return err
}

func (a *Article) DeleteByAid() error {
	return nil
}

func (a *Article) Comment() error {
	return nil
}

func (a *Article) Like() error {
	return nil
}
