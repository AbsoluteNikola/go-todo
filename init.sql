CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    item TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE
);
