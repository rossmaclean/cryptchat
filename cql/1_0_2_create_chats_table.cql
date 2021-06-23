CREATE TABLE chats
(
    chat_id   uuid,
    user      text,
    timestamp timestamp,
    PRIMARY KEY ((user), chat_id, timestamp)
) WITH CLUSTERING ORDER BY (timestamp DESC);

CREATE TABLE IF NOT EXISTS user
(
    user_id         varchar,
    username        varchar,
    email           varchar,
    hashed_password varchar,
    PRIMARY KEY ((username), user_id)
);

CREATE TABLE chats
(
    chat_id   text,
    user      text,
    timestamp text,
    PRIMARY KEY ((user), chat_id, timestamp)
) WITH CLUSTERING ORDER BY (chat_id DESC, timestamp DESC);

CREATE TABLE User1.User_information (
                                        User_name text PRIAMRY KEY,
                                        User_email text,
                                        User_password text,
                                        User_address text
);

CREATE MATERIALIZED VIEW User1.Users_by_User_email AS
SELECT User_email, User_password, User_address
FROM User_information
WHERE User_name IS NOT NULL AND User_email IS NOT NULL
PRIMARY KEY (User_email, User_name );

CREATE TABLE IF NOT EXISTS chat_messages
(
    message_id text,
    chat_id    text,
    user       text,
    message    text,
    timestamp  text,
    PRIMARY KEY ((chat_id), message_id, timestamp)
) WITH CLUSTERING ORDER BY (message_id DESC, timestamp DESC);