FROM java:8
VOLUME /tmp
ADD gs-spring-boot-docker-0.1.0.jar demo-boot-0.0.1-SNAPSHOT.jar
RUN bash -c 'touch /demo-boot-0.0.1-SNAPSHOT.jar'
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/demo-boot-0.0.1-SNAPSHOT.jar"]