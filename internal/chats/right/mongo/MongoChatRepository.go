package chatsrightmongo

import (
	"context"
	"cryptchat/internal"
	chatscoremodel "cryptchat/internal/chats/core/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Find = collection.Find

var collection mongo.Collection

type MongoChatRepository struct {
	collection mongo.Collection
}

func (r *MongoChatRepository) FindAllByUserId(userId string) ([]chatscoremodel.Chat, error) {
	filterCursor, err := Find(context.TODO(), bson.M{"userIds": userId})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// TODO for JSON I need the extra at the end, verify
	var chats []chatscoremodel.Chat = []chatscoremodel.Chat{}
	if err = filterCursor.All(context.TODO(), &chats); err != nil {
		log.Println(err)
		return nil, err
	}
	return chats, nil
}

func (r *MongoChatRepository) InitDb() {
	collectionName := "chats"

	log.Printf("Connecting to MongoDB %s", collectionName)
	p := internal.GetMongoProperties()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017/?authSource=%s",
		p.User, p.Password, p.Host, p.Database)

	clientOptions := options.Client().ApplyURI(uri)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Println(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println(err)
	}
	log.Printf("Connected to MongoDB %s", collectionName)
	collection = *client.Database(p.Database).Collection(collectionName)
}

//var GetChatsForUserMongo = func(userId string) ([]ChatMongo, error) {
//	filterCursor, err := Find(context.TODO(), bson.M{"userIds": userId})
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//	var chats []ChatMongo
//	if err = filterCursor.All(context.TODO(), &chats); err != nil {
//		log.Println(err)
//		return nil, err
//	}
//	return chats, nil
//}

//var GetChatMessagesMongo = func(chatId string) (ChatMessageMongo, error) {
//	var messages ChatMessageMongo
//	if err := messagesCollection.FindOne(context.TODO(), bson.M{"chatId": chatId}).Decode(&messages); err != nil {
//		log.Println(err)
//		return ChatMessageMongo{}, nil
//	}
//	return messages, nil
//}
//
//var SaveChatMongo = func(chat ChatMongo) error {
//	_, err := chatsCollection.InsertMany(context.TODO(), []interface{}{
//		bson.D{
//			{"chatId", chat.ChatId},
//			{"userIds", chat.UserIds},
//			{"timestamp", chat.Timestamp},
//		},
//	})
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	return nil
//}
//
//var SaveChatMessageMongo = func(chatMessage ChatMessageMongo) error {
//	_, err := chatsCollection.InsertMany(context.TODO(), []interface{}{
//		bson.D{
//			{"chatId", chatMessage.ChatId},
//			{"messages", chatMessage.Messages},
//		},
//	})
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	return nil
//}
