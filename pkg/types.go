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
