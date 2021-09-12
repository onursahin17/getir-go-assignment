package Database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// Initializes the mongodb connection
func InitDbConnection(mongodbUri string) (context.Context, *mongo.Client) {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbUri))
	if err != nil {
		log.Fatal(err)
	}
	// context with 1000 seconds timeout
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return ctx, client
}
