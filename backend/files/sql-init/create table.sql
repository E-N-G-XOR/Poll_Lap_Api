USE appdb;

CREATE TABLE apptable (
id int NOT NULL AUTO_INCREMENT,
name varchar(25),
timestamps  varchar(25),
updatestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
createstamp DATETIME DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (id)
);

/*  */