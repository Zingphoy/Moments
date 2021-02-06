package article

import (
	"Moments/biz/database"
	"Moments/pkg/log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArticleTestSuite struct {
	suite.Suite

	Client   database.DatabaseEngine
	Dbname   string
	TestData Article
}

func (s *ArticleTestSuite) SetupTest() {
	log.InitLogger(true)
	log.RedirectLogStd()

	s.Client = database.NewDatabaseClient()
	err := s.Client.Connect()
	if err != nil {
		return
	}

	s.Dbname = "article_0"
	s.TestData = Article{
		Aid:       int64(0),
		Uid:       int32(88888),
		Content:   "test",
		PhotoList: []string{},
	}

	err = s.Client.Insert("article_0", []interface{}{s.TestData})
	if err != nil {
		log.Fatal(nil, err.Error())
	}
}

func (s *ArticleTestSuite) TearDownTest() {
	defer s.Client.Disconnect()
	err := s.Client.Remove("article_0", map[string]interface{}{"uid": s.TestData.Uid})
	if err != nil {
		log.Fatal(nil, err.Error())
	}
}

func (s *ArticleTestSuite) TestIsAidExist() {
	ami := ArticleModelImpl{}
	err := ami.IsArticleNotExist(s.TestData.Aid)
	assert.NotNil(s.T(), err)
}

func (s *ArticleTestSuite) TestGetDetail() {
	ami := ArticleModelImpl{}
	ret, err := ami.GetArticleDetailByAid(s.TestData.Aid)
	expect := s.TestData.Content
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expect, ret.Content)
}

func TestArticleTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleTestSuite))
}

/****************************** Another TestSuite ******************************/

type ArticleAddDeleteTestSuite struct {
	suite.Suite

	Client              database.DatabaseEngine
	Dbname              string
	TestData            Article
	IsAddArticleSuccess bool
}

/*
 Test函数的执行顺序不是按照实现的顺序来的，而是按照Test函数的函数名在字母表中的顺序来执行的
 */
func (s *ArticleAddDeleteTestSuite) TestAddArticle() {
	ami := ArticleModelImpl{}
	s.TestData = Article{
		Aid:       int64(0),
		Uid:       int32(88888),
		Content:   "AddArticle() unit test",
		PhotoList: []string{},
	}
	err := ami.AddArticle(&s.TestData)
	if err != nil {
		s.IsAddArticleSuccess = false
	} else {
		s.IsAddArticleSuccess = true
	}
	assert.Nil(s.T(), err)
}

func (s *ArticleAddDeleteTestSuite) TestDeleteArticleSoftly() {
	if s.IsAddArticleSuccess {
		ami := ArticleModelImpl{}
		// clear test data at the same time
		err := ami.DeleteArticleSoftByUidAid(s.TestData.Uid, s.TestData.Aid)
		if err != nil {
			s.T().Errorf("need to clear test data in database manully")
		}
		assert.Nil(s.T(), err)
	} else {
		s.T().Errorf("need to clear test data in database manully")
	}
}

// 为了让这个单测在 TestDeleteArticleSoftly 后面运行，特意在函数名末加一个Z字母
func (s *ArticleAddDeleteTestSuite) TestDeleteArticleZ() {
	if s.IsAddArticleSuccess {
		ami := ArticleModelImpl{}
		// clear test data at the same time
		err := ami.DeleteArticleByUidAid(s.TestData.Uid, s.TestData.Aid)
		if err != nil {
			s.T().Errorf("need to clear test data in database manully")
		}
		assert.Nil(s.T(), err)
	} else {
		s.T().Errorf("need to clear test data in database manully")
	}
}

func TestArticleAddDeleteTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleAddDeleteTestSuite))
}
