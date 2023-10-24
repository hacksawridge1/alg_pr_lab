package main

import "fmt"

func main() {
	var inputValue int64
	var counter = 0
	fmt.Scan(&inputValue)

	for inputValue > 1 {
		if inputValue%2 == 0 {
			inputValue = inputValue / 2
		} else {
			inputValue = 3*inputValue + 1
		}

		counter++
	}

	fmt.Println(counter)
}
