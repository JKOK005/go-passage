# noinspection SqlNoDataSourceInspectionForFile
CREATE SCHEMA IF NOT EXISTS Passage_master;

# Passage Server tables. For local dev, user: passage / pass: passage
CREATE TABLE IF NOT EXISTS Passage_master.servers (
    id smallint unsigned not null auto_increment,
    url varchar(20) not null,
    port smallint unsigned not null,

    UNIQUE (url, port),
    PRIMARY KEY (id),
    INDEX (url, port)
) ENGINE = INNODB;

CREATE TABLE IF NOT EXISTS Passage_master.apps (
    id smallint unsigned not null auto_increment,
    serverID smallint unsigned not null,
    name varchar(20) not null,

    UNIQUE (serverID, name),
    PRIMARY KEY (id),
    FOREIGN KEY (serverID) REFERENCES Passage_master.servers(id) ON DELETE CASCADE,
    INDEX (name)
) ENGINE = INNODB;