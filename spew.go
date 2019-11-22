// spew: spew out text based on statistical analysis of some text.
package main

import (
	"fmt"
)

const NPREFIX = 2
const NOWORD = "\n"

type Prefix struct {
	prefix []string
}

type Suffix struct {
	suffix []string
}

func main() {
	fmt.Println("Ready to spew!")
	markovChain := build()
	// spew(markovChain, 50)
}

func build() map[Prefix]Suffix {
	// chain := make(map[string]Suffix)
	var chain map[Prefix]Suffix
	return chain
}

// func spew(chain map[Prefix]Suffix, count int) {
// 	fmt.Printf("chain: %v\n", chain)
// 	fmt.Printf("count: %v\n", counts)
// }