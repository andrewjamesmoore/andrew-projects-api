package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBNote struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	URL         string             `bson:"url"`
	Date        string             `bson:"date"`
}
