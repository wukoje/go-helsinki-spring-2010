package main

import f "fmt"

func main(){

	a := []int{1, 7, 2, 3, 9, 4, 5, 6}
	f.Println(a)
	var r []int
	r = reverse(a)
	f.Println(r)

	b := []int{1, 4, 6, 7, 2, 3, 9, 4, 5, 6, 4, 7, 676, 554, 65, 322, 9, 42}
	f.Println(b)
	var t []int
	t = reverse(b)
	f.Println(t)
	
	c := []int{322, 9, 42}
	f.Println(c)
	var u []int
	u = reverse(c)
	f.Println(u)
	

}

func reverse(s [] int) []int {
	//re		:= s[0:]
	var i int	= 0
	var sl int	= len(s)-1
	var apu int	= 0 // apu int jota kaytetaanfor-silmukassa
	for i<(len(s)/2){
		
		apu = s[i]
		s[i] = s[sl]
		s[sl] = apu
	
		i++
		sl = sl-1
	}

	return s
}
 