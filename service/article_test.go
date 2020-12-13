package service

import (
	"Moments/pkg/log"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	assert2 "github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

/*
	Test Data:
	Article{
		Aid:        ?,
		Uid:        90001,
		Post_time:  time.Now().Unix(),
		Content:    "Unit Test From Function TestArticle_Add()",
		Photo_list: bson.A{"http://www.baidu.com"},
		Privacy:    0,
	}
*/

var (
	globalTestAid int64
)

//func TestMain(m *testing.M) {
//	log.InitLogger(true)
//	exitCode := m.Run()
//	os.Exit(exitCode)
//}

func init() {
	log.InitLogger(true)
}

func TestRefreshTimeline(t *testing.T) {

}

func TestArticle_Add(t *testing.T) {
	td := Article{
		Uid:        90001,
		Post_time:  time.Now().Unix(),
		Content:    "Unit Test From Function TestArticle_Add()",
		Photo_list: bson.A{"http://www.baidu.com"},
		Privacy:    9,
	}
	if err := td.Add(); err != nil {
		assert2.Error(t, err, " add article failed")
		return
	}

	article := Article{Aid: td.Aid}
	err := article.GetDetailByAid()
	if err != nil {
		assert2.Error(t, err, " get article detail failed")
		return
	}
	globalTestAid = td.Aid
	assert.Equal(t, article.Content, td.Content)
}

func TestArticle_DeleteByAid(t *testing.T) {
	td := Article{
		Aid:     globalTestAid,
		Content: "Unit Test From Function TestArticle_Add()",
	}

	//soft-delete
	if err := td.DeleteByAid(true); err != nil {
		assert2.Error(t, err, "softly delete article failed ")
	}

	article := Article{Aid: td.Aid}
	err := article.GetDetailByAid()
	if err != nil {
		assert2.Error(t, err, " get article detail failed")
		return
	}
	assert.Equal(t, article.Content, td.Content)

	// delete pernamently
	if err := td.DeleteByAid(false); err != nil {
		assert2.Error(t, err, "permanently delete article failed ")
		return
	}
	err = td.GetDetailByAid() // expect variable err to be not nil
	assert2.NotNil(t, err)
}

func TestArticle_RefreshTimeline(t *testing.T) {

}

// db.article_0.find()
// db.album.find()
// db.article_0.remove({"is_deleted":0})
