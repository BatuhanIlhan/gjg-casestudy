CREATE TABLE IF NOT EXISTS "user"
(
    id              UUID PRIMARY KEY,
    points          DOUBLE PRECISION DEFAULT 0.0,
    country_code    VARCHAR(2),
    display_name    VARCHAR(100) NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at      TIMESTAMP WITH TIME ZONE NULL
);