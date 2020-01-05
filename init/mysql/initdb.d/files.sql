CREATE DATABASE IF NOT EXISTS files CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE files;

CREATE TABLE do_buckets (
    id                  int(10)      NOT NULL AUTO_INCREMENT, -- ent requires..
    bucket_name         VARCHAR(63)  UNIQUE NOT NULL,         -- unique name for the bucket
    endpoint            VARCHAR(35)  NOT NULL,                -- fra1.digitaloceanspaces.com
    cdn_endpoint        VARCHAR(100) NULL,                    -- {bucket_name}.{region}.cdn.digitaloceanspaces.com
    created_at          DATETIME(3)  NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE files (
    id             int(10)      NOT NULL AUTO_INCREMENT,
    filename        VARCHAR(100) NOT NULL,
    mime_type      VARCHAR(255) NOT NULL,
    file_size       int(10)      NOT NULL,
    do_bucket_id   int(10)      NULL,
    is_draft       BOOLEAN      DEFAULT false,
    is_deleted     BOOLEAN      DEFAULT false,
    FOREIGN KEY (do_bucket_id)  REFERENCES do_buckets (id),
    PRIMARY KEY (id)
);