/*
   Copyright 2020 Adam Chalkley

   https://github.com/atc0005/mysql2sqlite

   Licensed under the MIT License. See LICENSE file in the project root for
   full license information.

   See README.md for deployment path, etc.

   Purpose: Create a database used by mailserver setup: Postfix/Dovecot/MySQL

   WARNING: Requires MySQL >= 5.5.3 due to length of table/field comments

   Notes:

      * View column comments via:
          show full columns from TABLE_NAME;
          show create TABLE_NAME;

      * A comment for a column can be specified with the COMMENT option, up to
        1024 characters long (255 characters before MySQL 5.5.3).

      * A comment for the table, up to 2048 characters long (60 characters
        before MySQL 5.5.3).

      * As of MySQL 5.0.2 'show full tables' will display tables and views
        and show in an additional column the type of each

      * When a view is invoked, the DEFINER and SQL SECURITY clauses determine
        the security context (the account to use for access checking). The
        default is to use the account for the user who executes the CREATE VIEW
        statement.

*/


-- Disabled by default. Remove the comment prefix for the `DROP DATABASE`
-- statement below if you wish to use this in a test environment where
-- retaining the database content does not matter.
--
-- DROP DATABASE IF EXISTS mailserver;

-- https://dev.mysql.com/doc/refman/8.0/en/charset-unicode-utf8mb4.html
CREATE DATABASE mailserver CHARACTER SET utf8mb4;

USE mailserver;


CREATE TABLE `mailserver`.`virtual_domains` (
  `id` int(11) NOT NULL auto_increment,
  `name` varchar(50) NOT NULL,
  `ticket_number` INT UNSIGNED NOT NULL
    COMMENT 'Associated ticket for this entry (i.e., 12345)',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_modified` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comments` text,
  PRIMARY KEY (`id`),
  INDEX `idx_name` (`name`)
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COMMENT='Virtual domains that our servers are responsible for (primarily mail)'
;


CREATE TABLE `mailserver`.`virtual_users` (
  `id` int(11) NOT NULL auto_increment,
  `domain_id` int(11) NOT NULL,
  `password` varchar(106) NOT NULL,
  `email` varchar(100) NOT NULL,
  `enabled` TINYINT(1) NOT NULL DEFAULT 1
    COMMENT 'Controls whether rule is active',
  `ticket_number` INT UNSIGNED NOT NULL
    COMMENT 'Associated ticket for this entry (i.e., 12345)',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_modified` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comments` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  FOREIGN KEY (domain_id) REFERENCES virtual_domains(id) ON DELETE CASCADE,
  INDEX `idx_email` (`email`, `enabled`)
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COMMENT='Virtual Mailbox users'
;


CREATE TABLE `mailserver`.`virtual_aliases` (
  `id` int(11) NOT NULL auto_increment,
  `domain_id` int(11) NOT NULL,
  `source` varchar(100) NOT NULL,
  `destination` varchar(100) NOT NULL,
  `enabled` TINYINT(1) NOT NULL DEFAULT 1
    COMMENT 'Controls whether rule is active',
  `ticket_number` INT UNSIGNED NOT NULL
    COMMENT 'Associated ticket for this entry (i.e., 12345)',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_modified` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comments` text,
  PRIMARY KEY (`id`),
  FOREIGN KEY (domain_id) REFERENCES virtual_domains(id) ON DELETE CASCADE,
  INDEX `idx_source` (`source`, `enabled`),
  INDEX `idx_destination` (`destination`)
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COMMENT='Caution: These aliases can mask real addresses from other domains'
  ;


CREATE TABLE `mailserver`.`local_aliases` (
  `id` int(11) NOT NULL auto_increment,
  `source` varchar(100) NOT NULL,
  `destination` varchar(100) NOT NULL,
  `enabled` TINYINT(1) NOT NULL DEFAULT 1
    COMMENT 'Controls whether rule is active',
  `ticket_number` INT UNSIGNED NOT NULL
    COMMENT 'Associated ticket for this entry (i.e., 12345)',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_modified` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comments` text,
  PRIMARY KEY (`id`),
  INDEX `idx_source` (`source`, `enabled`),
  INDEX `idx_destination` (`destination`(100))
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COMMENT='Replaces local aliases db'
;



/****************************************************************
  Postfix Transport maps

  http://www.postfix.org/postconf.5.html#transport_maps
  http://www.postfix.org/transport.5.html

*****************************************************************/

CREATE TABLE `mailserver`.`transport_maps` (
  `id` int(11) NOT NULL auto_increment,
  `recipient` varchar(100) NOT NULL
    COMMENT 'user+extension, user, domain, subdomain or wildcards',
  `transport` varchar(100) NOT NULL
    COMMENT 'String composed of message delivery transport (built-in or custom) and optional nexthop in the format of transport:nexthop. See master.cf for service names that can be combined with valid nexthop destinations.',
  `enabled` TINYINT(1) NOT NULL DEFAULT 1
    COMMENT 'Controls whether the transport rule is active',
  `ticket_number` INT UNSIGNED NOT NULL
    COMMENT 'Associated ticket for this entry (i.e., 12345)',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_modified` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comments` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`recipient`),
  INDEX `idx_recipient` (`recipient`, `enabled`)
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COMMENT='Mappings from recipient address to (message delivery transport, next-hop destination).'
;



/****************************************************************
  Postfix Access Tables

  http://www.postfix.org/access.5.html
  http://www.postfix.org/postconf.5.html#check_client_access
  http://www.postfix.org/postconf.5.html#check_recipient_access
  http://www.postfix.org/postconf.5.html#check_sender_access

*****************************************************************/

CREATE TABLE `mailserver`.`access_check_clients` (
  `id` int(11) NOT NULL auto_increment,
  `client` varchar(100) NOT NULL
    COMMENT 'client hostname, parent domains, IP address, or networks',
  `action` varchar(100) NOT NULL
    COMMENT 'Action to take against email for matching clients',
  `enabled` TINYINT(1) NOT NULL DEFAULT 1
    COMMENT 'Controls whether rule is active',
  `ticket_number` INT UNSIGNED NOT NULL
    COMMENT 'Associated ticket for this entry (i.e., 12345)',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_modified` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comments` text,
  PRIMARY KEY (`id`),
  INDEX `idx_client` (`client`, `enabled`)
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COMMENT='Explicitly reject or allow clients (which includes other mail servers)'
;


CREATE TABLE `mailserver`.`access_check_recipients` (
  `id` int(11) NOT NULL auto_increment,
  `recipient` varchar(100) NOT NULL
    COMMENT 'the resolved RCPT TO address, domain, parent domains, or localpart@',
  `action` varchar(100) NOT NULL
    COMMENT 'Action to take against email for matching recipients',
  `enabled` TINYINT(1) NOT NULL DEFAULT 1
    COMMENT 'Controls whether rule is active',
  `ticket_number` INT UNSIGNED NOT NULL
    COMMENT 'Associated ticket for this entry (i.e., 12345)',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_modified` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comments` text,
  PRIMARY KEY (`id`),
  INDEX `idx_recipient` (`recipient`, `enabled`)
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COMMENT='Explicitly reject, hold or permit emails for specific addresses'
;


CREATE TABLE `mailserver`.`access_check_senders` (
  `id` int(11) NOT NULL auto_increment,
  `sender` varchar(100) NOT NULL
    COMMENT 'the MAIL FROM address, domain, parent domains, or localpart@',
  `action` varchar(100) NOT NULL
    COMMENT 'Action to take against email for matching senders',
  `enabled` TINYINT(1) NOT NULL DEFAULT 1
    COMMENT 'Controls whether rule is active',
  `ticket_number` INT UNSIGNED NOT NULL
    COMMENT 'Associated ticket for this entry (i.e., 12345)',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_modified` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comments` text,
  PRIMARY KEY (`id`),
  INDEX `idx_sender` (`sender`, `enabled`)
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COMMENT='Explicitly reject or allow senders (workaround bad mail server setups, etc)'
;


CREATE TABLE `mailserver`.`mail_relay_whitelist` (
  `id` int(11) NOT NULL auto_increment,
  `client` varchar(100) NOT NULL
    COMMENT 'client hostname, parent domains, IP address, or networks',
  `action` varchar(100) NOT NULL
    COMMENT 'Action to take against client attempting to relay mail',
  `enabled` TINYINT(1) NOT NULL DEFAULT 1
    COMMENT 'Controls whether rule is active',
  `ticket_number` INT UNSIGNED NOT NULL
    COMMENT 'Associated ticket for this entry (i.e., 12345)',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_modified` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comments` text,
  PRIMARY KEY (`id`),
  INDEX `idx_client` (`client`, `enabled`)
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COMMENT='Explicitly reject or allow mail relay access to clients (which includes other mail servers)'
;


CREATE TABLE `mailserver`.`sender_bcc_maps` (
  `id` int(11) NOT NULL auto_increment,
  `sender` varchar(100) NOT NULL
    COMMENT 'Source email address whose mail should be blind carbon-copied elsewhere',
  `bcc_recipient` varchar(100) NOT NULL
    COMMENT 'Recipient of mail that was blind carbon-copied for source address',
  `enabled` TINYINT(1) NOT NULL DEFAULT 1
    COMMENT 'Controls whether rule is active',
  `ticket_number` INT UNSIGNED NOT NULL
    COMMENT 'Associated ticket for this entry (i.e., 12345)',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_modified` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comments` text,
  PRIMARY KEY (`id`),
  INDEX `idx_sender` (`sender`, `enabled`),
  INDEX `idx_bcc_recipient` (`bcc_recipient`)
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COMMENT='Blind carbon-copy address lookup table, intended to be used by sender_bcc_maps Postfix directive'
;


CREATE TABLE `mailserver`.`recipient_bcc_maps` (
  `id` int(11) NOT NULL auto_increment,
  `original_recipient` varchar(100) NOT NULL
    COMMENT 'Original recipient email address whose mail should be blind carbon-copied elsewhere',
  `bcc_recipient` varchar(100) NOT NULL
     COMMENT 'Recipient of mail that was blind carbon-copied for source address',
  `enabled` TINYINT(1) NOT NULL DEFAULT 1
     COMMENT 'Controls whether rule is active',
  `ticket_number` INT UNSIGNED NOT NULL
     COMMENT 'Associated ticket for this entry (i.e., 12345)',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_modified` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comments` text,
  PRIMARY KEY (`id`),
  INDEX `idx_original_recipient` (`original_recipient`, `enabled`),
  INDEX `idx_bcc_recipient` (`bcc_recipient`)
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COMMENT='Blind carbon-copy address lookup table, intended to be used by sender_bcc_maps Postfix directive'
;


CREATE TABLE `mailserver`.`sender_dependent_default_transport_maps` (
  `id` int(11) NOT NULL auto_increment,
  `sender` varchar(100) NOT NULL
    COMMENT 'the MAIL FROM address, domain, parent domains, or localpart@',
  `transport` varchar(100) NOT NULL
    COMMENT 'The service name used to transport mail',
  `enabled` TINYINT(1) NOT NULL DEFAULT 1
    COMMENT 'Controls whether rule is active',
  `ticket_number` INT UNSIGNED NOT NULL
    COMMENT 'Associated ticket for this entry (i.e., 12345)',
  `created` timestamp DEFAULT CURRENT_TIMESTAMP,
  `last_modified` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comments` text,
  PRIMARY KEY (`id`),
  INDEX `idx_sender` (`sender`, `enabled`)
)
  ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COMMENT='Explicitly specify the service used to transport mail based on the sender address.'
;
