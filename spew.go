// spew: spew out text based on statistical analysis of some text.
package main

import (
	"os"
	"io"
	"strings"
	"fmt"
	"bufio"
	"log"
)

const NOWORD = "\n"
const NPREFIX = 2

type Phrase struct {
	prefix []string
	suffix []string
}

func (p *Phrase) GetKey() string {
	return strings.Join(p.prefix, "+")
}

func newPhrase(word string, nprefix int) *Phrase {
	var phrase Phrase
	phrase.prefix = make([]string, nprefix)
	for i := range phrase.prefix {
		phrase.prefix[i] = word
	}
	return &phrase
}

type MarkovChain struct {
	chain map[string]*Phrase
	phrase *Phrase
}

func (mc *MarkovChain) add(word string) {
	// Find existing phrase

}

func (mc *MarkovChain) Build(input io.Reader) {
	mc.chain = make(map[string]*Phrase)
	mc.phrase = newPhrase(NOWORD, NPREFIX)

	var s *bufio.Scanner = bufio.NewScanner(input)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		mc.add(s.Text()) 
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func (mc *MarkovChain) Generate(count int) {
	fmt.Println(count, mc.chain)
}

func main() {
	var chain MarkovChain
	chain.Build(os.Stdin)
	chain.Generate(50)
}
