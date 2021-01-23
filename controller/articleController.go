package controller

import (
	"Moments/model"
	"Moments/pkg/app"
	"Moments/pkg/hint"
	"Moments/pkg/log"
	"Moments/pkg/utils"
	"Moments/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleRequest struct {
	Aid       int64    `json:"aid" validate:"required,gtfield=0"`
	Uid       int32    `json:"uid" validate:"gtfield=0"`
	PostTime  int64    `json:"post_time" validate:"gtfield=0"`
	Content   string   `json:"content"`
	PhotoList []string `json:"photo_list"`

	// extra fields
	Privacy   int32 `json:"privacy" validate:"gtfield=0,ltfield=1000"`
	IsDeleted int32 `json:"is_deleted" validate:"gtefield=0,ltefield=1"`
}

type ArticleResponse struct {
	model.Article
}

// GetArticleDetail get detail of an article
func GetArticleDetail(c *gin.Context) {
	webapp := app.GinCtx{C: c}
	aid, err := utils.Str(c.DefaultQuery("aid", "0")).Int64()
	if err != nil || !utils.ValidAid(aid) {
		log.Info("invalid param aid")
		webapp.MakeJsonRes(http.StatusOK, hint.INVALID_PARAM, nil)
		return
	}

	article, err := service.DetailArticle(aid)
	if err != nil {
		log.Error("database error:", err.Error())
		webapp.MakeJsonRes(http.StatusInternalServerError, hint.INTERNAL_ERROR, err.Error())
		return
	}

	webapp.MakeJsonRes(http.StatusOK, hint.SUCCESS, article)
}

// 发布表入库相关信息，接着相册表完成入库，并将一个扩散朋友圈的消息添加到消息队列
// SendArticle called after api.UploadPicture, deal with users' moments
func SendArticle(c *gin.Context) {
	webapp := app.GinCtx{C: c}
	article := model.article{}
	err := c.BindJSON(&article)
	if err != nil {
		log.Error("data parse json error:", err.Error())
		webapp.MakeJsonRes(http.StatusInternalServerError, hint.INTERNAL_ERROR, err.Error())
		return
	}

	err = article.AddArticle()
	if err != nil {
		log.Error("error:", err.Error())
		webapp.MakeJsonRes(http.StatusOK, hint.INTERNAL_ERROR, err.Error())
		return
	}

	// todo
	//albumSrv := service.Album{
	//	Uid: article.Uid,
	//	Aid: article.Aid,
	//}
	//err = albumSrv.Append()
	//if err != nil {
	//	log.Error("error:", err.Error())
	//	webapp.MakeJsonRes(http.StatusOK, hint.INTERNAL_ERROR, err.Error())
	//	return
	//}

	webapp.MakeJsonRes(http.StatusOK, hint.SUCCESS, nil)
}

// DeleteArticle softly delete
func DeleteArticle(c *gin.Context) {
	webapp := app.GinCtx{C: c}
	article := model.article{}
	err := c.BindJSON(&article)
	if err != nil {
		log.Error("data parse json error:", err.Error())
		webapp.MakeJsonRes(http.StatusInternalServerError, hint.INTERNAL_ERROR, err.Error())
		return
	}

	err = article.DeleteArticleSoft()
	if err != nil {
		log.Error("model delete error")
		webapp.MakeJsonRes(http.StatusInternalServerError, hint.INTERNAL_ERROR, err.Error())
		return
	}

	webapp.MakeJsonRes(http.StatusOK, hint.SUCCESS, nil)
}

func CommentArticle(c *gin.Context) {

}

func LikeArticle(c *gin.Context) {

}

func Get() {

}

func Add() {

}

func Update() {

}

func Insert() {

}

func Remove() {

}
