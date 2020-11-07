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
	article, err := articleSrv.GetDetailByAid()
	if err != nil {
		log.Error("database error: ", err.Error())
		webapp.MakeJsonRes(http.StatusInternalServerError, hints.INTERNAL_ERROR, err.Error())
		return
	}

	webapp.MakeJsonRes(http.StatusOK, hints.SUCCESS, article)
}

func GetTimeline(c *gin.Context) {

}

// 发布表入库相关信息，接着相册表完成入库，并将一个扩散朋友圈的消息添加到消息队列
// SendArticle called after api.UploadPicture, deal with users' moments
func SendArticle(c *gin.Context) {
	webapp := app.GinCtx{C: c}
	//var (
	//	err       error
	//	uid       int32
	//	postTime  int64
	//	content   string
	//	photoList []string
	//	privacy   int8
	//)

	//uid, err = utils.Str(c.DefaultPostForm("uid", "-1")).Int32()
	//postTime, err = utils.Str(c.DefaultPostForm("post_time", "-1")).Int64()
	//content = c.DefaultPostForm("content", "")
	//photoList = c.DefaultPostForm("photo_list", "-1")
	//privacy, err = utils.Str(c.DefaultPostForm("privacy", "-1")).Int8()
	//if err != nil {
	//	log.Error("database error")
	//	webapp.MakeJsonRes(http.StatusInternalServerError, hints.INTERNAL_ERROR, err.Error())
	//	return
	//}

	articleSrv := service.Article{}
	err := c.BindJSON(&articleSrv)
	if err != nil {
		log.Error("data parse json error: ", err.Error())
		webapp.MakeJsonRes(http.StatusInternalServerError, hints.INTERNAL_ERROR, err.Error())
		return
	}

	err = articleSrv.Add()
	if err != nil {
		log.Error("error: ", err.Error())
		webapp.MakeJsonRes(http.StatusOK, hints.INTERNAL_ERROR, err.Error())
		return
	}

	albumSrv := service.Album{
		Uid: articleSrv.Uid,
		Aid: articleSrv.Aid,
	}
	err = albumSrv.Append()
	if err != nil {
		log.Error("error: ", err.Error())
		webapp.MakeJsonRes(http.StatusOK, hints.INTERNAL_ERROR, err.Error())
		return
	}

	webapp.MakeJsonRes(http.StatusOK, hints.SUCCESS, nil)
}

// todo 先完成这个接口，但是还不好测试，数据删了就没了
func DeleteArticle(c *gin.Context) {

}

func CommentArticle(c *gin.Context) {

}

func LikeArticle(c *gin.Context) {

}
