package model

import (
	"Moments/pkg/log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

type articleTest struct {
	Dbname  string
	Article Article
}

func mockTestData4Article() (*articleTest, error) {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	testData := articleTest{Dbname: "article_0",
		Article: Article{
			Aid:       int64(0),
			Uid:       int32(90000),
			Content:   "test",
			PhotoList: bson.A{},
		},
	}

	td, _ := bson.Marshal(testData.Article)
	data := bson.M{}
	_ = bson.Unmarshal(td, &data)

	collection := db.Collection("article_0")
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return &testData, nil
}

func clearTestData4Article() error {
	db, client, ctx, _ := ConnectDatabase()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("error while trying to disconnect database: ", err.Error())
		}
	}()

	collection := db.Collection("article_0")
	_, err := collection.DeleteOne(ctx, bson.M{"uid": 90000})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func TestMain(m *testing.M) {
	log.InitLogger(true)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestIsAidExist(t *testing.T) {
	data, err := mockTestData4Article()
	assert.Nil(t, err)

	r := IsAidExist(data.Dbname, data.Article.Aid)
	assert.True(t, r)

	err = clearTestData4Article()
	assert.Nil(t, err)
}

func TestGetDetail(t *testing.T) {
	data, err := mockTestData4Article()
	assert.Nil(t, err)

	expect := data.Article.Content
	ret, err := GetDetail(data.Dbname, bson.M{"aid": data.Article.Aid})
	assert.Nil(t, err)
	assert.Equal(t, expect, ret.Content)

	err = clearTestData4Article()
	assert.Nil(t, err)
}

func TestAddArticle(t *testing.T) {
	//AddArticle()
}

func TestDeleteArticle(t *testing.T) {

}
