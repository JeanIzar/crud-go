/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `compras_rf`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `compras_rf` (
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `idcompra` int NOT NULL,
  `compracod` int NOT NULL,
  `codigo` varchar(45) NOT NULL,
  `compra_fecha_emi` date NOT NULL,
  `compra_serie` varchar(45) NOT NULL,
  `compra_numer` varchar(45) NOT NULL,
  `compra_descuent` text NOT NULL,
  `compra_total_item` int DEFAULT NULL,
  `compra_tipo_cambio` decimal(20,3) DEFAULT NULL,
  `compra_vvb` decimal(20,2) DEFAULT NULL,
  `compra_oper_g` decimal(20,2) DEFAULT NULL,
  `compra_isc` decimal(20,2) DEFAULT NULL,
  `compra_vv` decimal(20,2) DEFAULT NULL,
  `compra_igv` decimal(20,2) DEFAULT NULL,
  `compra_pv` decimal(20,2) DEFAULT NULL,
  `compra_cargo` decimal(20,2) DEFAULT NULL,
  `compra_tot_v` decimal(20,2) DEFAULT NULL,
  `compra_oper_a` decimal(20,2) DEFAULT NULL,
  `compra_oper_i` decimal(20,2) DEFAULT NULL,
  `compra_oper_e` decimal(20,2) DEFAULT NULL,
  `id_estado` int NOT NULL,
  `id_proveedor` int NOT NULL,
  `txt_tipodocumento` varchar(15) NOT NULL,
  `txt_tipocomprobante` varchar(12) NOT NULL,
  `id_periodo` int NOT NULL,
  `txt_moneda` varchar(10) NOT NULL,
  `id_empresa` int NOT NULL,
  `id_centro_costo` int DEFAULT NULL,
  `compra_cant_asiento` int DEFAULT NULL,
  `igv` decimal(4,2) NOT NULL,
  `prov_ruc` varchar(11) DEFAULT NULL,
  `prov_nombre` varchar(150) DEFAULT NULL,
  `prov_direccion` varchar(15) DEFAULT NULL,
  `prov_contacto` varchar(50) DEFAULT NULL,
  `prov_telefono` varchar(20) DEFAULT NULL,
  `prov_correo` varchar(50) DEFAULT NULL,
  `fech_reg` datetime NOT NULL,
  `fech_act` datetime NOT NULL,
  `cod_usua_reg` int NOT NULL,
  `cod_usua_act` int NOT NULL,
  PRIMARY KEY (`idcompra`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `schema_migrations`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schema_migrations` (
  `version` varchar(128) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping routines for database 'emma_compras'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed

--
-- Dbmate schema migrations
--

LOCK TABLES `schema_migrations` WRITE;
INSERT INTO `schema_migrations` (version) VALUES
  ('20231220162116');
UNLOCK TABLES;
