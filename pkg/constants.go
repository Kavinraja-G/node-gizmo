package pkg

// required constants used across the CLI commands (mostly well-known labels, annotations etc.,)
const (
	AwsNodepoolLabel     = "eks.amazonaws.com/nodegroup"
	GkeNodepoolLabel     = "cloud.google.com/gke-nodepool"
	AksNodepoolLabel     = "kubernetes.azure.com/agentpool"
	KarpenterNodepool    = "karpenter.sh/provisioner-name"
	KarpenterNodepoolV1  = "karpenter.sh/nodepool"
	OpenshiftMachinepool = "hive.openshift.io/machine-pool"

	NodeInstanceTypeLabel = "node.kubernetes.io/instance-type"
	TopologyRegionLabel   = "topology.kubernetes.io/region"
	TopologyZoneLabel     = "topology.kubernetes.io/zone"
)
