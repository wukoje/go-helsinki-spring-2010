
package main
import (
      "fmt"
     "strings"
)

//hash function from the internetz, by Thomas Wang
func hash3(a int) int{
    a = (a ^ 61) ^ (a >> 16)
    a = a + (a << 3)
    a = a ^ (a >> 4)
    a = a * 0x27d4eb2d
    a = a ^ (a >> 15)
    return a;
}

func hash1(key []byte,salt int,le int) int {
   var rslt = salt
   for i :=0;i< len(key);i++{
   rslt += int(key[i])
  }
  rslt = hash3(rslt)%le
  //fmt.Println(rslt)
 return rslt
}

type bf struct {
    filter []bool
  }
  
func create(s int) *bf{
    z := make([]bool,s)
    return &bf{z}
  
}

func (a *bf) add (s string) {
  b := make([]byte,len(s))
  strings.NewReader(s).Read(b)

  a.filter[hash1(b,666,len(a.filter))] = true
  a.filter[hash1(b,999,len(a.filter))] = true
  a.filter[hash1(b,343,len(a.filter))] = true
  
}

func (a *bf) query (s string) bool{
  b := make([]byte,len(s))
  strings.NewReader(s).Read(b)
  return a.filter[hash1(b,666,len(a.filter))] && 
     a.filter[hash1(b,999,len(a.filter))] &&
     a.filter[hash1(b,343,len(a.filter))] 
}


func main () {
  fmt.Print("Initialized\n")
  a := create(18)
  //fmt.Println(a)
  a.add("s")
  a.add("anni")
  a.add("Gooogle")
  fmt.Println(a.query("zz"))
  fmt.Println(a.query("anni"))
  fmt.Println(a.query("Gooogle"))
 fmt.Println(a)
}
