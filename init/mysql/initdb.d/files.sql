CREATE DATABASE IF NOT EXISTS files CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE files;

-- one-to-many `files`
CREATE TABLE buckets (
    id                  BIGINT      NOT NULL AUTO_INCREMENT, -- ent requires..
    name                VARCHAR(63)  UNIQUE NOT NULL,         -- unique name for the bucket
    endpoint            VARCHAR(35)  NOT NULL,                -- fra1.digitaloceanspaces.com
    cdn_endpoint        VARCHAR(100) NULL,                    -- {bucket_name}.{region}.cdn.digitaloceanspaces.com
    created_at          DATETIME(3)  NOT NULL,
    PRIMARY KEY (id)
);

-- many-to-one `buckets`
CREATE TABLE files (
    id             BIGINT       NOT NULL AUTO_INCREMENT,
    filename        VARCHAR(100) NOT NULL,                     -- 'hello_world.txt'
    mime_type      VARCHAR(255) NOT NULL,                     -- 'plain/text'
    file_size       BIGINT       NOT NULL,                     -- '12'
    bucket_files    BIGINT       NULL,
    is_draft       BOOLEAN      DEFAULT false,                -- only visible for the self
    is_deleted     BOOLEAN      DEFAULT false,                -- 'deleted' from the self and anyone
    user_id        BIGINT       NOT NULL,                     -- userId is referenced from users-service (from the jwt-token)
    created_at     DATETIME(3)  NOT NULL,
    updated_at     DATETIME(3)  NOT NULL,
    FOREIGN KEY (bucket_files)   REFERENCES buckets (id),
    PRIMARY KEY (id)
);