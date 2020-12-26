package service

import (
	"Moments/model"
	"Moments/pkg/log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	log.InitLogger(true)
}

func mockTestData4ArticleService() *Article {
	testArticle := Article{
		Uid:       90000,
		PostTime:  int64(1608961415),
		Content:   "Unit Test From Function TestArticle_Add(), should be deleted after test",
		PhotoList: bson.A{"https://www.baidu.com"},
		Privacy:   999, // this is the sign of test
	}
	return &testArticle
}

func clearTestData4ArticleService(aid int64) error {
	db, client, ctx, _ := model.ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	dbname := getDatabaseName(aid)
	collection := db.Collection(dbname)
	_, err := collection.DeleteOne(ctx, bson.M{"privacy": 999})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func TestArticle_Add(t *testing.T) {
	testArticle := mockTestData4ArticleService()
	err := testArticle.Add()
	assert.Nil(t, err)
	if err != nil {
		return
	}

	article := Article{Aid: testArticle.Aid}
	err = article.GetDetailByAid()
	assert.Nil(t, err)
	assert.Equal(t, article.Content, testArticle.Content)

	err = clearTestData4ArticleService(testArticle.Aid)
	assert.Nil(t, err)
}

func TestArticle_DeleteByAid(t *testing.T) {
	//soft-delete
	testArticle := mockTestData4ArticleService()
	err := testArticle.Add()
	assert.Nil(t, err)
	if err != nil {
		return
	}

	err = testArticle.DeleteByAid(true)
	assert.Nil(t, err)
	if err != nil {
		return
	}

	// data should exist in database
	err = testArticle.GetDetailByAid()
	assert.Nil(t, err)
	assert.Equal(t, testArticle.Content, testArticle.Content)

	// delete permanently
	err = testArticle.DeleteByAid(false)
	assert.Nil(t, err)
	if err != nil {
		return
	}

	// check whether has been deleted
	err = testArticle.GetDetailByAid() // expect variable err to be not nil
	assert.NotNil(t, err)
}

func TestArticle_RefreshTimeline(t *testing.T) {
	// todo add test
	//RefreshTimeline("refresh")
	//RefreshTimeline("loadmore")
}
