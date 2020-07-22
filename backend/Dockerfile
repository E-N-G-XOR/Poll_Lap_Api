FROM mysql:latest

ENV MYSQL_DATABASE appdb
ENV MYSQL_ROOT_PASSWORD dbrootpass

EXPOSE 3306

COPY files/sql-init/ /docker-entrypoint-initdb.d/