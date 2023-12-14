package biz

import (
	"Moments/biz/database"
	"Moments/pkg/log"

	"github.com/gin-gonic/gin"
)

func GetDatabaseData(c *gin.Context) {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()

	resp := make([]interface{}, 0)
	dbname := []string{"article_0", "article_1", "article_2", "article_3", "album", "timeline", "friend"}
	for _, name := range dbname {
		ret, err := client.Query(name, nil)
		if err != nil {
			log.Error(c, err)
		}
		temp := make(map[string]interface{})
		temp["collection"] = name
		temp["data"] = ret
		resp = append(resp, temp)
	}
	c.JSON(200, resp)
}
