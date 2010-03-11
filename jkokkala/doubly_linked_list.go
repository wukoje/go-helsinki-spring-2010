package main
import fmt "fmt"

type node struct {
    next* node
    prev* node
    value string
}
type list struct {
    first* node
    size int
}
func (l *list) insert(i int, e string) {
    var newnode node
    newnode.value=e
    if(l.first == nil) {
    	l.size=1
        l.first=&newnode
        newnode.next=&newnode
        newnode.prev=&newnode
    } else {
    	l.size++
        var curnode* node = l.first
        for j:=0 ; j<i;j++ {
            curnode=curnode.next
        }
        curnode.prev.next=&newnode
        newnode.next=curnode
        newnode.prev=curnode.prev
        curnode.prev=&newnode
        if(i == 0)  {
            l.first = &newnode
        }
    }
}
func(l* list) delete(i int) {
	l.size--
    var curnode* node = l.first
    for j:=0 ; j<i;j++ {
        curnode=curnode.next
    }
    if(curnode.next == curnode) {
        l.first=nil
    } else {
    	fmt.Printf("Poistetaan %s\n",curnode.value)
        curnode.prev.next=curnode.next
        curnode.next.prev=curnode.prev
        if( curnode == l.first) { l.first=curnode.next }
    }
}

func printlist(l* list) {
	var curnode* node=l.first
	for i:=0 ; i<l.size ; i++ {
		fmt.Printf("%s\n",curnode.value)
		curnode=curnode.next
    }
}

func main() {
	var testlist list
	testlist.insert(0,"A")
	testlist.insert(1,"C")
	testlist.insert(1,"B")
	testlist.insert(3,"F")
	testlist.insert(3,"D")
	testlist.insert(4,"E")
	testlist.insert(0,"0")
	printlist(&testlist)
	testlist.delete(0)
	testlist.delete(1)
	testlist.delete(3)
	printlist(&testlist)
}
