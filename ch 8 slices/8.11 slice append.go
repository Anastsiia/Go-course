package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	/* max := 0
	for i := 0; i < len(costs); i++ {
		if costs[i].day > max {
			max = costs[i].day
		}
	}
	total := make([]float64, max+1)
	fmt.Println(max)
	fmt.Println(total)
	for i := 0; i < len(costs); i++ {
		total[costs[i].day] += costs[i].value
	}
	 return total
	*/
	// same, but with "append":

	total := []float64{}
	for i := 0; i < len(costs); i++ {
		for costs[i].day >= len(total) {
			total = append(total, 0.0)
		}
		total[costs[i].day] += costs[i].value
	}
	return total
}

// dont edit below this line

func test(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	test([]cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})
}
