BEGIN;

CREATE TABLE IF NOT EXISTS chats (
    id serial PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS messages (
    id serial PRIMARY KEY,
    chat_id INT NOT NULL,
    FOREIGN KEY (chat_id)
        REFERENCES chats (id)
);

COMMIT;