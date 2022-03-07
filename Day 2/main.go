package main 

import (
	"fmt"
	"math"
)

func solve(mealCost float64, tipPercent int, taxPercent int ) {
	tipCost:= float64(tipPercent)/100 * mealCost
	taxCost:= float64(taxPercent)/100 * mealCost

	totalCost:= math.Round(mealCost + tipCost + taxCost)
	fmt.Println(totalCost)
}

func main() {
	solve(12.0, 20, 8)
}