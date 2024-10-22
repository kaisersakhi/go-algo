package main

import (
	"fmt"
	"go-algo/algorithms/sorts"
	"go-algo/shared"
)

func main(){
    numbers := shared.NumberArray([]int{1, 3, 7, 5, 99, 0, 4})

    sorts.InsertionSort(numbers)

    fmt.Println(shared.ToString(numbers))
}
