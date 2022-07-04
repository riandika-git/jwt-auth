/*
SQLyog Community v12.2.0 (64 bit)
MySQL - 5.7.11-log : Database - test_reko
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`test_reko` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `test_reko`;

/*Table structure for table `customer` */

DROP TABLE IF EXISTS `customer`;

CREATE TABLE `customer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(100) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `point` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

/*Data for the table `customer` */

insert  into `customer`(`id`,`username`,`password`,`name`,`email`,`phone`,`address`,`point`) values 
(1,'user01','482c811da5d5b4bc6d497ffa98491e38','Johne Depp','johnyjohny@test.com','081111111','Jl. Sukamaju No. 110 Bandung',100),
(2,'user02','482c811da5d5b4bc6d497ffa98491e38','Michael Jordan','michael@test.com','082222222','Baker Street No. 25',50),
(3,'user03','482c811da5d5b4bc6d497ffa98491e38','Black Widow','widowmania@test.com','083333333','Atap Gedung Bertingkat No. 99',80);

/*Table structure for table `voucher_group` */

DROP TABLE IF EXISTS `voucher_group`;

CREATE TABLE `voucher_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `voucher_group_name` varchar(200) NOT NULL,
  `qty` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

/*Data for the table `voucher_group` */

insert  into `voucher_group`(`id`,`voucher_group_name`,`qty`) values 
(1,'Bli Bli',1000),
(2,'Traveloka',2000),
(3,'Wilopo',500);

/*Table structure for table `voucher_purchase` */

DROP TABLE IF EXISTS `voucher_purchase`;

CREATE TABLE `voucher_purchase` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) NOT NULL,
  `voucher_group_id` int(11) NOT NULL,
  `voucher_code` varchar(255) DEFAULT NULL,
  `purchase_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
