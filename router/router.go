package router

import (
	"Moments/biz/album"
	"Moments/biz/article"
	"Moments/middleware"
	"Moments/router/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() http.Handler {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.TrackingId())

	tools := r.Group("/tools")
	{
		tools.POST("/upload/pic", api.UploadPicture)
	}

	apiV1 := r.Group("/v1")
	{
		// moments refresh or load more
		//apiV1.GET("/moments/timeline", v1.GetTimeline)

		/********** Article module **********/
		apiV1.GET("/moments/article/detail", article.GetArticleDetail)
		apiV1.POST("/moments/article/post", article.SendArticle)
		apiV1.POST("/moments/article/delete", article.DeleteArticle)
		apiV1.POST("/moments/article/comment", article.CommentArticle)
		apiV1.POST("/moments/article/like", article.LikeArticle)

		/********** Album module **********/
		apiV1.GET("/moments/album/detail", album.GetAlbumDetail)

	}

	//apiTest := r.Group("/test")
	//apiTest.Use(middleware.AddTimeout())

	return r
}
