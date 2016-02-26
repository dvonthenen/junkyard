#!/bin/bash

#simple restart script that clears out all zookeeper cache to start fresh
vagrant ssh -c 'sudo stop marathon; sudo stop mesos-master; sudo stop zookeeper; sudo rm -rf /var/lib/zookeeper/version-2; sudo start zookeeper; sudo start mesos-master; sudo start marathon' mesos-master
vagrant ssh -c 'sudo stop mesos-slave; sudo rm -f /tmp/mesos/meta/slaves/latest; sudo start mesos-slave' mesos-slave1
vagrant ssh -c 'sudo stop mesos-slave; sudo rm -f /tmp/mesos/meta/slaves/latest; sudo start mesos-slave' mesos-slave2
vagrant ssh -c 'sudo stop mesos-slave; sudo rm -f /tmp/mesos/meta/slaves/latest; sudo start mesos-slave' mesos-slave3
vagrant ssh -c 'sudo stop mesos-slave; sudo rm -f /tmp/mesos/meta/slaves/latest; sudo start mesos-slave' mesos-slave4
