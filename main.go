package main

import (
    "fmt"
    "go-algo/algorithms/sorts"
)

func main(){
    numbers := NumberArray([]int{1, 3, 7, 5, 99, 0, 4})

    sorts.BubbleSort(numbers)

    fmt.Println(ToString(numbers))
}