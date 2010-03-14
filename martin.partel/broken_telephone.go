package main

/*
 * I did this perhaps a bit differently than described in the exercise.
 * I made each of the 20 goroutines handle one word at a time,
 * change it with a certain probability, and pass it on.
 * This makes the program work nicely as a stream processor when the length
 * of the entire message is not known.
 *
 * When I want to change a word, I find its levenshtein to every other
 * word in the dictionary. This is quite slow, but fast enough
 * with the default dictionary of only about 100 000 words.
 */

import (
	"fmt"
	"os"
	"bufio"
	"container/vector"
	"strings"
	"unicode"
	"rand"
	"time"
)

func levenshtein(s string, t string, max int) int {
	// Based on the dynamic programming approach presented in Wikipedia
	// Slightly optimized to stop immediately if the distance gets too big

	m := len(s)
	n := len(t)

	prev := make([]int, m+1)
	for i := 0; i <= m; i++ {
		prev[i] = i
	}

	next := make([]int, m+1)

	// For each iteration of j,
	// prev[i] is the distance between s[0:i] and t[0:j-1]
	for j := 1; j <= n; j++ {
		next[0] = j

		// For each iteration of i,
		// next[i] is set to the distance between s[0:i] and t[0:j]
		minInRow := max
		for i := 1; i <= m; i++ {
			if s[i-1] == t[j-1] {
				next[i] = prev[i-1]
			} else {
				// Three possibilities:
				// 1. If s[0:i-1] -> t[0:j] takes d steps, we can
				//    delete the last char in s[0:i] to get s[0:i-1]
				//    and reach t[0:j] in d+1 steps
				// 2. If s[0:i] -> t[0:j-1] takes d steps, we can
				//    append s[i] to t[0:j-1]
				//    and reach t[0:j] in d+1 steps
				// 3. If s[0:i-1] -> t[0:j-1] takes d steps, we can
				//    replace s[i] with t[j]
				//    and reach t[0:j] in d+1 steps
				next[i] = min3(next[i-1], prev[i], prev[i-1]) + 1
			}

			if next[i] < minInRow {
				minInRow = next[i]
			}
		}

		if minInRow >= max {
			return max
		}

		t := prev
		prev = next
		next = t
	}

	return prev[m]
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

type dictionary struct {
	words []string
}

func main() {
	words, err := readDictionary("/usr/share/dict/words")
	if err != nil {
		bail(err)
	}

	dict := &dictionary{words}

	wordsIn, err := startWordReader(os.Stdin)

	if err != nil {
		bail(err)
	}

	out := wireNodes(wordsIn, dict, 20, 0.01)

	for s := <-out; s != nil; s = <-out {
		fmt.Print(*s, " ")
	}
}

func wireNodes(firstChan <-chan *string, dict *dictionary, count int, breakage float) <-chan *string {
	bufSize := count * 2 // Works with 0, but won't scale as nicely

	nextInputChan := firstChan
	nextOutputChan := make(chan *string, bufSize)

	for i := 0; i < count; i++ {
		go telephoneNode(i, nextInputChan, nextOutputChan, dict, breakage)

		nextInputChan = nextOutputChan
		nextOutputChan = make(chan *string, bufSize)
	}
	lastChan := nextInputChan

	return lastChan
}

func telephoneNode(num int, in <-chan *string, out chan<- *string, dict *dictionary, breakage float) {
	rng := rand.New(rand.NewSource(time.Seconds() + int64(num)))
	for {
		s := <-in
		if s == nil {
			out <- nil
			break
		}

		if rng.Float() < breakage {
			s = findCloseWord(dict, s)
		}

		out <- s
	}
}

func findCloseWord(dict *dictionary, s *string) *string {
	words := make(vector.StringVector, 0, 1000)

	for _, word := range dict.words {
		if closeEnough(&word, s) {
			words.Push(word)
		}
	}

	if words.Len() > 0 {
		return &words[rand.Intn(words.Len())]
	}
	return s
}

func closeEnough(a *string, b *string) bool { return levenshtein(*a, *b, 3) < 3 }

func startWordReader(f *os.File) (<-chan *string, os.Error) {
	r := bufio.NewReader(f)

	runeChan := make(chan int)
	strChan := make(chan *string)

	go runeStreamer(r, runeChan)
	go wordTokenizer(runeChan, strChan)

	return strChan, nil
}

func runeStreamer(r *bufio.Reader, out chan<- int) {
	for {
		rune, _, err := r.ReadRune()
		if err != nil {
			out <- -1
			break
		}

		out <- rune
	}
}

func wordTokenizer(in <-chan int, out chan<- *string) {
	buf := make(vector.IntVector, 0, 100)
	for {
		c := <-in
		if c <= 0 {
			break
		}

		if !unicode.IsSpace(c) {
			buf.Push(c)
		} else {
			if len(buf) > 0 {
				s := string(buf)
				out <- &s
				buf.Resize(0, 100)
			}
		}
	}

	out <- nil
}


func readDictionary(filename string) ([]string, os.Error) {
	f, err := os.Open(filename, os.O_RDONLY, 0000)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	words := make(vector.StringVector, 0)

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == os.EOF {
			break
		} else if err != nil {
			return words, err
		}

		line = line[0 : len(line)-1] // remove delimiter
		words.Push(strings.ToLower(line))
	}
	return words, nil
}

func bail(err os.Error) {
	fmt.Println(err.String())
	os.Exit(1)
}
