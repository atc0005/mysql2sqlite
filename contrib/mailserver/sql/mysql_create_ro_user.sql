/*

   Copyright 2020 Adam Chalkley

   https://github.com/atc0005/mysql2sqlite

   Licensed under the MIT License. See LICENSE file in the project root for
   full license information.

*/

-- Create db user
GRANT SELECT,LOCK TABLES ON mailserver.*
    TO 'mysql2sqlite'@'127.0.0.1'
    IDENTIFIED BY 'qwerty';

FLUSH PRIVILEGES;
