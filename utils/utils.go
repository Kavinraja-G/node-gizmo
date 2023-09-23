package utils

import (
	"fmt"
	"math"
	"os"
)

// GetEnv Helper function for fetching envs with defaults
func GetEnv(env, defaults string) string {
	if val, ok := os.LookupEnv(env); ok {
		return val
	}
	return defaults
}

// PrettyByteSize converts the bytes to human-readable format
func PrettyByteSize(b int64) string {
	bf := float64(b)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", bf, unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%.1fYiB", bf)
}
