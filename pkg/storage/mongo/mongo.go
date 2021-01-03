package mongo

import (
	"context"
	"fmt"

	c "github.com/mrceylan/go-url-shortener/pkg/configuration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Client *mongo.Client
}

func MongoInit() *Connection {
	defer fmt.Println("Mongo client connected")
	clientOptions := options.Client().ApplyURI(c.Config.Mongo.Uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	return &Connection{client}
}
