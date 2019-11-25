// spew: spew out text based on statistical analysis of some text.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

const NOWORD = "\n"
const NPREF = 2
const MAXGEN = 100

type Prefix = [NPREF]string
type Suffix = []string

var chain map[Prefix]Suffix

func build(input io.Reader) {
	chain = make(map[Prefix]Suffix)

	prefix := Prefix{}
	for i := range prefix {
		prefix[i] = NOWORD
	}

	var s *bufio.Scanner = bufio.NewScanner(input)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		word := s.Text()
		chain[prefix] = append(chain[prefix], word)

		// Shift prefix elements to the left to make room for word
		copy(prefix[:], prefix[1:])
		prefix[NPREF-1] = word
	}

	// Add sentinel to suffix of "last" prefix
	chain[prefix] = append(chain[prefix], NOWORD)

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func spew(count int) {
	prefix := Prefix{}
	for i := range prefix {
		prefix[i] = NOWORD
	}

	// Build non-deterministic random number generator which we will use to
	// choose suffixs.
	rsrc := rand.NewSource(time.Now().UnixNano())
	rgen := rand.New(rsrc)

	for count > 0 {
		suffix := chain[prefix]
		word := suffix[rgen.Intn(len(suffix))]
		if word == NOWORD {
			break
		}
		fmt.Println(word)

		// Shift prefix elements to the left to make room for word
		copy(prefix[:], prefix[1:])
		prefix[NPREF-1] = word

		count--
	}
}

func main() {
	build(os.Stdin)
	spew(MAXGEN)
}
