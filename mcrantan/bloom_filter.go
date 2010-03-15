package main

import (
  "os"
  "fmt"
  "flag"
  "rand"
  "time"
  "math"
  "hash"
  "crypto/sha256"
)

type bloom struct {
  bits []uint64
  seed []byte
  h []hash.Hash
  k int
  lookup []uint64
}
 
func main() {

  var BITS, ELEM, PROB float64
  var ck, fk, k int

  flag.Float64Var(&BITS, "m", 10240, "Bloom filter size (default: 10240 bits)")
  flag.Float64Var(&PROB, "p", -1.0, "   Desired maximum probability of false positives\n            (default: 0.02)")
  flag.Float64Var(&ELEM, "n", 1024, " Desired number of elements to store (defult: 1024)")
  flag.Parse()

// If ELEM (-n) or PROB (-p) was set, let's adjust BITS (-m) to
// get better size requirements. 1.003 = 0.3% more for better PROB
// compliance (hive/vake-menetelmä). Clearly, the small difference
// in desired and resulting probability should be compensated using
// some log-function (the smaller the probability, the larger the
// error). Additionally, the numbers should probably be adjusted
// to minimize k (how many hashes are being used) to minimize time.
//
// Check/set values:
//  1. If desired n is different from 1024, adjust m to match p
//  2. If desired probability is set, adjust m
//  3. (I don't see clearly how to figure out if a particular
//      command line option has been set or is at its default
//      because flag package sets it)

  if (ELEM != 1024 || PROB > 0.0 ) && BITS == 10240 {
    if PROB <= 0 {
      PROB = 0.02                                                   // some sane default
    }
    var ratio float64 = math.Log(PROB) / math.Log(float64(0.6185))  // calculate m/n ratio for optimal k
    BITS = math.Ceil(1.003 * ELEM * ratio)                          //
    fmt.Printf("\nAdjusted filter size (-m) to %g.\n", BITS)
  }

  if PROB <= 0 {
    PROB = 0.02
  }
    
  var optk float64 = math.Ln2 * BITS / ELEM                         // Optimal k (number of hash functions)
  var minp float64 = math.Pow(0.6185, (BITS / ELEM))                // Best achievable probability
  
// Check if options make sense.
// Probability of a false positive >= 0.39 is considered useless.
// The ratio m/n is clearly too small; this is achievable by for example m/n=2, k=1.
// See http://pages.cs.wisc.edu/~cao/papers/summary-cache/node8.html , but note
// that the formula is incorrect (exponent kn/m should read -kn/m).

  if PROB < 0 || ELEM < 2 || (BITS / ELEM) < 2 {
    fmt.Printf("\nUse positive value for p, at least 2 items and m/n ratio >= 2.\n")
    os.Exit(0)
  }

  if PROB >= 0.39 {
    fmt.Printf("\nProbability limit is set to about 0.39 (39% chance of a false positive).\n")
    fmt.Printf("Please consider using smaller value.\n")
    os.Exit(0)
  }

  if optk < math.Ln2 * float64(2.0) {        // m/n=2, optimal optk=1.39, min p ~ 0.393
    fmt.Printf("\nCalculated optimal amount of hash functions as %g.\n", optk)
    fmt.Printf("Please choose less items or more bits (bigger filter size).\n")
    os.Exit(0)
  }

  if BITS < 64 {
    fmt.Printf("\nPlease use at least 64 bits.\n")
    os.Exit(0)
  }

  ck = int(math.Ceil(optk))
  fk = int(math.Floor(optk))
  k = chooseK(ck, fk, ELEM, PROB, BITS)

  if minp > PROB {
    fmt.Printf("\nMinimum probability possible with given constraints\n")
    fmt.Printf("(m bits=%g, n elements=%g) is about %0.5f (%e).\n", BITS, ELEM, minp, minp)
    fmt.Printf("This is greater than wanted probability %g.\n", PROB)
    fmt.Printf("Good bye.\n\n")
    os.Exit(0)
  }

// Could do some pretty-printing but that would probably require
// float64 conversion to string, then determining leading zeroes
// and then choosing how many numbers to show, based on user-entered
// or default PROB value.

  fmt.Printf("\nUsing k=%d for probability p=%g for false positives.\n", k, calcProb(k, ELEM, BITS))
  fmt.Printf("Size is m=%g bits (%d uint64) for n=%g elements.\n", BITS, uint(math.Ceil(BITS / 64)), ELEM)

// Seed RNG and create Bloom filter with predetermined values.
// Store seed for individual hash-functions and initialized them.
// Initialize lookup table for OR-operations.

  rand.Seed(time.Nanoseconds())
  mybloom := &bloom{ make([]uint64, uint(math.Ceil(BITS / 64))), make([]byte, k * 8), make([]hash.Hash, k), k, make([]uint64, 64) }
  mybloom.initLookup()

// For each hash function we've calculated we need, set
// 8 bytes of pseudorandomness as their precalculated seed.
// With SHA256, the hash function needs to be initialized
// with this value after each addition, so the seed must
// be stored within the Bloom struct.

  for i := 0; i < mybloom.k; i++ {
    var seed int64 = rand.Int63()
    mybloom.h[i] = sha256.New()
    for sb := 0; sb < 8; sb++ {
      mybloom.seed[i * 8 + sb] = byte(seed >> uint(sb * 8))
    }
  }
  mybloom.hashReset(mybloom.k)

// Demonstrate use. Put in n random byte[] permutations, then
// check with another n random byte[] and see how many false
// positives we get; repeat 5 times.

  var elems int = int(ELEM)
  var putbytes []byte = make([]byte, elems * 8)
  putBytes(putbytes, elems, mybloom)

  fmt.Printf("\nInserted %d elements into filter.\n", elems)

  checkBytes(putbytes, elems, mybloom)

  fmt.Printf("\nGenerated another %d random elements. Check how many of them\n", elems)
  fmt.Printf("are reported as being in the filter (there may be duplicates,\n")
  fmt.Printf("but we are using %d elements out of 64 bits of number space).\n", elems)
  fmt.Printf("\nRepeat five times.\n");

  for i := 0; i < 5; i++ {
    initBytes(putbytes, elems)
    checkBytes(putbytes, elems, mybloom)
  }

  fmt.Printf("Good bye.\n")
}



// Helper / struct functions begin here

func calcProb(k int, n float64, m float64) float64 {
  var p float64 = math.Pow((1 - math.Pow(math.E, float64(float64(-1 * k) * n / m))), float64(k))
  return p
}

func chooseK(ck int, fk int, ELEM float64, PROB float64, BITS float64) int {

// Choose less computatively expensive k, if it
// satisfies the desired probability of false positives
// or is better

  if fk > 0 {
    if calcProb(fk, ELEM, BITS) <= PROB || calcProb(fk, ELEM, BITS) < calcProb(ck, ELEM, BITS) {
      return fk
    } else {
      return ck
    }
  }

  return ck
}

func putBytes(putbytes []byte, count int, mybloom *bloom) {

// Shortcut bot both creating byte array of random 8 bit numbers
// and 'inserting' 8 byte slices into the filter.

  for i := 0; i < count; i++ {
    for j := 0; j < 8; j++ {
       putbytes[i * 8 + j] = byte(rand.Intn(256))
    }
    mybloom.add(putbytes[i * 8:i * 8 + 7])
  }
}

func initBytes(putbytes []byte, count int) {

// Fill given byte array with random 8 bit numbers.

  for i := 0; i < count; i++ {
    for j := 0; j < 8; j++ {
       putbytes[i * 8 + j] = byte(rand.Intn(256))
    }
  }
}

func getSum(sum []byte) uint64 {

// Convert byte array to uint64.

  var retval uint64 = 0
  var i int
  for i = 0; i < len(sum) - 1 && i < 8; i++ {
    retval |= uint64(sum[i])
    retval = retval << 8
  }
 
  retval |= uint64(sum[i]) 
  return retval
}

func checkBytes(putbytes []byte, count int, mybloom *bloom) {

// From given byte array, take slices and check if they are
// caught by the filter.

  var matches, nomatches int

  for i := 0; i < count; i++ {
    if mybloom.query(putbytes[i * 8:i * 8 + 7]) {
      matches++
    } else {
      nomatches++                    // nomatches used to make sure; count - matches could be used, too
    }
  }
  fmt.Printf("Checked %d elements. %d positive matches, %d negative matches.\n", count, matches, nomatches)
}

func (b *bloom) hashReset(k int) {

// Reset hash-functions back to their original state
// (hashing should typically be done with general hash function
// such as http://www.partow.net/programming/hashfunctions/#Download
// not requiring reset). This is trivial to change.
//
// When a value is 'put' into filter, all hash functions must be reset.
// When a query is made, only those that have been permuted need
// to be reset (when first 'not-in-filter' is seen, we can quit
// because false negatives are not allowed); that's why we have
// k int as parameter.

  for i := 0; i < k; i++ {
    b.h[i].Reset()
    for sb := 0; sb < 8; sb++ {
      b.h[i].Write(b.seed[i * sb:i * sb + 7])
    }
  }
}

func (b *bloom) initLookup() {

// The lookup table just saves ROL instructions and costs memory
// accesses (latency; though this should be somewhere near in L1/2).
// The table contains the AND/OR values (64 uin64's, with each
// bit set in turn).

  var orvalue uint64 = 1

  for i := 0; i < len(b.lookup); i++ {
    b.lookup[i] = orvalue
    orvalue = orvalue << 1
  }
}
    
func (b *bloom) add(element []byte) {

// Add to filter. Calculates the position in the uint64 array
// and the bit number within the uint64 to set.

  bits := len(b.bits) * 64

  for i := 0; i < b.k; i++ {
    b.h[i].Write(element)
    sum := getSum(b.h[i].Sum())
    posbyte := sum % uint64(bits) / 64
    posbit := sum % uint64(bits) % 64
    b.bits[posbyte] |= b.lookup[posbit]
  }

  b.hashReset(b.k) 
}

func (b *bloom) query(element []byte) bool {

// Query the filter if given byte array has been caught
// by the filter (with given probability). If we get negative
// reply, we reset those hashes we've inquired this
// far, otherwise reset all (and return true for presence).

  bits := uint64(len(b.bits) * 64)
  count := 0

  for count = 0; count < b.k; count++ {
    b.h[count].Write(element)
    sum := getSum(b.h[count].Sum())
    posbyte := (sum % bits) / 64
    posbit := sum % bits % 64
    if b.bits[posbyte] & b.lookup[posbit] == 0 {
      b.hashReset(count)
      return false
    }
  }

  b.hashReset(b.k)
  return true
}

