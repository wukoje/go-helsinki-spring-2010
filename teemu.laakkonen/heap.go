//Binary heap

package main

import "fmt"

func main(){

	fmt.Println("Binary heap")

	bunk := newBheap(20)

	bunk.insert(45)
	bunk.insert(50)
	bunk.insert(13)
	bunk.insert(23)
	bunk.insert(12)
	bunk.insert(48)
	bunk.insert(53)
	bunk.insert(41)

	print(bunk)

	fmt.Println("pop: ",bunk.pop())

	print(bunk)

	laskeArvoa(bunk, 3,5)

	print(bunk)

}

type bheap struct{
	taulu [] int
	koko int // kertoo ensimmaisen tyhjan alkion

}


func newBheap (tila int) *bheap{

	jimmy := new(bheap)
	jimmy.koko=1
	jimmy.taulu = make ([] int, tila)
	jimmy.taulu[0]=42 // ensimmainen alkio ei kuulu kekoon
	return jimmy

}

func (h *bheap) insert (i int) {

	if h.koko<len(h.taulu) {

		h.taulu[h.koko]=i
		h.koko++
		nosta(h, h.koko-1)

	}

}

//poistaa koen paalimmaisen alkion

func (h *bheap) pop () int{

	poppi := h.taulu[1]

	h.taulu[1] = h.taulu[h.koko-1]
	h.koko--
	laske(h, 1)

	return poppi

}

func laskeArvoa (h *bheap, alkio, arvo int){

	h.taulu[alkio] = arvo
	nosta(h, alkio)

}

// nostaa annettua alkiota keossa tarpeen mukaan
func nosta(h *bheap, alk int){

	p := alk/2

	if p>0{
		if h.taulu[p] > h.taulu[alk]{
			h.taulu[p], h.taulu[alk] = h.taulu[alk], h.taulu[p]
			nosta(h, p)
		}
	}

}

//laskee annettua alkiota keossa tarpeen mukaan
func laske(h *bheap, alkio int){

	switch {

		case h.koko<=lC(alkio):

		case h.koko<=rC(alkio):
			if h.taulu[lC(alkio)]<h.taulu[alkio]{
				h.taulu[alkio], h.taulu[lC(alkio)] = h.taulu[lC(alkio)], h.taulu[alkio]
				laske(h, lC(alkio))
			}

		default:
			if h.taulu[lC(alkio)]<h.taulu[rC(alkio)]{
				if h.taulu[lC(alkio)]<h.taulu[alkio]{
					h.taulu[lC(alkio)], h.taulu[alkio] = h.taulu[alkio], h.taulu[lC(alkio)]
					laske(h, lC(alkio))
				}
			}else{
				if h.taulu[rC(alkio)]<h.taulu[alkio]{
					h.taulu[rC(alkio)], h.taulu[alkio] = h.taulu[alkio], h.taulu[rC(alkio)]
					laske(h, rC(alkio))
				}
			}
		}

}

//palauttaa alkion vanhemman
func prt(i int) int{
	return i/2
}

func lC (i int) int{
	return (i*2)
}

func rC (i int) int{
	return (i*2)+1
}

func print(h *bheap){

	i := 1

	for i<h.koko {
		fmt.Print(h.taulu[i], " ")
		i++
	}
	fmt.Println()

}























