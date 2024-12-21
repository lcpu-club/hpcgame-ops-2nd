# Labels for Nodes

`kubernetes.io/hostname`
`topology.kubernetes.io/region`
`topology.kubernetes.io/zone`
`kubernetes.io/arch`
`kubernetes.io/os`

`hpc.lcpu.dev/partition` (`control`, `x86`, `gpu80g`, `gpu40g`, `npu`, `arm`)

`hpc.lcpu.dev/lustre-capable`: (`true`, `false`)

# Taints for Nodes

`node-role.kubernetes.io/control-plane:NoSchedule`
`node-role.hpc.lcpu.dev/control-plane:NoSchedule`

<!-- `node-type.hpc.lcpu.dev/arch-arm64:PreferNoSchedule` -->
<!-- `node-type.hpc.lcpu.dev/accelerator-gpu:PreferNoSchedule` -->
<!-- `node-type.hpc.lcpu.dev/accelerator-npu:PreferNoSchedule` -->
