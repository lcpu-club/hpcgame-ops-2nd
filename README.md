# HPCGame 2nd Infra Ops

This repository contains the SDI (Software Defined Infrastructure) for the HPCGame 2nd Infrastructure Operations.

Based on the Kubernetes and Ansible platform, we aim to make the full cluster deployment and management reproduceable and easy to maintain.

## TODOs

- Containerd: DEPRECATION: The `mirrors` property of `[plugins."io.containerd.grpc.v1.cri".registry]` is deprecated since containerd v1.5 and will be removed in containerd v2.1. Use `config_path` instead.

## Notes

Cross building: `docker run --rm --privileged tonistiigi/binfmt --install arm64,arm,riscv64`.

Then: `docker build . -t tag --platform=linux/amd64,arm64`.

If building golang, use golang's cross build to accelerate: `FROM --platform=${BUILDPLATFORM} golang:1.22 as builder`.
