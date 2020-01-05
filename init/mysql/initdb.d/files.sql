CREATE DATABASE IF NOT EXISTS files CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE files;

CREATE TABLE buckets (
    id                  int(10)      NOT NULL AUTO_INCREMENT,
    name                VARCHAR(63)  UNIQUE NOT NULL,    -- bucket_name
    region              VARCHAR(10)  NOT NULL,           -- region.digitaloceanspaces.com --> fra1.digitaloceanspaces.com
    endpoint            VARCHAR(20)  NOT NULL,           -- digitaloceanspaces.com
    location            VARCHAR(12)  NULL,               -- us-east-1
    is_cdn              BOOLEAN      DEFAULT false,      -- fra1.digitaloceanspaces.com --> fra1.cdn.digitaloceanspaces.com
    PRIMARY KEY (id)
);

CREATE TABLE files (
    id          int(10)      NOT NULL AUTO_INCREMENT,
    filename     VARCHAR(100) NOT NULL,
    mime_type   VARCHAR(255) NOT NULL,
    file_size    int(10)      NOT NULL,
    bucket_id   int(10)      NULL,
    is_public   BOOLEAN      DEFAULT false,
    FOREIGN KEY (bucket_id)  REFERENCES buckets (id),
    PRIMARY KEY (id)
);