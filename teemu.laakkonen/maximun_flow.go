// Maximun Flow

// Palauttaa verkon lapi kulkevan reitin, jolla on suurin kapasiteetti
// Solmut on numeroitu int:ein ja kaarien kapasiteetit ovat int:jä
// Verkossa ei saa olla silmukoita

/* Algoritmi käy verkon ensin läpi syvyys syyntaisella haulla ja
tallentaa kaikki reitit sourcesta sinkiin. Tämän jälkeen reiteistä 
haetaan se, jonka kapasiteetti on suurin. En käyttänyt mitään 
erityistä algoritmia*/

package main

import "fmt"

func main(){

	fmt.Println("Maximun Flow - verkko:")
	fmt.Println("solmu1 -> solmu2 ja kaaren kapasiteetti")

	lester :=createGraph(6)
	insert(lester, 1, 2, 15)
	insert(lester, 1, 4, 10)
	insert(lester, 2, 5, 9)
	insert(lester, 4, 3, 11)
	insert(lester, 4, 5, 5)
	insert(lester, 3, 5, 10)

	fmt.Println(lester.e)

	cedric := createPK(20)
	kima :=  createPath(7)

	fmt.Println("Etsinta alkaa");

	mfs(lester, 1, 1, 5, kima, cedric)

	fmt.Println("etsinta paattynyt")

	fmt.Println("loytyneet reitit ovat")

	printKasa(cedric)

	greek:= maxFlowPath(cedric)

	//turha insert
	insertPath(greek, 1,1,1)

/*
	ziggy := createPath(6)
	insertPath(ziggy, 3, 5, 10)
	insertPath(ziggy, 4,3,11)
	fmt.Print("ziggy")
	fmt.Println(ziggy.e)
	nick := dublicatePath(ziggy)
	fmt.Print("nick")
	fmt.Println(nick.e)
	herc := createPK(3)
	insertPK(herc, ziggy)
	insertPK(herc, nick)
	printKasa(herc)
*/
}

type graph struct{
	e [][3] int // yksi rivi kuvaa yhta verkon kaarta
	koko int // ensimmainen alkio, jossa ei ole tavaraa
}

func createGraph (ekoko int) *graph{
	avon := new(graph)
	avon.e = make([][3] int, ekoko)
	avon.koko = 0

	return avon
}

// lisaa verkkoon solmun. ei tarkista onko verkossa tilaa
func insert (g *graph, lah, saap, cap int){
	g.e[g.koko][0]=lah // kaaren lahto solmu
	g.e[g.koko][1]=saap  // kaaren paatto solmu
	g.e[g.koko][2]=cap // kaaren kapasiteetti
	g.koko++
}

type path struct{

	e [][3]int // reittiin kuuluvat kaaret
	pituus int // ensimmainen alkio, jossa ei ole tavaraa

}

func createPath(pituus int) *path{

	sabotka := new(path)
	sabotka.e = make([][3] int, pituus)
	sabotka.pituus=0

	return sabotka

}

func dublicatePath(p *path) *path{

	stringer := new(path)
	stringer.e = make([][3]int, len(p.e))

	 i := 0
	for i<p.pituus {

		stringer.e[i][0] = p.e[i][0]
		stringer.e[i][1] = p.e[i][1]
		stringer.e[i][2] = p.e[i][2]
		i++
	}

	stringer.pituus = p.pituus

	return stringer

}

func insertPath(p *path, alk, lop, cap int){

	p.e[p.pituus][0]=alk
	p.e[p.pituus][1]=lop
	p.e[p.pituus][2]=cap
	p.pituus++

}

type pathKasa struct{

	k [] *path
	koko int

}

func createPK(i int) *pathKasa{

	omar := new(pathKasa)
	omar.k = make([] *path, i)
	omar.koko = 0

	return omar

}

func insertPK (k *pathKasa, p *path){

	k.k[k.koko] = p
	k.koko++

}

func (p *path) printPath(){

	for i:=0; i<p.pituus; i++{
		fmt.Print(p.e[i])
	}

}

func printKasa(k *pathKasa){

	for i:=0; i<k.koko; i++{

		k.k[i].printPath()
		fmt.Println()

	}

}

//palauttaa pathKasan, jossa on kaikki verkon pathit sourcesta sinkiin
//source on lopullisessa versiossa turha
func mfs (g *graph, node, source, sink int, p *path, k *pathKasa){

	//fmt.Println(node)
	//p.printPath()
	
	if node == sink{

		insertPK(k, p)

	}else{

		for i:=0; i<g.koko;i++{

			if g.e[i][0] == node{

				prez := dublicatePath(p)
				insertPath(prez, g.e[i][0], g.e[i][1], g.e[i][2])

				mfs(g, g.e[i][1], source, sink, prez, k)

			}


		}

	}

}

//palauttaa kasasta pathin, jolla on suurin max flow
func maxFlowPath(k *pathKasa) *path{

	result :=new(path)
	max := -1 //pathin max flow

	for i:=0;i<k.koko;i++{

		if pathFlow(k.k[i])>max{
			max = pathFlow(k.k[i])
			result = k.k[i]
		}

	}

	fmt.Println("Verkon suurin kapasiteetti on reitilla: ")
	result.printPath()
	fmt.Print(" jonka kantokyky on ")
	fmt.Println(max)

	return result


}

func pathFlow(p *path) int{

	min := 2147483647

	for i:=0;i<p.pituus;i++{
		if p.e[i][2]<min{
			min = p.e[i][2]
		}
	}

	return min

}





















