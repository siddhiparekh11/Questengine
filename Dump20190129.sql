CREATE DATABASE  IF NOT EXISTS `QuestEngine` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `QuestEngine`;
-- MySQL dump 10.13  Distrib 5.6.17, for osx10.6 (i386)
--
-- Host: localhost    Database: QuestEngine
-- ------------------------------------------------------
-- Server version	5.7.16

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Games`
--

DROP TABLE IF EXISTS `Games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Games` (
  `idGames` int(11) NOT NULL AUTO_INCREMENT,
  `questOrder` varchar(45) NOT NULL,
  `nameGame` varchar(45) NOT NULL,
  PRIMARY KEY (`idGames`)
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Games`
--

LOCK TABLES `Games` WRITE;
/*!40000 ALTER TABLE `Games` DISABLE KEYS */;
INSERT INTO `Games` VALUES (101,'11,12,13','BlackJack'),(103,'12,13','las vegas'),(104,'13,14','pokemon'),(105,'13,14','Treasure hunt'),(107,'11,12,13','HelloKitty');
/*!40000 ALTER TABLE `Games` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Players`
--

DROP TABLE IF EXISTS `Players`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Players` (
  `idPlayers` int(11) NOT NULL AUTO_INCREMENT,
  `namePlayer` varchar(45) NOT NULL,
  `chipsAmount` int(11) NOT NULL,
  `totalNoChips` varchar(45) NOT NULL,
  PRIMARY KEY (`idPlayers`)
) ENGINE=InnoDB AUTO_INCREMENT=1008 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Players`
--

LOCK TABLES `Players` WRITE;
/*!40000 ALTER TABLE `Players` DISABLE KEYS */;
INSERT INTO `Players` VALUES (1001,'John',400,'4'),(1002,'Tom',200,'2'),(1006,'Ish',300,'3');
/*!40000 ALTER TABLE `Players` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Progress`
--

DROP TABLE IF EXISTS `Progress`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Progress` (
  `idProgress` int(11) NOT NULL AUTO_INCREMENT,
  `idGame` int(11) NOT NULL,
  `idQuest` int(11) NOT NULL,
  `idPlayer` int(11) NOT NULL,
  `questPointsEarned` int(11) NOT NULL,
  `totalQuestComPercent` int(11) NOT NULL,
  `lastMilestoneInd` int(11) NOT NULL,
  `createdTimestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`idProgress`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Progress`
--

LOCK TABLES `Progress` WRITE;
/*!40000 ALTER TABLE `Progress` DISABLE KEYS */;
INSERT INTO `Progress` VALUES (1,101,11,1001,203,100,1,'2019-01-24 11:14:07'),(2,101,12,1001,28,50,2,'2019-01-25 02:09:03'),(5,101,12,1001,56,100,3,'2019-01-25 08:30:03');
/*!40000 ALTER TABLE `Progress` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Quests`
--

DROP TABLE IF EXISTS `Quests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Quests` (
  `idQuests` int(11) NOT NULL AUTO_INCREMENT,
  `questName` varchar(45) NOT NULL,
  `milestonesOrder` varchar(45) NOT NULL,
  PRIMARY KEY (`idQuests`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Quests`
--

LOCK TABLES `Quests` WRITE;
/*!40000 ALTER TABLE `Quests` DISABLE KEYS */;
INSERT INTO `Quests` VALUES (11,'level1','1'),(12,'level2','2,3'),(13,'level3','4,5'),(14,'level4','4,5'),(15,'level8','2,3'),(16,'level9','1,4,5');
/*!40000 ALTER TABLE `Quests` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-01-29 17:32:40
