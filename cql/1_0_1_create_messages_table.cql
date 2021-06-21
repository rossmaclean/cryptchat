CREATE TABLE IF NOT EXISTS chat_messages
(
    message_id uuid,
    from_user  text,
    to_user    text,
    message    text,
    timestamp  timestamp,
    PRIMARY KEY ((from_user, to_user), timestamp)
) WITH CLUSTERING ORDER BY (timestamp DESC);