package utils

import "testing"

func TestPrettyByteSize(t *testing.T) {
	type args struct {
		b int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ValidByteInPrettyFormatInKiB",
			args: args{
				b: 1024,
			},
			want: "1.0KiB",
		}, {
			name: "ValidByteInPrettyFormatInGiB",
			args: args{
				b: 1073741824,
			},
			want: "1.0GiB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrettyByteSize(tt.args.b); got != tt.want {
				t.Errorf("PrettyByteSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
