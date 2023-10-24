package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		p      float64 = 0.5 * 365
		oak    float64 = p / 20
		poplar float64 = p / 32
	)

	fmt.Println(p, int(math.Ceil(poplar)), int(math.Ceil(oak)))
}
