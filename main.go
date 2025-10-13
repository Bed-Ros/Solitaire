package main

import (
	"Solitaire/solitaire"
	"bufio"
	"os"
)

func main() {
	q := solitaire.NewDiagonal()
	q.Print()
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
