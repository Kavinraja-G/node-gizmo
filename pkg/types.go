package pkg

type GenericNodeInfo struct {
	NodeName           string
	K8sVersion         string
	Image              string
	Os                 string
	OsArch             string
	NodeStatus         string
	Taints             []string
	NodeProvider       string
	NodeTopologyRegion string
	NodeTopologyZone   string
}

type OutputOptsForGenericNodeInfo struct {
	ShowTaints           bool
	ShowNodeProviderInfo bool
	ShowNodeTopologyInfo bool
}

type NodeCapacities struct {
	NodeName         string
	CPU              string // cores
	Memory           int64  // bytes
	Disk             int64  // bytes
	EphemeralStorage int64  // bytes
	PodCapacity      string // count
}

type GenericNodepoolInfo struct {
	NodepoolID   string
	Nodes        []string
	Provider     string
	InstanceType string
	Region       string
	Zone         string
}
