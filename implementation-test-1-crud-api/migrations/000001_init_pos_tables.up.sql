CREATE TABLE IF NOT EXISTS `customer` (
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

CREATE TABLE IF NOT EXISTS `invoice` (
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

CREATE TABLE IF NOT EXISTS `item` (
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

CREATE TABLE IF NOT EXISTS `invoice_item` (
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
