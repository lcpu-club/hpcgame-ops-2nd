#!/bin/bash

my_hostname="$(hostname -s)"
QCOW2_IMAGE="/mnt/nfs/_system/$my_hostname.qcow2"
MOUNT_POINT="/var/lib/containerd"

modprobe nbd max_part=16

qemu-nbd --connect=/dev/nbd0 "$QCOW2_IMAGE"
mount /dev/nbd0 "$MOUNT_POINT"
