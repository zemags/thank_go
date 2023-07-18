package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
	"unicode"
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

func sliceString() {
	var s string = "apple"
	var width int = 3

	bytes := []byte(s)
	sliced := bytes[0:width]

	fmt.Printf("%v\n", string(sliced))

}

func countDigit() {

	var number int = 1234233

	counter := make(map[int]int)

	for number > 0 {
		digit := number % 10
		counter[digit]++
		number /= 10
	}

	fmt.Println(counter, number)
}

func firstLetters() {
	phrase := "apple-orange 3f carrot"
	var abbr string
	phraseSplit := strings.Fields(phrase)
	for _, word := range phraseSplit {
		runeWordLst := []rune(word)
		runeWordLst[0] = unicode.ToUpper(runeWordLst[0])
		if unicode.IsLetter(runeWordLst[0]) {
			abbr = abbr + string(runeWordLst[0])
		}
	}
	fmt.Println(abbr)
	// AC
}

func filter(predicate func(int) bool, iterable []int) []int {
	var res []int
	for _, i := range iterable {
		if predicate(i) {
			res = append(res, i)
		}
	}
	return res
}

func filterEven() {
	var lst []int = []int{1, 2, 3, 4, 5, 6}
	res := filter(func(i int) bool {
		if i%2 == 0 {
			return true
		} else {
			return false
		}
	}, lst)
	fmt.Println(res)
}

func shuffle(nums []int) {
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
}

func main() {
	printTime()
	euclid()
	concateString()
	lang()
	sliceString()
	countDigit()
	firstLetters()
	filterEven()
	lst := []int{1, 2, 3, 4, 5}
	shuffle(lst)
	fmt.Println(lst)
}
