package authright

import (
	"context"
	"cryptchat/auth/core"
	properties2 "cryptchat/properties"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var collection = getMongoCollection()

func getMongoCollection() *mongo.Collection {
	p := properties2.GetMongoProperties()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017/?authSource=%s",
		p.User, p.Password, p.Host, p.Database)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	return client.Database(p.Database).Collection(p.AuthCollection)
}

var SaveUserMongo = func(user authcore.User) error {
	_, insertErr := collection.InsertOne(context.TODO(), user)
	return insertErr
}

var FindUserMongo = func(username string) (*authcore.User, error) {
	filter := bson.D{{"username", username}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result = new(authcore.User)
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return result, nil
}
