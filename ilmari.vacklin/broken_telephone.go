package main

import "fmt"
import "rand"
import "strings"
import "strconv"
import "io/ioutil"
import "container/list"
import "json"
import "os"

const use_hamming = false

func soundex(s string) string {
    cs := make([]int, len(s), len(s))

    i := 0
    for _, c := range s {
        cs[i] = c
        i++
    }

    /* If "h" or "w" separate two consonants with the same soundex code,
       change the consonant on the right to "h" and repeat until the codes
       differ. */
    for i := 0; i < len(cs)-2; i++ {
        c := cs[i]
        n := cs[i+1]

        if n == 'h' || n == 'w' {
            j := i+2
            if j >= len(cs) { break }
            m := cs[j]
            LOOP: for j < len(cs)-1 && (m == 'w' || m == 'h') {
                j++
                m = cs[j]
            }
            if j > i && j < len(cs)-1 && code(m) == code(c) && code(m) != 0 {
                cs[j] = 'h'
                m = 'h'
                goto LOOP
            }
        }
    }

    for i, v := range cs {
        cs[i] = code(v)
    }

    ds := make([]int, len(cs), len(cs))

    for i := 0; i < len(cs); {
        j := i
        for j < len(cs) && cs[j] == cs[i] { j++ }
        if j < len(cs) {
            ds[i] = cs[j]
        }
        i = j
    }

    es := make([]int, len(ds), len(ds))

    j := 0

    for i := 0; i < len(ds); i++ {
        if ds[i] != 0 {
            es[j] = ds[i]
            j++
        }
    }

    l := len(es)
    rs := make([]string, 4)
    if l < 4 {
        for i := 0; i < l; i++ {
            rs[i] = strconv.Itoa(es[i])
        }
        for i := l; i < 4; i++ {
            rs[i] = "0"
        }
        return strings.Join(rs, "")
    }

    for i := 0; i < 4; i++ {
        rs[i] = strconv.Itoa(es[i])
    }

    return strings.Join(rs, "")
}

func code(c int) int {
    switch c {
        case 'b', 'f', 'p', 'v': return 1
        case 'c', 'g', 'j', 'k', 'q', 's', 'x', 'z': return 2
        case 'd', 't': return 3
        case 'l': return 4
        case 'm', 'n': return 5
        case 'r': return 6
    }
    return 0
}

func telephone(in chan string, valid map[string] []string) chan string {
    out := make(chan string)
    go func() {
        msg := <-in
    
        words := strings.Fields(msg)

        i := int(rand.Float() * float(len(words)))
        w := words[i]
        s := soundex(w)

        if ws, ok := valid[s]; ok {
            var c string
            if use_hamming {
                best := -1
                besth := -1
                for j := 0; j < len(ws); j++ {
                    if best == -1 && ws[j] != w {
                        best = j
                        besth = hamming(ws[j], w)
                        continue
                    } 
                    if b := hamming(w, ws[j]); ws[j] != w && b < besth {
                        best = j
                        besth = b
                    }
                }
                c = ws[best]
            } else {
                j := int(rand.Float() * float(len(ws)))
                c = ws[j]
            }
                
            fmt.Println("Replacing", w, "with", c)
            words[i] = c

            out <- strings.Join(words, " ")
        } else {
            out <- msg
        }
    }()
    return out
}

func hamming(a string, b string) int {
    t := 0
    la := len(a)
    lb := len(b)
    if lb < la {
        lb, la = la, lb
        a, b = b, a
    }
    lendiff := lb - la
    
    ca := strings.Split(a, "", 0)
    cb := strings.Split(b, "", 0)

    for i := 0; i < len(ca); i++ {
        if i >= len(cb) { break }
        if ca[i] != cb[i] { t++ }
    }

    return t + lendiff
}

func main() {
    rand.Seed(int64(os.Getpid()))
    in := make(chan string)
    c, ok := ioutil.ReadFile("/usr/share/dict/words")
    if ok != nil { return }
    cs := strings.Fields(string(c))
    ss := make(map[string] []string)
    if f, _ := ioutil.ReadFile("soundex.json"); f != nil {
        fmt.Println("Loading wordlist...")
        if ok, err := json.Unmarshal(string(f), ss); !ok {
            fmt.Println("Can't open soundex.json:", err)
            return
        }
    } else {
        fmt.Println("Soundexing wordlist...")
        sl := make(map[string] *list.List)
        l := len(cs)
        for i, c := range cs {
            c = strings.ToLower(c)
            if i % 20 == 0 {
                fmt.Printf("\r%d/%d", i, l)
            }
            sc := soundex(c)
            if _, ok := sl[sc]; !ok {
                sl[sc] = list.New()
            }
            sl[sc].PushBack(c)
        }

        sc := make(map[string] int)
        for k, v := range sl {
            l = v.Len()
            if _, ok := ss[k]; !ok { ss[k] = make([]string, l, l) }
            if _, ok := sc[k]; !ok { sc[k] = 0 }
            for n := v.Front(); n != nil; n = n.Next() {
                w := n.Value.(string)
                ss[k][sc[k]] = w
                sc[k]++
            }
        }

        fmt.Print("\rWriting soundex.json... ")
        if f, _ := os.Open("soundex.json", os.O_WRONLY | os.O_CREAT, 0600); f != nil {
            json.Marshal(f, ss)
        } 
    }
    fmt.Println("done.")
    out := telephone(in, ss)
    for i := 0; i < 20; i++ {
        out = telephone(out, ss)
    }
    b, _ := ioutil.ReadAll(os.Stdin)
    in <- string(b)
    msg := <-out
    
    fmt.Println(msg)
}
