CREATE TABLE `fizzbuzz`
(
    `id`         bigint PRIMARY KEY AUTO_INCREMENT,
    `useragent`  varchar(255) NOT NULL,
    `message`    varchar(255) NOT NULL,
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP
);