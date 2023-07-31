package main

import (
	"fmt"
	"time"
)

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// начало решения
	result := make([]any, len(funcs))
	ch := make(chan any, len(funcs))
	order := make(chan int, len(funcs))

	numWorkers := len(funcs)

	for i := 0; i < numWorkers; i++ {
		f := funcs[i]
		go func(index int) {
			ch <- f()
			order <- index
		}(i)
	}

	// gather results
	for i := 0; i < numWorkers; i++ {
		index := <-order
		result[index] = <-ch
	}

	// конец решения
	return result
}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	//funcs := []func() any{squared(2), squared(3), squared(4)}
	//[5 2 1 3 4]
	funcs := []func() any{squared(5), squared(2), squared(1), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
