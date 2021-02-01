package model

import (
	"Moments/pkg/log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

type articleTest struct {
	Dbname  string
	Article Article
}

func init() {
	log.InitLogger(true)
	log.RedirectLogStd()
}

func mockTestData4Article() (*articleTest, error) {
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

	err := insert("article_0", data)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return &testData, nil
}

func clearTestData4Article() error {
	err := remove("article_0", map[string]interface{}{"uid": 90000})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
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
