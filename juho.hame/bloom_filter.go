/**
2 bloom_filter.go
Implement a bloom filter that supports add and query
*/

package main

import "fmt"
import hash "hash/adler32"
import hash2 "hash/crc32"

type bloomfilter struct {
	bits	[]bool
}

func (bf *bloomfilter) add(s string) {
	sum := make([]uint8, len(s))
	for i := 0; i < len(s); i++ {
		sum[i] = s[i]
	}

	bf.bits[hash.Checksum(sum)/1000] = true
	bf.bits[hash2.ChecksumIEEE(sum)/100000] = true
}

func (bf *bloomfilter) query(s string) bool {
	sum := make([]uint8, len(s))
	for i := 0; i < len(s); i++ {
		sum[i] = s[i]
	}

	if bf.bits[hash.Checksum(sum)/1000] && bf.bits[hash2.ChecksumIEEE(sum)/100000] {
		return true
	} 

	return false
}

func main() {

	bf := new(bloomfilter)
	bf.bits = make([]bool, 1000000)
	for i,_ := range bf.bits {
		bf.bits[i] = false
	}

	bf.add("juho")
	bf.add("great")

	if bf.query("juho") {
		fmt.Printf("juho found\n")
	}

	if bf.query("ohju") {
		fmt.Printf("ohju found\n")
	}

	if !bf.query("yuho") {
		fmt.Printf("yuho not found\n")
	}
}
