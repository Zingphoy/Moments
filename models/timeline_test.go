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

	var templist []int64
	for i := 91004; i >= 91000; i-- {
		templist = append(templist, int64(i))
	}
	testData := bson.M{
		"uid":      int32(90000),
		"aid_list": templist,
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

	start := 3
	aids, err := GetTimelineRefreshByUid(testData["uid"].(int32), testData["aid_list"].([]int64)[start])
	assert2.Nil(t, err)
	assert2.Equal(t, aids, testData["aid_list"].([]int64)[0:start])

	err = clearTestData()
	assert2.Nil(t, err)
}

func TestModels_GetTimelineLoadMoreByUid(t *testing.T) {
	testData, err := mockTestData()
	assert2.Nil(t, err)

	start := 0
	aids, err := GetTimelineLoadMoreByUid(testData["uid"].(int32), testData["aid_list"].([]int64)[start])
	assert2.Nil(t, err)
	assert2.Equal(t, aids, testData["aid_list"].([]int64)[start+1:])

	err = clearTestData()
	assert2.Nil(t, err)
}
