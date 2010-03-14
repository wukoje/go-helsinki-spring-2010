package main

import (
	"fmt"
	"os"
	"strings"
	"rand"
	"time"
	"bytes"
	"strconv"
)

const (
	MAX_PHONES = 20
	DICTIONARY_FILENAME = "/usr/share/dict/web2"
	DICTIONARY_SIZE = 500000 // words (strings)
	BUFFER_SIZE = 100 // bytes
)

var soundexDict []soundex

type soundex struct {
	word string // plain word
	code string // soundex code of word
}


// values for soundex conversion
var soundexValue = map[byte] byte {
	'b': '1',
	'f': '1',
	'p': '1',
	'v': '1',
	'c': '2',
	'g': '2',
	'j': '2',
	'k': '2',
	'q': '2',
	's': '2',
	'x': '2',
	'z': '2',
	'd': '3',
	't': '3',
	'l': '4',
	'm': '5',
	'n': '5',
	'r': '6',
}

func hamming(word string) string {
	return hammingRecursion(word, 0)
}

func hammingRecursion(word string, recDepth int) string {
	if recDepth > 1 { // limit recursion
		return word
	}

	codeByte := stringToByte(soundexCode(word))
	if recDepth > 0 { // mutate a soundex code a little
		y := rand.Intn(6) // new soundex digit
		p := rand.Intn(3) // new soundex digit position
		for i := 0; i < 4; i++ {
			if i == p {
				codeByte[i] = strconv.Itoa(y)[0]
			}
		}
	}

	table := make([]string, len(soundexDict)/2) // reserve space for tmp table
	counter := 0
	for i := 0; i < len(soundexDict); i++ {
		if bytes.Compare(stringToByte(soundexDict[i].code), codeByte) == 0 && len(soundexDict[i].word) == len(word) {
			table[counter] = soundexDict[i].word
			counter++
		}
	}
	table = table[0:counter]

	x := rand.Intn(counter)

	if counter < 2 {
		// recursively search for a new word with slightly less strict rules
		return hammingRecursion(word, recDepth+1)
	}

	return table[x] // return random matching word
	
}

func sort(s []soundex) {
	quicksort(s, 0, len(s)-1)
}

func quicksort(s []soundex, left int, right int) {
	if right>left {
		// choose pivot point from the middle
		pivot := partition(s, left, right, (left+right)/2)
		quicksort(s, left, pivot-1)
		quicksort(s, pivot+1, right)
	}
}

// partition for quicksort
func partition(s []soundex, left int, right int, pivot int) int {
	pivotVal := stringToByte(s[pivot].code)
	s[pivot], s[right] = s[right], s[pivot]
	store := left
	for i := left; i< right; i++ {
		soundexVal := stringToByte(s[i].code)

		if bytes.Compare(soundexVal, pivotVal) <=0 {
			s[i], s[store] = s[store], s[i]
			store++
		}
	}
	s[store], s[right] = s[right], s[store]
	return store
}

// calculate soundex code for the given string
func soundexCode(word string) string {
	w := stringToByte(strings.ToLower(word))
	for i := 1; i < len(w); i++ { // replace consonants with their soundex code
		tmp, ok := soundexValue[w[i]]
		if ok {
			w[i] = tmp
		} else {
			w[i] = '0'
		}
	}

	w2 := make([]byte, 50)
	counter := 0 
	var prevChar uint8 = '*'
	for i := 0; i < len(w); i++ { // strip repeating digits (not 0's)
		if w[i] != '0' && w[i] == prevChar {
			// skip character
		} else {
			w2[counter] = w[i]
			prevChar = w[i]
			counter++
		}
	}
	w2 = w2[0:counter] // trim w2 to correct size

	w = w2
	w2 = make([]byte, 50)
	counter = 0
	for i := 0; i < len(w); i++ { // strip all 0's
		if w[i] == '0' {
			// skip character
		} else {
			w2[counter] = w[i]
			counter++
		}
	}
	w2 = w2[0:counter] // trim w2 to correct size

	counter = 0
	for i := 0; i < len(w2); i++ { // copy w2->w, don't copy vowels
		if w2[i] != '*' {
			w[counter] = w2[i]
			counter++
		}
	}
	w = w[0:counter] // trim w to correct size

	for ;len(w) < 4; { // extend the size to 4 digits if necessary
		w = w[0:len(w)+1]
		w[len(w)-1] = '0'
	}

	w = w[0:4] // limit the soundex code to 4 characters

	return string(w)
}

// converts the given string into a byte array
func stringToByte(s string) []byte { // borrowed from golan.org
	cap := len(s) + 5
	b := make([]byte, len(s), cap)
	for i := 0; i < len(s); i++ {
		b[i] = s[i]
	}
	return b
}

// locates a word from the dictionary
func search(word string) int {
	low := 0
	up := len(soundexDict)-1
	var mid int

	for low < up {
	        mid = (up + low) / 2
		bWord := stringToByte(strings.ToLower(soundexCode(word)))
		bDictWord := stringToByte(strings.ToLower(soundexDict[mid].code))

		comp := bytes.Compare(bWord, bDictWord)
		if comp < 0 {
			up = mid
		} else if comp > 0 {
			low = mid + 1
		} else {
			return mid
		}
    	}
	return -1 // word not found from dictionary
}

func phone(input chan []string, output chan []string, phoneNumber int) {
	message := <-input
	fmt.Printf ("phone %d received a message: %v\n", phoneNumber, message)
	mutateMessage(message)
	output <- message
}

func createMessage() []string {
	tmp := "The quick brown fox jumps over the lazy dog"

	return strings.Split(strings.ToLower(tmp), " ", 0)
}

func mutateMessage(message []string) {
	i := rand.Intn(len(message)) // choose a word to be mutated
	message[i] = hamming(message[i]) // get new word using soundex && hamming code
}	

func readDictionaryFile () {
	soundexDict = make([]soundex, DICTIONARY_SIZE)
	SEPARATOR := "\n"
	dictionarySize := len(soundexDict)
	counter := 0
	name := DICTIONARY_FILENAME
	permissions := 0666
	file, err := os.Open(name, os.O_RDONLY, permissions)
	defer file.Close()

	if err == nil {
		buffer := make([]byte, BUFFER_SIZE*2) 
		n, err := file.Read(buffer[0:BUFFER_SIZE])
		startIndex := 0
		j := 0
		
		for err == nil && counter<dictionarySize {
			array := strings.Split(string(buffer[0:startIndex+n]), SEPARATOR, 0)
			wordLength := 0
			for i := 0; i < len(array)-1; i++ {
				if counter < dictionarySize {
					wordLength += len(array[i])+1
					word := strings.ToLower(array[i])
					soundexDict[counter] = soundex{word, soundexCode(word)}
					if counter>dictionarySize-1 {
						break
					}
				}
				counter++
			}
			if wordLength<n+j {
				tmp := j
				for j = 0; j < n-wordLength+tmp; j++ { // copy leftover to the beginning
					buffer[j] = buffer[wordLength+j]
				}
				startIndex = j // don't fill buffer from beginning
			} else {
				j = 0
				startIndex = 0 // fill buffer from beginning
			}
			
			n, err = file.Read(buffer[startIndex:startIndex+BUFFER_SIZE-j])
		}
	}

	fmt.Printf("dictionary contains %d words\n", counter)	
	soundexDict = soundexDict[0:counter-1] // shrink soundexDict to actual size
	sort(soundexDict)
}

func main() {
	rand.Seed(time.Nanoseconds())
	readDictionaryFile()

	message := createMessage()

	var firstInputChannel chan []string
	var lastOutputChannel chan []string
	var previousOutputChannel chan []string

	for i := 0; i<MAX_PHONES; i++ {
		var inputChannel chan []string

		outputChannel := make(chan []string)

		if i==0 {
			inputChannel = make(chan []string)
			firstInputChannel = inputChannel
			previousOutputChannel = outputChannel
		} else if i==MAX_PHONES-1 {
			lastOutputChannel = outputChannel
		}

		if i>0 {
			inputChannel = previousOutputChannel
			previousOutputChannel = outputChannel
		}

		// start a phone
		go phone(inputChannel, outputChannel, i)
	}
	
	fmt.Printf ("sending message to the first phone: %v\n", message)
	firstInputChannel <- message
	reply := <-lastOutputChannel

	fmt.Printf ("reply: %v\n", reply)
	
}
