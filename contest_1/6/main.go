package main

import "fmt"

func main() {
	var inputValue, fivet, onet, fiveh, twoh, oneh int64
	fmt.Scan(&inputValue)

	for inputValue > 1 {
		if inputValue >= 5000 {
			fivet = inputValue / 5000
			inputValue = inputValue % 5000
		} else if inputValue >= 1000 {
			onet = inputValue / 1000
			inputValue = inputValue % 1000
		} else if inputValue >= 500 {
			fiveh = inputValue / 500
			inputValue = inputValue % 500
		} else if inputValue >= 200 {
			twoh = inputValue / 200
			inputValue = inputValue % 200
		} else if inputValue >= 100 {
			oneh = inputValue / 100
			inputValue = inputValue % 100
		}
	}

	fmt.Println(fivet, onet, fiveh, twoh, oneh)
}
