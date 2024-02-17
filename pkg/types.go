package pkg

// GenericNodeInfo struct used for 'node' info command
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

// OutputOptsForGenericNodeInfo options set via flags, mostly used in the table-writer
// for output format selections
type OutputOptsForGenericNodeInfo struct {
	ShowTaints           bool
	ShowNodeProviderInfo bool
	ShowNodeTopologyInfo bool
}

// NodeCapacities required node capacity related info. Used by 'node capacity' command
type NodeCapacities struct {
	NodeName         string
	CPU              string // cores
	Memory           int64  // bytes
	Disk             int64  // bytes
	EphemeralStorage int64  // bytes
	PodCapacity      string // count
}

// GenericNodepoolInfo required node pool info which is used 'nodepool' command
type GenericNodepoolInfo struct {
	NodepoolID   string
	Node         string
	Provider     string
	InstanceType string
	Region       string
	Zone         string
	K8sVersion   string
}
