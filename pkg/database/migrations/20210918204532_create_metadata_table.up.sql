create table files_metadata(
    id BIGSERIAL NOT NULL,
    file_id BIGINT NOT NULL,
    key CHARACTER VARYING(255) NOT NULL,
    value CHARACTER VARYING(2048) NOT NULL,
    CONSTRAINT file_metadata_pkey PRIMARY KEY (id),
    CONSTRAINT files_fkey FOREIGN KEY (file_id) REFERENCES files (id) ON DELETE CASCADE ON UPDATE CASCADE
);
