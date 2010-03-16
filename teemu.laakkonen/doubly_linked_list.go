//Doubly linked list

package main

import "fmt"

func main(){
	frank:=createAlkio("linkedlist");
	frank.insert(1, "eka")
	frank.insert(2, "toka")
	frank.insert(3, "kolmas")
	frank.insert(3, "minapas")

	listPrint(frank)

	frank.delete(2)

	listPrint(frank)

}

type listAlkio struct{

	s string //Alkion sisalto

	n *listAlkio // seuraava
	p *listAlkio // edellinen
}

func createAlkio (s string) *listAlkio{
	la := new(listAlkio)
	la.s=s
	return la
}

func (l *listAlkio) insert (monesko int, s string){

	switch {

	case monesko==1:
		luusi := createAlkio(s)
		
		luusi.n=l.n
		l.n=luusi
		luusi.p=l
		luusi.s=s

	case l.n==nil:

	default: l.n.insert(monesko-1, s)
	}

}

func (l *listAlkio) delete (i int){
	if i == 1{
		l.n=l.n.n
	}else{
		l.n.delete(i-1)
	}
}

func listPrint(l *listAlkio){
	fmt.Println(l.s)
	if l.n!=nil{
		listPrint(l.n)
	}
}



















