package compare

import "fmt"

func oneEditInsert(s1, s2 string) bool {
	index1 := 0
	index2 := 0
	for index2 < len(s2) && index1 < len(s1) {
		if s1[index1] != s2[index2] {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}

	return true
}