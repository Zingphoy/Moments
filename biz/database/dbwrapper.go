package database

import (
	"Moments/pkg/hint"
	"Moments/pkg/log"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	URI              = "mongodb://localhost:8080" // uri := "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"
	DATABASE         = "Moments"
	DATABASE_TIMEOUT = 5 * time.Second
	Client           DatabaseEngine
	ENGINE_TYPE      = "mongo"
)

type Map = map[string]interface{}

type DatabaseEngine interface {
	Query(string, Map) ([]Map, error)
	Update(string, Map, Map) error
	Insert(string, []interface{}) error
	Remove(string, Map) error
	Connect() error
	Disconnect() error
}

func NewDatabaseClient() DatabaseEngine {
	switch ENGINE_TYPE {
	case "mongo":
		return &mongoEngine{}
	default:
		return &mongoEngine{}
	}
}

type mongoEngine struct {
	client *mongo.Client
}

func (e *mongoEngine) Connect() error {
	var err error
	clientOptions := options.Client().ApplyURI(URI)
	e.client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Error(nil, "mongodb connect failed:", err.Error())
		return hint.CustomError{
			Code: hint.CONNECT_FAILED,
			Err:  err,
		}
	}
	return nil
}

func (e *mongoEngine) Disconnect() error {
	if e != nil {
		return e.client.Disconnect(context.TODO())
	}
	return nil
}

func (e *mongoEngine) Query(collection string, filter Map) ([]Map, error) {
	db := e.client.Database(DATABASE)
	ctx, cancelFun := context.WithTimeout(context.Background(), DATABASE_TIMEOUT)
	defer cancelFun()

	c := db.Collection(collection)
	cur, err := c.Find(ctx, filter)
	if err != nil {
		log.Error(nil, "mongodb query failed:", err.Error())
		return nil, hint.CustomError{
			Code: hint.QUERY_INTERNAL_ERROR,
			Err:  err,
		}
	}
	var rows []bson.M
	if err = cur.All(ctx, &rows); err != nil {
		log.Error(nil, "traverse query data failed:", err.Error())
		return nil, hint.CustomError{
			Code: hint.QUERY_INTERNAL_ERROR,
			Err:  err,
		}
	}

	data := make([]Map, 0, 128)
	for _, row := range rows {
		data = append(data, BsonMToMap(row))
	}
	return data, nil
}

func (e *mongoEngine) Update(collection string, filter Map, new Map) error {
	db := e.client.Database(DATABASE)
	ctx, cancelFun := context.WithTimeout(context.Background(), DATABASE_TIMEOUT)
	defer cancelFun()

	old := bson.D{}
	for k, v := range filter {
		old = append(old, bson.E{k, v})
	}

	var arg = bson.D{}
	var up = bson.D{{"$set", nil}}
	for k, v := range new {
		arg = append(arg, bson.E{k, v})
	}
	up[0].Value = arg

	c := db.Collection(collection)
	_, err := c.UpdateOne(ctx, old, up)
	if err != nil {
		log.Error(nil, "mongodb update data failed:", err.Error())
		return hint.CustomError{
			Code: hint.UPDATE_INTERNAL_ERROR,
			Err:  err,
		}
	}
	return nil
}

func (e *mongoEngine) Insert(collection string, data []interface{}) error {
	db := e.client.Database(DATABASE)
	ctx, cancelFun := context.WithTimeout(context.Background(), DATABASE_TIMEOUT)
	defer cancelFun()

	c := db.Collection(collection)
	_, err := c.InsertMany(ctx, data)
	if err != nil {
		log.Error(nil, "mongodb insert data failed:", err.Error())
		return hint.CustomError{
			Code: hint.INSERT_INTERNAL_ERROR,
			Err:  err,
		}
	}
	return nil
}

func (e *mongoEngine) Remove(collection string, filter Map) error {
	db := e.client.Database(DATABASE)
	ctx, cancelFun := context.WithTimeout(context.Background(), DATABASE_TIMEOUT)
	defer cancelFun()

	c := db.Collection(collection)
	_, err := c.DeleteMany(ctx, filter)
	if err != nil {
		log.Error(nil, "mongo delete data failed:", err.Error())
		return hint.CustomError{
			Code: hint.DELETE_INTERNAL_ERROR,
			Err:  err,
		}
	}
	return nil
}

func MapToBsonM(data Map) bson.M {
	ret := make(bson.M)
	for k, v := range data {
		ret[k] = v
	}
	return ret
}

func BsonMToMap(data bson.M) Map {
	ret := make(Map)
	for k, v := range data {
		ret[k] = v
	}
	return ret
}

func BsonAToSliceString(data interface{}) []string {
	val, ok := data.(bson.A)
	if !ok || val == nil {
		return make([]string, 0, 0)
	}
	strSlice := make([]string, len(val))
	for _, v := range val {
		s := v.(string)
		strSlice = append(strSlice, s)
	}
	return strSlice
}

func BsonToSliceInt64(data interface{}) []int64 {
	val, ok := data.(bson.A)
	if !ok || val == nil {
		return make([]int64, 0, 0)
	}
	intSlice := make([]int64, len(val))
	for _, v := range val {
		s := v.(int64)
		intSlice = append(intSlice, s)
	}
	return intSlice
}
