package algorithms

import "strings"

func StringOverlap(a, b string) int {
	var overlap int

	split := strings.Split(b, "")
	for _, s := range split {
		if strings.Contains(a, s) {
			overlap++
		}
	}

	return overlap
}
