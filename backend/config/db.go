//Create the db.go file in the config folder to handle the MongoDB connection:
// backend/config/db.go
package config

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() *mongo.Database {
    clientOptions := options.Client().ApplyURI("mongodb+srv://gouser:gopassword@cluster0.zichifw.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")
    DB = client.Database("poc-go-db")
    return DB
}

