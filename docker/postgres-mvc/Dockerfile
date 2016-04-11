FROM postgres:9.3.11
MAINTAINER David vonThenen <david.vonthenen@emc.com>

ENV DB_NAME demo
ENV DB_USER dev
ENV DB_PASS vmware

ADD setup-remote.sh /docker-entrypoint-initdb.d/setup-remote.sh
RUN chmod 755 /docker-entrypoint-initdb.d/setup-remote.sh
ADD init.sql /docker-entrypoint-initdb.d/init.sql
