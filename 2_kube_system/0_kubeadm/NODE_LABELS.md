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

kubectl label node/l08c49n1 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c49n2 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c49n3 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c49n4 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c50n1 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c50n2 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c50n3 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c50n4 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c51n1 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c51n2 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c52n1 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c52n2 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c52n3 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c52n4 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c54n1 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c54n2 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c54n3 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c54n4 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c55n1 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c55n2 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c55n3 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c56n2 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c56n3 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true
kubectl label node/l08c56n4 hpc.lcpu.dev/partition=x86 topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-wm2 hpc.lcpu.dev/nfs-capable=true

kubectl label node/sc2 hpc.lcpu.dev/partition=x86_amd topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/sc4 hpc.lcpu.dev/partition=x86_amd topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main

kubectl label node/pku-arm-a01 hpc.lcpu.dev/partition=arm topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main hpc.lcpu.dev/outdated=true
kubectl label node/pku-arm-a02 hpc.lcpu.dev/partition=arm topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main hpc.lcpu.dev/outdated=true
kubectl taint node pku-arm-a01 hpc.lcpu.dev/outdated=true:NoSchedule
kubectl taint node pku-arm-a02 hpc.lcpu.dev/outdated=true:NoSchedule

kubectl label node/pku-arm-b01 hpc.lcpu.dev/partition=arm topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/pku-arm-b02 hpc.lcpu.dev/partition=arm topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/pku-arm-b03 hpc.lcpu.dev/partition=arm topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/pku-arm-b04 hpc.lcpu.dev/partition=arm topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/pku-arm-b05 hpc.lcpu.dev/partition=arm topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main

kubectl label node/pku-arm-c01 hpc.lcpu.dev/partition=npu topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/pku-arm-c02 hpc.lcpu.dev/partition=npu topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/pku-arm-c03 hpc.lcpu.dev/partition=npu topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main

kubectl label node/pku-arm-d01 hpc.lcpu.dev/partition=npu_inf topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/pku-arm-d02 hpc.lcpu.dev/partition=npu_inf topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/pku-arm-d03 hpc.lcpu.dev/partition=npu_inf topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/pku-arm-d04 hpc.lcpu.dev/partition=npu_inf topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
kubectl label node/pku-arm-d05 hpc.lcpu.dev/partition=npu_inf topology.kubernetes.io/region=cn-bj-pku topology.kubernetes.io/zone=cn-bj-pku-main
```
