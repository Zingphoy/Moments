package models

import (
	"Moments/pkg/log"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	TIMEOUT  time.Duration = 8
	URI                    = "mongodb://localhost:8080" // uri := "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"
	DATABASE               = "Moments"
	db       *mongo.Database
	op       *options.ClientOptions
)

// initialize some basic values
func init() {
	op = options.Client().ApplyURI(URI)
	op.SetMaxPoolSize(5)
}

func Ping() (string, error) {
	_, client, ctx, _ := ConnectDatabase()
	err := client.Ping(ctx, readpref.Primary())
	return "success", err
}

// ConnectDatabase call this first when you need to communicate with database
// call client.Disconncet(ctx) when finished database query
func ConnectDatabase() (*mongo.Database, *mongo.Client, context.Context, error) {
	ctx, _ := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	client, err := mongo.Connect(ctx, op)
	if err != nil {
		log.Error("error while connecting to database")
		return nil, nil, nil, err
	}
	db = client.Database(DATABASE)
	return db, client, ctx, nil
}
