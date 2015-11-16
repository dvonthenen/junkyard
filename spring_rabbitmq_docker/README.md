## Example Spring RabbitMQ Consumer/Producer Project using Docker Compose

**Prerequisites:**  
STS - https://spring.io/tools/sts/all  
Apache Maven - https://maven.apache.org/download.cgi  
RabbitMQ - http://www.rabbitmq.com/  
RabbitMQ Docker Image - https://hub.docker.com/_/rabbitmq/

This was just a exercise for creating a Docker Compose project that uses RabbitMQ in a message producer/consumer model. Since this is a Spring project, I used the STS for my IDE with Apache Maven as my package manager.

There are three docker images that make up this project:  
1. RabbitMQ Docker Container - This takes the existing the RabbitMQ Docker Image and simply extends the version they have in order to insert a custom configuration when the container is created. This is located on Docker Hub at https://hub.docker.com/r/dvonthenen/rabbitmq/
2. Producer Docker Container - This is a Spring message producer for RabbitMQ's message bus. It's a simple project that publishes a message on the bus every 3 seconds. This is located on Docker Hub at https://hub.docker.com/r/dvonthenen/producer/
3. Consumer Docker Container - This is a Spring message consumer. It will takes the messages published on the bus by the producer and print them to the console. This is located on Docker Hub at https://hub.docker.com/r/dvonthenen/consumer/

You can launch these components separately by doing the following:  
docker run -t -i dvonthenen/rabbitmq  
docker run -t -i dvonthenen/producer  
docker run -t -i dvonthenen/consumer

Or you can create a yaml file named docker-compose.yml with the following contents:  
```
rabbit:
  image: dvonthenen/rabbitmq
  hostname: rabbitmq
  ports:
    - "5672:5672"
    - "15672:15672"

consumer:
  image: dvonthenen/consumer
  hostname: consumer
  links:
    - rabbit:rabbit

producer:
  image: dvonthenen/producer
  hostname: producer
  links:
    - rabbit:rabbit
    - consumer:consumer
```

Then run using Docker Compose by executing:  
docker-compose up

All 3 containers will be started as a single unit.
