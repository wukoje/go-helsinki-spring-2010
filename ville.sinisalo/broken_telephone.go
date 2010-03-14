package main

import (
	"fmt"
	"strings"
	"os"
	"container/vector"
	"bufio"
	"rand"
	"time"
	"sort"
	"flag"
	"unicode"
)

// Returns soundex value of string s. bool will be false if 
// the word contained any non alphabetic characters.
func soundex(s string) (string, bool) {
	ok := true
	if len(s) == 0 {
		return " 000", false
	}
	s = strings.ToUpper(s)
	ret := make([]byte, 4)
	ret[0] = s[0]
	if !('a' <= ret[0] && ret[0] <= 'z') &&
	   !('A' <= ret[0] && ret[0] <= 'Z') {
		ok = false
	}
	c,i := 1,1
	for ; i<len(s) && c<=3; i++ {
		switch s[i] {
			case 'B': fallthrough;
			case 'F': fallthrough;
			case 'P': fallthrough;
			case 'V':
				if c == 1 || ret[c-1] != '1' {
					ret[c] = '1'
					c++
				}
			case 'C': fallthrough;
			case 'G': fallthrough;
			case 'J': fallthrough;
			case 'K': fallthrough;
			case 'Q': fallthrough;
			case 'S': fallthrough;
			case 'X': fallthrough;
			case 'Z':
				if c == 1 || ret[c-1] != '2' {
					ret[c] = '2'
					c++
				}
			case 'D': fallthrough;
			case 'T':
				if c == 1 || ret[c-1] != '3' {
					ret[c] = '3'
					c++
				}
			case 'L':
				if c == 1 || ret[c-1] != '4' {
					ret[c] = '4'
					c++
				}
			case 'M': fallthrough;
			case 'N':
				if c == 1 || ret[c-1] != '5' {
					ret[c] = '5'
					c++
				}
			case 'R':
				if c == 1 || ret[c-1] != '6' {
					ret[c] = '6'
					c++
				}
			case 'A': fallthrough;
			case 'E': fallthrough;
			case 'H': fallthrough;
			case 'I': fallthrough;
			case 'O': fallthrough;
			case 'U': fallthrough;
			case 'Y':
				;
			default:
				ok = false
		}
	}
	if ok {
		for ; i<len(s); i++ {
			if 'A' <= s[i] && s[i] <= 'Z' {
				ok = false
				break
			}
		}
	}
	for ; c<=3; c++ {
		ret[c] = '0'
	}
	return string(ret), ok
}

// Read a simple dictionary that has one entry per line
func readDictionary(r *bufio.Reader) map[string] *vector.StringVector {
	ret := make(map[string] *vector.StringVector)
	for s, ok := r.ReadString('\n'); ok == nil; s, ok = r.ReadString('\n') {
		s = s[0:len(s)-1]
		sd, ok := soundex(s)
		if !ok {
			continue;
		}
		v, ok := ret[sd]
		if !ok {
			v = new(vector.StringVector)
			ret[sd] = v
		}
		v.Push(s)
	}
	return ret
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// Own type for sorting (string length compared to original)
type Candidates struct {
	vector.StringVector
	origlen int // The word's length that we compare to
}

func (c *Candidates) Less(i, j int) bool {
	return abs(c.origlen - len(c.At(i))) < abs(c.origlen - len(c.At(j)))
}

// Reads string slices from channel c1, outputs same to c2,
// performing the soundex based manipulation in between. Also, 
// factors in the length of strings in deciding a "close neighbor" 
// of a word. Uses dictionary d.
func participant(c1 chan []string, c2 chan []string, d map[string] *vector.StringVector) {
	for {
		in := <- c1
		
		// Decide mangled word by random
		x := rand.Intn(len(in))
		w := in[x]
		b := w[0]
		for ; !('a' <= b && b <= 'z') && !('A' <= b && b <= 'Z'); {
			// If this is not a word, choose the next token, as it,
			// or the one after, must be a word (in case of last token)
			x = (x+1) % len(in)
			w = in[x]
			b = w[0]
		}
	
		// Look for neighbors with no soundex hamming distance
		sd, _ := soundex(w)
		v, ok := d[sd]
		all := new(Candidates)
		all.origlen = len(w)
		primary := 0
		if ok {
			all.AppendVector(v)
			primary = v.Len()
			sort.Sort(all)
		}
		
		if !ok || all.Len() < 10 {
			// More candidates. Seek words with soundex hamming 
			// distance == 1 (numeric part only)
			extra := new(Candidates)
			extra.origlen = len(w)
			seek := []byte(sd)
			for i:=1; i<4; i++ {
				var j byte
				for j=0; j<10; j++ {
					if sd[i] != '0'+j {
						seek[i] = '0'+j
						v, ok = d[string(seek)]
						if ok {
							extra.AppendVector(v)
						}
					}
				}
				seek[i] = sd[i]
			}
			sort.Sort(extra)
			all.AppendVector(&extra.StringVector)
		}

		if all.Len() != 0 {
			// Pick a candidate
			pick := int(rand.ExpFloat64()/distrib * float64(all.Len()))
			if pick >= all.Len() {
				pick = all.Len() -1
			}
			rep := all.At(pick)

			// Equalize case
			if unicode.IsUpper(int(rep[0])) != unicode.IsUpper(int(w[0])) {
				t := []int(rep)
				if unicode.IsUpper(int(w[0])) {
					t[0] = unicode.ToUpper(t[0])
				} else {
					t[0] = unicode.ToLower(t[0])
				}
				rep = string(t)
			}
			in[x] = rep

			if debug {
				s,_ := soundex(rep)
				if primary == all.Len() {
					fmt.Printf("%s(%s): %d all, choosing %s(%s)=%d Δlen=%d\n",
					w, sd, primary, rep, s, pick, len(rep)-len(w))
				} else {
					fmt.Printf("%s(%s): %d all(%d dist=1), choosing %s(%s)=%d Δlen=%d\n",
					w, sd, all.Len(), primary, rep, s, pick, len(rep)-len(w))
				}
			}

		} else if debug {
			fmt.Printf("%s can not replace, no words within soundex hamming distance <=1!\n", w)
		}

		c2 <- in
	}
}

// Tokenize by letters a-z. bool is false iff no letters were found.
func tokenize(s string) ([]string, bool) {
	v := new(vector.StringVector)
	ok := false
	for i:=0; i<len(s); {
		begin := i
		for ; i<len(s); {
			if ('a' <= s[i] && s[i] <= 'z') ||
			   ('A' <= s[i] && s[i] <= 'Z') {
				i++
				ok = true
			} else {
				break;
			}
		}
		if begin < i {
			v.Push(s[begin:i])
		}
		begin = i
		for ; i<len(s); {
			if !('a' <= s[i] && s[i] <= 'z') &&
			   !('A' <= s[i] && s[i] <= 'Z') {
				i++
			} else {
				break;
			}
		}
		if begin < i {
			v.Push(s[begin:i])
		}
	}
	return v.Data(), ok
}

var debug bool
var distrib float64
func main() {
	rand.Seed(time.Nanoseconds())

	participants := flag.Int("par",
				 20, 
				 "Number of participants")
	dict := flag.String("dict",
			    "/usr/share/dict/american-english-small",
			    "Dictionary to use")
	flag.Float64Var(&distrib, 
		   "dist",
		   10.0, 
		   "Random distribution decay rate. Less = more distributed.")
	flag.BoolVar(&debug, 
		     "debug",
		     false, 
		     "Print info about choosing of replacements.")
	flag.Parse()
	
	f, err := os.Open(*dict, os.O_RDONLY, 0)
	if err != nil {
		panic("Could not open dictionary file ", *dict)
	}
	b := bufio.NewReader(f)

	sdict := readDictionary(b)

	enter := make(chan []string)
	var exit chan []string
	for i, last := 0, enter; i<*participants; i++ {
		exit = make(chan []string)
		go participant(last, exit, sdict)
		last = exit
	}

	input := bufio.NewReader(os.Stdin)
	for {
		s, err := input.ReadString('\n')
		if err != nil || s == "exit\n" || s=="\n" {
			os.Exit(0)
		}
		tok, ok := tokenize(s)
		if !ok {
			os.Exit(0)
		}
		enter <- tok
		result := <-exit
		for i:=0; i<len(result); i++ {
			fmt.Printf("%s", result[i])
		}
	}
}
