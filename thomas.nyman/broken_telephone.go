/*
	Distorts sentence given as argument similary to the children's game of Broken telephone.
	Uses soundex and hamming distance to find similary sounding words from a wordlist given as argument.
	The specified wordlist should be formatted with single words separated by the character '\n'.
	The specified sentence should be surrounded by quotes and individual words separated by spaces.
	Punctuation might be discarded and no word filters are used so usage of this program might note be
	entirely family friendly depending on the word list used.

	Example usage: ./broken_phone /usr/share/dict/words "a quick brown fox jumped over the lazy dog"
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"rand"
	"time"
	"strings"
	"io/ioutil"
	"container/vector"
)

func main() {
	const passes = 20

	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Fprintln(os.Stderr, "USAGE:", "broken_telephone", "WORDLIST", "SENTENCE")
		os.Exit(1)
	}

	m, err := readWordList(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "broken_telephone: error reading from %s: %s\n", flag.Arg(0), err)
		os.Exit(1)
	}

	rand.Seed(time.Nanoseconds())

	c := make(chan string)
	for i := 0; i <= passes; i++ {
		go brokenPhone(c, m)
	}
	c <- flag.Arg(1)
	s := <-c
	fmt.Println(s)
}

// distorts random word in string received from c and sends it back
func brokenPhone(c chan string, m *map[string]*vector.StringVector) {
	s := <-c
	tmp := strings.Split(s, " ", 0)
	i := rand.Int() % len(tmp)
	tmp[i] = lookUpWord(tmp[i], m)
	s = strings.Join(tmp, " ")
	c <- s
}

// looks up a similar word to s from m
func lookUpWord(s string, m *map[string]*vector.StringVector) string {
	if words, ok := (*m)[soundex(s)]; ok {
		for i := 0; i < len(*words); i++ {
			if hammingDistance(s, words.At(i)) > 0 { // don't allow words of different length
				s = words.At(i)
				break
			}
		}
	}
	return s
}

// reads a word list from a file and returns the list m
func readWordList(filename string) (*map[string]*vector.StringVector, os.Error) {
	// Read file contents
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	m := make(map[string]*vector.StringVector)

	// Parse words from file contents and store them in m
	s := ""
	for i := 0; i < len(contents); i++ {
		if string(contents[i]) != "\n" {
			s += string(contents[i])
		} else {
			soundex := soundex(s)
			switch _, ok := m[soundex]; ok {
			case false:
				tmp := make(vector.StringVector, 0)
				m[soundex] = &tmp
				fallthrough
			case true:
				words := m[soundex]
				words.Push(s)
				s = ""
			}
		}
	}
	return &m, err
}

// return hamming distance of s and t or -1 if string length differs
func hammingDistance(s string, t string) int {
	if len(s) != len(t) {
		return -1
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			sum++
		}
	}
	return sum
}

// return soundex code for s
func soundex(s string) string {
	s = strings.ToUpper(s)
	if tmp := strings.Split(s, "", 2); len(tmp) == 1 {
		s = tmp[0] + "000"
	} else {
		s = tmp[0] + strings.Map(mapSoundex, tmp[1])
		s = collapseSoundex(s)
		s = padSoundex(s)
	}
	return s
}

// maps rune to corresponding integer according to soundex spec
func mapSoundex(rune int) int {
	switch rune {
	case 'A', 'E', 'I', 'O', 'U', 'H', 'W', 'Y':
		return '0'
	case 'B', 'F', 'P', 'V':
		return '1'
	case 'C', 'G', 'J', 'K', 'Q', 'S', 'X', 'Z':
		return '2'
	case 'D', 'T':
		return '3'
	case 'L':
		return '4'
	case 'M', 'N':
		return '5'
	case 'R':
		return '6'
	}
	return '0'
}

// removes zeroes and collapses adjacent identical digits into a single digit
func collapseSoundex(s string) string {
	t := ""
	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			t += string(s[i])
		}
	}

	s = ""
	for i := 0; i < len(t); i++ {
		if i != len(t)-1 && t[i] == t[i+1] {
			continue
		}
		s += string(t[i])
	}
	return s
}

// pads or cuts s if needed to a valid soundex code
func padSoundex(s string) string {
	switch {
	case len(s) > 4:
		return s[0:4]
	case len(s) < 4:
		return s + strings.Repeat(string('0'), 4-len(s))
	}
	return s
}
