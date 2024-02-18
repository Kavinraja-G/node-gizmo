# node-gizmo
A CLI utility for your Kubernetes nodes.

[![Release](https://github.com/Kavinraja-G/node-gizmo/actions/workflows/release.yml/badge.svg)](https://github.com/Kavinraja-G/node-gizmo/actions/workflows/release.yml)
[![Go Coverage](https://github.com/Kavinraja-G/node-gizmo/wiki/coverage.svg)](https://raw.githack.com/wiki/Kavinraja-G/node-gizmo/coverage.html)

### Installation
nodegizmo kubectl plugin is available in [krew](https://krew.sigs.k8s.io/) plugin manager. Anyone can install with the following steps:
1. Install `krew` for kuebctl using the following [doc](https://krew.sigs.k8s.io/docs/user-guide/setup/install/).
2. ```bash
    > kubectl krew install nodegizmo
    ```

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
  - K8sVersion
  - Nodepool provider (supported: EKS/AKS/GKE/Karpenter)
<p align="center"><img src="/assets/nodegizmo-nodepool.png" alt="Nodegizmo node "/></p>

##### nodegizmo exec nodeName

Exec into any node by spawning a `nsenter` pod automatically based on the node selection.