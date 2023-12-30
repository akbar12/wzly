-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: localhost    Database: point_of_sales
-- ------------------------------------------------------
-- Server version	5.7.44-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `customer`
--

DROP TABLE IF EXISTS `customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_no` varchar(45) DEFAULT NULL,
  `customer_name` varchar(100) DEFAULT NULL,
  `detail_address` varchar(255) DEFAULT NULL,
  `created_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(45) DEFAULT NULL,
  `modified_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `modified_by` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `customer_no_UNIQUE` (`customer_no`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer`
--

LOCK TABLES `customer` WRITE;
/*!40000 ALTER TABLE `customer` DISABLE KEYS */;
INSERT INTO `customer` VALUES (1,'00001','Budi Susanto','Jl Medan Merdeka Selatan No.2','2023-12-15 22:30:50','admin','2023-12-15 22:30:50','admin'),(2,'00002','Arik Susanto','Jl Aris Munandar No.1','2023-12-15 22:30:50','admin','2023-12-15 22:30:50','admin'),(3,'00003','Julianto','Jl Letjen Sutoyo No.34','2023-12-15 22:30:50','admin','2023-12-15 22:30:50','admin');
/*!40000 ALTER TABLE `customer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoice`
--

DROP TABLE IF EXISTS `invoice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `invoice` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `invoice_no` int(8) unsigned zerofill DEFAULT NULL,
  `issued_date` datetime DEFAULT NULL,
  `due_date` datetime DEFAULT NULL,
  `status` int(1) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `subject` varchar(100) DEFAULT NULL,
  `total_item` int(11) DEFAULT NULL,
  `sub_total` decimal(11,2) DEFAULT NULL,
  `tax` decimal(11,2) DEFAULT NULL,
  `grand_total` decimal(11,2) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `created_by` varchar(45) DEFAULT NULL,
  `modified_date` datetime DEFAULT NULL,
  `modified_by` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `invoice_no_UNIQUE` (`invoice_no`),
  KEY `status_INDEX` (`status`),
  KEY `total_item_INDEX` (`total_item`),
  KEY `customer_FK_idx` (`customer_id`),
  CONSTRAINT `customer_FK` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoice`
--

LOCK TABLES `invoice` WRITE;
/*!40000 ALTER TABLE `invoice` DISABLE KEYS */;
INSERT INTO `invoice` VALUES (17,00000017,'2023-12-16 00:00:00','2023-12-16 00:00:00',0,2,'Food for Arik Susanto',0,0.00,0.00,0.00,'2023-12-16 01:05:44','JohnDoe','2023-12-16 01:05:44','JohnDoe'),(18,00000018,'2023-12-16 00:00:00','2023-12-16 00:00:00',0,2,'Food for Arik Susanto',0,0.00,4.60,50.62,'2023-12-16 01:09:44','JohnDoe','2023-12-16 01:09:44','JohnDoe'),(19,00000019,'2023-12-22 00:00:00','2023-12-22 00:00:00',1,3,'Food for Julianto',2,529.05,52.91,581.96,'2023-12-16 01:12:37','JohnDoe','2023-12-16 06:04:48','JohnDoe'),(20,00000020,'2023-12-16 00:00:00','2023-12-16 00:00:00',1,2,'Food for Arik Susanto',5,46.02,4.60,50.62,'2023-12-16 01:16:24','JohnDoe','2023-12-16 01:16:24','JohnDoe'),(27,00000027,'2023-12-01 00:00:00','2023-12-01 00:00:00',1,1,'Food for Budi 123',2,105.81,10.58,116.39,'2023-12-16 07:16:57','JohnDoe','2023-12-16 07:21:00','JohnDoe');
/*!40000 ALTER TABLE `invoice` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoice_item`
--

DROP TABLE IF EXISTS `invoice_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `invoice_item` (
  `invoice_item_id` int(11) NOT NULL AUTO_INCREMENT,
  `invoice_id` int(11) DEFAULT NULL,
  `item_id` int(11) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `unit_price` decimal(11,2) DEFAULT NULL,
  `amount` decimal(11,2) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `created_by` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`invoice_item_id`),
  UNIQUE KEY `invoice_item_id_UNIQUE` (`invoice_item_id`),
  KEY `invoice_FK_idx` (`invoice_id`),
  KEY `item_FK_idx` (`item_id`),
  CONSTRAINT `invoice_FK` FOREIGN KEY (`invoice_id`) REFERENCES `invoice` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `item_FK` FOREIGN KEY (`item_id`) REFERENCES `item` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoice_item`
--

LOCK TABLES `invoice_item` WRITE;
/*!40000 ALTER TABLE `invoice_item` DISABLE KEYS */;
INSERT INTO `invoice_item` VALUES (7,17,5,2,5.76,11.52,'2023-12-16 01:05:45','John Doe'),(8,17,1,3,11.50,34.50,'2023-12-16 01:05:45','JohnDoe'),(9,18,5,2,5.76,11.52,'2023-12-16 01:09:44','JohnDoe'),(10,18,1,3,11.50,34.50,'2023-12-16 01:09:44','JohnDoe'),(13,20,5,2,5.76,11.52,'2023-12-16 01:16:24','JohnDoe'),(14,20,1,3,11.50,34.50,'2023-12-16 01:16:24','JohnDoe'),(35,19,5,5,5.76,28.80,'2023-12-16 06:04:48','JohnDoe'),(36,19,4,5,100.05,500.25,'2023-12-16 06:04:48','JohnDoe'),(39,27,5,1,5.76,5.76,'2023-12-16 07:21:00','JohnDoe'),(40,27,4,1,100.05,100.05,'2023-12-16 07:21:00','JohnDoe');
/*!40000 ALTER TABLE `invoice_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item`
--

DROP TABLE IF EXISTS `item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item_no` varchar(45) COLLATE utf8_swedish_ci DEFAULT NULL,
  `item_name` varchar(100) COLLATE utf8_swedish_ci DEFAULT NULL,
  `item_type` varchar(25) COLLATE utf8_swedish_ci DEFAULT NULL,
  `unit_price` decimal(11,2) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `created_by` varchar(45) COLLATE utf8_swedish_ci DEFAULT NULL,
  `modified_date` datetime DEFAULT NULL,
  `modified_by` varchar(45) COLLATE utf8_swedish_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `item_no_UNIQUE` (`item_no`),
  KEY `item_type_INDEX` (`item_type`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COLLATE=utf8_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item`
--

LOCK TABLES `item` WRITE;
/*!40000 ALTER TABLE `item` DISABLE KEYS */;
INSERT INTO `item` VALUES (1,'001','Milk','food',11.50,'2023-12-15 22:30:50','admin','2023-12-15 22:30:50','admin'),(2,'002','Egg','food',25.25,'2023-12-15 22:30:50','admin','2023-12-15 22:30:50','admin'),(3,'003','Broom','houseware',35.00,'2023-12-15 22:30:50','admin','2023-12-15 22:30:50','admin'),(4,'004','Beef','food',100.05,'2023-12-15 22:30:50','admin','2023-12-15 22:30:50','admin'),(5,'005','Banana','fruit',5.76,'2023-12-15 22:30:50','admin','2023-12-15 22:30:50','admin');
/*!40000 ALTER TABLE `item` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-16 14:29:01
