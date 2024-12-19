# Hardware Orchestration

This document describes the hardwares used in the HPCGame 2nd Infrastructure Operations.

We haven't applied the cutting-edge robot technology yet, so we still need to manually do such operations. However, we try to write the steps down and make them reproducible. We call such method "Human Reconciliation".

## Hardware List

### Overview

We have three clusters, located in three different datacenters. Each cluster has different hardware specifications.

The first one is Weiming 2 Supercomputer, located in the Xin Yan Yuan Campus of Peking University, which is geo-located in Changping, Beijing. We utilize over 20 elastic nodes, including CPU (Intel Xeon Platinum 8358) and GPU (NVIDIA A100) ones. They are all x86_64 architecture. We have a control-plane 1GbE network and a data-plane 100GbE RoCE-capable network. Each node get a Mellanox ConnectX-5 NIC, connected to a large spine-leaf L2 underlay network. The network infrastructure is to some extent unmodifiable, so we have to adapt to it using SDN technologies.

The second one is located inside Yan Yuan Campus, the main campus of Peking University. It is a heterogeneous cluster, with x86_64 and ARM64 nodes. The x86_64 nodes are equipped with AMD EPYC 9654 CPUs and NVIDIA Tesla V100 GPUs. The ARM64 nodes are equipped with Ascend Atlas A2 Training Series NPUs. The x86_64 nodes are interconnected using 100Gbps Infiniband, while the ARM64 nodes are connected using 100Gbps RoCE. The network infrastructure is modifiable to some extent. Besides that, we have a cloud based on OpenStack, which can share the same control-plane network (10Gbps) with the cluster.

The third one is provided by Huawei Technologies, on the Huawei Cloud Platform. It is a cloud cluster, with Huawei's most capable ARM64 CPU machines.

In the beginning, we even proposed to utilize LoongArch / RISC-V machines. However, due to the lack of testing time, we have to give up such crazy ideas.

The first and the second cluster are in the same PKU LAN, between which is a reliable (almost lossless) 10Gbps fibre link. However, there are strict firewall rules between them, so we have to use encapsulation technologies to connect them. The Huawei cluster can only be accessed through the Internet, with only a unreliable 1Gbps network.

### How to maintain the cluster

With such a hyper-converged, heterogeneous and diversed infrastructure, we definitely shall employ newest measures the maintain the cluster. Plus, we want to make the cluster free from known security flaws, so the traditional HPC software stack does not suite us.

The great complexity pushes us to utilize the cutting-edge SDI technologies. After careful investigation, we finally chose Kubernetes and the Operator Framework as the base of our SDI. We also use Ansible and MAAS for the hardware provisioning and configuration management when Kubernetes has not been deployed.

## The infrastructure for orchestration

We have several control-plane virtual machines in the first cluster. Besides these VMs, we also have the ability to elasticly deploy control-plane VMs on the Clab Cloud (the OpenStack cloud, built by Linux Club of Peking University under the support of PKU ITS). The outgoing traffic gateway and the websites are hosted on LCPU's central server, which cannot be changed due to the school's safety policy.

We deployed MAAS, a bare-metal provisioning system by Canonical, on control-plane VMs and Clab. Thus, we can easily install the OS on the nodes and configure the network. We also deployed Ansible on the control-plane VMs, which can be used to configure the nodes after the OS is installed.

## The network

The network is the most important part of the infrastructure. We have to make sure that the network is reliable and fast enough. We have to make sure that the network is secure and isolated. We have to make sure that the network is easy to manage and monitor.

We use WireGuard to connect the clusters, forming a large interconnected LAN. We have dedicated gateway nodes to terminate the WireGuard tunnels, making sure inner cluster network's performance. This is called our control-plane network. For shared storage across the clusters, as the OSS services already provides encryption via TLS, and in our tests WireGuard cannot make full use of the 10Gbps interlink, we give it a routable Internet IP address so it does not need to pass WireGuard.

We also have a data-plane network, which is used for high-speed, low-latency communication between the nodes. We use the 100Gbps RoCE network in the first cluster, and the 100Gbps Infiniband network in the second cluster. The Huawei cluster is on the cloud, so we use the VPC network.
