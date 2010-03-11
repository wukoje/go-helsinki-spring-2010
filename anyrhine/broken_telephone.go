/*
 *	the database is read from broken_telephone.words which i is /lib/words
 *	from plan9.
 */
package main


import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"sort"
	"rand"
)

type Sdxmap struct {
	sdx string
	s   string
}

type SdxArray []Sdxmap

var sdxtab SdxArray

func (s SdxArray) Len() int { return len(s) }

func (s SdxArray) Less(i, j int) bool { return s[i].sdx < s[j].sdx }

func (s SdxArray) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s SdxArray) findsimilar(sdx string) int {
	m, i, j := 0, 0, len(s)
	for i < j {
		m = int(uint(i+j) / 2)
		if s[m].sdx < sdx {
			i = m + 1
		} else if s[m].sdx > sdx {
			j = m
		} else {
			break
		}
	}
	return m
}

var smap = map[int]int{'b': '1', 'p': '1', 'v': '1',
	'c': '2', 'g': '2', 'j': '2', 'k': '2', 'q': '2', 's': '2', 'x': '2', 'z': '2',
	'd': '3', 't': '3', 'l': '4', 'm': '5', 'n': '5', 'r': '6'}

func soundex(s string) string {
	code := make([]int, 4)
	i := 0
	for _, c := range strings.ToLower(s) {
		if i == 0 {
			code[i] = c
			i++
		} else {
			if tcode, ok := smap[c]; ok && tcode != code[i-1] {
				code[i] = tcode
				i++
			}
		}
		if i >= len(code) {
			break
		}
	}
	for ; i < len(code); i++ {
		code[i] = '0'
	}
	return string(code)
}

func keyb(out chan string) {
	r := bufio.NewReader(os.Stdin)
	for {
		var s string
		var err os.Error
		if s, err = r.ReadString('\n'); err != nil {
			break
		}
		for _, t := range strings.Fields(s) {
			out <- t
		}
		out <- "\n"
	}
}

func disp(in chan string) {
	for {
		s := <-in
		if s != "\n" {
			fmt.Printf("%s ", s)
		} else {
			fmt.Print("\n")
		}
	}
}

func sentraalisantra(in chan string, out chan string) {
	var s, sdx string
	for {
		s = <-in
		if s != "\n" {
			sdx = soundex(s)
			i := sdxtab.findsimilar(sdx)
			if sdxtab[i].sdx == sdx {
				for i > 0 && sdxtab[i-1].sdx == sdx {
					i--
				}
				j := i
				for j < len(sdxtab) && sdxtab[j].sdx == sdx {
					j++
				}
				i = i + (rand.Int() % (j - i))
			}
			s = sdxtab[i].s
		}
		out <- s
	}
}

func main() {
	sdxtab = make(SdxArray, 0, 10000)
	if dbfile, err := os.Open("broken_telephone.words", os.O_RDONLY, 0); err == nil {
		r := bufio.NewReader(dbfile)
		for {
			if s, err := r.ReadString('\n'); err == nil {
				if len(sdxtab) == cap(sdxtab) {
					nsdxtab := make(SdxArray, len(sdxtab), cap(sdxtab)*2)
					for i, sdx := range sdxtab {
						nsdxtab[i] = sdx
					}
					sdxtab = nsdxtab
				}
				sdxtab = sdxtab[0 : len(sdxtab)+1]
				sdx := &sdxtab[len(sdxtab)-1]
				sdx.s = s[0 : len(s)-1]
				sdx.sdx = soundex(s)
			} else {
				break
			}
		}
	} else {
		fmt.Printf("open:%s\n", err)
		os.Exit(1)
	}
	sort.Sort(sdxtab)

	a := make(chan string)
	first := a
	for i := 0; i < 20; i++ {
		b := make(chan string)
		go sentraalisantra(a, b)
		a = b
	}
	last := a
	go disp(last)
	keyb(first)
}
