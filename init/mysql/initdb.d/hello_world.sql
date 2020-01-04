CREATE DATABASE IF NOT EXISTS images CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE images;
-- https://fra1.cdn.digitaloceanspaces.com/simo.
CREATE TABLE buckets (
    id                  int(10) NOT NULL AUTO_INCREMENT,
    name                UNIQUE  VARCHAR(50) NOT NULL,   -- bucket_name
    region              VARCHAR(10) NOT NULL,           -- region.digitaloceanspaces.com --> fra1.digitaloceanspaces.com
    endpoint            VARCHAR(20) NOT NULL            -- digitaloceanspaces.com
    location            VARCHAR(12)  NULL               -- us-east-1
    is_cdn              BOOLEAN      DEFAULT false      -- fra1.cdn.digitaloceanspaces.com --> fra1.cdn.digitaloceanspaces.com
    cdn_custom_domain   VARCHAR(20)  NULL               -- default 'null'
    PRIMARY KEY (id)
);

CREATE TABLE files (
    id int(10) NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    mime_type VARCHAR(255),
    bucket_id int(10) NOT NULL, -- many-to-one
    FOREIGN KEY (bucket_id) REFERENCES buckets (id),
    PRIMARY KEY (id)
);