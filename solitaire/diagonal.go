package solitaire

import (
	"Solitaire/basic"
	"github.com/fatih/color"
	"strings"
)

type Diagonal struct {
	Deck         basic.Deck
	Stacks       []basic.Stack
	ReserveStack basic.Stack
}

func NewDiagonal() Diagonal {
	result := Diagonal{Stacks: make([]basic.Stack, 9)}
	result.Deck = append(basic.New52Deck(), basic.New52Deck()...)
	result.Deck.Shuffle()
	openFrom := 8
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			curCard := result.Deck[0]
			if j >= openFrom {
				curCard.Open = true
			}
			result.Stacks[j] = append(result.Stacks[j], curCard)
			result.Deck = result.Deck[1:]
		}
		openFrom--
	}
	return result
}

func (d Diagonal) Print() error {
	blankLine13 := strings.Repeat(basic.BlankCell, 13)
	desc := color.New(basic.DescColor)
	_, err := desc.Println(blankLine13)
	if err != nil {
		return err
	}
	var index int
	for {
		_, err = desc.Print(basic.BlankCell)
		if err != nil {
			return err
		}
		stacksPrinted := false
		for _, stack := range d.Stacks {
			printed, err := stack.Print(index)
			if printed {
				stacksPrinted = printed
			}
			if err != nil {
				return err
			}
		}
		_, err = desc.Print(basic.BlankCell)
		if err != nil {
			return err
		}
		printed, err := d.ReserveStack.Print(index)
		if printed {
			stacksPrinted = printed
		}
		if err != nil {
			return err
		}
		_, err = desc.Println(basic.BlankCell)
		if err != nil {
			return err
		}
		index++
		if !stacksPrinted {
			break
		}
	}
	blankLine11 := strings.Repeat(basic.BlankCell, 11)
	_, err = desc.Print(blankLine11)
	if err != nil {
		return err
	}
	err = d.Deck.Print()
	if err != nil {
		return err
	}
	_, err = desc.Println(basic.BlankCell)
	if err != nil {
		return err
	}
	_, err = desc.Println(blankLine13)
	if err != nil {
		return err
	}
	return nil
}
