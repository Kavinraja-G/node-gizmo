package pkg

// GetNodeTopologyInfo retrieves region and zone info from topology labels
func GetNodeTopologyInfo(labels map[string]string) (string, string) {
	var region string
	var zone string

	if val, ok := labels[TopologyRegionLabel]; ok {
		region = val
	}
	if val, ok := labels[TopologyZoneLabel]; ok {
		zone = val
	}
	return region, zone
}
