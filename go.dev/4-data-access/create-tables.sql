DROP TABLE IF EXISTS album;

CREATE TABLE album (
    id int AUTO_INCREMENT NOT NULL,
    title varchar(128) NOT NULL,
    artist varchar(255) NOT NULL,
    price DECIMAL(5, 2) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO album (title, artist, price)
    VALUES ('Blue Train', 'John Coltrane', 56.99),
    ('Giant Steps', 'John Coltrane', 63.99),
    ('Jeru', 'Gerry Mulligan', 17.99),
    ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
