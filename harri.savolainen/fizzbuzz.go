package main
import fmt "fmt" 

func main() { 
	for i := 1; i<=100; i++ {        
	switch i=i; true {
	case  i % 3 == 0 && i % 5 == 0:
	fmt.Printf("%d FizzBuzz\n",i)
	case  i % 5 == 0:
	fmt.Printf("%d Buzz\n",i)
	case  i % 3 == 0:   
	fmt.Printf("%d Fizz\n",i)
	default:
	fmt.Printf("%d\n",i)
	}
        }
} 
