package main

import "bytes"
import "crypto/sha1"
import "encoding/binary"
import "fmt"
import "hash"
import "rand"
import "time"

const (
  bitsPerElem int = 10; // bits per planned element in the filter
  seedCount int = 5; // amount of different seeded hash functions
)
var defaultEndianness binary.ByteOrder = binary.LittleEndian

// implementation of Bloom filter with SHA1 hashes
type Bloom struct {
  filter []byte
  seeds [seedCount]uint32
  size int
}

// adds an element to Bloom filter
func (bf *Bloom) Add (element []byte) {
  hashFunc := sha1.New()
  var truncHash uint64
  var byteBuffer *bytes.Buffer
  for i := 0; i < seedCount; i++ {
    hashCode := seededHash(hashFunc, element, bf.seeds[i])
    byteBuffer = bytes.NewBuffer(hashCode)
    binary.Read(byteBuffer, defaultEndianness, &truncHash)
    mappedBit := truncHash % uint64(bf.size)
    updateFilter(bf.filter, mappedBit)
  }
}

// Checks if given element has a chance of being in current filter
func (bf *Bloom) Query (element []byte) (found bool) {
  hashFunc := sha1.New()
  var truncHash uint64
  var byteBuffer *bytes.Buffer
  found = true
  for i := 0; i < seedCount; i++ {
    hashCode := seededHash(hashFunc, element, bf.seeds[i])
    byteBuffer = bytes.NewBuffer(hashCode)
    binary.Read(byteBuffer, defaultEndianness, &truncHash)
    
    mappedBit := truncHash % uint64(bf.size)
    found = checkFilter(bf.filter, mappedBit)
    if !found {
      break
    }
  }
  return found
}

// Updates the new bit to filter, if needed
func updateFilter (filter []byte, mappedBit uint64) {
  oldFilterByte := filter[findByte(mappedBit)]
  filter[findByte(mappedBit)] = oldFilterByte | filterByte(mappedBit)
}

// Helper; creates a byte with only the mapped bit true
func filterByte(mappedBit uint64) (retByte byte) {
  bitPosition := mappedBit % 8
  if bitPosition == 0 { // the leftmost bit is true
    retByte = 1 << 7
  } else {
    retByte = 1 << (uint(bitPosition) - 1)
  }
  return retByte
}

// Helper; finds, which byte has the mapped bit and returns its index
func findByte(mappedBit uint64) (byteInd uint64) {
  if mappedBit % 8 == 0 {
    byteInd = mappedBit / 8 - 1
  } else {
    byteInd = (mappedBit - (mappedBit % 8)) / 8
  }
  return byteInd
}

// Helper; calculates hash for given element with given seed
func seededHash(hashFunc hash.Hash, element []byte, seed uint32) []byte{
  hashFunc.Reset()
  binary.Write(hashFunc, defaultEndianness, seed)
  hashFunc.Write(element)
  hashCode := hashFunc.Sum()
  return hashCode
}

// Checks, if the filter contains the mapped bit; if found, returns true
func checkFilter(filter []byte, mappedBit uint64) (found bool) {
  // first we OR the bit with existing filter byte
  checkByte := filter[findByte(mappedBit)] | filterByte(mappedBit)
  // if the bytes are equal, the bit existed in the filter already
  return (checkByte == filter[findByte(mappedBit)])
}

// Creates a new Bloom filter.
// elemCount is an estimate of the amount of elements to be added in the filter
// if elemCount <= 0, returns nil
func New(elemCount int) *Bloom {
  if elemCount <= 0 {
    return nil
  }
  // initialize filter and bit array size
  var bloom *Bloom = new(Bloom)
  bloom.size = elemCount * bitsPerElem - (elemCount * bitsPerElem % 8) // whole bytes
  // initialize hash seeds  
  randomizer := rand.New(rand.NewSource(time.Seconds()))
  for i := 0; i < seedCount; i++ {
    bloom.seeds[i] = randomizer.Uint32()
  }
  // initialize bit array 
  bloom.filter = make([]byte, bloom.size / 8)
  
  return bloom
}

func main() {
  fmt.Println("Started")
  var testSlice = []byte{22, 33, 44}
  filter := New(30)
  fmt.Println("Created new filter!")
  filter.Add(testSlice)
  fmt.Println("Added following byte array:")
  fmt.Println(testSlice)
  fmt.Println("Checking if filter contains it")
  if filter.Query(testSlice) {
    fmt.Println("Success!")
  } else {
    fmt.Println("Failure")
  }
}
