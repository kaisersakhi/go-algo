package sorts

import "go-algo/shared"

func BubbleSort[T shared.Comparable[T]](list []T) []T{
	for i := 0; i < len(list); i ++ {
		for j := 0; j < len(list) - i  - 1; j ++ {
			if list[j].Compare(list[j + 1]) == shared.GreaterThan {
				temp := list[j]

				list[j] = list[j + 1]
				list[j + 1] = temp
			}
		}
	}
	return list
}