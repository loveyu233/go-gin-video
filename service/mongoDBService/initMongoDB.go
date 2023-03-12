package mongoDBService

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	db *mongo.Database
}

func (m *MongoClient) Comment() *mongo.Collection {
	return m.db.Collection("comment")
}
func (m *MongoClient) Reply() *mongo.Collection {
	return m.db.Collection("reply")
}

var mongoClient *MongoClient

func InitMongoDb() error {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?%s",
		viper.GetString("mongodb.username"),
		viper.GetString("mongodb.password"),
		viper.GetString("mongodb.host"),
		viper.GetString("mongodb.port"),
		viper.GetString("mongodb.param"),
	)
	connect, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	mongoClient = &MongoClient{db: connect.Database(viper.GetString("mongodb.dbname"))}
	return nil
}
