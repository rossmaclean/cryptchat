package chatsright

import (
	"context"
	"cryptchat/chats/core"
	properties2 "cryptchat/properties"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var chatsCollection = getMongoCollection(chatsCollectionName)
var messagesCollection = getMongoCollection(messagesCollectionName)

var chatsCollectionName string
var messagesCollectionName string

func getMongoCollection(collection string) *mongo.Collection {
	p := properties2.GetMongoProperties()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017/?authSource=%s",
		p.User, p.Password, p.Host, p.Database)

	clientOptions := options.Client().ApplyURI(uri)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	return client.Database(p.Database).Collection(collection)
}

var GetChatsForUserMongo = func(userId string) ([]chatscore.ChatMongo, error) {
	filterCursor, err := chatsCollection.Find(context.TODO(), bson.M{"userIds": userId})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var chats []chatscore.ChatMongo
	if err = filterCursor.All(context.TODO(), &chats); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return chats, nil
}

var GetChatMessagesMongo = func(chatId string) (chatscore.ChatMessageMongo, error) {
	var messages chatscore.ChatMessageMongo
	if err := messagesCollection.FindOne(context.TODO(), bson.M{"chatId": chatId}).Decode(&messages); err != nil {
		log.Fatal(err)
		return chatscore.ChatMessageMongo{}, nil
	}
	return messages, nil
}

var SaveChatMongo = func(chat chatscore.ChatMongo) error {
	_, err := chatsCollection.InsertMany(context.TODO(), []interface{}{
		bson.D{
			{"chatId", chat.ChatId},
			{"userIds", chat.UserIds},
			{"timestamp", chat.Timestamp},
		},
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

var SaveChatMessageMongo = func(chatMessage chatscore.ChatMessageMongo) error {
	_, err := chatsCollection.InsertMany(context.TODO(), []interface{}{
		bson.D{
			{"chatId", chatMessage.ChatId},
			{"messages", chatMessage.Messages},
		},
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
