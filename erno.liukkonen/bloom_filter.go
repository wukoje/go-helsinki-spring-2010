package main

import "fmt"

type Filter struct {
	data []bool
	num uint
}

func main() {
	var bloom *Filter = newFilter(887) 
	bloom.addValue("test")
	fmt.Printf("%v\n", bloom.queryFilter("test"))
}

func newFilter(s int) *Filter {
	num := uint(float(s)*0.1*0.7)
	return &Filter{make([]bool, s), num}
}

func (bloom *Filter) addValue(d string) {
	for i := uint(1); i <= bloom.num; i++ {
		bloom.data[hash(d, uint(len(bloom.data)), i)] = true
	}
}

func (bloom *Filter) queryFilter(d string) (hasit bool) {
	hasit = true
	for i := uint(1); i <= bloom.num; i++ {
		hasit = hasit && bloom.data[hash(d, uint(len(bloom.data)), i)]
	}
	return
}

func hash(d string, n, num uint) uint {
	var temp uint = 0
	for _, b := range []byte(d) {
		temp = temp + uint(b)
	}
	temp2 := 1 + (temp % (n-1))
	return (temp % n + num * temp2) % n
}

