CREATE TABLE proxy
(
    `id`       int          NOT NULL AUTO_INCREMENT,
    `ip`       varchar(15)  NOT NULL,
    `domain`   varchar(255) NOT NULL,
    `port`     int          NOT NULL,
    `login`    varchar(255)                           DEFAULT NULL,
    `password` varchar(1024)                          DEFAULT NULL,
    `alive`    tinyint(1)   NOT NULL,
    `protocol` varchar(10)  NOT NULL                  DEFAULT 'socks5',
    `location` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `active`   tinyint(1)   NOT NULL                  DEFAULT '1',
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4
  COLLATE `utf8mb4_unicode_ci`
  ENGINE = InnoDB;

INSERT INTO proxy (ip, `domain`, port, login, password, alive, protocol, location, active)
VALUES ('109.172.98.1', 'test.http', 63040, 'test', 'test', 1, 'http', 'AR', 1),
       ('77.91.1.1', 'test.pxl', 64250, 'test', 'test', 1, 'socks5', 'NP', 1);
