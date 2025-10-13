package solitaire

import (
	"Solitaire/basic"
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
	printer := basic.NewPrinter()
	//Пустая строка
	printer.Repeat(13).Blank().Ln()
	//Основное поле со стаками
	var index int
	for {
		printer.Blank()
		stacksPrinted := false
		for _, stack := range d.Stacks {
			stacksPrinted = printer.Stack(stack, index) || stacksPrinted
		}
		printer.Blank()
		stacksPrinted = printer.Stack(d.ReserveStack, index) || stacksPrinted
		printer.Blank().Ln()
		index++
		if !stacksPrinted {
			break
		}
	}
	//Строка с колодой
	printer.Repeat(11).Blank().Deck(d.Deck).Blank().Ln()
	//Пустая строка
	return printer.Repeat(13).Blank().Ln().Error()
}
