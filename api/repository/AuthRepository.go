package repository

import (
	"cryptchat/model"
	"github.com/gocql/gocql"
	"log"
)

var session gocql.Session

const userAuthTable = "user_auth"

func InitDbSession() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "cryptchat"
	cluster.Consistency = gocql.Quorum
	s, err := cluster.CreateSession()
	if err != nil {
		s.Close()
		log.Fatal(err)
	}
	session = *s
	log.Print("Connected to Cassandra")
}

var SaveUserAuth = func(userAuth model.UserAuth) error {
	if err := session.Query(`INSERT INTO user_auth (username, hashed_password) VALUES (?, ?)`,
		userAuth.Username, userAuth.HashedPassword).Exec(); err != nil {
		return err
	}
	return nil
}

var GetUserAuth = func(username string) (*model.UserAuth, error) {
	userAuth := new(model.UserAuth)
	if err := session.Query(`SELECT username, hashed_password FROM user_auth WHERE username = ? LIMIT 1`,
		username).Consistency(gocql.One).Scan(&userAuth.Username, &userAuth.HashedPassword); err != nil {
		return nil, err
	}
	return userAuth, nil
}
