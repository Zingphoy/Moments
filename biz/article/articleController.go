package article

import (
	"Moments/biz/album"
	"Moments/pkg/app"
	"Moments/pkg/hint"
	"Moments/pkg/log"
	"Moments/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleRequest struct {
	Aid       int64    `form:"aid" json:"aid" validate:"required,gtfield=0"`
	Uid       int32    `json:"uid" validate:"gtfield=0"`
	PostTime  int64    `json:"post_time" validate:"gtfield=0"`
	Content   string   `json:"content"`
	PhotoList []string `json:"photo_list"`

	// extra fields
	Privacy   int32 `json:"privacy" validate:"gtfield=0,ltfield=1000"`
	IsDeleted int32 `json:"is_deleted" validate:"gtefield=0,ltefield=1"`
}

type ArticleResponse struct {
	*Article
}

// GetArticleDetail get detail of an article
func GetArticleDetail(c *gin.Context) {
	webapp := app.GinCtx{C: c}
	aid, err := utils.Str(c.DefaultQuery("aid", "0")).Int64()
	if err != nil {
		log.Info(c, "invalid param aid,", err.Error())
		webapp.MakeFailedJsonRes(http.StatusOK, err)
		return
	}

	srv := NewArticleService(&Article{}, &ArticleModelImpl{})
	srv.Data.Aid = aid
	err = srv.DetailArticle(c)
	if err != nil {
		webapp.MakeFailedJsonRes(0, err)
		return
	}
	webapp.MakeSuccessJsonRes(http.StatusOK, srv.Data)
}

// 发布表入库相关信息，接着相册表完成入库，并将一个扩散朋友圈的消息添加到消息队列
// SendArticle called after api.UploadPicture, deal with users' moments
func SendArticle(c *gin.Context) {
	webapp := app.GinCtx{C: c}
	articleData := Article{}
	err := c.BindJSON(&articleData)
	if err != nil {
		webapp.MakeFailedJsonRes(http.StatusOK, hint.CustomError{
			Code: hint.INVALID_PARAM,
			Err:  err,
		})
		return
	}

	srv := ArticleService{Data: &articleData, Impl: &ArticleModelImpl{}}
	err = srv.AddArticle(c)
	if err != nil {
		webapp.MakeFailedJsonRes(http.StatusOK, err)
		return
	}

	albumSrv := album.NewAlbumService(&album.Album{}, &album.AlbumModelImpl{})
	albumSrv.Data.Uid = srv.Data.Uid
	albumSrv.Data.AidList = []int64{srv.Data.Aid}
	err = albumSrv.AppendAlbum(c)
	if err != nil {
		webapp.MakeFailedJsonRes(http.StatusOK, err.Error())
		return
	}
	webapp.MakeSuccessJsonRes(http.StatusOK, nil)
}

// DeleteArticle softly delete
func DeleteArticle(c *gin.Context) {
	webapp := app.GinCtx{C: c}
	articleData := Article{}
	err := c.BindJSON(&articleData)
	if err != nil {
		webapp.MakeFailedJsonRes(http.StatusOK, hint.CustomError{
			Code: hint.INVALID_PARAM,
			Err:  err,
		})
		return
	}

	srv := ArticleService{Data: &articleData, Impl: &ArticleModelImpl{}}
	err = srv.DeleteArticle(c)
	if err != nil {
		webapp.MakeFailedJsonRes(http.StatusOK, err)
		return
	}
	webapp.MakeSuccessJsonRes(http.StatusOK, nil)
}

func CommentArticle(c *gin.Context) {

}

func LikeArticle(c *gin.Context) {

}
