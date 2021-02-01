package timeline

import (
	"Moments/pkg/log"
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
	log.RedirectLogStd()
}

// 老的单测代码，暂时先保留，等新单测代码ok后即可废弃
//func mockTestData4Timeline() (*timelineTest, error) {
//	var tempList []int64
//	var length = 5
//	for count, i := 0, 91004; count < length; count++ {
//		tempList = append(tempList, int64(i-count))
//	}
//	testData := timelineTest{
//		Uid:     int32(90000),
//		AidList: tempList,
//	}
//
//	td, _ := bson.Marshal(testData)
//	data := bson.M{}
//	_ = bson.Unmarshal(td, &data)
//
//	err := insert("timeline", data)
//	if err != nil {
//		log.Fatal(err.Error())
//		return nil, err
//	}
//	return &testData, nil
//}
//
//func clearTestData4Timeline() error {
//	err := remove("timeline", map[string]interface{}{"uid": 90000})
//	if err != nil {
//		log.Fatal(err.Error())
//		return err
//	}
//	return nil
//}
//
//func TestModel_GetTimelineRefreshByUid(t *testing.T) {
//	testData, err := mockTestData4Timeline()
//	assert.Nil(t, err)
//
//	uid := testData.Uid
//	aidList := testData.AidList
//
//	// normal case
//	start := 3
//	aids, err := GetTimelineRefreshByUid(uid, aidList[start])
//	assert.Nil(t, err)
//	assert.Equal(t, aids, aidList[0:start])
//
//	// corner aid
//	aids, err = GetTimelineRefreshByUid(uid, aidList[0])
//	assert.Nil(t, err)
//	assert.Empty(t, aids)
//
//	err = clearTestData4Timeline()
//	assert.Nil(t, err)
//}
//
//func TestModel_GetTimelineLoadMoreByUid(t *testing.T) {
//	testData, err := mockTestData4Timeline()
//	assert.Nil(t, err)
//
//	uid := testData.Uid
//	aidList := testData.AidList
//
//	// normal case
//	start := 0
//	aids, err := GetTimelineLoadMoreByUid(uid, aidList[start])
//	assert.Nil(t, err)
//	assert.Equal(t, aidList[start+1:], aids)
//
//	// corner case
//	last := len(aidList) - 1
//	aids, err = GetTimelineLoadMoreByUid(uid, aidList[last])
//	assert.Nil(t, err)
//	assert.Empty(t, aids)
//
//	err = clearTestData4Timeline()
//	assert.Nil(t, err)
//}
