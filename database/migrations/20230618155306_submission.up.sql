CREATE TABLE IF NOT EXISTS submission
(
    id                    UUID              PRIMARY KEY,
    user_id               UUID              NOT NULL,
    score                 DOUBLE PRECISION            NOT NULL,
    created_at            TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at            TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at            TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT fk_user    FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE
    );