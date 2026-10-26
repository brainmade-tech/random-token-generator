/* Author : Yassine Chaoui */

CREATE DATABASE token_db;

USE token_db;

/*DROP TABLE IF EXISTS `tokens`;*/

CREATE TABLE tokens (
  id int(11) NOT NULL AUTO_INCREMENT,
  value varchar(45) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=81088380 DEFAULT CHARSET=latin1;
