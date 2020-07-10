BEGIN;
CREATE TABLE IF NOT EXISTS users
(
    user_id uuid DEFAULT uuid_generate_v4(),
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL
);
COMMIT;