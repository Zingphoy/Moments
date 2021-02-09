package article

import (
	"Moments/pkg/log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ArticleServiceSuite struct {
	suite.Suite
	TestData Article

	GeneratedAid int64
}

func (s *ArticleServiceSuite) SetupSuite() {
	log.InitLogger(true)
	//mq.InitMQ()
}

func (s *ArticleServiceSuite) TearDownSuite() {
	//mq.StopMQ()
}

func (s *ArticleServiceSuite) BeforeTest(suiteName string, testName string) {
	switch testName {
	case "TestServiceDeleteArticle":
		// 保证AddArticle正常运作完成
		require.NotEqualf(s.T(), 0, s.GeneratedAid, "Delete article failed, please delete aid: %d it manually", s.GeneratedAid)
		require.NotEqualf(s.T(), -1, s.GeneratedAid, "Delete article failed, please delete aid: %d it manually", s.GeneratedAid)
	}
}

func (s *ArticleServiceSuite) AfterTest(suiteName string, testName string) {
	switch testName {
	case "TestServiceAddArticle":
		s.GeneratedAid = s.TestData.Aid
	}
}

func (s *ArticleServiceSuite) SetupTest() {
	s.TestData = Article{
		Aid:       int64(0),
		Uid:       int32(88888),
		Content:   "AddArticle() unit test",
		PhotoList: []string{},
	}

	ami := ArticleModelImpl{}
	err := ami.AddArticle(&s.TestData)
	if err != nil {
		s.T().Fatal("add test article failed, unit test shall be stoped")
	}
}

func (s *ArticleServiceSuite) TearDownTest() {
	ami := ArticleModelImpl{}
	err := ami.DeleteArticleByUidAid(s.TestData.Uid, s.TestData.Aid)
	if err != nil {
		s.T().Fatal("clear test article failed, unit test shall be stoped, please clear test data by yourself, uid is 88888")
	}
}

func (s *ArticleServiceSuite) TestServiceAddArticle() {
	srv := NewArticleService(&Article{}, &ArticleModelImpl{})
	srv.Data = &s.TestData
	err := srv.AddArticle(nil)
	assert.Nil(s.T(), err)

	if err != nil {
		s.GeneratedAid = -1
	} else {
		s.GeneratedAid = srv.Data.Aid
	}

	// todo 检查album、自己的timeline以及好友的timeline
}

func (s *ArticleServiceSuite) TestServiceDeleteArticle() {
	srv := NewArticleService(&Article{}, &ArticleModelImpl{})
	srv.Data = &s.TestData
	srv.Data.Aid = s.GeneratedAid // 这里会使得TearDownTest中使用同一个aid，最后新增的文章会被直接删除
	log.Info(nil, srv.Data)
	err := srv.DeleteArticle(nil)
	assert.Nil(s.T(), err)
}

func (s *ArticleServiceSuite) TestServiceDetailArticle() {
	srv := NewArticleService(&Article{}, &ArticleModelImpl{})
	srv.Data = &s.TestData
	err := srv.DetailArticle(nil)
	assert.Nil(s.T(), err)
}

func TestArticleServiceSuite(t *testing.T) {
	suite.Run(t, new(ArticleServiceSuite))
}

// db.article_3.remove({"aid":888881612362979})
// db.article_3.find()
