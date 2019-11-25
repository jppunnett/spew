# Spew
Use Markov chain to spew text. Inspired by the book, "The Practice of Programming", by Brian W. Kernighan & Rob Pike.

## Install

`go get github.com/jppunnett/spew`

## Sample Input Files

Use your own input file or one of the following:
- brooks.txt: Some wisdom from Frederick P. Brookes, Jr.
- kjb.txt: King James Bible, from Project Gutenberg
- hod.txt: Heart of Darkness, from Project Gutenberg

## Example Usage

Powershell (ugh!): `PS> @(Get-Content .\brooks.txt | .\spew.exe) -join " "`

Linux: `$> ./spew < ./brooks.txt | fmt`
