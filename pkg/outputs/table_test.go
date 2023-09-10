package outputs

import "testing"

func TestTableOutput(t *testing.T) {
	type args struct {
		headers    []string
		outputData [][]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "GenericOutputTable",
			args: args{
				headers: []string{"NAME", "VERSION", "OS", "ARCHITECTURE"},
				outputData: [][]string{
					{
						"Node01",
						"1.25.6",
						"linux",
						"amd64",
					}, {
						"Node02",
						"1.25.6",
						"linux",
						"arm64",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TableOutput(tt.args.headers, tt.args.outputData)
		})
	}
}
