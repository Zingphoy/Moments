package main

import (
	"Moments/biz/database"
	"Moments/pkg/log"
)

/*
 This is the utility file for test.
 To make test more convenient, here will offer some database testing functions for inserting/deleting/updating data.
*/

func init() {
	log.InitLogger(true)
	log.RedirectLogStd()
}

// write test code heret
func main2() {
	dbname := "article_3"
	filter := map[string]interface{}{
		"uid": int32(88888),
		"aid": int64(888881612363047),
	}
	up := map[string]interface{}{
		"is_deleted": 1,
	}

	c := database.NewDatabaseClient()
	c.Connect()
	defer c.Disconnect()
	c.Update(dbname, filter, up)

}
