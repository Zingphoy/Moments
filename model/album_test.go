package model

//
//type albumTest struct {
//	Album service.Album
//}
//
//func init() {
//	log.InitLogger(true)
//	log.RedirectLogStd()
//}
//func mockTestData4Album() (*albumTest, error) {
//	testData := albumTest{
//		service.Album{
//			Uid:     90000,
//			AidList: bson.A{[]int64{5, 4, 3, 2, 1}},
//		},
//	}
//
//	td, _ := bson.Marshal(testData.Album)
//	data := bson.M{}
//	_ = bson.Unmarshal(td, &data)
//
//	err := insert("album", data)
//	if err != nil {
//		log.Fatal(err.Error())
//		return nil, err
//	}
//	return &testData, nil
//}
//
//func clearTestData4Album() error {
//	err := remove("album", bson.M{"uid": 90000})
//	if err != nil {
//		log.Fatal("clear album test data failed,", err.Error())
//		return err
//	}
//	return nil
//}
//
//// todo
//func TestModel_AddAlbum(t *testing.T) {
//	//data, err := mockTestData4Album()
//	//assert.Nil(t, err)
//
//	//err = AddAlbumByUid(data.Album.Uid)
//	//assert.Nil(t, err)
//
//	//err = clearTestData4Album()
//	//assert.Nil(t, err)
//}
//
//func TestModel_AppendAlbum(t *testing.T) {
//
//}
//
//func TestModel_DeleteAlbum(t *testing.T) {
//
//}
