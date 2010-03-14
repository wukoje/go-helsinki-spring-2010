/*
 KNOWN BUGS:
 - This software completely ignores the fact that an unicode glyph can be more
   than one byte wide.
 OTHER NOTICES:
 - Needs the dictionary file in /usr/share/dict/british-english
*/
package main

import (
	"fmt"
	"io/ioutil"
	"container/list"
	"os"
	"rand"
	"strings"
)

const (
	dictFile = "/usr/share/dict/british-english"
	numTelephones = 20
	minHammingDist = 1
	maxHammingDist = 3
)

var soundexMap = map[string]string {
	"b": "1",
	"f": "1",
	"p": "1",
	"v": "1",
	"c": "2",
	"g": "2",
	"j": "2",
	"k": "2",
	"q": "2",
	"s": "2",
	"x": "2",
	"z": "2",
	"d": "3",
	"t": "3",
	"l": "4",
	"m": "5",
	"n": "5",
	"r": "6",
}

func readDictionary(dictFile string) ([]string, os.Error) {
	file, err := ioutil.ReadFile(dictFile)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(file), "\n", 0), nil
}

func toSoundex(s string) (soundex string) {
	soundex += s[0:1] // FIXME: not all glyphs are 1 byte long.
	s = s[1:len(s)] // Skip first letter
	lastDigit := ""
	for _, ch := range s {
		c := string(ch)
		if v, present := soundexMap[c]; present {
			if lastDigit != "" || v != lastDigit {
				soundex += v
				lastDigit = v
			}
		}
	}
	for len(soundex) < 4 {
		soundex += "0"
	}
	soundex = soundex[0:4]
	return
}

// Ignores empty strings in dict.
func buildSoundexMap(dict []string) map[string](*list.List) {
	smap := make(map[string](*list.List))
	for _, s := range dict {
		if (s == "") {
			continue
		}
		soundex := toSoundex(s)
		if _, present := smap[soundex]; !present {
			smap[soundex] = list.New()
		}
		smap[soundex].PushBack(s)
	}
	return smap
}

func hamDist(s1, s2 string) (dist int, countable bool) {
	if len(s1) != len(s2) {
		return 0, false
	}
	for i := 0; i < len(s1); i++ {
		if (s1[i] != s2[i]) {
			dist++
		}
	}
	return dist, true
}

func pickWord(original string, words *list.List) (word string, found bool) {
	elem := words.Front()
	for elem != nil {
		word := elem.Value.(string)
		dist, ok := hamDist(original, word)
		if ok && minHammingDist <= dist && dist <= maxHammingDist {
			return word, true
		}
		elem = elem.Next()
	}
	return "", false
}

func telephone(in, out chan string, smap map[string](*list.List)) {
	msg := <-in
	words := strings.Fields(msg)
	n := rand.Intn(len(words)) // Pick a word at random
	if choises, present := smap[toSoundex(words[n])]; present {
		if word, found := pickWord(words[n], choises); found {
			words[n] = word
		}
	}
	out <- strings.Join(words, " ")
}

func main() {
	dict, err := readDictionary(dictFile)
	if (err != nil) {
		fmt.Println(err.String())
		return
	}
	smap := buildSoundexMap(dict)
	_, nsec, _ := os.Time()
	rand.Seed(nsec)
	msg := "If I was in charge we'd never see grass between October and May"
	input := make(chan string)
	var in, out chan string
	in = input
	for i := 0; i < numTelephones; i++ {
		out = make(chan string)
		go telephone(in, out, smap)
		in = out
	}
	fmt.Println("Original message: ", msg)
	input <- msg
	fmt.Println("Result: ", <-out)
}
