/*
 * broken_telephone.go: The children's game of "broken telephone"
 * implemented with goroutines passing messages read from stdin.
 * Each goroutine mutates the message by replacing one word before
 * passing it on.
 *
 * For Introduction to Go, Spring 2010
 * Kimmo Kulovesi <kkuloves@cs.helsinki.fi>
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"rand"
	"sort"
	"strings"
	"unicode"
	"utf8"
)

var (
	numberOfLinks = flag.Int("l", 20,
		"Number of unreliable links to pass the message through")
	wordList = flag.String("w", "/usr/share/dict/words",
		"Filename of the word list to use")
	printIntermediateStates = flag.Bool("i", false,
		"Print intermediate states of the message")
	debugMetaphone = flag.Bool("d", false,
		"Print double metaphone algorithm debug output (very verbose!)")
	allowNamesInWordList = flag.Bool("n", false,
		"Allow names to be read from the word list")
	dictionary metaphoneDict
)

// Minimum length of a word to consider it for substitution.
// TODO: Implement stopwords instead of using this limit.
const minWordLength = 4

// Default dicitionary capacity.
const dictCapacity = 100000

// The distance (in either direction) at which to "fudge" the word
// substitutes - increasing this gives more random matches.
const fudgeDistance = 4

// The maximum length of the phonetic representations to use, 0 for unlimited.
// The traditional metaphone algorithm uses 4.
const doubleMetaphoneMaxLen = 4

// mutate copies m but randomly substitutes one word.
func mutate(m message) message {
	if len(m) == 0 {
		return m
	}
	n := make(message, len(m))
	copy(n, m)
	for tries := 0; tries < 3; tries++ {
		i := rand.Intn(len(m))
		word := m[i]
		if len(word.original) >= minWordLength {
			n[i] = dictionary.randomSubstituteFor(word)
			return n
		}
	}
	return n
}

// telephone passes mutated messages between from and to.
func telephone(from, to chan message) {
	for m := <-from; m != nil; m = <-from {
		if *printIntermediateStates {
			fmt.Printf("Heard: %s\n", m)
		}
		to <- mutate(m)
	}
	to <- nil
}

// main reads the dictionary, creates the telephone system and then passes
// messages read from stdin to it.
func main() {
	flag.Parse()

	if file, err := os.Open(*wordList, os.O_RDONLY, 0); err == nil {
		fmt.Fprintf(os.Stderr, "Reading words from \"%s\"...\n", *wordList)
		dictionary = readWords(file)
		file.Close()
	} else {
		fmt.Fprintf(os.Stderr, "%s: %s\n", *wordList, err)
	}
	if dictionary.Len() < 1 {
		fmt.Fprintf(os.Stderr, "No words read, aborting.")
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "%d words in dictionary\n", dictionary.Len())

	send := make(chan message)
	receive := send
	for i := *numberOfLinks; i > 0; i-- {
		c := make(chan message)
		go telephone(receive, c)
		receive = c
	}

	input := bufio.NewReader(os.Stdin)
	for {
		if line, err := input.ReadString('\n'); err == nil {
			words := strings.Fields(line[0 : len(line)-1])
			msg := make(message, len(words))
			for i, word := range words {
				msg[i] = doubleMetaphone(word)
			}
			send <- msg
			fmt.Println(<-receive)
		} else {
			if err != os.EOF {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			}
			break
		}

	}
	send <- nil
}

type metaphoneWord struct {
	original, literal, metaphone, secondary string
}

type message []*metaphoneWord

func (m message) String() (s string) {
	for i, word := range m {
		s += word.original
		if i != len(m)-1 {
			s += " "
		}
	}
	return
}

type metaphoneDict []metaphoneWord

func newMetaphoneDict() metaphoneDict      { return make(metaphoneDict, 0, dictCapacity) }
func (d metaphoneDict) Len() int           { return len(d) }
func (d metaphoneDict) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d metaphoneDict) Less(i, j int) bool { return d[i].metaphone < d[j].metaphone }

// phoneticLocation returns the index where metaphone is or would be sorted.
func (d metaphoneDict) phoneticLocation(metaphone string) (i int) {
	left, right := 0, len(d)
	for left < right {
		i = left + ((right - left) / 2)
		if d[i].metaphone < metaphone {
			left = i + 1
		} else {
			right = i
		}
	}
	return
}

// matches returns a slice containing all exact matches in d for metaphone.
func (d metaphoneDict) matches(metaphone string) metaphoneDict {
	var l, r int
	i := d.phoneticLocation(metaphone)
	for r = i; r < len(d) && d[r].metaphone == metaphone; r++ {
	}
	for l = i; l >= 0 && d[l].metaphone == metaphone; l-- {
	}
	l++
	if r-l < fudgeDistance*2 {
		l -= fudgeDistance
		r += fudgeDistance
	}
	if l < 0 {
		l = 0
	}
	if r > len(d) {
		r = len(d)
	}
	return d[l:r]
}

// randomNonEqual returns a randomly selected word in d which is not
// literally equal to w. Returns -1 if no such word exists.
// (The assumption is that the same word literal appears at most twice.)
func (d metaphoneDict) randomNonEqual(w string) *metaphoneWord {
	if len(d) == 0 {
		return nil
	}
	i := rand.Intn(len(d))
	switch {
	case d[i].literal != w:
	case i > 0 && d[i-1].literal != w:
		i--
	case i < len(d)-1 && d[i+1].literal != w:
		i++
	case i == 0 && len(d) > 2 && d[i+2].literal != w:
		i += 2
	case i == len(d)-1 && len(d) > 2 && d[i-2].literal != w:
		i -= 2
	default:
		return nil
	}
	return &d[i]
}

// randomSubstituteFor returns a semi-random but usually phonetically
// close substitute for w.
func (d metaphoneDict) randomSubstituteFor(w *metaphoneWord) *metaphoneWord {
	mp := w.metaphone
	if rand.Intn(2) == 0 {
		mp = w.secondary
	}
	match := d.matches(mp).randomNonEqual(w.literal)
	if match != nil {
		return match
	}
	return w
}

// Push adds mp to the end of d, reallocating space if necessary.
func (d *metaphoneDict) Push(mp metaphoneWord) {
	dict := *d
	if len(dict) < cap(dict) {
		dict = dict[0 : len(dict)+1]
	} else {
		dict = make(metaphoneDict, len(dict)+1, cap(dict)+dictCapacity)
		copy(dict, *d)
	}
	dict[len(dict)-1] = mp
	*d = dict
}

// readWords creates a metaphone dictionary from input of one word per line.
// The dictionary is sorted by phonetic representation.
func readWords(input io.Reader) metaphoneDict {
	dict := newMetaphoneDict()
	rd := bufio.NewReader(input)
	for {
		if word, err := rd.ReadString('\n'); err == nil {
			word = strings.TrimSpace(word)
			if !*allowNamesInWordList {
				rune, _ := utf8.DecodeRuneInString(word)
				if unicode.IsUpper(rune) {
					continue
				}
			}
			if strings.Index(word, " ") != -1 || len(word) < minWordLength {
				continue
			}
			mp := *doubleMetaphone(word)
			dict.Push(mp)
			if mp.metaphone != mp.secondary {
				// Secondary phonetic representation
				dict.Push(metaphoneWord{word, mp.literal,
					mp.secondary, mp.metaphone})
			}
		} else {
			if err != os.EOF {
				fmt.Printf("Error: %s\n", err)
			}
			break
		}
	}
	sort.Sort(dict)
	return dict
}

// doubleMetaphone is like doubleMetaphoneLimited but uses the default maxLen.
func doubleMetaphone(word string) (result *metaphoneWord) {
	if doubleMetaphoneMaxLen > 0 {
		return doubleMetaphoneLimited(word, doubleMetaphoneMaxLen)
	}
	return doubleMetaphoneLimited(word, len(word)*2)
}

// doubleMetaphone returns two phonetic representations of an English word.
// The secondary representation may equal primary. The implementation is
// currently not fully complete with all special cases. The phonetic
// representations are limited to maxLen length.
func doubleMetaphoneLimited(word string, maxLen int) (result *metaphoneWord) {
	result = new(metaphoneWord)
	result.original = word
	word = strings.ToUpper(word)
	// TODO: Strip punctuation
	result.literal = word
	prev, skip, last, slavoGermanic := 0, 0, len(word)-1, false

testSlavoGermanic: for pos, c := range word {
		switch c {
		case 'C':
			if pos == last || word[pos+1] != 'Z' {
				break
			}
			fallthrough
		case 'W', 'K':
			slavoGermanic = true
			break testSlavoGermanic
		}
	}

	word += "     " // Allow indexing beyond the end
	for pos, c := range word {
		if c == ' ' {
			break
		}
		if skip > 0 {
			prev = 0
			skip--
			continue
		}
		mp, ms := "", "-"

		switch c {
		case 'A', 'E', 'I', 'O', 'U', 'Y', 'Ü', 'Ä', 'Ö', 'Å', 'É', 'È', 'Ï':
			if pos == 0 {
				// Initial vowel
				mp = "A"
			} else if pos == 1 && prev == 'W' {
				// W + vowel at the start of the word
				mp, ms = "A", "F"
			}
		case 'B':
			if prev == 'M' && pos > 1 && word[pos-2] == 'U' &&
				(pos == last || (word[pos+1] == 'E' &&
					word[pos+2] == 'R')) {
				// e.g. dumb, thumb
				break
			}
			if prev != 'B' {
				mp = "P"
			}
		case 'C':
			if prev == 'X' {
				break
			}
			if pos == 0 && strings.HasPrefix(word, "CAESAR") {
				mp = "S"
				skip = 1
				break
			}
			next := word[pos+1]
			if next == 'H' {
				skip = 1
				n2, n3 := word[pos+2], word[pos+3]
				if pos > 0 {
					if n2 == 'A' && n3 == 'E' {
						// michael
						mp, ms = "K", "X"
						break
					}
					if (pos == 1 && (prev == 'M' || prev == 'S')) ||
						n2 == 'T' || n2 == 'S' {
						// Mc, Sch, -cht, -chs
						mp = "K"
						break
					}
					if (prev == 'A' || prev == 'O' || prev == 'U' ||
						prev == 'E') && (n2 == 'L' || n2 == 'R' || n2 == 'N' ||
						n2 == 'M' || n2 == 'B' || n2 == 'B' ||
						n2 == 'H' || n2 == 'F' || n2 == 'V' ||
						n2 == 'W') {
						// e.g. wachtler, wechsler, but not tichner
						mp = "K"
						break
					}
					if pos > 1 {
						p2 := word[pos-2]
						if prev == 'R' &&
							((p2 == 'O' && n2 == 'E' && n3 == 'S') ||
								(p2 == 'O' && n2 == 'I' && n3 == 'D') ||
								(p2 == 'A' && n2 == 'I' && n3 == 'T')) {
							// orchestra, orchid, architect (but not arch)
							mp = "K"
							break
						}
					}
				} else {
					// pos == 0
					n4, n5 := word[pos+4], word[pos+5]
					if (n2 == 'A' && n3 == 'R' && ((n4 == 'A' && n5 == 'C') ||
						(n4 == 'I' && n5 == 'S'))) ||
						(n2 == 'E' && n3 == 'M') || (n2 == 'Y' && n3 == 'M') ||
						(n2 == 'I' && n3 == 'A') ||
						(n2 == 'O' && n3 == 'R' && (n4 != 'O' || n5 != 'E')) {
						// e.g. character, charisma, chorus, chemistry
						// but not "chore"
						mp = "K"
					} else {
						switch n2 {
						case 'L', 'R', 'N', 'M', 'B', 'H', 'F', 'V', 'W', ' ':
							mp = "K"
						default:
							mp = "X"
						}
					}
					break
				}
				mp, ms = "X", "K"
				break
			} else if next == 'Z' {
				if pos < 2 || word[pos-1] != 'I' || word[pos-2] == 'W' {
					// cz, not wicz
					mp, ms = "S", "X"
					skip = 1
					break
				}
			} else if next == 'C' {
				n2 := word[pos+2]
				if n2 == 'I' && word[pos+3] == 'A' {
					// -ccia, e.g. focaccia
					mp = "X"
					skip = 2
					break
				}
				if pos != 2 || prev != 'M' {
					// -cc, but not e.g. McClellan
					if n2 == 'I' || n2 == 'E' ||
						(n2 == 'H' && word[pos+3] != 'U') {
						// e.g. bellocchio, but not bacchus
						skip = 3
						if pos == 1 && prev == 'A' {
							// e.g. accident
							mp = "KS"
							break
						} else if prev == 'U' && n2 == 'E' &&
							(word[pos+4] == 'S' || word[pos+4] == 'E') {
							// succeed, success
							mp = "KS"
							break
						}
						mp = "X"
						break
					}
				}
				if n2 != 'I' && n2 != 'E' {
					skip = 1
				}
			} else if next == 'K' || next == 'Q' {
				skip = 1
			} else if next == 'I' {
				mp = "S"
				skip = 1
				n2 := word[pos+2]
				if n2 == 'O' || n2 == 'E' || n2 == 'A' {
					// cio, cie, cia
					ms = "X"
				}
				break
			} else if next == 'E' || next == 'Y' {
				skip = 1
				mp = "S"
				break
			}
			mp = "K"
		case 'D':
			if prev != 'D' && prev != 'T' {
				if word[pos+1] == 'G' {
					skip = 1
					switch word[pos+2] {
					case 'E', 'I', 'Y':
						// e.g. "edge"
						mp = "J"
					default:
						// e.g. "edgar"
						mp = "K"
					}
					break
				}
				mp = "T"
			}
		case 'F', 'V':
			if prev != c {
				mp = "F"
			}
		case 'G':
			next := word[pos+1]
			if next == 'H' {
				skip = 1
				if !isVowel(prev) {
					mp = "K"
					break
				}
				if pos == 0 {
					if word[pos+2] == 'I' {
						mp = "J"
					} else {
						mp = "K"
					}
					break
				}
				if pos > 1 {
					if word[pos-2] == 'B' || word[pos-2] == 'H' ||
						word[pos-2] == 'D' {
						// e.g. hugh
						break
					}
					if pos > 2 {
						p3 := word[pos-3]
						if p3 == 'B' || p3 == 'H' || p3 == 'D' {
							// e.g. bough
							break
						}
						if pos > 3 && (word[pos-4] == 'B' || word[pos-4] == 'H') {
							// e.g. brought
							break
						}
						if prev == 'U' && (p3 == 'C' || p3 == 'G' ||
							p3 == 'L' || p3 == 'R' ||
							p3 == 'T') {
							// e.g. laugh, cough, rough, tough
							mp = "F"
							break
						}
					}
				}
				if prev != 'I' {
					mp = "K"
				}
				break
			}
			if next == 'N' {
				skip = 1
				if !slavoGermanic {
					if pos == 1 && isVowel(prev) {
						mp, ms = "KN", "N"
						break
					} else if word[pos+2] != 'E' || word[pos+3] != 'Y' {
						// not e.g. cagney
						mp, ms = "N", "KN"
						break
					}
				}
				mp = "KN"
				break
			}
			if next == 'L' {
				if word[pos+2] == 'I' && !slavoGermanic {
					// e.g. tagliaro
					mp, ms = "KL", "L"
					skip = 1
				}
			} else if next == 'E' || next == 'I' || next == 'Y' || next == 'G' {
				skip = 1
				n2 := word[pos+2]
				if next != 'G' {
					if pos == 0 {
						if (next == 'E' && (n2 == 'S' || n2 == 'P' ||
							n2 == 'B' || n2 == 'L' ||
							n2 == 'Y' || n2 == 'I' ||
							n2 == 'R')) || next == 'Y' ||
							(next == 'I' && (n2 == 'L' || n2 == 'N')) {
							skip = 1
							mp, ms = "K", "J"
							break
						}
					}
					if !(next == 'I' || prev == 'I' || prev == 'E' ||
						(next == 'Y' && (prev == 'R' || prev == 'O')) ||
						(next == 'E' && pos > 0 && word[pos-1] != 'R')) {
						// -ger-, -gy-
						mp = "K"
						if !(pos == 3 && next == 'E' &&
							strings.HasPrefix(word, "DANGER") ||
							strings.HasPrefix(word, "RANGER") ||
							strings.HasPrefix(word, "MANGER")) {
							ms = "J"
						}
						break
					}
				} else if !(n2 == 'I' && (prev == 'A' || prev == 'O')) {
					// not -aggi -oggi
					mp = "K"
					break
				}
				if !strings.HasPrefix(word, "SCH") ||
					(next == 'E' && n2 == 'T') {
					// obvious Germanic
					mp = "K"
				} else if next == 'I' && pos == (last-3) &&
					strings.HasSuffix(word, "ER") {
					// -gier suffix
					mp = "J"
				} else {
					mp, ms = "J", "K"
				}
				break
			}
			mp = "K"
		case 'H':
			if pos == 0 || isVowel(prev) {
				next, _ := utf8.DecodeRuneInString(word[pos+1 : len(word)])
				if isVowel(next) {
					// H between two vowels, or at the beginning followed by a vowel
					mp = "H"
					skip = 1
				}
			}
		case 'J':
			if prev == 'S' || prev == 'K' || prev == 'L' || prev == 'J' {
				break
			}
			next := word[pos+1]
			if pos == 0 {
				if next == 'O' && word[pos+2] == 'S' && word[pos+3] == 'E' {
					if word[pos+4] == ' ' {
						// Jose
						mp = "H"
					} else {
						mp, ms = "J", "H"
					}
					break
				}
				mp, ms = "J", "A"
			} else if !slavoGermanic && isVowel(prev) &&
				(next == 'A' || next == 'O') {
				mp, ms = "J", "H"
			} else if pos == last {
				mp, ms = "J", ""
			} else {
				switch next {
				case 'L', 'T', 'K', 'S', 'N', 'M', 'B', 'Z':
					// NOP
				default:
					mp = "J"
				}
			}
		case 'Q', 'K':
			if prev != c && prev != 'C' {
				mp = "K"
			}
		case 'L':
			if word[pos+1] == 'L' {
				skip = 1
				if pos > 0 && ((word[pos+3] == ' ' &&
					(((word[pos+2] == 'O' || word[pos+2] == 'A') &&
						word[pos-1] == 'I') || (word[pos+2] == 'E' &&
						word[pos-1] == 'A'))) ||
					((word[last] == 'S' && (word[last-1] == 'A' ||
						word[last-1] == 'O')) ||
						(word[last] == 'A' || word[last] == 'O') &&
							(word[pos-1] == 'A' && word[pos+2] == 'E'))) {
					// Spanish, -illo, -illa, -alle
					ms = ""
				}
			}
			mp = "L"
		case 'M':
			if prev != 'M' {
				mp = "M"
			}
		case 'N':
			if pos == 1 && (prev == 'K' || prev == 'G' || prev == 'P') {
				// Skip GN, KN, PN at the start of the word
				result.metaphone, result.secondary = "", ""
				break
			}
			fallthrough
		case 'Ñ':
			if prev != c {
				mp = "N"
			}
		case 'P':
			next := word[pos+1]
			if next == 'H' {
				mp = "F"
				skip = 1
				break
			}
			if next == 'S' && pos == 0 {
				// Ignore PS at the start of the word
				skip = 1
				break
			}
			if next == 'P' || next == 'B' {
				skip = 1
			}
			mp = "P"
		// case 'Q': is grouped with K
		case 'R':
			if prev == 'R' {
				break
			}
			if pos == last && !slavoGermanic && prev == 'E' && pos > 1 &&
				word[pos-2] == 'I' && (pos < 4 || word[pos-4] != 'M' ||
				!(word[pos-3] == 'E' || word[pos-3] == 'A')) {
				// French, e.g. rogier, but not e.g. hochmeier
				mp, ms = "", "R"
			} else {
				mp = "R"
			}
		case 'S', 'ß', 'Š':
			if prev == 'S' {
				break
			}
			next := word[pos+1]
			if (prev == 'I' || prev == 'Y') && next == 'L' {
				// isl, ysl, e.g. island, isle, carlysle
				break
			}
			if pos == 0 {
				if next == 'M' || next == 'N' || next == 'L' || next == 'W' {
					mp, ms = "S", "X"
					break
				}
				if strings.HasPrefix(word, "SUGAR") {
					mp, ms = "X", "S"
					break
				}
			}
			if next == 'H' {
				if word[pos+2] == 'O' {
					if (word[pos+3] == 'E' && word[pos+4] == 'K') ||
						(word[pos+3] == 'L' && (word[pos+4] == 'M' ||
							word[pos+4] == 'Z')) {
						// holm, holz, hoek
						mp = "S"
						break
					}
				} else if word[pos+2] == 'E' && word[pos+3] == 'I' &&
					word[pos+4] == 'M' {
					// heim
					mp = "S"
					break
				}
				mp = "X"
				skip = 1
			} else if next == 'I' && (word[pos+2] == 'O' || word[pos+2] == 'A') {
				// sio, sia
				mp = "S"
				if !slavoGermanic {
					ms = "X"
				}
				skip = 2
			} else if next == 'Z' {
				mp, ms = "S", "X"
				skip = 1
			} else if next == 'C' {
				skip = 2
				if word[pos+2] == 'H' {
					n3, n4 := word[pos+3], word[pos+4]
					if (n3 == 'O' && n4 == 'O') || (n3 == 'U' && n4 == 'Y') ||
						(n3 == 'E' && (n4 == 'D' || n4 == 'M')) {
						// Dutch origin, e.g. "school", "schooner"
						mp = "SK"
					} else if n3 == 'E' && (n4 == 'R' || n4 == 'N') {
						mp, ms = "X", "SK"
					} else {
						mp = "X"
						if pos == 0 && !isVowel(int(word[3])) && word[3] != 'W' {
							ms = "S"
						}
					}
				} else if word[pos+2] == 'I' || word[pos+2] == 'E' ||
					word[pos+2] == 'Y' {
					mp = "S"
				} else {
					mp = "SK"
					skip = 1
					// TODO: Check correctness of skip
				}
			} else if pos == last && prev == 'I' {
				if pos > 1 && (word[pos-2] == 'A' || word[pos-2] == 'O') {
					// French, e.g. artois
					ms = "S"
				} else {
					mp = "S"
				}
			} else {
				mp = "S"
			}
		case 'T':
			if prev == 'T' {
				if word[pos+1] == 'H' {
					// tth
					mp, ms = "0", "T"
					skip = 1
				} else {
					mp = "T"
				}
				break
			}
			if prev == 'D' {
				break
			}
			next := word[pos+1]
			if next == 'I' {
				if word[pos+2] == 'A' ||
					(word[pos+2] == 'O' && word[pos+3] == 'N') {
					// tia, tion
					mp = "X"
					skip = 2
				}
			} else if next == 'C' && word[pos+2] == 'H' {
				// tch
				mp = "X"
				skip = 2
			} else if next == 'H' {
				skip = 1
				if word[pos+3] == 'M' {
					if word[pos+2] == 'O' || word[pos+2] == 'A' {
						mp = "T"
						break
					}
				}
				mp, ms = "0", "T"
			} else if next != 'T' {
				mp = "T"
			}
		// case 'V': is grouped with F
		case 'W':
			next := word[pos+1]
			if next == 'R' {
				if pos != 0 {
					mp = "R"
				}
				skip = 1
				break
			}
			if pos == 0 {
				if next == 'H' {
					mp = "A"
				}
				break
			}
			if (pos == last && isVowel(prev)) ||
				strings.HasPrefix(word, "SCH") {
				ms = "F"
				break
			}
			n2, n3 := word[pos+2], word[pos+3]
			if (prev == 'E' || prev == 'O') && next == 'S' && n2 == 'K' &&
				(n3 == 'I' || n3 == 'Y') {
				// -ewski, -ewsky, -owski, -owsky
				ms = "F"
			} else if next == 'I' && n3 == 'Z' && (n2 == 'C' || n2 == 'T') {
				// -wicz, -witz
				mp, ms = "TS", "FX"
				skip = 3
			}
		case 'X':
			if pos == 0 {
				// Initial X pronounced like a Z, e.g. Xavier
				mp = "S"
			} else if prev != 'X' {
				if pos == last && prev == 'U' && pos > 1 &&
					(word[pos-2] == 'A' || word[pos-2] == 'O') {
					// French, e.g. breaux
					break
				}
				mp = "KS"
			}
		case 'Z':
			if prev == 'S' || prev == 'Z' {
				break
			}
			if word[pos+1] == 'H' {
				// Chinese, e.g. Zhao
				mp = "J"
				skip = 1
				break
			}
			if word[pos+1] == 'I' || word[pos+1] == 'O' || word[pos+1] == 'A' ||
				(slavoGermanic && prev != 'T' && pos > 0) {
				ms = "TS"
			}
			fallthrough
		case 'Ç':
			mp = "S"
		default:
		}
		prev = c
		result.metaphone += mp
		if ms == "-" {
			ms = mp
		}
		result.secondary += ms

		if *debugMetaphone {
			fmt.Fprintf(os.Stderr, "\t%c -> [%s] [%s]\n", c, mp, ms)
		}

		if len(result.metaphone) >= maxLen && len(result.secondary) >= maxLen {
			break
		}
	}
	if len(result.metaphone) > maxLen {
		result.metaphone = result.metaphone[0:maxLen]
	}
	if len(result.secondary) > maxLen {
		result.secondary = result.secondary[0:maxLen]
	}
	if result.secondary == result.metaphone {
		result.secondary = result.metaphone
	}
	if *debugMetaphone {
		fmt.Fprintf(os.Stderr, "%s: [%s] [%s]\n", result.literal, result.metaphone,
			result.secondary)
	}
	return
}

func isVowel(c int) bool {
	// TODO: Non-English support is rather limited
	return c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' || c == 'Y' ||
		c == 'Ä' || c == 'Ö' || c == 'Ü' || c == 'Å' || c == 'É' || c == 'È' ||
		c == 'Ï'
}
