package mongo

import (
	"context"
	"time"

	c "github.com/mrceylan/go-url-shortener/pkg/configuration"
	"github.com/mrceylan/go-url-shortener/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
)

func (conn *Connection) AddLog(log logging.RedirectLog) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Config.Mongo.DefaultTimeOut*time.Second)
	defer cancel()
	collection := conn.Client.Database(c.Config.Mongo.LogDbName).Collection(c.Config.Mongo.LogCollectionName)
	_, err := collection.InsertOne(ctx, log)
	if err != nil {
		return err
	}
	return nil
}

func (conn *Connection) GetLogs() ([]logging.RedirectLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Config.Mongo.DefaultTimeOut*time.Second)
	defer cancel()
	collection := conn.Client.Database(c.Config.Mongo.LogDbName).Collection(c.Config.Mongo.LogCollectionName)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var results []logging.RedirectLog
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
