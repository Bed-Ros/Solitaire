package main

import (
	"Solitaire/solitaire"
	"bufio"
	"os"
	"os/exec"
)

func Clear() error {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func WaitInput() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}

func main() {
	diagonal := solitaire.NewDiagonal()
	steps := diagonal.FindSteps()
	for _, step := range steps {
		err := diagonal.Print()
		if err != nil {
			panic(err)
		}
		err = step.Print()
		if err != nil {
			panic(err)
		}
		WaitInput()
		err = Clear()
		if err != nil {
			panic(err)
		}
	}
}
