package properties

import (
	"fmt"
	"github.com/magiconair/properties"
	"os"
)

var p *properties.Properties

func init() {
	filename := "./properties/"
	switch GetEnv() {
	case "staging":
		filename += "staging.properties"
		break
	case "prod":
		filename += "prod.properties"
		break
	default:
		os.Setenv("env", "local")
		filename += "local.properties"
		break
	}
	p = properties.MustLoadFile(filename, properties.UTF8)
}

func GetEnv() string {
	return os.Getenv("ENV")
}

type MongoProperties struct {
	User               string
	Password           string
	Host               string
	Database           string
	ChatsCollection    string
	MessagesCollection string
	AuthCollection     string
}

func GetMongoProperties() MongoProperties {
	mongoPassword := p.GetString("mongo.password", os.Getenv("MONGO_PASSWORD"))
	p := MongoProperties{
		User:               p.MustGetString("mongo.user"),
		Password:           mongoPassword,
		Host:               p.MustGetString("mongo.host"),
		Database:           p.MustGetString("mongo.database"),
		ChatsCollection:    p.MustGetString("mongo.collections.chats"),
		MessagesCollection: p.MustGetString("mongo.collections.messages"),
		AuthCollection:     p.MustGetString("mongo.collections.auth"),
	}
	fmt.Println(p)
	return p
}
