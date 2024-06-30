package main

import "fmt"
import "go-algo/algorithms/sorts"

func main(){
	x := sorts.BubbleSort([]int{1})
	fmt.Printf(string(x[0]))
}