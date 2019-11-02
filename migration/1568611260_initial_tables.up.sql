CREATE TABLE IF NOT EXISTS products (
    id bigint PRIMARY KEY NOT NULL, -- IDENTIFIER
    name varchar(255) NOT NULL,
    image_closed varchar(255) NOT NULL,
    image_open varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    story text NOT NULL,
    sourcing_values text NOT NULL,
    ingredients text NOT NULL,
    allergy_info varchar(255) NULL DEFAULT NULL,
    dietary_certifications varchar(50) NULL DEFAULT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone NULL DEFAULT NULL
);