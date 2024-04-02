// Инвестиционный калькулятор CLI
package main

import (
	"fmt"
	"math"
)

const inflation = 2.5

func main() {
	investAmount := 1000.0
	expectReturnRate := 5.5
	yars := 10.0
	futureValue := investAmount * math.Pow(1+expectReturnRate/100, yars)
	realInvest := futureValue / math.Pow(1+inflation/100, yars)
	fmt.Printf("Реальный доход с инвестиций = %.2f", realInvest)
}
