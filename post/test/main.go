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
)

func main() {
	opts := options.Client().ApplyURI(DBURI)

	client, err := mongo.NewClient(opts)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), DBTimeout)

	if err := client.Connect(ctx); err != nil {
		log.Fatal()
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	DB = client.Database("tutorial")

	post.Collection = DB.Collection("posts")

	p := post.New("XXX", "To many text ...")
	if err := p.Save(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(p)

	// p2, err := post.Find(p.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(p2)

	p2, err := post.FindByTitle(p.Title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("p2", p2)
}
