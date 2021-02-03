package article

import (
	"Moments/pkg/log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArticleServiceSuite struct {
	suite.Suite
	TestData Article
}

func (s *ArticleServiceSuite) SetupTest() {
	log.InitLogger(true)
	ami := ArticleModelImpl{}
	s.TestData = Article{
		Aid:       int64(0),
		Uid:       int32(88888),
		Content:   "AddArticle() unit test",
		PhotoList: []string{},
	}
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

func (s *ArticleServiceSuite) TestServiceDetailArticle() {
	h := NewArticleHandler()
	h.Data = &s.TestData
	err := h.DetailArticle(nil)
	assert.Nil(s.T(), err)
}

func (s *ArticleServiceSuite) TestServiceCommentArticle() {

}

func (s *ArticleServiceSuite) TestServiceLikeArticle() {

}

func TestArticleServiceSuite(t *testing.T) {
	suite.Run(t, new(ArticleServiceSuite))
}

/****************************** Another TestSuite ******************************/

type ArticleServiceAddDeleteSuite struct {
	suite.Suite

	IsAddSuccess bool
	h            *ArticleHandler
}

func (s *ArticleServiceAddDeleteSuite) SetupSuite() {
	log.InitLogger(true)
	//mq.InitMQ()

	// AddArticle中生成的aid，要沿用到DeleteArticle
	s.h = NewArticleHandler()
	s.h.Data = &Article{
		Aid:       int64(0),
		Uid:       int32(88888),
		Content:   "AddArticle() unit test",
		PhotoList: []string{},
	}
}

func (s *ArticleServiceAddDeleteSuite) TearDownSuite() {
	//defer mq.StopMQ()	// consumer要先start才能shutdown，否则panic
}

func (s *ArticleServiceAddDeleteSuite) TestServiceAddArticle() {
	err := s.h.AddArticle(nil)
	assert.Nil(s.T(), err)

	if err != nil {
		s.IsAddSuccess = false
	} else {
		s.IsAddSuccess = true
	}
	log.Error(nil, "我在这里1", s.h.Data)

	// todo 检查album、自己的timeline以及好友的timeline
}

func (s *ArticleServiceAddDeleteSuite) TestServiceDeleteArticle() {
	assert.Equal(s.T(), true, s.IsAddSuccess)
	if s.IsAddSuccess == true {
		err := s.h.DeleteArticle(nil)
		assert.Nil(s.T(), err)
	}
}

func TestArticleServiceAddDeleteSuite(t *testing.T) {
	suite.Run(t, new(ArticleServiceAddDeleteSuite))
}

// db.article_3.remove({"aid":888881612362979})
// db.article_3.find()

// todo 为什么delete文章，update is_delete字段不成功
