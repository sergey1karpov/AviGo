CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL,
    role VARCHAR(20) DEFAULT 'user'
);