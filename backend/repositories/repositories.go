// backend/repositories/pocRepository.go
package repositories

import (
    "context"
    "github.com/klnswamy1702/poc-go-app/backend/config"
    "github.com/klnswamy1702/poc-go-app/backend/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type POCRepository struct {
    Collection *mongo.Collection
}

func NewPOCRepository() *POCRepository {
    return &POCRepository{
        Collection: config.DB.Collection("poc"),
    }
}

func (repo *POCRepository) CreatePOC(poc *models.POC) error {
    _, err := repo.Collection.InsertOne(context.TODO(), poc)
    return err
}

func (repo *POCRepository) GetAllPOCs() ([]models.POC, error) {
    var pocs []models.POC
    cursor, err := repo.Collection.Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())
    for cursor.Next(context.TODO()) {
        var poc models.POC
        err := cursor.Decode(&poc)
        if err != nil {
            return nil, err
        }
        pocs = append(pocs, poc)
    }
    return pocs, nil
}

func (repo *POCRepository) GetPOCByID(id primitive.ObjectID) (*models.POC, error) {
    var poc models.POC
    err := repo.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&poc)
    return &poc, err
}

func (repo *POCRepository) UpdatePOC(id primitive.ObjectID, poc *models.POC) error {
    _, err := repo.Collection.UpdateOne(
        context.TODO(),
        bson.M{"_id": id},
        bson.M{"$set": poc},
    )
    return err
}

func (repo *POCRepository) DeletePOC(id primitive.ObjectID) error {
    _, err := repo.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
    return err
}
