--- create table
CREATE TABLE data (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	value TEXT
);

--- insert demo data
INSERT INTO data("value")
	VALUES("hello mom!");

INSERT INTO data("value")
	VALUES("Lorem ipsum dolor sit amet");

INSERT INTO data("value")
	VALUES("Programing is fun");
