package router

import (
	"Moments/middleware"
	"Moments/router/api"
	v1 "Moments/router/api/v1"
	"context"
	"fmt"
	"net/http"
	"time"

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

	apiTest := r.Group("/test")
	//apiTest.Use(middleware.AddTimeout())
	apiTest.GET("/hello", testFunc)

	return r
}

// todo 起协程来处理业务逻辑，在外头使用timeContext，如果timeout后是否能手动中止协程？
func testFunc(c *gin.Context) {
	fmt.Println(time.Now().Second())

	t, cancelFunc := context.WithTimeout(c, 2*time.Second)
	defer cancelFunc()

	done := make(chan int)

	go func(chan int) {
		time.Sleep(5 * time.Second)
		done <- 0
	}(done)

	select {
	case <-done:
		fmt.Println("ok")
	case <-t.Done():
		fmt.Println("timeout in go routine: ", time.Now().Second())
	}

	c.JSON(http.StatusOK, "hello world")
}
