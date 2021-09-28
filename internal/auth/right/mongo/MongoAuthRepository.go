package authrightmongo

import (
	"context"
	"cryptchat/internal"
	"cryptchat/internal/auth/core/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client mongo.Client
var collection mongo.Collection

type MongoAuthRepository struct {
	client     mongo.Client
	collection mongo.Collection
}

func (r *MongoAuthRepository) Ping() error {
	return client.Ping(context.TODO(), nil)
}

func (r *MongoAuthRepository) InitDb() {
	collectionName := "users"

	log.Printf("Connecting to MongoDB %s", collectionName)
	p := internal.GetMongoProperties()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017/?authSource=%s",
		p.User, p.Password, p.Host, p.Database)

	clientOptions := options.Client().ApplyURI(uri)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clnt, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Println(err)
	}
	err = clnt.Ping(context.TODO(), nil)

	if err != nil {
		log.Println(err)
	}
	log.Printf("Connected to MongoDB %s", collectionName)
	client = *clnt
	collection = *clnt.Database(p.Database).Collection(collectionName)
}

func (r *MongoAuthRepository) SaveOne(user authcoremodel.User) error {
	_, insertErr := collection.InsertOne(context.TODO(), user)
	return insertErr
}

func (r *MongoAuthRepository) FindOneByEmail(email string) (authcoremodel.User, error) {
	filter := bson.D{{"email", email}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result = new(authcoremodel.User)
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return authcoremodel.User{}, nil
	} else if err != nil {
		return authcoremodel.User{}, err
	}
	return *result, nil
}
