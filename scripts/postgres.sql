CREATE TABLE delegation
(
    id SERIAL PRIMARY KEY,
    delegator VARCHAR(36),
    amount VARCHAR(64),
    block_high VARCHAR(10),
    timestamp TIMESTAMP
);
