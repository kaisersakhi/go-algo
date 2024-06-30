package searches

func linearSearch(list []int, element int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == element {
			return i
		}
	}
	
	return -1
}