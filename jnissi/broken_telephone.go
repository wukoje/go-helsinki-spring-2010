package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"rand"
	"regexp"
	"strings"
	"time"
)

var dict = map[string]*list.List{}

func soundex(word string) string {
	word = strings.ToLower(word)
	firstLetter := word[0:1]
	word = word[1:]
	//I need a better regexp lib. Could make this that much faster.
	searches := []string{
		"[bfpv]", "[cgjkqsxz]", "[dt]", "[l]", "[mn]", "[r]",
		"1[hw]1", "2[hw]2", "3[hw]3", "4[hw]4", "5[hw]5", "6[hw]6",
		"11+", "22+", "33+", "44+", "55+", "66+",
		"[^0-9]",
	}
	replaces := []string{
		"1", "2", "3", "4", "5", "6",
		"1h", "2h", "3h", "4h", "5h", "6h",
		"1", "2", "3", "4", "5", "6",
		"",
	}

	for i, r := range searches {
		re := regexp.MustCompile(r)
		word = re.ReplaceAllString(word, replaces[i])
	}

	return (firstLetter + word + strings.Repeat("0", 3))[0:4]
}

func hammingDistance(s1 string, s2 string) int {
	i, ret := len(s1)-1, 0
	if len(s1) != len(s2) {
		return math.MaxInt32
	}
	for ; i >= 0; i-- {
		if s1[i:i+1] != s2[i:i+1] {
			ret++
		}
	}
	return ret
}

func findReplacement(word string) string {
	ret := word
	s := soundex(word)
	if l, ok := dict[s]; ok {
		dist := math.MaxInt32
		for s := range l.Iter() {
			n := hammingDistance(word, s.(string))
			if n < dist && s.(string) != word {
				dist = n
				ret = s.(string)
			}
		}
	}
	return ret
}

func relay(in chan string, out chan string) {
	w := strings.Fields(<-in)
	r := rand.Intn(len(w) - 1)
	w[r] = findReplacement(w[r])
	out <- strings.Join(w, " ")
}

func main() {
	rand.Seed(time.Nanoseconds())
	msg := "Hate hate hate hatred for all one and all. No matter what you believe. Don't believe in you and that's true. Yeah. We hate everyone We hate everyone."
	file, err := os.Open("/usr/share/dict/words", os.O_RDONLY, 0)
	if err != nil {
		panic(err.String())
	}

	bufReader := bufio.NewReader(file)
	for {
		w, e := bufReader.ReadString('\n')
		if e != nil {
			break
		}
		if w[0] < 'a' || w[0] > 'z' {
			continue
		}
		if len(w) > 1 {
			w = w[0 : len(w)-1]
			soundex := soundex(w)
			if _, ok := dict[soundex]; !ok {
				dict[soundex] = list.New()
			}
			dict[soundex].PushBack(w)
		}
	}
	in := make(chan string)
	start := in
	var out chan string
	for i := 0; i < 20; i++ {
		out = make(chan string)
		go relay(in, out)
		in = out
	}
	start <- msg
	relayed := <-out
	fmt.Printf(" original:\n%s\n relayed:\n%s\n", msg, relayed)
}
