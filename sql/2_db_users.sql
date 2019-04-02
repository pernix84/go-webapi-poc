CREATE USER 'app_authservice'@'%' IDENTIFIED BY 'Qwer1234';

GRANT ALL PRIVILEGES ON `authservice`.* TO 'app_authservice'@'%';


FLUSH PRIVILEGES;