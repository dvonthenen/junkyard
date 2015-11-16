## Example Spring REST End-Point as a Docker Image

**Prerequisites:**  
STS - https://spring.io/tools/sts/all  
Apache Maven - https://maven.apache.org/download.cgi  
Docker Maven Plugin - https://github.com/spotify/docker-maven-plugin

This was just a exercise for creating a DEB package for Ubuntu. Since this is a spring project, I used the STS for my IDE with Apache Maven as my package manager. The docker image creation process is hooked into STS using using Spotify's docker-maven-plugin.

The image for this can be found on Docker Hub located at:  
Docker Hub URL: https://hub.docker.com/r/dvonthenen/demo-boot/

To launch the example:  
docker run -t -i dvonthenen/demo-boot
