package article

import (
	"Moments/biz/database"
	"Moments/pkg/hint"
	"Moments/pkg/utils"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// article data structure
type Article struct {
	// basic fields of Article, it should not be edited
	Aid       int64    `bson:"aid" json:"aid"`
	Uid       int32    `bson:"uid" json:"uid"`
	PostTime  int64    `bson:"post_time" json:"post_time"`
	Content   string   `bson:"content" json:"content"`
	PhotoList []string `bson:"photo_list" json:"photo_list"`

	// extra fields
	Privacy   int32 `bson:"privacy" json:"privacy"`
	IsDeleted int32 `bson:"is_deleted" json:"is_deleted"`
}

type ArticleModel interface {
	GenerateAid(uid int32) (int64, error)
	GetArticleDatabase(aid int64) string
	IsArticleNotExist(aid int64) error
	GetArticleDetailByAid(aid int64) (*Article, error)
	AddArticle(article *Article) error
	//UpdateArticle(filter Map, data Map) error
	DeleteArticleByUidAid(uid int32, aid int64) error
	DeleteArticleSoftByUidAid(uid int32, aid int64) error
}

/********** Real Implement **********/

type ArticleModelImpl struct {
}

// GenerateAid also restrict article sending frequency of one user to 1 time per second because of redundant aid
func (a *ArticleModelImpl) GenerateAid(uid int32) (int64, error) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	tmp := strconv.FormatInt(int64(uid), 10) + ts
	aid := utils.Str(tmp).MustInt64()
	err := a.IsArticleNotExist(aid)
	if err != nil {
		return 0, err
	}
	return aid, nil
}

// GetArticleDatabase articles has been split into 4 collections, find the correct collection
func (a *ArticleModelImpl) GetArticleDatabase(aid int64) string {
	dbname := "article_" + strconv.Itoa(int(aid%4))
	return dbname
}

// IsArticleNotExist check if specific aid is already existed
func (a *ArticleModelImpl) IsArticleNotExist(aid int64) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		// shall not return false, otherwise aid would be redundant
		return err
	}
	defer client.Disconnect()

	rows, err := client.Query(a.GetArticleDatabase(aid), database.Map{"aid": aid})
	if cErr, ok := err.(hint.CustomError); ok && cErr.Code == hint.QUERY_INTERNAL_ERROR {
		return err
	}
	if len(rows) > 0 {
		return hint.CustomError{
			Code: hint.AID_ALREADY_EXIST,
			Err:  errors.New("aid already exists"),
		}
	}
	return nil
}

// GetArticleDetailByAid get detail of an Article with specific filter
func (a *ArticleModelImpl) GetArticleDetailByAid(aid int64) (*Article, error) {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect()

	dbname := a.GetArticleDatabase(aid)
	rows, err := client.Query(dbname, database.Map{"aid": aid})
	if cErr, ok := err.(hint.CustomError); ok && cErr.Code == hint.QUERY_INTERNAL_ERROR {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, hint.CustomError{
			Code: hint.AID_NOT_FOUND,
			Err:  errors.New("article not found"),
		}
	}

	row := rows[0]
	article := Article{}
	article.Aid = row["aid"].(int64)
	article.Uid = row["uid"].(int32)
	article.PostTime = row["post_time"].(int64)
	article.Content = row["content"].(string)
	article.PhotoList = database.BsonAToSliceString(row["photo_list"])
	article.Privacy = row["privacy"].(int32)
	article.IsDeleted = 0
	return &article, nil
}

// AddArticle add Article to database, and expand this Article into friends' timeline
func (a *ArticleModelImpl) AddArticle(art *Article) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	err = client.Insert(a.GetArticleDatabase(art.Aid), []interface{}{art})
	if cErr, ok := err.(hint.CustomError); ok && cErr.Code == hint.INSERT_INTERNAL_ERROR {
		return err
	}
	return nil
}

// DeleteArticleByUidAid delete Article permanently
func (a *ArticleModelImpl) DeleteArticleByUidAid(uid int32, aid int64) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	err = client.Remove(a.GetArticleDatabase(aid), database.Map{"aid": aid, "uid": uid})
	if cErr, ok := err.(hint.CustomError); ok && cErr.Code == hint.DELETE_INTERNAL_ERROR {
		return err
	}
	return nil
}

// DeleteArticleSoftByUidAid delete Article softly
func (a *ArticleModelImpl) DeleteArticleSoftByUidAid(uid int32, aid int64) error {
	client := database.NewDatabaseClient()
	err := client.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect()

	filter := database.Map{"aid": aid, "uid": uid}
	updatedData := database.Map{"is_deleted": 1}
	err = client.Update(a.GetArticleDatabase(aid), filter, updatedData)
	if cErr, ok := err.(hint.CustomError); ok && cErr.Code == hint.UPDATE_INTERNAL_ERROR {
		return err
	}
	return nil
}

func (a *ArticleModelImpl) Update() error {
	return nil
}
