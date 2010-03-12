package main
import (
	"fmt";
 	"strings";
	"rand";
	"time";
	"os";
	"bytes";
	"io";
)
var chars = map[string] string {
	"A": "0",	"B": "1",	"C": "2",
	"D": "3",	"E": "0",	"F": "1",
	"G": "2",	"H": "0",	"I": "0",
	"J": "2",	"K": "2",	"L": "4",
	"M": "5",	"N": "5",	"O": "0",
	"P": "1",	"Q": "2",	"R": "6",
	"S": "2",	"T": "3",	"U": "0",
	"V": "1",	"W": "0",	"X": "2",
	"Y": "0",	"Z": "2",
}
var channel = make(chan string)
var back = make(chan string)

func main() {
	message := "A common mistake that people make when trying to design something completely foolproof is to underestimate the ingenuity of complete fools"
	wordList := readText() 
	rand.Seed(time.Nanoseconds())
	for i:=0; i<20;i++ {
		go child(wordList)
	}
	for i:=0; i<20;i++ {
		channel <- message
		fmt.Println(message)
		message = <- back
	}
}

/* replaces one word in the message and passes it on to the next goruotine*/
func child(wordList []string) {
	message := <-channel;
	mesArray := strings.Fields(message)
	wordNumb := rand.Intn(len(mesArray))
	newWord := getNewWord(mesArray[wordNumb],wordList)
	if newWord != "" {
		listWordArr := strings.Fields(newWord)
		mesArray[wordNumb] = listWordArr[1]
	}
	ret := strings.Join(mesArray," ")
	back <- ret
}

/* 	return a new word to replace the old one if no word is found nothing is changed*/
func getNewWord(oldWord string,wordList []string) string {
	searchIndex := searchSoundex(wordList,calcSoundex(oldWord))
	if searchIndex != -1 {
		return wordList[searchIndex]
	} else {
		searchIndex := searchHamm(wordList,calcSoundex(oldWord))
		
		if searchIndex != -1 {
			return wordList[searchIndex]
		}
	}
	return ""
}

/*	Read text from file to create wordlist the list is formed as "soundexCode, realword"
	The array is sorted before it is returned */
func readText() []string{
	r, err := os.Open("alice.txt", os.O_RDONLY, 0600);
	defer r.Close()
	if err != nil {	
		return nil
	}
	var b bytes.Buffer
	io.Copy(&b, r)
	data := b.String();
	bookArray := strings.Fields(data)
	retSlice := make([]string,len(bookArray))
	for i:=0;i<len(retSlice);i++ {
		retSlice[i] = calcSoundex(bookArray[i])+" "+bookArray[i]
	}
	quicksort(retSlice,0,len(retSlice)-1)
	return retSlice 
}

/* 	calculate soundex code for and old word maximum lenght 4 chars */
func calcSoundex(oldWord string) string {
	word := strings.ToUpper(oldWord)
	wordArr := strings.Split(word,"",len(oldWord))
	soundex := make([]string,4)
	prev := ""
	j:=1

	if wordArr[0] != "[" && wordArr[0] != "*" && wordArr[0] != "#"{
		soundex[0] = wordArr[0]
	}
	for i:=1; i<len(oldWord); i++ {
		if wordArr[i] != prev {
			prev = wordArr[i]
			numb,_ := chars[wordArr[i]]
			if "0" != numb {
				soundex[j] = numb
				j++
				if  j==4 {
					break
				}
			}
		}
	}
	//add zeros if needed
	for k:=0; k<len(soundex);k++ {
		if soundex[k] == "" {
			soundex[k] = "0"
		}
	}
	ret := strings.Join(soundex,"")
	return ret
}
/* 	calculate the hamming distance between two words dos 
	not take into account multiple occurrences since the 
	functions will be used on soundex values where no  
	char occurs twice */
func hammingDist(a string, b string) int {
	retval :=0
	aArr := strings.Split(a,"",len(a))
	for i:=0;i<len(aArr);i++ {
		if strings.Count(b,aArr[i])!=1 {
			retval++
		}
	}
	return retval
}
/*	basic quicksort for strings in the wordlist*/
func quicksort(A []string,l int, r int) {
	if l<r {
		pIndex := (l+r)/2
		newIndex := partition(A,l,r,pIndex)
		quicksort(A,l,newIndex-1)
		quicksort(A,newIndex+1,r)
	}
}
func partition(A []string,l int, r int,index int) int{
	pValue := A[index]
	temp := A[index]
	A[index] = A[r]
	A[r] = temp
	sIndex := l
	for i:=l; i<r; i++ {
		if A[i] < pValue {
			temp = A[i]
			A[i] = A[sIndex]
			A[sIndex] = temp
			sIndex++
		}
	}
	temp = A[sIndex] 
	A[sIndex] = A[r]
	A[r] = temp
	return sIndex
}
/* 	search for a matching word in the sorted wordlist */
func searchSoundex(a []string, searchWord string) int {
	true := 0
	i:=0
	for i=0;i<len(a);i++ {
		listWordArr := strings.Fields(a[i])
		listCharArr := strings.Split(listWordArr[0],"",4)
		searchWordArr := strings.Split(searchWord,"",4)
		true = 0
		for j:=0;j<4;j++ {
			if listCharArr[j] == searchWordArr[j] {
				true++
			} else {
				true--
			}
		}
		if true == 4 {
			return i
		}
	}
	fmt.Println("fail soundex")
	return -1
}
func searchHamm(a []string, searchWord string) int {
	for i:=0; i<len(a);i++ {
		current := strings.Fields(a[i])
		val := hammingDist(current[0],searchWord)
		if  val == 0 || val == 1 {	
			return i
		}
	}
	fmt.Println("fail hamm")
	return -1
}





