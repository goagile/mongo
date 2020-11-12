package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/goagile/mongo/post"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBTimeout = 10 * time.Second
	DBURI     = "mongodb://127.0.0.1:27017"
	DB        *mongo.Database
	DBClient  *mongo.Client
)

func DBConnectClient(ctx context.Context) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Connect(ctx); err != nil {
		log.Fatal()
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	return client
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), DBTimeout)
	DBClient := DBConnectClient(ctx)
	defer DBClient.Disconnect(ctx)

	DB = DBClient.Database("tutorial")
	post.Collection = DB.Collection("posts")

	p := post.New("XXX", "To many text ...")
	if err := p.Save(); err != nil {
		log.Fatal(err)
	}

	p2, err := post.FindByTitle(p.Title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("p2", p2)
}
