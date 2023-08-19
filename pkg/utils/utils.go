package utils

import "os"

// GetEnv Helper function for fetching envs with defaults
func GetEnv(env, defaults string) string {
	if val, ok := os.LookupEnv(env); ok {
		return val
	}
	return defaults
}
