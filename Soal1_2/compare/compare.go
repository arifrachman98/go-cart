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

func oneEditReplace(s1, s2 string) bool {
	foundDiff := false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if foundDiff {
				return false
			}
			foundDiff = true
		}
	}
	return true
}

