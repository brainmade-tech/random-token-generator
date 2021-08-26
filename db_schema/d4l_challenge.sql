/* Author : Yassine Chaoui */

CREATE DATABASE d4l_challenge;

USE d4l_challenge;

/*DROP TABLE IF EXISTS `tokens`;*/

CREATE TABLE tokens (
  id int(11) NOT NULL AUTO_INCREMENT,
  value varchar(45) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=81088380 DEFAULT CHARSET=latin1;