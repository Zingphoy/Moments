package timeline

import (
	"Moments/pkg/log"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TimelineService struct {
	Data *Timeline
	Impl TimelineModel
}

func NewTimelineService(data *Timeline, impl *TimelineModelImpl) *TimelineService {
	return &TimelineService{
		Data: data,
		Impl: impl,
	}
}

//func (t *TimelineService) GetRefreshTimeline(c *gin.Context, uid int32, boundaryAid int64, schema string) ([]article.Article, error) {
//	var err error
//	var aids []int64
//	switch schema {
//	case "refresh":
//		aids, err = t.Impl.GetTimelineRefreshByUidAid(uid, boundaryAid)
//	case "loadmore":
//		aids, err = t.Impl.GetTimelineLoadMoreByUidAid(uid, boundaryAid)
//	}
//	if err != nil {
//		log.Error(c, fmt.Sprintf("refresh timeline failed, uid=%d", uid))
//		return nil, err
//	}
//
//	articleSrv := article.NewArticleService(&article.Article{}, &article.ArticleModelImpl{})
//	articleList := make([]article.Article, 0, 10)
//	for _, aid := range aids {
//		articleSrv.Data.Aid = aid
//		err = articleSrv.DetailArticle(c)
//		if err != nil {
//			log.Info(c, fmt.Sprintf("get article detail failed, aid=%d", aid))
//			continue
//		}
//		articleList = append(articleList, *articleSrv.Data)
//	}
//	return articleList, nil
//}

func (t *TimelineService) GetRefreshTimeline(c *gin.Context, uid int32, boundaryAid int64, schema string) ([]int64, error) {
	var err error
	var aids []int64
	switch schema {
	case "refresh":
		aids, err = t.Impl.GetTimelineRefreshByUidAid(uid, boundaryAid)
	case "loadmore":
		aids, err = t.Impl.GetTimelineLoadMoreByUidAid(uid, boundaryAid)
	}
	if err != nil {
		log.Error(c, fmt.Sprintf("refresh timeline failed, uid=%d", uid))
		return nil, err
	}
	return aids, nil
}

func (t *TimelineService) AppendArticleIntoTimeline(c *gin.Context, uid int32, aid int64) error {
	err := t.Impl.AppendArticleIntoTimelineByUid(uid, aid)
	if err != nil {
		return err
	}
	return nil
}
