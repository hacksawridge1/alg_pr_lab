package main

import (
	"fmt"
)

func main() {
	var array []float64
	var num float64
	var a int
	var count int
	fmt.Scan(&count)
	for a < count {
		fmt.Scan(&num)
		array = append(array, num)
		a++
	}
	fmt.Printf("%.15f ", array[0])
	for i := 1; i < len(array)-1; i++ {
		fmt.Printf("%.15f ", (array[i-1]+array[i]+array[i+1])/3)
	}
	fmt.Printf("%.15f", array[len(array)-1])
}
