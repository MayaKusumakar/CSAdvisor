package mongoApi

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"slices"
)

func GetMongoClient() (*mongo.Client, error) {
	//certPath := "../../backend/secrets/X509-cert-4597202730823277483.pem"
	uri := "mongodb+srv://swipengo.wchxnfh.mongodb.net/?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&appName=SwipeNGo&tlsCertificateKeyFile=/Users/ricorodriguez/Documents/GitHub/Swipe-Go/backend/cmd/api/X509-cert-4597202730823277483.pem"
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client, err
	//defer client.Disconnect(context.TODO())
	//collection := client.Database("sample_mflix").Collection("comments")
	//docCount, err := collection.CountDocuments(context.TODO(), bson.D{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(docCount)
}

func GetCollection(client *mongo.Client, dbName string, collectionName string) (*mongo.Collection, error) {
	// Check if database exists
	dbNames, err := client.ListDatabaseNames(context.Background(), nil)
	if err != nil {
		fmt.Println("HERE1")
		return nil, err
	}
	if !slices.Contains(dbNames, dbName) {
		return nil, fmt.Errorf("database {%s} not found", dbName)
	}

	database := client.Database(dbName)
	collectionNames, err := database.ListCollectionNames(context.Background(), nil)
	if err != nil {
		fmt.Println("HERE")
		return nil, err
	}
	if !slices.Contains(collectionNames, collectionName) {
		return nil, fmt.Errorf("collection {%s} not found in database {%s}", collectionName, dbName)
	}
	collection := client.Database(dbName).Collection(collectionName)
	return collection, nil
}
