-- Create upload table
CREATE TABLE upload (
    id           BIGINT UNSIGNED AUTO_INCREMENT
        PRIMARY KEY,
    filename     VARCHAR(512)                       NOT NULL,
    content_type VARCHAR(255)                       NOT NULL,
    size         BIGINT                             NOT NULL,
    upload_time  DATETIME DEFAULT CURRENT_TIMESTAMP NULL,
    user         VARCHAR(255)                       NOT NULL,
    user_agent   VARCHAR(512)                       NOT NULL,
    created_at   DATETIME DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at   DATETIME DEFAULT CURRENT_TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
);

-- Create user table
CREATE TABLE user (
    id         BIGINT UNSIGNED AUTO_INCREMENT
        PRIMARY KEY,
    username   VARCHAR(255)                       NOT NULL,
    password   VARCHAR(255)                       NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT username
        UNIQUE (username)
); 