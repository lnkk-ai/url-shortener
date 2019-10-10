package store

import (
	"context"
	"fmt"
	syslog "log"
	"time"

	"github.com/majordomusio/commons/pkg/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var _client *mongo.Client
var _db *mongo.Database

func initM() {

	var databaseURI string

	// mongodb://user:password@myhost1.mydomain.com:27017/my_db

	host := env.Getenv("DATABASE_SERVICE_NAME", "localhost")
	port := env.Getenv("DATABASE_PORT", "27017")
	database := env.Getenv("MONGODB_DATABASE", "shadowman")

	user := env.Getenv("MONGODB_USER", "")
	if user == "" {
		databaseURI = fmt.Sprintf("mongodb://%s:%s/%s", host, port, database)
	} else {
		password := env.Getenv("MONGODB_PASSWORD", "")
		databaseURI = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", user, password, host, port, database)
	}

	// create a client and connect to it
	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURI))
	if err != nil {
		syslog.Fatalf("Store: can't create client. Resons: %v ", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		syslog.Fatalf("Store: can't connect to database. Resons: %v", err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		syslog.Fatalf("Store: Can't ping database. Resons: %v", err)
	}

	_client = client
	_db = _client.Database(database)

	//log.Info("store", "init", "initializing")
}

// Close discconects from the database
func Close() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := _client.Disconnect(ctx)
	if err != nil {
		//log.Error("store", "close", err)
	}
	_client = nil
	_db = nil

	//log.Info("store", "close", "disconnected")
}
