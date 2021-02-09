package timeline

import (
	"Moments/biz/article"
	"Moments/pkg/log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TimelineServiceSuite struct {
	suite.Suite
	TestData struct {
		Uid      int32
		Articles []article.Article
	}
}

func (s *TimelineServiceSuite) SetupSuite() {
	log.InitLogger(true)
	s.TestData.Uid = 88888
	s.TestData.Articles = []article.Article{
		{
			Uid:     88888,
			Content: "Timeline Unit Test1",
		}, {
			Uid:     88888,
			Content: "Timeline Unit Test2",
		}, {
			Uid:     88888,
			Content: "Timeline Unit Test3",
		},
	}
	ami := article.ArticleModelImpl{}
	for i, _ := range s.TestData.Articles {
		articleSrv := article.NewArticleService(&s.TestData.Articles[i], &ami)
		err := articleSrv.AddArticle(nil)
		require.Nil(s.T(), err)
		time.Sleep(time.Second)
	}
}

func (s *TimelineServiceSuite) TearDownSuite() {
	ami := article.ArticleModelImpl{}
	for _, a := range s.TestData.Articles {
		articleSrv := article.NewArticleService(&a, &ami)
		err := articleSrv.Delete(nil, false)
		require.Nil(s.T(), err)
	}
}

func (s *TimelineServiceSuite) SetupTest() {
	var err error
	tli := TimelineModelImpl{}
	err = tli.CreateTimelineByUid(s.TestData.Uid)
	assert.Nil(s.T(), err)
	for _, a := range s.TestData.Articles {
		err = tli.AppendArticleIntoTimelineByUid(s.TestData.Uid, a.Aid)
		require.Nil(s.T(), err)
	}
}

func (s *TimelineServiceSuite) TearDownTest() {
	var err error
	tli := TimelineModelImpl{}
	for _, a := range s.TestData.Articles {
		err = tli.RemoveArticleFromTimelineByUid(s.TestData.Uid, a.Aid)
		assert.Nil(s.T(), err)
	}
	err = tli.DeleteTimelineByUid(s.TestData.Uid)
	assert.Nil(s.T(), err)
}

func (s *TimelineServiceSuite) TestGetRefreshTimeline() {
	tlSrv := NewTimelineService(&Timeline{}, &TimelineModelImpl{})
	var err error
	var articleList []article.Article
	articleList, err = tlSrv.GetRefreshTimeline(nil, s.TestData.Uid, s.TestData.Articles[0].Aid, "refresh")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(articleList))
	// latest article will be shown first
	assert.Equal(s.T(), s.TestData.Articles[2].Content, articleList[0].Content)
	assert.Equal(s.T(), s.TestData.Articles[1].Content, articleList[1].Content)

	articleList, err = tlSrv.GetRefreshTimeline(nil, s.TestData.Uid, s.TestData.Articles[2].Aid, "loadmore")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(articleList))
	// earliest article will be shown last
	assert.Equal(s.T(), s.TestData.Articles[1].Content, articleList[0].Content)
	assert.Equal(s.T(), s.TestData.Articles[0].Content, articleList[1].Content)
}

func TestTimelineServiceSuite(t *testing.T) {
	suite.Run(t, new(TimelineServiceSuite))
}

// 问题
// 1. album 没有删除？

//db.article_1.remove({"uid":88888})
//db.article_0.remove({"uid":88888})
//db.article_2.remove({"uid":88888})
//db.article_3.remove({"uid":88888})
//db.album.remove({"uid":88888})
//db.timeline.remove({"uid":88888})
