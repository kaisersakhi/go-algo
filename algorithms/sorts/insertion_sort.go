package sorts

import "go-algo/shared"

func InsertionSort[T shared.Comparable[T]](list []T) []T{
	for i := 0; i < len(list); i++ {
		j := i 
		for j > 0 && list[(j - 1)].Compare(list[j]) == shared.GreaterThan {		
			temp := list[j]
			list[j] = list[j - 1]
			list[j - 1] = temp
			j--
		}
	}

	return list
}
