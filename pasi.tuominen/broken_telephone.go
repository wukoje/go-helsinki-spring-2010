package main

import (
 "fmt"
 "io/ioutil"
 "strings"
 "os"
 "rand"
 "flag"
)


type word struct {
	written string
	soundex string
}


func soundex(w string) string {

	if (len(w) < 1) {
		return "0000"
	}

	w =  strings.ToLower(w)
	soundexCode := make([]byte, 4, 4)
	soundexCode[0] = w[0]
	
	replaceMap := map[byte] byte {
	 'b':'1', 'p':'1', 'f':'1', 'v':'1',
	 'c':'2', 's':'2', 'k':'2', 'g':'2', 'j':'2', 'q':'2', 'x':'2', 'z':'2',
	 'd':'3', 't':'3',
	 'l':'4',
	 'm':'5', 'n':'5',
	 'r':'6' }
	
	pos := 1
	for i:=1;i < len(w) && pos < 4;i++ {
		newChar, ok := replaceMap[w[i]]
		if (ok) {
			if (soundexCode[pos-1] != newChar) { //Collapse adjacent digits
				soundexCode[pos] = newChar
				pos++
			}
		} 
	}

	for pos < 4 { 
		soundexCode[pos] = '0'
		pos++
	}
	return string(soundexCode);
}


func create(w string) *word {

	newWord := new(word)
	newWord.written = strings.ToLower(w)
	newWord.soundex = soundex(newWord.written)
	return newWord
}


func append(s []word, w *word) []word {
	l := len(s)

	if len(s) == cap(s) {
		newSlice := make([]word, l*2)
		for i, c := range s {
			newSlice[i] = c
		}
		s = newSlice
	}
	s = s[0 : l+1]
	s[l] = *w
	return s
}

func findSimilarWords(w string) []word {
	w = strings.ToLower(w)
	soundex := soundex(w)
	
	similar := make([]word,0,50)
	
	
	for _,v := range words {
		distance := len(v.written) - len(w);
		if (distance < 0) { 
			distance *= -1
		}
		
		if (v.soundex == soundex && distance <= 1) {
			similar = append(similar, &v);
		}
	}
	return similar
}




func createDictionary(dictFile string) []word {

	dictionary := make([]word,0,100000);

	data,error := ioutil.ReadFile(dictFile);
	if (error != nil) {
		fmt.Fprintf(os.Stderr, "Error: %s\n\n", error.String());
		return nil
	}

	wc :=0
	wordStart := 0
	
	for k,v := range data {
		if (v == '\n') {
			wc++
			wordString := string(data[wordStart:k])
			if (len(wordString) > 1) {
				dictionary = append(dictionary, create(wordString));

			}
			wordStart = k+1
		}
	}
	
	return dictionary;
}


var words []word
var similarWords [][]word



func main() {

	flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
    flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "Everything after flags are considered as being the used sentence.\n")
}

	defaultDictionary := "/usr/share/dict/words";

	dictFile := flag.String("d", defaultDictionary, "a dictionary file having one word per line.")
  
	flag.Parse()

	words = createDictionary(*dictFile);
	if (words == nil) {
		flag.Usage();
		return
	}
	
	
	
	var sentence []string
	
	if (len(flag.Args()) != 0) {
		sentence = flag.Args()
		fmt.Printf("In: ");
		for _,v := range sentence {
			fmt.Printf("%s ", v);
		}
		fmt.Printf("\n\n");
	
	} else {
		sentence = []string{"The", "man", "ate", "his", "dinner"}
		fmt.Printf("In: \"The man ate his dinner\" (default, provide your own sentence as arguments.)\n\n")
	}
	
	
	similarWords = make([][]word, len(sentence), len(sentence))
	
	for k,v := range sentence {
	
		similarWords[k] = findSimilarWords(v);
	
	}

		
	in := make(chan []string)
	head := in
	var out chan []string
	
	for i:=0;i<20;i++ {
	
		out = make(chan []string)
		go person(in, out)
		in = out
		
	}

	head <- sentence
	
	sentence = <-in
	
	fmt.Printf("Out: ");
	for _,v := range sentence {
		fmt.Printf("%s ", v);
	}
	fmt.Printf("\n");
}



func person(in chan []string, out chan []string) {

	sentence := <- in
	
	_,nsec,_ := os.Time()
	gen := rand.New(rand.NewSource(nsec))

	randomIndex := int( gen.Float() * float(len(sentence)))

	similarCount := len(similarWords[randomIndex])
	if (similarCount > 0) {
		randWordIndex := int( gen.Float() * float(similarCount))
		sentence[randomIndex] = similarWords[randomIndex][randWordIndex].written
	}
	
	out <- sentence

}
