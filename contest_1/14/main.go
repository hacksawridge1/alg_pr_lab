package main

import (
	"fmt"
)

func main() {
	var rows, cols int
	fmt.Scan(&rows, &cols)

	for i := 1; i <= (rows + 2); i++ {
		if i == 1 {
			for l := 1; l <= (cols + 1); l++ {
				if l == 1 {
					fmt.Print("    |")
				} else {
					fmt.Printf("%4d", l-1)
				}
			}
		} else if i == 2 {
			for k := 1; k <= (cols*2 + 2); k++ {
				if k == 1 {
					fmt.Print("   ")
				} else {
					fmt.Print("--")
				}
			}
		} else {
			for j := 1; j <= (cols + 1); j++ {
				if j == 1 {
					fmt.Printf("%4d", i-2)
					fmt.Print("|")
				} else {
					fmt.Printf("%4d", (i-2)*(j-1))
				}
			}
		}
		fmt.Print(" ")
		fmt.Print("\n")
	}
}
