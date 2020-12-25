package model

import (
	"Moments/pkg/log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

/*
 Uid 90000 is a test account
*/

type timelineTest struct {
	Uid     int32   `bson:"uid"`
	AidList []int64 `bson:"aid_list"`
}

func init() {
	log.InitLogger(true)
}

func mockTestData4Timeline() (*timelineTest, error) {
	var tempList []int64
	var length = 5
	for count, i := 0, 91004; count < length; count++ {
		tempList = append(tempList, int64(i-count))
	}
	testData := timelineTest{
		Uid:     int32(90000),
		AidList: tempList,
	}

	td, _ := bson.Marshal(testData)
	data := bson.M{}
	_ = bson.Unmarshal(td, &data)

	err := insert("timeline", data)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return &testData, nil
}

func clearTestData4Timeline() error {
	err := remove("timeline", map[string]interface{}{"uid": 90000})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func TestModels_GetTimelineRefreshByUid(t *testing.T) {
	testData, err := mockTestData4Timeline()
	assert.Nil(t, err)

	uid := testData.Uid
	aidList := testData.AidList

	// normal case
	start := 3
	aids, err := GetTimelineRefreshByUid(uid, aidList[start])
	assert.Nil(t, err)
	assert.Equal(t, aids, aidList[0:start])

	// corner aid
	aids, err = GetTimelineRefreshByUid(uid, aidList[0])
	assert.Nil(t, err)
	assert.Empty(t, aids)

	err = clearTestData4Timeline()
	assert.Nil(t, err)
}

func TestModels_GetTimelineLoadMoreByUid(t *testing.T) {
	testData, err := mockTestData4Timeline()
	assert.Nil(t, err)

	uid := testData.Uid
	aidList := testData.AidList

	// normal case
	start := 0
	aids, err := GetTimelineLoadMoreByUid(uid, aidList[start])
	assert.Nil(t, err)
	assert.Equal(t, aidList[start+1:], aids)

	// corner case
	last := len(aidList) - 1
	aids, err = GetTimelineLoadMoreByUid(uid, aidList[last])
	assert.Nil(t, err)
	assert.Empty(t, aids)

	err = clearTestData4Timeline()
	assert.Nil(t, err)
}
