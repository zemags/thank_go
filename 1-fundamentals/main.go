package main

import (
	"fmt"
	"math"
	"time"
)

func printTime() {
	var s string = "1h50m"
	d, _ := time.ParseDuration(s)
	fmt.Printf("%v = %v min\n", s, d.Minutes())
}

func euclid() {
	var x1, x2, y1, y2 float64 = 1.0, 2.0, 1.0, 2.0
	d := math.Sqrt(math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2))
	fmt.Println(d)
}

func concateString() {
	var source string = "abc"
	var times int = 3
	var result string
	for i := 0; i < times; i++ {
		result += source
	}

	fmt.Println(result)
}

func lang() {
	var code string = "en"
	var lang string
	if code == "en" {
		lang = "English"
	} else if code == "fr" {
		lang = "French"
	} else {
		lang = "Unknown"
	}

	fmt.Println(lang)
}

func main() {
	printTime()
	euclid()
	concateString()
	lang()
}
