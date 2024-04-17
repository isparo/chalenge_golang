# Dockerfile-mysql
FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD: root
ENV MYSQL_DATABASE: user_service
ENV MYSQL_USER: user123
ENV MYSQL_PASSWORD: pass123

COPY database.sql /docker-entrypoint-initdb.d/database.sql