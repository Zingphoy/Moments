package timeline

//func RefreshTimeline(uid int32, latestAid int64, schema string) ([]*model.ArticleHandler, error) {
//	var err error
//	var aids []int64
//	tl := model.NewTimelineObject(map[string]interface{}{"": 1})
//
//	switch schema {
//	case "refresh":
//		aids, err = tl.GetUserTimelineRefresh(latestAid)
//	case "loadmore":
//		aids, err = tl.GetUserTimelineLoadMore(latestAid)
//	}
//	if err != nil {
//		log.Error("get timeline failed")
//		return nil, err
//	}
//
//	// todo 这里真的是返回一个timeline结构？看起来更像是需要返回article详细列表吧
//	articleList := make([]*model.ArticleHandler, 0, 10)
//	for _, aid := range aids {
//		article := model.NewArticleObject(map[string]interface{}{"aid": aid})
//		err = article.GetArticleDetail()
//		if err != nil {
//			log.Error("get article detail failed")
//			return nil, err
//		}
//		articleList = append(articleList, &article)
//	}
//
//	log.Info("get timeline success, aid list:", aids)
//	return articleList, nil
//}
