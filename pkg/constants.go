package pkg

const (
	AwsNodepoolLabel      = "eks.amazonaws.com/nodegroup"
	GkeNodepoolLabel      = "cloud.google.com/gke-nodepool"
	AksNodepoolLabel      = "kubernetes.azure.com/agentpool"
	NodeInstanceTypeLabel = "node.kubernetes.io/instance-type"
	TopologyRegionLabel   = "topology.kubernetes.io/region"
	TopologyZoneLabel     = "topology.kubernetes.io/zone"
)
