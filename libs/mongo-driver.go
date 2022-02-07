package libs

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Client *mongo.Client
	ConErr error
)

func StartCon() {
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	Client, ConErr = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if ConErr != nil {
		panic(ConErr)
	}

	if ConErr = Client.Ping(context.TODO(), readpref.Primary()); ConErr != nil {
		panic(ConErr)
	}

	fmt.Println("Successfully connected to mongoDB and pinged.")
}

func CloseCon() {
	if ConErr = Client.Disconnect(context.TODO()); ConErr != nil {
		panic(ConErr)
	}
}
