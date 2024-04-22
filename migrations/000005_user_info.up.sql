CREATE TABLE user_info (
    id BIGSERIAL PRIMARY KEY ,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOT(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOT(),
    user_name varchar(255),
    user_surname varchar(255),
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    role varchar(50),
    activated bool NOT NULL,
    version INTEGER NOT NULL DEFAULT 1
);