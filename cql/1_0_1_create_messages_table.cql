CREATE TABLE IF NOT EXISTS chat_messages
(
    message_id text,
    user       text,
    message    text,
    timestamp  timestamp,
    PRIMARY KEY ((user), message_id, timestamp)
) WITH CLUSTERING ORDER BY (timestamp DESC);