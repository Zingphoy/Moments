package routers

import (
	"Moments/routers/api"
	v1 "Moments/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	tools := r.Group("/tools")
	{
		tools.POST("/upload/pic", api.UploadPicture)
	}

	apiV1 := r.Group("/v1")
	{
		// moments refresh or load more
		apiV1.GET("/moments/timeline", v1.GetTimeline)

		// get article detail
		apiV1.GET("/moments/article/detail", v1.GetArticleDetail)

		// post a new article
		apiV1.POST("/moments/post", v1.SendArticle)

		// delete an article softly
		apiV1.POST("/moments/delete", v1.DeleteArticle)

		// comment an article
		apiV1.POST("/moments/comment", v1.CommentArticle)

		// like an article
		apiV1.POST("/moments/like", v1.LikeArticle)
	}

	return r
}
