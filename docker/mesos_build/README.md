## Mesos Docker Container Images

### dvonthenen/mesos-dev
Description: Docker image for a mesos build. Mesos is already pre-compiled within these images. Currently provide versions 0.24.1, 0.25.0 and an intermediate version of 0.26.0 (latest). Pull based on the tag to get the desired version.  
URL: https://hub.docker.com/r/dvonthenen/mesos-dev/

**Example**  
Pull mesos build environment for 0.24.1:  
docker pull dvonthenen/mesos-dev:0.24.1


### dvonthenen/mesos-dev-dvdi
Description: Docker image for a mesos build that will also build the EMC {code} DVDI isolator. Currently provide versions 0.24.1 and 0.25.0. Pull based on the tag to get the desired version.  
URL: https://hub.docker.com/r/dvonthenen/mesos-dev-dvdi/

**Requires DVDI Source Code**  
You can get the DVDI source code at https://github.com/emccode/mesos-module-dvdi

**Syntax**  
docker run -ti -v <location of your DVDI source>:/isolator dvonthenen/mesos-dev-dvdi:<mesos version> /bin/bash

**Example**  
Pull dvdi build environment for 0.24.1:  
docker run -ti -v /home/dev/git-workspace/dvonthenen/mesos-module-dvdi/isolator/:/isolator dvonthenen/mesos-dev:0.24.1 /bin/bash
