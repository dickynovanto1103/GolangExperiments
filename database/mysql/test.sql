DROP DATABASE IF EXISTS testDB;
CREATE DATABASE testDB;

USE testDB;

CREATE TABLE Person (
    ID INT NOT NULL AUTO_INCREMENT,
    Name VARCHAR(100),
    Age INT,
    PRIMARY KEY(ID)
);

INSERT INTO Person(Name, Age) VALUES('dicky', 22);
INSERT INTO Person(Name, Age) VALUES('angel', 20);