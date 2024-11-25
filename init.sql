-- init.sql
DROP TABLE IF EXISTS pengguna;

CREATE TABLE pengguna (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    username VARCHAR(50) UNIQUE,
    password VARCHAR(255)
);

-- Pre-hashed password for "password123" using bcrypt
INSERT INTO pengguna (name, email, username, password) VALUES
('Alice', 'alice@example.com', 'alice123', '$2a$10$7s0N5K2c5/JGFmM5IZPEe.q9mTCF3J1KoV5zPzA6V3I8e1Uy/2s2a');