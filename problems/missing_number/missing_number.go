package missingnumber

import "sort"

func FindMissingNumber(arr []int) int {
	if len(arr) == 1 {
		if arr[0] == 0 {
			return 1
		} else {
			return 0
		}
	}

	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] > i {
			return arr[i] - 1
		}
	}

	return len(arr)
}
