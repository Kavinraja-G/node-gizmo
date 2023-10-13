package outputs

import (
	"k8s.io/utils/strings/slices"
	"sort"
	"strings"
)

// getSortKeyIdxFromHeader retrieves the index of the sortKey in the header slice
func getSortKeyIdxFromHeader(headers []string, sortKey string) int {
	// defaults to first column always (usually node/nodepool name)
	idx := 0
	if slices.Contains(headers, strings.ToUpper(sortKey)) {
		idx = slices.Index(headers, strings.ToUpper(sortKey))
	}

	return idx
}

// SortOutputBasedOnHeader sorts output based on the Header key provided in the flags
func SortOutputBasedOnHeader(headers []string, sortSlices [][]string, sortKey string) {
	sortByHeaderIndex := getSortKeyIdxFromHeader(headers, sortKey)

	sort.SliceStable(sortSlices, func(i, j int) bool {
		return sortSlices[i][sortByHeaderIndex] < sortSlices[j][sortByHeaderIndex]
	})
}
