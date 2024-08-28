// backend/models/pocModel.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type POC struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Title       string             `bson:"title" json:"title"`
    Description string             `bson:"description" json:"description"`
    Completed   bool               `bson:"completed" json:"completed"`
}


