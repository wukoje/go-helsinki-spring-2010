package main 

import (
	"fmt"
    "rand"
)

func soundex(s string) string {
	mutate := s[0:1]
	skip := false
	for i,j := 1, 0; j < 3 && i < len(s); i++{	
		if !skip {		
			switch 	s[i] {
				case 'b', 'f', 'p', 'v':
					mutate += "1"
					j++			
				case 'c', 'g', 'j', 'k', 'q', 's', 'x', 'z':
					mutate += "2"
					j++
				case 'd', 't':
					mutate += "3"
					j++
				case 'l':
					mutate += "4"
					j++
				case 'm', 'n':
					mutate += "5"
					j++
				case 'r':
					mutate += "6"
					j++
				case 'h', 'w':
					skip = true			
			}
		} else {
			skip = false;
		}
	}
	for len(mutate) != 4{
		mutate += "0"
	} 
	return mutate
}

func hamming(orig string, subs string) int {
	if len(orig) != len(subs) {
		return 999
	}
	diff := 0
	for i := 0; i < len(orig); i++ {
		if orig[i] != subs[i] {
			diff++
		} 
	}
	return diff
}

var wait = make(chan int , 10)
var sem = make(chan int,1)
var channel chan []string = make(chan []string, 8)


func playGame(number int) {
	sem <- 1	   		// Wait for active queue to drain.	
	s := <-channel
	s = changeWord(s)   // May take a long time.
	channel <- s
	wait<-1				//this process is done	
	<-sem		  		// Done; enable next request to run.
}

//change
func changeWord(s []string) []string{
    if s != nil {    
		//access random word        
		random := rand.Int()%len(s)
		//get soundex    
        word := soundex(s[random])
		//get right stringTable    	
		wordList := Worlds[word]
		var wordRandom = 0		
		
		//find suitable word		
		for {
			wordRandom = rand.Int()%len(wordList)
			result := hamming(wordList[wordRandom], s[random])
			if result <= len(wordList)/3 {
				s[random] = wordList[wordRandom]
				return s
			}
		}
	}
    return s
}

var Worlds map[string][]string

func main(){

	Worlds = initMap()

	test := []string{"introduction", "in" ,"to", "go", "is", "very","interesting", "course", }
	sem <- 1	   		// Wait for active queue to drain.		
	channel<-test
	<-sem
	fmt.Printf("ORIGINAL\n ")
	for i,_ := range test{
		fmt.Printf("%s ", test[i])		
	}
	
	fmt.Printf("\n")		
	for i:=0; i < 10; i++{
		go playGame(i)
	}

	//wait to finnish
	for i := 0; i < 10; i++ {
        <-wait    // wait for one task to complete
    }

	test = <-channel	
	fmt.Printf("BROKEN MESSAGE\n ")
	for i,_ := range test{
		fmt.Printf("%s ", test[i])		
	}
	fmt.Printf("\n")		
}
