package getmiddlestring

import (
	"strings"
)

func Get(src string, left string, right string, more ...bool) string {
	if src == "" || left == "" || right == "" {
		return ""
	}

	leftIndex := strings.Index(src, left)
	if leftIndex == -1 {
		return ""
	}

	leftIndex += len(left)
	sub := src[leftIndex:]
	var rightIndex int
	if len(more) > 0 && more[0] {
		rightIndex = strings.LastIndex(sub, right)
	} else {
		rightIndex = strings.Index(sub, right)
	}

	if rightIndex == -1 {
		return ""
	}

	return src[leftIndex : leftIndex+rightIndex]
}
