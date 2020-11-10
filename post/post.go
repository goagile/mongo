package post

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Collection *mongo.Collection
)

//
// New
//
func New(title, body string) *Post {
	return &Post{Title: title, Body: body}
}

//
// Post
//
type Post struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title,omitempty"`
	Body  string             `bson:"body,omitempty"`
}

//
// String
//
func (p *Post) String() string {
	return fmt.Sprintf("Post:\n\t%v\n\t%v\n\t%v\n", p.ID, p.Title, p.Body)
}

//
// Save
//
func (p *Post) Save() error {
	res, err := Collection.InsertOne(context.TODO(), p)
	if err != nil {
		return err
	}
	p.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//
// Find
//
func Find(id primitive.ObjectID) (*Post, error) {
	filter := bson.M{"_id": id}
	ctx := context.TODO()

	cursor, err := Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("Collection.Find:%v", err)
	}

	defer cursor.Close(ctx)

	var p *Post
	if err := cursor.Decode(&p); err != nil {
		return nil, fmt.Errorf("Collection.Decode:%v", err)
	}

	return p, nil
}

//
// FindByTitle
//
func FindByTitle(title string) (*Post, error) {
	ctx := context.TODO()
	filter := bson.M{"title": title}
	cursor, err := Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("Collection.Find:%v", err)
	}
	var posts []*Post
	if err := cursor.All(ctx, &posts); err != nil {
		return nil, fmt.Errorf("cursor.All:%v", err)
	}
	if len(posts) < 1 {
		return nil, fmt.Errorf("Len posts < 1")
	}
	return posts[0], nil
}
