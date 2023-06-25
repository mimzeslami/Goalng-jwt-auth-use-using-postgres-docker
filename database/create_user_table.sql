CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public."user" (
	id uuid NOT null DEFAULT uuid_generate_v4() PRIMARY KEY,
    email varchar(255) NOT null UNIQUE,
    first_name varchar(255) NOT null,
    last_name varchar(255) NOT null,
    password varchar(255) NOT null,
    updated_at timestamp NOT null DEFAULT now(),
    created_at timestamp NOT null DEFAULT now()
    deleted_at timestamp
    is_deleted boolean NOT null DEFAULT false
);
