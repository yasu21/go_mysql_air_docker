FROM mysql:8.0.28

COPY my.cnf /etc/mysql/conf.d/my.cnf
RUN chmod 644 /etc/mysql/conf.d/my.cnf

ADD init.sql /docker-entrypoint-initdb.d
