BEGIN
CREATE TABLE chats (
    id serial PRIMARY KEY
);

CREATE TABLE messages (
    id serial PRIMARY KEY,
    chat_id INT NOT NULL,
    FOREIGN KEY (chat_id)
        REFERENCES chats (id)
)
COMMIT