# node-gizmo
A CLI utility for your Kubernetes nodes

### Features
- Generic node related information
  - NodeName
  - K8sVersion
  - Image
  - OS & Architecture info
  - NodeStatus (Ready/NotReady)
  - Taints
  - Node Provider (AWS/Azure/GCP)
  - Topology info (Region & Zone)
- Node Capacity information
  - CPU
  - Memory
  - Disk
  - Ephemeral storage
  - Pod capacities
- Nodepool related information
  - Grouped by NodePool ID
  - Node list
  - Topology info (Region & Zone)
  - Instance type
  - Nodepool provider (supported: EKS/AKS/GKE)
- Exec into any node by spawning a `nsenter` pod automatically based on the node selection