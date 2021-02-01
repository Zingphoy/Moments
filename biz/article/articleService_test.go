package service

import (
	"Moments/mocks"
	"Moments/model"
	"Moments/pkg/log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.InitLogger(true)
	//mq.InitMQ()
	//mq.StopMQ()	// consumer要先start才能shutdown，否则panic
}

func TestArticleSrv_DetailArticle(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	mockArticle := mocks.NewMockArticleModel(m)

	test := struct {
		model.Article
		error
	}{
		Article: model.Article{Aid: 110, Uid: 120, Content: "test mock"},
		error:   nil,
	}
	mockArticle.EXPECT().GetArticleDetailByAid(gomock.Any()).Return(&test.Article, test.error)
	handler := ArticleHandler{Data: &test.Article, Impl: mockArticle}
	err := handler.DetailArticle()
	assert.Nil(t, err)
	assert.Equal(t, test.Article, test.Article)
}

func TestArticleSrv_AddArticle(t *testing.T) {
	//m := gomock.NewController(t)
	//defer m.Finish()
	//mockArticle := mocks.NewMockArticleModel(m)
	//
	//test := struct {
	//}{}
	//mockArticle.EXPECT().AddArticle(gomock.Any()).Return(nil)
	//handler := ArticleHandler{Data: &test.Article, Impl: mockArticle}
	//err := handler.AddArticle()
	//assert.Nil(t, err)

}

func TestArticleSrv_DeleteArticle(t *testing.T) {

}

// 老单测代码废弃，但是暂时保留注释
//func mockTestData4ArticleService() *ArticleService {
//	testArticle := ArticleService{
//		Uid:       90000,
//		PostTime:  time.Now().Unix(),
//		Content:   "Unit Test From Function TestArticle_Add(), should be deleted after test",
//		PhotoList: bson.A{"https://www.baidu.com"},
//		Privacy:   999, // this is the sign of test
//	}
//	return &testArticle
//}
//
//func clearTestData4ArticleService(aid int64, deleteAlbum bool) error {
//	db, client, ctx, _ := model.ConnectDatabase()
//	defer func() {
//		if err := client.Disconnect(ctx); err != nil {
//			log.Error("error while trying to disconnect database: ", err.Error())
//		}
//	}()
//
//	dbname := getDatabaseName(aid)
//	collection := db.Collection(dbname)
//	_, err := collection.DeleteOne(ctx, bson.M{"privacy": 999})
//	if err != nil {
//		log.Fatal(err.Error())
//		return err
//	}
//
//	if deleteAlbum {
//		_, err = db.Collection("album").DeleteOne(ctx, bson.M{"uid": 90000})
//		if err != nil {
//			log.Fatal(err.Error())
//			return err
//		}
//	}
//	return nil
//}
//
//func TestArticle_Add(t *testing.T) {
//	testArticle := mockTestData4ArticleService()
//	err := testArticle.Add()
//	assert.Nil(t, err)
//	if err != nil {
//		return
//	}
//
//	// single add
//	article := ArticleService{Aid: testArticle.Aid}
//	err = article.GetDetailByAid()
//	assert.Nil(t, err)
//	assert.Equal(t, article.Content, testArticle.Content)
//
//	album := Album{Uid: testArticle.Uid, AidList: bson.A{testArticle.Aid}}
//	err = album.Detail()
//	assert.Nil(t, err)
//	assert.Equal(t, testArticle.Aid, album.AidList[0].(int64))
//
//	// double add
//	time.Sleep(1 * time.Second)
//	testArticle2 := mockTestData4ArticleService()
//	err = testArticle2.Add()
//	assert.Nil(t, err)
//	if err != nil {
//		return
//	}
//	err = album.Detail()
//	assert.Nil(t, err)
//	assert.Equal(t, testArticle2.Aid, album.AidList[1].(int64))
//
//	err = clearTestData4ArticleService(testArticle.Aid, false)
//	assert.Nil(t, err)
//	err = clearTestData4ArticleService(testArticle2.Aid, true)
//	assert.Nil(t, err)
//}
//
//func TestArticle_DeleteByAid(t *testing.T) {
//	//soft-delete
//	testArticle := mockTestData4ArticleService()
//	err := testArticle.Add()
//	assert.Nil(t, err)
//	if err != nil {
//		return
//	}
//
//	err = testArticle.DeleteByAid(true)
//	assert.Nil(t, err)
//	if err != nil {
//		return
//	}
//
//	// data should exist in database
//	err = testArticle.GetDetailByAid()
//	assert.Nil(t, err)
//	assert.Equal(t, testArticle.Content, testArticle.Content)
//
//	// todo 还要检查album，album如果已经被删除，还要做容错，先不搞这么细节
//
//	// delete permanently
//	err = testArticle.DeleteByAid(false)
//	assert.Nil(t, err)
//	if err != nil {
//		return
//	}
//
//	// check whether has been deleted
//	err = testArticle.GetDetailByAid() // expect err to be not nil
//	assert.NotNil(t, err)
//}
//
//func TestArticle_RefreshTimeline(t *testing.T) {
//	// todo add test
//	//RefreshTimeline("refresh")
//	//RefreshTimeline("loadmore")
//}
