DROP TABLE IF EXISTS country;
CREATE TABLE country (
    `uid` integer primary key autoincrement,
    `name` varchar(260) not null,
    `abbr_name` varchar(10) null
);

DROP TABLE IF EXISTS person;
CREATE TABLE person (
    `uid` integer primary key autoincrement,
    `name` varchar(64) not null,
    `created` date default CURRENT_DATE,
    `country` integer not null,
    FOREIGN KEY (`country`) REFERENCES country(uid)
);