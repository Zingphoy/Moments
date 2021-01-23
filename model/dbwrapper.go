package model

import (
	"context"
	"time"

	"github.com/pkg/errors"
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
		return errors.Wrap(err, "mongodb connect failed")
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
	db = e.client.Database(DATABASE)
	ctx, cancelFun := context.WithTimeout(context.Background(), DATABASE_TIMEOUT)
	defer cancelFun()

	c := db.Collection(collection)
	cur, err := c.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "mongodb query failed")
	}

	var rows []bson.M
	if err = cur.All(ctx, &rows); err != nil {
		return nil, err
	}
	data := make([]Map, 0, 128)
	for _, row := range rows {
		data = append(data, BsonMToMap(row))
	}
	return data, nil
}

func (e *mongoEngine) Update(collection string, filter Map, new Map) error {
	db = e.client.Database(DATABASE)
	ctx, cancelFun := context.WithTimeout(context.Background(), DATABASE_TIMEOUT)
	defer cancelFun()

	c := db.Collection(collection)
	_, err := c.UpdateOne(ctx, filter, new)
	if err != nil {
		return errors.Wrap(err, "mongodb update failed")
	}
	return nil
}

func (e *mongoEngine) Insert(collection string, data []interface{}) error {
	db = e.client.Database(DATABASE)
	ctx, cancelFun := context.WithTimeout(context.Background(), DATABASE_TIMEOUT)
	defer cancelFun()

	c := db.Collection(collection)
	_, err := c.InsertMany(ctx, data)
	if err != nil {
		return errors.Wrap(err, "mongodb insert failed")
	}
	return nil
}

func (e *mongoEngine) Remove(collection string, filter Map) error {
	db = e.client.Database(DATABASE)
	ctx, cancelFun := context.WithTimeout(context.Background(), DATABASE_TIMEOUT)
	defer cancelFun()

	c := db.Collection(collection)
	_, err := c.DeleteMany(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "mongodb delete failed")
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
