CREATE TABLE users(
    id BIGSERIAL,
    token CHARACTER VARYING(255) NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

