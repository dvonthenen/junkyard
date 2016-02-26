#!/bin/bash

#IMPORTANT: DONT USE THIS SCRIPT. THIS WAS ONLY FOR TESTING AND REVERSE ENGINEERING PURPOSES
#IT TURNS OUT HOSTNAMES EVEN WITH A PRIVATE DNS SERVER DOES NOT WORK WITH
#VIRTUALBOX BECAUSE OF ALL THE NICs THAT GET PLACED ON THE INSTANCES

#general
vagrant ssh -c 'echo "set nocompatible" | sudo tee $HOME/.vimrc' mesos-master
vagrant ssh -c 'echo "set nocompatible" | sudo tee $HOME/.vimrc' mesos-slave1
vagrant ssh -c 'echo "set nocompatible" | sudo tee $HOME/.vimrc' mesos-slave2
vagrant ssh -c 'echo "set nocompatible" | sudo tee $HOME/.vimrc' mesos-slave3
vagrant ssh -c 'echo "set nocompatible" | sudo tee $HOME/.vimrc' mesos-slave4

#master
vagrant ssh -c 'echo "10.141.141.10" | sudo tee /etc/mesos-master/ip' mesos-master
vagrant ssh -c 'echo "mesos-master.demo.com" | sudo tee /etc/hostname' mesos-master
vagrant ssh -c 'echo "mesos-master.demo.com" | sudo tee /etc/mesos-master/hostname' mesos-master
vagrant ssh -c 'echo "zk://10.141.141.10:2181/mesos" | sudo tee /etc/mesos/zk' mesos-master
vagrant ssh -c 'echo "1" | sudo tee /etc/zookeeper/conf/myid' mesos-master
vagrant ssh -c 'sudo sed -i -- ''s/\#server\.1\=zookeeper1/server\.1\=10.141.141.10/g'' /etc/zookeeper/conf/zoo.cfg' mesos-master
vagrant ssh -c 'sudo mkdir -p /etc/marathon/conf; sudo cp /etc/mesos-master/hostname /etc/marathon/conf' mesos-master
vagrant ssh -c 'sudo cp /etc/mesos/zk /etc/marathon/conf/master' mesos-master
vagrant ssh -c 'sudo cp /etc/marathon/conf/master /etc/marathon/conf/zk' mesos-master
vagrant ssh -c 'sudo sed -i -- ''s/mesos/marathon/g'' /etc/marathon/conf/zk' mesos-master

#slaves
vagrant ssh -c 'echo "mesos-slave1.demo.com" | sudo tee /etc/mesos-slave/hostname' mesos-slave1
vagrant ssh -c 'echo "mesos-slave1.demo.com" | sudo tee /etc/mesos-slave/hostname' mesos-slave1
vagrant ssh -c 'echo "zk://10.141.141.10:2181/mesos" | sudo tee /etc/mesos/zk' mesos-slave1
#vagrant ssh -c 'echo "52.36.116.225" | sudo tee /etc/resolv.conf' mesos-slave1
vagrant ssh -c 'echo "mesos-slave2.demo.com" | sudo tee /etc/mesos-slave/hostname' mesos-slave2
vagrant ssh -c 'echo "mesos-slave2.demo.com" | sudo tee /etc/mesos-slave/hostname' mesos-slave2
vagrant ssh -c 'echo "zk://10.141.141.10:2181/mesos" | sudo tee /etc/mesos/zk' mesos-slave2
#vagrant ssh -c 'echo "52.36.116.225" | sudo tee /etc/resolv.conf' mesos-slave2
vagrant ssh -c 'echo "mesos-slave3.demo.com" | sudo tee /etc/mesos-slave/hostname' mesos-slave3
vagrant ssh -c 'echo "mesos-slave3.demo.com" | sudo tee /etc/mesos-slave/hostname' mesos-slave3
vagrant ssh -c 'echo "zk://10.141.141.10:2181/mesos" | sudo tee /etc/mesos/zk' mesos-slave3
#vagrant ssh -c 'echo "52.36.116.225" | sudo tee /etc/resolv.conf' mesos-slave3
vagrant ssh -c 'echo "mesos-slave4.demo.com" | sudo tee /etc/mesos-slave/hostname' mesos-slave4
vagrant ssh -c 'echo "mesos-slave4.demo.com" | sudo tee /etc/mesos-slave/hostname' mesos-slave4
vagrant ssh -c 'echo "zk://10.141.141.10:2181/mesos" | sudo tee /etc/mesos/zk' mesos-slave4
#vagrant ssh -c 'echo "52.36.116.225" | sudo tee /etc/resolv.conf' mesos-slave4

#DNS
#vagrant ssh -c 'sudo grep -q master /etc/hosts || echo "10.141.141.10 mesos-master" | sudo tee -a /etc/hosts' mesos-master
#vagrant ssh -c 'sudo grep -q slave1 /etc/hosts || echo "10.141.141.11 mesos-slave1" | sudo tee -a /etc/hosts' mesos-master
#vagrant ssh -c 'sudo grep -q slave2 /etc/hosts || echo "10.141.141.12 mesos-slave2" | sudo tee -a /etc/hosts' mesos-master
#vagrant ssh -c 'sudo grep -q slave3 /etc/hosts || echo "10.141.141.13 mesos-slave3" | sudo tee -a /etc/hosts' mesos-master
#vagrant ssh -c 'sudo grep -q slave4 /etc/hosts || echo "10.141.141.14 mesos-slave4" | sudo tee -a /etc/hosts' mesos-master
#vagrant ssh -c 'sudo grep -q master /etc/hosts || echo "10.141.141.10 mesos-master" | sudo tee -a /etc/hosts' mesos-slave1
#vagrant ssh -c 'sudo grep -q slave1 /etc/hosts || echo "10.141.141.11 mesos-slave1" | sudo tee -a /etc/hosts' mesos-slave1
#vagrant ssh -c 'sudo grep -q slave2 /etc/hosts || echo "10.141.141.12 mesos-slave2" | sudo tee -a /etc/hosts' mesos-slave1
#vagrant ssh -c 'sudo grep -q slave3 /etc/hosts || echo "10.141.141.13 mesos-slave3" | sudo tee -a /etc/hosts' mesos-slave1
#vagrant ssh -c 'sudo grep -q slave4 /etc/hosts || echo "10.141.141.14 mesos-slave4" | sudo tee -a /etc/hosts' mesos-slave1
#vagrant ssh -c 'sudo grep -q master /etc/hosts || echo "10.141.141.10 mesos-master" | sudo tee -a /etc/hosts' mesos-slave2
#vagrant ssh -c 'sudo grep -q slave1 /etc/hosts || echo "10.141.141.11 mesos-slave1" | sudo tee -a /etc/hosts' mesos-slave2
#vagrant ssh -c 'sudo grep -q slave2 /etc/hosts || echo "10.141.141.12 mesos-slave2" | sudo tee -a /etc/hosts' mesos-slave2
#vagrant ssh -c 'sudo grep -q slave3 /etc/hosts || echo "10.141.141.13 mesos-slave3" | sudo tee -a /etc/hosts' mesos-slave2
#vagrant ssh -c 'sudo grep -q slave4 /etc/hosts || echo "10.141.141.14 mesos-slave4" | sudo tee -a /etc/hosts' mesos-slave2
#vagrant ssh -c 'sudo grep -q master /etc/hosts || echo "10.141.141.10 mesos-master" | sudo tee -a /etc/hosts' mesos-slave3
#vagrant ssh -c 'sudo grep -q slave1 /etc/hosts || echo "10.141.141.11 mesos-slave1" | sudo tee -a /etc/hosts' mesos-slave3
#vagrant ssh -c 'sudo grep -q slave2 /etc/hosts || echo "10.141.141.12 mesos-slave2" | sudo tee -a /etc/hosts' mesos-slave3
#vagrant ssh -c 'sudo grep -q slave3 /etc/hosts || echo "10.141.141.13 mesos-slave3" | sudo tee -a /etc/hosts' mesos-slave3
#vagrant ssh -c 'sudo grep -q slave4 /etc/hosts || echo "10.141.141.14 mesos-slave4" | sudo tee -a /etc/hosts' mesos-slave3
#vagrant ssh -c 'sudo grep -q master /etc/hosts || echo "10.141.141.10 mesos-master" | sudo tee -a /etc/hosts' mesos-slave4
#vagrant ssh -c 'sudo grep -q slave1 /etc/hosts || echo "10.141.141.11 mesos-slave1" | sudo tee -a /etc/hosts' mesos-slave4
#vagrant ssh -c 'sudo grep -q slave2 /etc/hosts || echo "10.141.141.12 mesos-slave2" | sudo tee -a /etc/hosts' mesos-slave4
#vagrant ssh -c 'sudo grep -q slave3 /etc/hosts || echo "10.141.141.13 mesos-slave3" | sudo tee -a /etc/hosts' mesos-slave4
#vagrant ssh -c 'sudo grep -q slave4 /etc/hosts || echo "10.141.141.14 mesos-slave4" | sudo tee -a /etc/hosts' mesos-slave4

#this is just an optimization is not needed
vagrant ssh -c 'sudo docker pull dvonthenen/elasticsearch-scheduler; sudo docker pull dvonthenen/elasticsearch-executor' mesos-slave1
vagrant ssh -c 'sudo docker pull dvonthenen/elasticsearch-scheduler; sudo docker pull dvonthenen/elasticsearch-executor' mesos-slave2
vagrant ssh -c 'sudo docker pull dvonthenen/elasticsearch-scheduler; sudo docker pull dvonthenen/elasticsearch-executor' mesos-slave3
vagrant ssh -c 'sudo docker pull dvonthenen/elasticsearch-scheduler; sudo docker pull dvonthenen/elasticsearch-executor' mesos-slave4
vagrant ssh -c 'sudo docker pull dvonthenen/elasticsearch-scheduler-test; sudo docker pull dvonthenen/elasticsearch-executor-test' mesos-slave1
vagrant ssh -c 'sudo docker pull dvonthenen/elasticsearch-scheduler-test; sudo docker pull dvonthenen/elasticsearch-executor-test' mesos-slave2
vagrant ssh -c 'sudo docker pull dvonthenen/elasticsearch-scheduler-test; sudo docker pull dvonthenen/elasticsearch-executor-test' mesos-slave3
vagrant ssh -c 'sudo docker pull dvonthenen/elasticsearch-scheduler-test; sudo docker pull dvonthenen/elasticsearch-executor-test' mesos-slave4

#restart
vagrant ssh -c 'sudo stop marathon; sudo stop mesos-master; sudo stop zookeeper; sudo rm -rf /var/lib/zookeeper/version-2; sudo start zookeeper; sudo start mesos-master; sudo start marathon' mesos-master
vagrant ssh -c 'sudo stop mesos-slave; sudo rm -f /tmp/mesos/meta/slaves/latest; sudo start mesos-slave' mesos-slave1
vagrant ssh -c 'sudo stop mesos-slave; sudo rm -f /tmp/mesos/meta/slaves/latest; sudo start mesos-slave' mesos-slave2
vagrant ssh -c 'sudo stop mesos-slave; sudo rm -f /tmp/mesos/meta/slaves/latest; sudo start mesos-slave' mesos-slave3
vagrant ssh -c 'sudo stop mesos-slave; sudo rm -f /tmp/mesos/meta/slaves/latest; sudo start mesos-slave' mesos-slave4
