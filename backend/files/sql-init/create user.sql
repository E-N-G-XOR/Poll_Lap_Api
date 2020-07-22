CREATE USER 'appuser'@'%' IDENTIFIED BY 'apppass';

GRANT ALL PRIVILEGES ON appdb.* TO 'appuser'@'%';