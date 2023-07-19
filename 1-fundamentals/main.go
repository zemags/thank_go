package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
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

type result byte

const (
	win  result = 'W'
	draw result = 'D'
	loss result = 'L'
)

type team byte
type match struct {
	first, second team
	result        result
}
type rating map[team]int
type tournament []match

func (trn *tournament) calcRating() rating {
	var r rating = map[team]int{}
	for _, match := range *trn {
		if match.result == win {
			r[match.first] += 3
			r[match.second] += 0
		} else if match.result == loss {
			r[match.second] += 3
			r[match.first] += 0
		} else {
			// draw
			r[match.first] += 1
			r[match.second] += 1
		}
	}
	return r
}

type validator func(s string) bool

func digits(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func letters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func minlen(length int) validator {
	return func(s string) bool {
		return utf8.RuneCountInString(s) >= length
	}
}

func and(funcs ...validator) validator {
	var result bool = true
	return func(s string) bool {
		for _, val := range funcs {
			if !val(s) {
				result = false
			}
		}
		return result
	}
}

func or(funcs ...validator) validator {
	var result bool = false
	return func(s string) bool {
		for _, val := range funcs {
			if val(s) {
				return true
			}
		}
		return result
	}
}

type password struct {
	value string
	validator
}

func (p *password) isValid() bool {
	return p.validator(p.value)
}

func validatePassword() {
	var s string = "hello123"
	validator := or(and(digits, letters), minlen(10))
	p := password{s, validator}
	fmt.Println(p.isValid())
}

type element interface{}
type weightFunc func(element) int
type iterator interface {
	next() bool
	val() element
}
type intIterator struct {
	src   []int
	index int
}

func (i *intIterator) next() bool {
	if i.index+1 < len(i.src) {
		i.index++
		return true
	}
	return false
}

func (i *intIterator) val() element {
	if i.index >= 0 && i.index < len(i.src) {
		return i.src[i.index]
	}
	return nil
}

func newIntIterator(src []int) *intIterator {
	return &intIterator{src, -1}
}

func getMax() {
	nums := []int{1, 2, 4, 5, 3, 7, 2, 3}
	it := newIntIterator(nums)
	weight := func(el element) int {
		return el.(int)
	}
	m := max(it, weight)
	fmt.Println(m)
}

func max(it iterator, weight weightFunc) element {
	var maxEl element
	for it.next() {
		curr := it.val()
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
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

	tour := tournament{
		{first: 'A', second: 'B', result: win},
		{first: 'C', second: 'D', result: draw},
		{first: 'B', second: 'C', result: loss},
	}
	tour.calcRating()
	fmt.Println(tour)
	validatePassword()
	getMax()
}
