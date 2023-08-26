package pkg

type GenericNodeInfo struct {
	NodeName   string
	K8sVersion string
	Image      string
	Os         string
	OsArch     string
	NodeStatus string
	Taints     []string
}

type OutputOptsForGenericNodeInfo struct {
	ShowTaints bool
}

type NodeCapacities struct {
	NodeName         string
	CPU              string // cores
	Memory           int64  // bytes
	Disk             int64  // bytes
	EphemeralStorage int64  // bytes
	PodCapacity      string // count
}

type DetailedNodeInfo struct {
	NodeName           string
	NodeAddresses      []string
	NodeHostname       string
	NodeTopologyRegion string
	NodeTopologyZone   string
	NodeCloudProvider  string
	NodeGroup          string
}
