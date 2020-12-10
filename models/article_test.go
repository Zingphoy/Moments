package models

import (
	"Moments/pkg/log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

type testData struct {
	dbname  string
	aid     int64
	uid     int32
	content string
}

// test data that stores in collection article_0
var td = testData{
	dbname:  "article_0",
	aid:     0,
	uid:     0,
	content: "test",
}

func TestMain(m *testing.M) {
	log.InitLogger(true)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestIsAidExist(t *testing.T) {
	r := IsAidExist(td.dbname, td.aid)
	assert.True(t, r)
}

func TestGetDetail(t *testing.T) {
	expect := td.content
	ret, err := GetDetail(td.dbname, bson.M{"aid": td.aid})
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, expect, ret.Content)
}

func TestAddArticle(t *testing.T) {
	//AddArticle()
}

func TestDeleteArticle(t *testing.T) {

}
