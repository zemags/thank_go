package main

import "fmt"

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	total := Sum(1, 2, 3)
	fmt.Println(total)
}
