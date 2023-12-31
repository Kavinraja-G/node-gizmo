package pkg

import "testing"

func TestGetNodeTopologyInfo(t *testing.T) {
	type args struct {
		labels map[string]string
	}
	tests := []struct {
		name   string
		args   args
		region string
		zone   string
	}{
		{
			name: "WithRegionAndZoneInfo",
			args: args{
				labels: map[string]string{
					TopologyRegionLabel: "us-east-1",
					TopologyZoneLabel:   "us-east-1a",
				},
			},
			region: "us-east-1",
			zone:   "us-east-1a",
		}, {
			name: "NoRegionAndZoneLabels",
			args: args{
				labels: map[string]string{},
			},
			region: "",
			zone:   "",
		}, {
			name: "WithOnlyRegionLabel",
			args: args{
				labels: map[string]string{
					TopologyRegionLabel: "us-east-1",
				},
			},
			region: "us-east-1",
			zone:   "",
		}, {
			name: "WithOnlyZoneLabel",
			args: args{
				labels: map[string]string{
					TopologyZoneLabel: "us-east-1a",
				},
			},
			region: "",
			zone:   "us-east-1a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetNodeTopologyInfo(tt.args.labels)
			if got != tt.region {
				t.Errorf("GetNodeTopologyInfo() got = %v, region %v", got, tt.region)
			}
			if got1 != tt.zone {
				t.Errorf("GetNodeTopologyInfo() got1 = %v, zone %v", got1, tt.zone)
			}
		})
	}
}
