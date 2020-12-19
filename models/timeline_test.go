package models

import (
	"Moments/pkg/log"
	"testing"

	assert2 "github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

/*
 Uid 90000 is a test account
*/

func init() {
	log.InitLogger(true)
}

func mockTestData() (bson.M, error) {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	var tempList []int64
	var length = 5
	for count, i := 0, 91004; count < length; count++ {
		tempList = append(tempList, int64(i-count))
	}
	testData := bson.M{
		"uid":      int32(90000),
		"aid_list": tempList,
	}

	collection := db.Collection("timeline")
	_, err := collection.InsertOne(ctx, testData)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return testData, nil
}

func clearTestData() error {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	collection := db.Collection("timeline")
	_, err := collection.DeleteOne(ctx, bson.M{"uid": 90000})
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return err
}

func TestModels_GetTimelineRefreshByUid(t *testing.T) {
	testData, err := mockTestData()
	assert2.Nil(t, err)

	uid := testData["uid"].(int32)
	aidList := testData["aid_list"].([]int64)

	// normal case
	start := 3
	aids, err := GetTimelineRefreshByUid(uid, aidList[start])
	assert2.Nil(t, err)
	assert2.Equal(t, aids, aidList[0:start])

	// corner aid
	aids, err = GetTimelineRefreshByUid(uid, aidList[0])
	assert2.Nil(t, err)
	assert2.Empty(t, aids)

	err = clearTestData()
	assert2.Nil(t, err)
}

func TestModels_GetTimelineLoadMoreByUid(t *testing.T) {
	testData, err := mockTestData()
	assert2.Nil(t, err)

	uid := testData["uid"].(int32)
	aidList := testData["aid_list"].([]int64)

	// normal case
	start := 0
	aids, err := GetTimelineLoadMoreByUid(uid, aidList[start])
	assert2.Nil(t, err)
	assert2.Equal(t, aidList[start+1:], aids)

	// corner case
	last := len(aidList) - 1
	aids, err = GetTimelineLoadMoreByUid(uid, aidList[last])
	assert2.Nil(t, err)
	assert2.Empty(t, aids)

	err = clearTestData()
	assert2.Nil(t, err)
}
