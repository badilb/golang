CREATE TABLE teachers_info (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    module_id INTEGER REFERENCES module_info(id)
);
