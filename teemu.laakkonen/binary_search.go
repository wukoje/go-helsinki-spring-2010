// Binary search

package main

import f "fmt"

func main(){
	a := []int{1, 2, 3, 4, 5, 6, 8, 10, 12, 13, 18, 43, 234, 753}
	e := 13
	var v int = binSrh(a, e)

	f.Println("Lista: ",a)
	f.Println("Etsittävä numero on ", e)
	if v==-1{
		f.Println("mutta numeroa ei löytynyt listasta")
	} else {
		f.Println(" ja se (",a[v],") oli listan alkiossa ", v)
	}

	a2 := []int{1, 2, 3, 4, 5, 6, 8, 10, 12, 13, 18, 43, 234, 753}
	e2 := 17
	var v2 int = binSrh(a2, e2)

	
	f.Println("Etsittävä numero on ", e2)
	if v2==-1{
		f.Println("mutta numeroa ei löytynyt listasta")
	} else {
		f.Println(" ja se (",a2[v2],") oli listan alkiossa ", v2)
	}

}

/* s on lista, josta etsitaan; etsit on etsittava numero ja palautta alkion*/
func binSrh(s [] int, etsit int) int{
	var r int = -1

	var vas int = 0
	var oik int = len(s)-1
	var kesk int

	for vas <= oik{
		kesk = vas + ((oik - vas)/2)

		switch {

		case etsit == s[kesk]:
			r = kesk
			vas = 45
			oik = 3

		case etsit < s[kesk]:
			oik = kesk -1

		case etsit > s[kesk]:
			vas = kesk +1

		}

	}

	return r
}
