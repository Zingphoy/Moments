package service

import (
	"Moments/model"
	"Moments/pkg/log"
	"Moments/pkg/utils"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type Map = map[string]interface{}

// generateAid generate global unique aid，rule as uid + timestamp
// also restrict article sending frequency of one user to 1 time per second
func generateAid(uid int32) (int64, error) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	tmp := strconv.FormatInt(int64(uid), 10) + ts
	aid := utils.Str(tmp).MustInt64()

	a := model.Article{}
	a.Aid = aid
	yes := a.IsArticleExist()
	if yes {
		log.Info("aid already existed")
		return 0, errors.New("aid already existed")
	}
	log.Info("generated aid:", aid)
	return aid, nil
}
