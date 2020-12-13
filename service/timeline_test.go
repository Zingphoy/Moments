package service

import (
	"Moments/models"
	"Moments/pkg/log"
	"testing"

	"github.com/magiconair/properties/assert"
	"go.mongodb.org/mongo-driver/bson"
)

/*
	Test Data:

    [ 90000 ] user is a test account
	testData := struct {
		Uid      int32
		Aid_list []int64
	}{
		Uid:      90000,
		Aid_list: []int64{900011604900532, 900011604900531, 900011604900530, 900011604900529},
	}
*/

func init() {
	log.InitLogger(true)
}

func TestTimeline_RefreshTimeline(t *testing.T) {
	testData := struct {
		Uid      int32
		Aid_list []int64
	}{
		Uid:      90000,
		Aid_list: []int64{900011604900532, 900011604900531, 900011604900530, 900011604900529},
	}

	// prepare test data
	err := models.InsertNewTimeline(testData)
	if err != nil {
		log.Fatal(err.Error())
	}

	// refresh test
	tl := Timeline{Uid: testData.Uid}
	uid := testData.Uid
	aidList := testData.Aid_list
	aid := aidList[len(aidList)-1]
	err = tl.RefreshTimeline(uid, aid, "refresh")
	if err != nil {
		log.Fatal(err.Error())
	}

	var ret []int64
	for _, article := range tl.Articles {
		ret = append(ret, article.Aid)
	}
	assert.Equal(t, ret, testData.Aid_list[0:len(aidList)-1])

	// loadmore test
	//tl.RefreshTimeline(uid, aid, "loadmore")
	//log.Info("loadmore:", tl)

	// clear test data
	err = models.DeleteRowTimeline(bson.M{"uid": testData.Uid})
	if err != nil {
		log.Fatal("delete test data failed, please delete Timeline test data manually")
	}
}