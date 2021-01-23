package v1

//// GetTimeline call when user refreshes his timeline, return all articles after/before specific time
//// only show articles with correct access
//func GetTimeline(c *gin.Context) {
//	// aid、uid、type
//	webapp := app.GinCtx{C: c}
//	var param map[string]interface{}
//	err := c.BindJSON(&param)
//	if err != nil {
//		log.Error("data parse json error:", err.Error())
//		webapp.MakeJsonRes(http.StatusInternalServerError, hint.INTERNAL_ERROR, err.Error())
//		return
//	}
//
//	articleList, err := service.RefreshTimeline(int32(param["uid"].(float64)), int64(param["aid"].(float64)), param["schema"].(string))
//	if err != nil {
//		log.Error(err.Error())
//		webapp.MakeJsonRes(http.StatusOK, hint.INTERNAL_ERROR, err.Error())
//	}
//	webapp.MakeJsonRes(http.StatusOK, hint.SUCCESS, articleList)
//}
