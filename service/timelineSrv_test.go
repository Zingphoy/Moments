package service

import (
	"Moments/model"
	"Moments/pkg/log"
	"testing"

	"github.com/stretchr/testify/assert"
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
	if err := model.InsertNewTimeline(testData.Uid, testData.Aid_list); err != nil {
		log.Fatal(err.Error())
	}

	// refresh test
	tl := TimelineService{Uid: testData.Uid}
	uid := testData.Uid
	aidList := testData.Aid_list
	aid := aidList[len(aidList)-1]
	err := tl.RefreshTimeline(uid, aid, "refresh")
	if err != nil {
		log.Fatal(err.Error())
	}
	var ret []int64
	for _, article := range tl.Articles {
		ret = append(ret, article.Aid)
	}
	assert.Equal(t, ret, testData.Aid_list[0:(len(aidList)-1)])

	// loadmore test
	tl2 := TimelineService{Uid: testData.Uid}
	aid = aidList[0]
	err = tl2.RefreshTimeline(uid, aid, "loadmore")
	if err != nil {
		log.Fatal(err.Error())
	}
	var ret2 []int64
	for _, article := range tl2.Articles {
		ret2 = append(ret2, article.Aid)
	}
	assert.Equal(t, ret2, testData.Aid_list[1:])

	// clear test data
	if err = model.DeleteRowTimeline(bson.M{"uid": testData.Uid}); err != nil {
		log.Fatal("delete test data failed, please delete Timeline test data manually")
	}
}
