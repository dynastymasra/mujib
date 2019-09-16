CREATE TABLE IF NOT EXISTS articles (
    id uuid PRIMARY KEY NOT NULL, -- IDENTIFIER
    title varchar(255) NOT NULL,
    content text NOT NULL,
    author varchar(50) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone NULL DEFAULT NULL
);

