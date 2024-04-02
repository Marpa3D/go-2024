// Инвестиционный калькулятор CLI
package main

import (
	"fmt"
	"math"
)

func main() {
	investAmount := 1000
	expectReturnRate := 5.5
	yars := 10
	futureValue := float64(investAmount) * math.Pow(1+expectReturnRate/100, float64(yars))

	fmt.Printf("%.2f", futureValue)
}
