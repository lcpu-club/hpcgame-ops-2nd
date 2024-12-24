FROM ubuntu:24.04

RUN apt update && apt install -y lxcfs

ENTRYPOINT ["/usr/bin/lxcfs", "/var/lib/lxcfs", "--enable-loadavg", "--enable-cfs"]
