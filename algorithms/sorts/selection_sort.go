package sorts

import "go-algo/shared"

func SelectionSort[T shared.Comparable[T]](list []T) []T {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[j].Compare(list[i]) == shared.LessThan {
				temp := list[j]

				list[j] = list[i]
				list[i] = temp
			}
		}
	}

	return list
}
