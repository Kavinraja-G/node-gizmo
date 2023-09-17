# node-gizmo
A CLI utility for your Kubernetes nodes.

### Features
##### nodegizmo node
Generic node related information
  - NodeName
  - K8sVersion
  - Image
  - OS & Architecture info
  - NodeStatus (Ready/NotReady)
  - Taints
  - Node Provider (AWS/Azure/GCP)
  - Topology info (Region & Zone)
<p align="center"><img src="/assets/nodegizmo-node.png" alt="Nodegizmo node "/></p>

##### nodegizmo node capacity
Node Capacity information
  - CPU
  - Memory
  - Disk
  - Ephemeral storage
  - Pod capacities
- Nodepool related information
<p align="center"><img src="/assets/nodegizmo-node-cp.png" alt="Nodegizmo node "/></p>

##### nodegizmo nodepool
Nodepool related information
  - Grouped by NodePool ID
  - Node list
  - Topology info (Region & Zone)
  - Instance type
  - Nodepool provider (supported: EKS/AKS/GKE)
<p align="center"><img src="/assets/nodegizmo-nodepool.png" alt="Nodegizmo node "/></p>

##### nodegizmo node exec nodeName
> Note: This feature is still in beta

Exec into any node by spawning a `nsenter` pod automatically based on the node selection.