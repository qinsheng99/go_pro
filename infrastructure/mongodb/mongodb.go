package mongodb

import (
	"context"
	"fmt"

	"github.com/qinsheng99/go-domain-web/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgo *mongo.Client

type Mongo struct {
	mo         *mongo.Client
	database   *mongo.Database
	collection string
}

func Init(cfg *config.MongoConfig) error {
	var err error
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port))

	// 连接到MongoDB
	mgo, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	// 检查连接
	err = mgo.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	return nil
}

func GetMongo() *mongo.Client {
	return mgo
}
