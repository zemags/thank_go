package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
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

type label string

type command label

var (
	eat  = command("eat")
	take = command("take")
	talk = command("talk to")
)

type thing struct {
	name    label
	actions map[command]string
}

func (t thing) supports(action command) bool {
	_, ok := t.actions[action]
	return ok
}

func (t thing) String() string {
	return string(t.name)
}

var (
	apple = thing{"apple", map[command]string{
		eat:  "mmm, delicious!",
		take: "you have an apple now",
	}}
	bob = thing{"bob", map[command]string{
		talk: "Bob says hello",
	}}
	coin = thing{"coin", map[command]string{
		take: "you have a coin now",
	}}
	mirror = thing{"mirror", map[command]string{
		take: "you have a mirror now",
		talk: "mirror does not answer",
	}}
	mushroom = thing{"mushroom", map[command]string{
		eat:  "tastes funny",
		take: "you have a mushroom now",
	}}
)

type step struct {
	cmd command
	obj thing
}

func (s step) isValid() bool {
	return s.obj.supports(s.cmd)
}

func (s step) String() string {
	return fmt.Sprintf("%s %s", s.cmd, s.obj)
}

type player struct {
	nEaten    int
	nDialogs  int
	inventory []thing
}

func (p *player) has(obj thing) bool {
	for _, got := range p.inventory {
		if got.name == obj.name {
			return true
		}
	}
	return false
}

type objectLimitExceededError struct {
	obj   thing
	limit int
}

func (err objectLimitExceededError) Error() string {
	return fmt.Sprintf("don't be greedy, %d %s is enough", err.limit, err.obj)
}

type commandLimitExceededError struct {
	cmd   command
	limit int
}

func (err commandLimitExceededError) Error() string {
	return fmt.Sprintf("%s less", err.cmd)
}

type notEnoughObjectsError struct {
	obj thing
}

func (err notEnoughObjectsError) Error() string {
	return fmt.Sprintf("be careful with scarce %ss", err.obj)
}

type invalidStepError struct {
	st step
}

func (err invalidStepError) Error() string {
	return fmt.Sprintf("things like '%s' are impossible", err.st)
}

type gameOverError struct {
	nSteps int
	cause  error
}

func (err gameOverError) Error() string {
	return fmt.Sprintf("%s", err.cause)
}

func (err gameOverError) Unwrap() error {
	return err.cause
}

func (p *player) do(cmd command, obj thing) error {
	switch cmd {
	case eat:
		if p.nEaten > 1 {
			return commandLimitExceededError{cmd, p.nEaten}
		}
		p.nEaten++
	case take:
		if p.has(obj) {
			return objectLimitExceededError{obj, 1}
		}
		p.inventory = append(p.inventory, obj)
	case talk:
		if p.nDialogs > 0 {
			return commandLimitExceededError{cmd, p.nDialogs}
		}
		p.nDialogs++
	}
	return nil
}

func newPlayer() *player {
	return &player{inventory: []thing{}}
}

type game struct {
	player *player
	things map[label]int
	nSteps int
}

func (g *game) has(obj thing) bool {
	count := g.things[obj.name]
	return count > 0
}

func (g *game) execute(st step) error {
	if !st.isValid() {
		err := invalidStepError{st}
		return gameOverError{g.nSteps, err}
	}

	if st.cmd == take || st.cmd == eat {
		if !g.has(st.obj) {
			err := notEnoughObjectsError{st.obj}
			return gameOverError{g.nSteps, err}
		}
		g.things[st.obj.name]--
	}

	if err := g.player.do(st.cmd, st.obj); err != nil {
		return gameOverError{g.nSteps, err}
	}

	g.nSteps++
	return nil
}

func newGame() *game {
	p := newPlayer()
	things := map[label]int{
		apple.name:    2,
		coin.name:     3,
		mirror.name:   1,
		mushroom.name: 1,
	}
	return &game{p, things, 0}
}

func giveAdvice(err error) string {
	return fmt.Sprintf("%s", err)
}

func playGame() {
	gm := newGame()
	steps := []step{
		{eat, apple},
		{eat, apple},
		{eat, mushroom},
	}

	for _, st := range steps {
		if err := tryStep(gm, st); err != nil {
			fmt.Println(err)
			giveAdvice(err)
			os.Exit(1)
		}
	}
	fmt.Println("You win!")
}

func tryStep(gm *game, st step) error {
	fmt.Printf("trying to %s %s... ", st.cmd, st.obj.name)
	if err := gm.execute(st); err != nil {
		fmt.Println("FAIL")
		return err
	}
	fmt.Println("OK")
	return nil
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
	playGame()
}
