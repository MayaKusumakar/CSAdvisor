package mongoApi

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"slices"
)

const (
	mongoURI         = "mongodb+srv://swipengo.wchxnfh.mongodb.net/?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority&appName=SwipeNGo&tlsCertificateKeyFile=/Users/ricorodriguez/Documents/GitHub/Swipe-Go/backend/secrets/X509-cert-3424146718248376778.pem"
	eventsDb         = "main"
	eventsCollection = "events"
	usersDb          = "main"
	usersCollection  = "users"
)

type Event struct {
	//Id           primitive.ObjectID `bson:"_id,omitempty"`
	OwnerId      string   `bson:"ownerId"`
	OwnerName    string   `bson:"ownerName"`
	Title        string   `bson:"title"`
	Location     Location `bson:"location"`
	StartTime    string   `bson:"startTime"`
	EndTime      string   `bson:"endTime"`
	Description  string   `bson:"description"`
	NumAttending int      `bson:"numAttending"`
}

type Location struct {
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
	Address   string  `bson:"address"`
}

func GetMongoClient() (*mongo.Client, error) {
	//certPath := "../../backend/secrets/X509-cert-4597202730823277483.pem"
	uri := mongoURI
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client, err
}

func GetCollection(client *mongo.Client, dbName string, collectionName string) (*mongo.Collection, error) {
	// Check if database exists
	dbNames, err := client.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if !slices.Contains(dbNames, dbName) {
		return nil, fmt.Errorf("database {%s} not found", dbName)
	}

	database := client.Database(dbName)
	collectionNames, err := database.ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if !slices.Contains(collectionNames, collectionName) {
		return nil, fmt.Errorf("collection {%s} not found in database {%s}", collectionName, dbName)
	}
	collection := client.Database(dbName).Collection(collectionName)
	return collection, nil
}

func AddEvent(collection *mongo.Collection, event Event) error {
	res, err := collection.InsertOne(context.Background(), event, options.InsertOne())
	if err != nil {
		log.Fatal(err)
		return err
	}

	if res.Acknowledged {
		return nil
	}
	return fmt.Errorf("event {%s} not acknowledged", res.InsertedID)
}

func GetEvents() ([]*Event, error) {
	var events []*Event

	// Connect to database
	client, err := GetMongoClient()
	if err != nil {
		log.Fatal(err)
		return events, err
	}

	defer client.Disconnect(context.Background())

	// Grab events collection
	coll := client.Database(eventsDb).Collection(eventsCollection)

	// Format events collection to slice of Events
	cursor, err := coll.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
		return events, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var event Event
		err := cursor.Decode(&event)
		if err != nil {
			log.Fatal(err)
			return events, err
		}
		events = append(events, &event)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return events, err
	}
	return events, nil
}
