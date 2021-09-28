CREATE TABLE IF NOT EXISTS user
(
    user_id         varchar,
    username        varchar,
    email           varchar,
    hashed_password varchar,
    PRIMARY KEY ((username), user_id)
);