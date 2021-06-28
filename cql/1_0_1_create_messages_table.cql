CREATE TABLE IF NOT EXISTS chat_messages
(
    message_id text,
    chat_id    text,
    user_id    text,
    message    text,
    timestamp  text,
    PRIMARY KEY ((chat_id), user_id, message_id, timestamp)
) WITH CLUSTERING ORDER BY (user_id ASC, message_id ASC, timestamp DESC);


CREATE MATERIALIZED VIEW chats AS
SELECT message_id, user_id, message_id, timestamp, chat_id
FROM chat_messages
WHERE user_id IS NOT NULL
  AND chat_id IS NOT NULL
  AND message_id IS NOT NULL
  AND timestamp IS NOT NULL
PRIMARY KEY ((user_id), chat_id, message_id, timestamp)
        WITH CLUSTERING ORDER BY (chat_id ASC, message_id ASC, timestamp DESC);


// Get all chats which user is a member of
SELECT * FROM chats WHERE user_id='ross123';
// Get all messages in chat
SELECT * FROM chat_messages WHERE chat_id='c1242';