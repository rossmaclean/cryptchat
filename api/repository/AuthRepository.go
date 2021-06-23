package repository

import (
	"cryptchat/model"
	"github.com/gocql/gocql"
	"log"
)

var session gocql.Session

func GetDbSession() gocql.Session {
	return session
}

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

var SaveUser = func(user model.User) error {
	if err := session.Query(`INSERT INTO user (user_id, username, email, hashed_password) VALUES (?, ?, ?, ?)`,
		user.UserId, user.Username, user.Email, user.HashedPassword).Exec(); err != nil {
		return err
	}
	return nil
}

var GetUser = func(username string) (*model.User, error) {
	user := new(model.User)
	if err := session.Query(`SELECT user_id, username, email, hashed_password FROM user WHERE username = ? LIMIT 1`,
		username).Consistency(gocql.One).Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.HashedPassword); err != nil {
		return nil, err
	}
	return user, nil
}
