CREATE TABLE chats
(
    chat_id   uuid,
    from_user text,
    to_user   text,
    timestamp timestamp,
    PRIMARY KEY ((from_user, to_user), timestamp)
) WITH CLUSTERING ORDER BY (timestamp DESC);