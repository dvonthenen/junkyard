FROM dvonthenen/mesos-dev:latest
MAINTAINER David vonThenen (EMCCODE)

ENTRYPOINT cd /isolator && \
    export ISOLATOR_VERSION=0.99.9 && \
    export LD_LIBRARY_PATH=LD_LIBRARY_PATH:/usr/local/lib && \
    ./bootstrap && \
    rm -Rf build && \
    mkdir build && \
    cd build && \
    ../configure --with-mesos-root=/mesos --with-mesos-build-dir=/mesos && \
    make all
