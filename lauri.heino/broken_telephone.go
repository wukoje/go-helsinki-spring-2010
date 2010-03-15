package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"container/list"
	"math"
	"rand"
	"time"
)

var soundexDigits = [...]int{0, 1, 2, 3, 0, 1, 2, 0, 0, 2, 2, 4, 5, 5, 0, 1, 2, 6, 2, 3, 0, 1, 0, 2, 0, 2}

// soundex
func soundex(s string) string {
	var code = make([]int, 4)

	i := 0

	for _, c := range strings.ToUpper(s) {
		if c >= 'A' && c <= 'Z' {
			if i == 0 {
				code[0] = c
				i++
			} else {
				digit := '0' + soundexDigits[(c-'A')]
				if digit != code[i-1] && digit != '0' {
					code[i] = digit
					i++
				}
			}
		}

		if i > len(code)-1 {
			break
		}

	}

	for i < len(code) {
		code[i] = '0'
		i++
	}

	return string(code)
}

// hamming returns the hamming distance of the strings s1 and s2
func hamming(s1, s2 string) int {
	dist := 0

	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			dist++
		}
	}

	if len(s1) < len(s2) {
		dist = dist + (len(s2) - len(s1))
	} else if len(s1) > len(s2) {
		dist = dist + (len(s1) - len(s2))
	}

	return dist
}

// substitute retruns the closest substitute for s
func substitute(s string, m map[string]*list.List) string {
	subst := s
	minDist := math.MaxInt32
	sx := soundex(s)

	// no substitutes available
	if _, ok := m[sx]; !ok {
		return s
	}

	words := m[sx]

	for elem := words.Front(); elem != nil; elem = elem.Next() {
		dist := hamming(strings.ToLower(s), strings.ToLower(elem.Value.(string)))
		if dist < minDist && dist > 0 {
			subst = elem.Value.(string)
			minDist = hamming(s, subst)
		}
	}

	return subst
}

// initDict builds the words-by-soundex map
func initDict() map[string]*list.List {
	wordsBySoundex := map[string]*list.List{}

	file, err := os.Open("/usr/share/dict/words", 0, os.O_RDONLY)

	if err != nil {
		panic(err.String())
	}

	io := bufio.NewReader(file)
	for {
		word, err := io.ReadString('\n')

		if err != nil {
			break
		}

		sx := soundex(word)

		if _, ok := wordsBySoundex[sx]; !ok {
			wordsBySoundex[sx] = list.New()
		}

		wordsBySoundex[sx].PushBack(word[0 : len(word)-1])
	}

	return wordsBySoundex
}

// printSoundexMap prints a words-by-soundex map
func printSoundexMap(m map[string]*list.List) {
	for key, _ := range m {
		fmt.Printf("\n")
		fmt.Printf("Soundex: %v\n", key)
		for elem := m[key].Front(); elem != nil; elem = elem.Next() {
			fmt.Printf("%v ", elem.Value)
		}
		fmt.Printf("\n")
	}
}

// passMessage simulates a child
func passMessage(m map[string]*list.List, in, out chan string) {
	message := strings.Fields(<-in)
	changeAt := rand.Intn(len(message))
	message[changeAt] = substitute(message[changeAt], m)
	out <- strings.Join(message, " ")
}

func main() {
	fmt.Printf("*** Broken Telephone ***\n")

	originalMsg := "Implement the children's game of Broken Telephone using 20 goroutines that pass the message using channels. Each goroutine should change one word in the message to another valid word. The recommended way is to use soundex and the hamming distance for finding a substitute word."

	originalSrc := make(chan string)
	in := originalSrc
	var out chan string

	m := initDict()
	rand.Seed(time.Seconds())

	for i := 0; i < 20; i++ {
		out = make(chan string)
		go passMessage(m, in, out)
		in = out
	}

	originalSrc <- originalMsg
	fmt.Printf("\nThe first child was told:\n%v\n\nBut the last child understood:\n%v\n\n", originalMsg, <-out)
}
