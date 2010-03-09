package main

import (
	"fmt"
	"rand"
	"time"
	"os"
	"strings"
	"container/list"
	"math"
	"bufio"
	"unicode"
)

const NChildren = 20

// dictionary of words indexed by soundex value
var dict = map[string]*list.List{}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// hamming distance of two strings
func hamming(s1 string, s2 string) int {
	dist := abs(len(s1) - len(s2))
	minlen := min(len(s1), len(s2))

	for i := 0; i < minlen; i++ {
		if s1[i] != s2[i] {
			dist = dist + 1
		}
	}
	return dist
}


// look up a word from the dictionary that is closest to it (but not it if possible)
func closest(s string) string {
	if entries, ok := dict[soundex(s)]; ok {
		candidate := s
		dist := math.MaxInt32
		for entry := range entries.Iter() {
			newdist := hamming(candidate, entry.(string))
			if newdist < dist && entry.(string) != s {
				dist = newdist
				candidate = entry.(string)
			}
		}
		return candidate
	}
	return s
}

var soundex_digits = []int{0, 1, 2, 3, 0, 1, 2, 0, 0, 2, 2, 4, 5, 5, 0, 1, 2, 6, 2, 3, 0, 1, 0, 2, 0, 2}
var soundex_buffer = make([]int, 1024)

// returns the soundex for a string.
func soundex(word string) string {
	i := 0

	for _, ch := range strings.ToUpper(word) {
		if ch >= 'A' && ch <= 'Z' {
			if i == 0 {
				soundex_buffer[i] = ch
				i++
			} else {
				d := soundex_digits[ch-'A']
				if soundex_buffer[0] == 0 || d != soundex_buffer[i-1] {
					soundex_buffer[i] = d + '0'
					i++
				}
			}
		}
	}
	soundex_buffer[i] = 0

	drop0fun := func(ch int) int {
		if ch == '0' {
			return -1
		}
		return ch
	}
	return (strings.Map(drop0fun, string(soundex_buffer)) + strings.Repeat("0", 4))[0:4]
}

// returns the character from string at given index
func charat(s string, i int) int {
	for j, c := range s {
		if i == j {
			return c
		}
	}
	return -1
}

func build_dict() {
	file, err := os.Open("/usr/share/dict/words", 0, os.O_RDONLY)
	if err != nil {
		panic(err.String())
	}

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		if unicode.IsUpper(charat(str, 0)) {
			continue
		}

		sex := soundex(str)
		if _, ok := dict[sex]; !ok {
			dict[sex] = list.New()
		}
		dict[sex].PushBack(str)
	}
}

func child(in chan string, out chan string) {
	// mutates a random word
	words := strings.Fields(<-in)
	rnd := rand.Intn(len(words) - 1)
	words[rnd] = closest(words[rnd])
	out <- strings.Join(words, " ")
}

func main() {
	build_dict()

	rand.Seed(time.Seconds())

	in := make(chan string)
	firstin := in
	var out chan string

	for i := 0; i < NChildren; i++ {
		out = make(chan string)
		go child(in, out)
		in = out
	}

	firstin <- "the first player whispers a phrase or sentence to the next player. Each player successively whispers what that player believes he or she heard to the next. The last player announces the statement to the entire group. Errors typically accumulate in the retellings, so the statement announced by the last player differs significantly, and often amusingly, from the one uttered by the first."
	fmt.Println(<-out)
}
