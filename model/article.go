package model

import (
	"Moments/pkg/log"
	"fmt"
	"strconv"
)

type ArticleModel interface {
	GetArticleDatabase() string
	IsArticleExist() bool
	GetArticleDetailByAid() error
	AddArticle() error
	//UpdateArticle(filter Map, data Map) error
	DeleteArticleByAid() error
	DeleteArticleSoftByAid() error
}

// article data structure
type Article struct {
	// basic fields of Article, it should not be edited
	Aid       int64    `bson:"aid" json:"aid"`
	Uid       int32    `bson:"uid" json:"uid"`
	PostTime  int64    `bson:"post_time" json:"post_time"`
	Content   string   `bson:"content" json:"content"`
	PhotoList []string `bson:"photo_list" json:"photo_list"`

	// extra fields
	Privacy   int32 `bson:"privacy" json:"privacy"`
	IsDeleted int32 `bson:"is_deleted" json:"is_deleted"`
}

// GetArticleDatabase articles has been split into 4 collections, find the correct collection
func (a *Article) GetArticleDatabase() string {
	dbname := "article_" + strconv.Itoa(int(a.Aid%4))
	return dbname
}

// IsArticleExist check if specific aid is already existed
func (a *Article) IsArticleExist() bool {
	client := NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		// shall not return false, otherwise aid would be redundant
		return true
	}
	defer client.Disconnect()

	rows, err := client.Query(a.GetArticleDatabase(), Map{"aid": a.Aid})
	if err != nil {
		log.Info("aid not exists")
		return false
	}
	if len(rows) <= 0 {
		log.Info("aid not exists")
		return false
	}
	return true
}

// GetArticleDetailByAid get detail of an Article with specific filter
func (a *Article) GetArticleDetailByAid() error {
	client := NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	dbname := a.GetArticleDatabase()
	rows, err := client.Query(dbname, Map{"aid": a.Aid})
	if err != nil {
		log.Info(fmt.Sprintf("Find() method in databese \"%s\" error: %s", dbname, err.Error()))
		return err
	}
	row := rows[0]
	a.Aid = row["aid"].(int64)
	a.Uid = row["uid"].(int32)
	a.PostTime = row["post_time"].(int64)
	a.Content = row["content"].(string)
	a.PhotoList = row["photo_list"].([]string)
	a.Privacy = row["privacy"].(int32)
	a.IsDeleted = 0
	return nil
}

// AddArticle add Article to database, and expand this Article into friends' timeline
func (a *Article) AddArticle() error {
	client := NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	err = client.Insert(a.GetArticleDatabase(), []interface{}{a})
	if err != nil {
		log.Error("add Article failed,", err.Error())
	}
	return err
}

// DeleteArticleByUidAid delete Article permanently
func (a *Article) DeleteArticleByUidAid() error {
	client := NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	err = client.Remove(a.GetArticleDatabase(), Map{"aid": a.Aid, "uid": a.Uid})
	if err != nil {
		log.Error("permanently delete Article failed, error:", err.Error())
	}
	log.Info("delete Article success")
	return err
}

// DeleteArticleSoftByUidAid delete Article softly
func (a *Article) DeleteArticleSoftByUidAid() error {
	client := NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	filter := Map{"aid": a.Aid, "uid": a.Uid}
	data := Map{"is_deleted": 1}
	err = client.Update(a.GetArticleDatabase(), filter, data)
	if err != nil {
		log.Error("softly delete Article failed, error:", err.Error())
	}
	log.Info("softly delete Article success")
	return err
}
