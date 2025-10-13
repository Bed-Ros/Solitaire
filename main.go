package main

import (
	"Solitaire/solitaire"
	"bufio"
	"os"
)

func main() {
	q := solitaire.NewDiagonal()
	err := q.Print()
	if err != nil {
		panic(err)
	}
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
