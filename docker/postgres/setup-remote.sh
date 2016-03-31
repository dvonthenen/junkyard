#!/bin/bash

sed -i '/host    all             all             0.0.0.0\/0            md5/a host    all             all             172.17.0.1\/16            trust' /var/lib/postgresql/data/pg_hba.conf
