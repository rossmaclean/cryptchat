# CryptChat

## Goal

The goal of this project is to create chat service which is encrypted e2e. The application will be web base, use a Go
backend and Cassandra for storage.

## Design

### Database

I need to be able to perform the following queries:

* Get users' hashed_password by username
* Get all the chats the user is part of
* Get the messages for a chat the user is part of
* Send messages to chats
* Leave group chats
* Create chats
* Create group chats
* Enter group chats

Initial Table Ideas

```cassandraql
CREATE TABLE user
(
    user_id         varchar,
    username        varchar,
    email           varchar,
    hashed_password varchar,
    PRIMARY KEY ((username), user_id)
);
```

```cassandraql
CREATE TABLE chat_messages
(
    message_id uuid,
    from_user  text,
    to_user    text,
    message    text,
    timestamp  timestamp,
    PRIMARY KEY ((from_user, to_user), timestamp)
) WITH CLUSTERING ORDER BY (timestamp DESC);
```