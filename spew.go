// spew: spew out text based on statistical analysis of some text.
package main

import (
	"os"
	"io"
	"fmt"
	"bufio"
	"log"
)

const NOWORD = "\n"
const NPREF = 2

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
		s := prefix[:]
		copy(s, s[1:])
		s[NPREF-1] = word
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	build(os.Stdin)
	fmt.Println(chain)
}
