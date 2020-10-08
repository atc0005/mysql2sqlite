/*

   Copyright 2020 Adam Chalkley

   https://github.com/atc0005/mysql2sqlite

   Licensed under the MIT License. See LICENSE file in the project root for
   full license information.

   See README.md for deployment path, etc.

*/

-- Notes:
--
--  * Skipping insertion for columns with 'default' values
--  * Intentionally inserting table entries with leading
--    and trailing spaces to force applications to deal with
--    problematic data.

INSERT INTO `mailserver`.`virtual_domains`
  (`id`, `name`, `ticket_number`)
VALUES
  ('1', 'example.com', 12345),
  ('2', 'example.net', 12345)
;


--
INSERT INTO `mailserver`.`virtual_users`
  (`domain_id`, `password`, `email`, `ticket_number`, `comments`)
VALUES
  (1, 'pickles', 'bobbo@example.com', 12345, 'Bob in many guises'),

  -- Intentional bug: trailing space
  (1, 'tuna', 'dorathecat@example.com ', 12345, 'Popular feline'),

  (1, 'cookies', 'martin@example.com', 12345, 'Somebody, somewhere')
;


INSERT INTO `mailserver`.`virtual_aliases`
  (`domain_id`, `source`, `destination`, `ticket_number`, `comments`)
VALUES
  (1, 'bob@example.com', 'bob@example.net', 12345, 'Bob is everywhere'),
  (1, 'alice@example.com', 'alice@example.net', 12345, 'Something snarky'),

  -- Intentional bug: trailing space
  (1, 'dora@example.com ', 'dora@example.net', 12345, 'meow')
;


INSERT INTO `mailserver`.`local_aliases`
  (`source`, `destination`, `ticket_number`, `comments`)
VALUES
  ('root', 'bob', 12345, 'Bob knows all'),
  ('postmaster', 'root', 12345, 'Basic system alias - this MUST be present'),
  ('MAILER-DAEMON', 'root', 12345, 'Basic system alias - this MUST be present'),
  ('adm', 'root', 12345, 'General redirections for pseudo accounts'),
  ('bin', 'root', 12345, 'General redirections for pseudo accounts'),
  ('mail', 'root', 12345, 'General redirections for pseudo accounts'),
  ('nobody', 'root', 12345, 'General redirections for pseudo accounts'),
  ('postfix', 'root', 12345, 'General redirections for pseudo accounts')

;

INSERT INTO `mailserver`.`transport_maps`
  (`recipient`, `transport`, `ticket_number`, `comments`)
VALUES
  ('bob@example.net', 'nodelay-smtp', 12345, 'Must, deliver, slow'),
  ('postfix.org', 'ipv4-only', 12345, 'Force mail via IPv4 transport since our IPv6 transport was wrongly tagged as a source of SPAM'),
  ('angrybirds@example.com', 'nodelay-smtp', 12345, NULL),
  ('silentbob@example.com', 'nodelay-smtp', 12345, '')
;


INSERT INTO `mailserver`.`access_check_clients`
  (`client`, `action`, `ticket_number`, `comments`)
VALUES
  (' example.xyz', 'reject Blocked due to history of spam', 12345, 'They spammed Bob; nobody spams Bob.'),

  -- Intentional bug: trailing space
  ('example.com ', 'reject Blocked due to history of spam', 12345, 'They also spammed Bob; nobody spams Bob.'),

  -- Intentional bug: leading space
  (' gog.com', 'OK', 12345, 'These guys are awesome'),

  ('apress.com', 'OK', 12345, 'Good books found here'),
  ('nostarch.com', 'OK', 12345, 'Good books here too'),
  ('informit.com', 'OK', 12345, 'Ditto re good books')
;


INSERT INTO `mailserver`.`access_check_recipients`
  (`recipient`, `action`, `ticket_number`, `comments`)
VALUES
  ('funwithtribbles@example.com', 'reject Account terminated', 12345, 'Really bad idea'),
  ('bob@example.com', 'reject Account terminated', 12345, 'No longer employed here'),
  ('ilovefleas@example.com', 'reject Invalid account', 12345, 'Spammers found this test email address, had to shut it down'),
  ('tarpit@example.com', 'OK', 12345, 'Network security team wanted an open alias that allows everything through')
;


INSERT INTO `mailserver`.`access_check_senders`
  (`sender`, `action`, `ticket_number`, `comments`)
VALUES
  ('apple.com', 'OK', 12345, 'Bypass zen.spamhaus.org SBL entry'),
  ('no-reply@drivethrustuff.com', 'OK', 12345, 'Whitelist: cannot find your reverse hostame')
;


INSERT INTO `mailserver`.`mail_relay_whitelist`
  (`client`, `action`, `ticket_number`, `comments`)
VALUES
  ('192.168.34.5', 'OK', 12345, 'test entry 1'),
  ('192.168.34.6', 'OK', 12345, 'test entry 2'),
  ('192.168.34.7', 'OK', 12345, 'test entry 3'),

  -- Intentional bug: trailing space
  ('192.168.34.8 ', 'OK', 12345, 'test entry 4'),

  ('192.168.34.9', 'OK', 12345, 'test entry 5'),
  ('192.168.34.10', 'OK', 12345, 'test entry 6')
;


INSERT INTO `mailserver`.`sender_bcc_maps`
  (`sender`, `bcc_recipient`, `ticket_number`, `comments`)
VALUES
  ('notifications@example.com', 'notifications@example.com', 12345, 'Redirect a copy of the notifications back for review'),
  ('help@example.com', 'archive@example.com', 12345, 'Archive a copy of all ticketing system traffic')
;


INSERT INTO `mailserver`.`recipient_bcc_maps`
  (`original_recipient`, `bcc_recipient`, `ticket_number`, `comments`)
VALUES
  ('problem_user@example.xyz', 'legal@example.com', 12345, 'Redirect a copy of all mail to this customer to legal team'),
  ('abuse@example.com', 'legal@example.com', 12345, 'Legal team requested an automatic bcc copy of all abuse reports')
;


-- Only a few entries for the most high-volume senders. This excludes those
-- sender email addresses from rate limiting that is applied to other senders.
INSERT INTO `mailserver`.`sender_dependent_default_transport_maps`
  (`sender`, `transport`, `ticket_number`, `comments`)
VALUES
  ('help@example.com', 'nodelay-smtp', 12345, 'Help requests submitted here'),
  ('notifications@example.com', 'nodelay-smtp', 12345, 'Customer alerts sent from this address'),
  ('nagios@example.net', 'nodelay-smtp', 12345, 'Nagios alerts for staff'),
  ('rsyslog-alerts@example.net', 'nodelay-smtp', 12345, 'Rsyslog alerts for staff')
;
