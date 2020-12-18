package v1

import (
	"Moments/pkg/app"
	"Moments/pkg/hints"
	"Moments/pkg/log"
	"Moments/pkg/utils"
	"Moments/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO 使用swagger生成Api

// GetArticleDetail get detail of an article
func GetArticleDetail(c *gin.Context) {
	webapp := app.GinCtx{C: c}
	aid, err := utils.Str(c.DefaultQuery("aid", "0")).Int64()
	if err != nil || !utils.ValidAid(aid) {
		log.Info("invalid param aid")
		webapp.MakeJsonRes(http.StatusOK, hints.INVALID_PARAM, nil)
		return
	}

	articleSrv := service.Article{Aid: aid}
	err = articleSrv.GetDetailByAid()
	if err != nil {
		log.Error("database error:", err.Error())
		webapp.MakeJsonRes(http.StatusInternalServerError, hints.INTERNAL_ERROR, err.Error())
		return
	}

	webapp.MakeJsonRes(http.StatusOK, hints.SUCCESS, articleSrv)
}

// GetTimeline call when user refreshes his timeline, return all articles after/before specific time
// only show articles with correct access
func GetTimeline(c *gin.Context) {
	// aid、uid、type
	webapp := app.GinCtx{C: c}
	var param map[string]interface{}
	err := c.BindJSON(&param)
	if err != nil {
		log.Error("data parse json error:", err.Error())
		webapp.MakeJsonRes(http.StatusInternalServerError, hints.INTERNAL_ERROR, err.Error())
		return
	}

	log.Info(param)
	tlSrv := service.Timeline{}
	err = tlSrv.RefreshTimeline(int32(param["uid"].(float64)), int64(param["aid"].(float64)), param["schema"].(string))
	if err != nil {
		log.Error(err.Error())
		webapp.MakeJsonRes(http.StatusOK, hints.INTERNAL_ERROR, err.Error())
	}
	webapp.MakeJsonRes(http.StatusOK, hints.SUCCESS, tlSrv)
}

// 发布表入库相关信息，接着相册表完成入库，并将一个扩散朋友圈的消息添加到消息队列
// SendArticle called after api.UploadPicture, deal with users' moments
func SendArticle(c *gin.Context) {
	webapp := app.GinCtx{C: c}
	articleSrv := service.Article{}
	err := c.BindJSON(&articleSrv)
	if err != nil {
		log.Error("data parse json error:", err.Error())
		webapp.MakeJsonRes(http.StatusInternalServerError, hints.INTERNAL_ERROR, err.Error())
		return
	}

	err = articleSrv.Add()
	if err != nil {
		log.Error("error:", err.Error())
		webapp.MakeJsonRes(http.StatusOK, hints.INTERNAL_ERROR, err.Error())
		return
	}

	albumSrv := service.Album{
		Uid: articleSrv.Uid,
		Aid: articleSrv.Aid,
	}
	err = albumSrv.Append()
	if err != nil {
		log.Error("error:", err.Error())
		webapp.MakeJsonRes(http.StatusOK, hints.INTERNAL_ERROR, err.Error())
		return
	}

	webapp.MakeJsonRes(http.StatusOK, hints.SUCCESS, nil)
}

// DeleteArticle softly delete
func DeleteArticle(c *gin.Context) {
	webapp := app.GinCtx{C: c}
	articleSrv := service.Article{}
	err := c.BindJSON(&articleSrv)
	if err != nil {
		log.Error("data parse json error:", err.Error())
		webapp.MakeJsonRes(http.StatusInternalServerError, hints.INTERNAL_ERROR, err.Error())
		return
	}

	err = articleSrv.DeleteByAid(true)
	if err != nil {
		log.Error("models delete error")
		webapp.MakeJsonRes(http.StatusInternalServerError, hints.INTERNAL_ERROR, err.Error())
		return
	}

	webapp.MakeJsonRes(http.StatusOK, hints.SUCCESS, nil)
}

func CommentArticle(c *gin.Context) {

}

func LikeArticle(c *gin.Context) {

}
