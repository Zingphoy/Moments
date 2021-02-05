package album

import (
	"Moments/pkg/log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AlbumServiceSuite struct {
	suite.Suite

	TestData Album
}

func (s *AlbumServiceSuite) SetupSuite() {
	log.InitLogger(true)
}

func (s *AlbumServiceSuite) TearDownSuite() {
	ami := AlbumModelImpl{}
	err := ami.DeleteAlbumByUid(80808)
	if err != nil {
		s.T().Errorf("delete new album failed: %d, please delete it manully", 80808)
		s.T().FailNow()
	}
}

func (s *AlbumServiceSuite) SetupTest() {
	var err error
	ami := AlbumModelImpl{}
	s.TestData.Uid = 88888
	err = ami.CreateAlbumByUid(s.TestData.Uid)
	if err != nil {
		s.T().Errorf("create new album failed")
		s.T().FailNow()
	}

	err = ami.AppendAlbumByUidAid(s.TestData.Uid, int64(888881000000002))
	if err != nil {
		s.T().Errorf("append album failed")
		s.T().FailNow()
	}
}

func (s *AlbumServiceSuite) TearDownTest() {
	ami := AlbumModelImpl{}
	err := ami.DeleteAlbumByUid(s.TestData.Uid)
	if err != nil {
		s.T().Errorf("delete album failed")
		s.T().FailNow()
	}
}

func (s *AlbumServiceSuite) TestServiceAppendAlbum() {
	srv := NewAlbumService()
	srv.Data.Uid = s.TestData.Uid
	srv.Data.AidList = []int64{888881000000001}
	err := srv.AppendAlbum(nil)
	assert.Nil(s.T(), err)
}

func (s *AlbumServiceSuite) TestServiceCreateAlbum() {
	srv := NewAlbumService()
	srv.Data.Uid = 80808
	err := srv.CreateAlbum(nil)
	assert.Nil(s.T(), err)
}

func (s *AlbumServiceSuite) TestServiceDeleteArticleInAlbum() {
	srv := NewAlbumService()
	srv.Data.Uid = 80808
	srv.Data.AidList = []int64{888881000000001}
	err := srv.DeleteArticleInAlbum(nil)
	assert.Nil(s.T(), err)
}

func (s *AlbumServiceSuite) TestServiceDetailAlbum() {
	srv := NewAlbumService()
	srv.Data.Uid = 88888
	err := srv.DetailAlbum(nil)
	assert.Nil(s.T(), err)
}

func TestAlbumServiceSuite(t *testing.T) {
	suite.Run(t, new(AlbumServiceSuite))
}
