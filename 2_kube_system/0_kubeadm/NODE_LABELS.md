# Labels for Nodes

`kubernetes.io/hostname`
`topology.kubernetes.io/region`
`topology.kubernetes.io/zone`
`kubernetes.io/arch`
`kubernetes.io/os`

`hpc.lcpu.dev/partition` (`control`, `x86`, `x86_amd`, `gpu-a100`, `gpu-l40s`, `npu`, `arm`)

`hpc.lcpu.dev/lustre-capable`: (`true`, `false`)

# Taints for Nodes

`node-role.kubernetes.io/control-plane:NoSchedule`
`node-role.hpc.lcpu.dev/control-plane:NoSchedule`

<!-- `node-type.hpc.lcpu.dev/arch-arm64:PreferNoSchedule` -->
<!-- `node-type.hpc.lcpu.dev/accelerator-gpu:PreferNoSchedule` -->
<!-- `node-type.hpc.lcpu.dev/accelerator-npu:PreferNoSchedule` -->

# Command

```bash
kubectl label node/xcat hpc.lcpu.dev/partition=control topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2
kubectl label node/control2 hpc.lcpu.dev/partition=control topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2
kubectl label node/control3 hpc.lcpu.dev/partition=control topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2

# kubectl label node/xcat hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2
kubectl label node/sc2 hpc.lcpu.dev/partition=x86_amd topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/sc4 hpc.lcpu.dev/partition=x86_amd topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main

kubectl label node/pku-arm-a01 hpc.lcpu.dev/partition=arm topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/pku-arm-a02 hpc.lcpu.dev/partition=arm topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main

kubectl label node/pku-arm-c01 hpc.lcpu.dev/partition=npu topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
```
