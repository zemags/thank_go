package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

func makePool(n int, handler func(int, string)) (func(string), func()) {
	pool := make(chan int, n)
	for idx := 1; idx <= n; idx++ {
		pool <- idx
	}

	handle := func(phrase string) {
		id := <-pool
		go func() {
			handler(id, phrase)
			pool <- id
		}()
	}

	wait := func() {
		for i := 0; i < n; i++ {
			<-pool
		}
	}

	return handle, wait
}

func main() {
	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}

	handle, wait := makePool(2, say)
	for _, phrase := range phrases {
		handle(phrase)
	}
	wait()
}
