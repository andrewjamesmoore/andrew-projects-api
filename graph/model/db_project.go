package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBProject struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Status      string             `bson:"status"`
	Link        bool               `bson:"link"`
	URL         string             `bson:"url"`
	Giturl      *string            `bson:"giturl,omitempty"`
	Tags        []string           `bson:"tags"`
}
