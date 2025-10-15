package solitaire

import (
	"Solitaire/basic"
)

type Diagonal struct {
	Deck         basic.Deck
	Stacks       []basic.Stack
	ReserveStack basic.Stack
}

func (d Diagonal) Copy() Diagonal {
	result := Diagonal{
		Deck:         d.Deck,
		ReserveStack: d.ReserveStack,
	}
	copy(result.Deck, d.Deck)
	for _, stack := range d.Stacks {
		result.Stacks = append(result.Stacks, stack)
	}
	copy(result.ReserveStack, d.ReserveStack)
	return result
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

func (d Diagonal) TryAddToMainStack(stackIndex int, toAdd basic.Stack) (Diagonal, bool) {
	//Пропускаем пустой добавляемый стак и неправильный индекс
	if toAdd == nil || stackIndex < 0 || stackIndex >= len(d.Stacks) {
		return Diagonal{}, false
	}

	curStack := d.Stacks[stackIndex]
	curStackLastCard := curStack[len(curStack)-1]

	//В пустой стак можно положить стак, начинающийся с любого короля
	ruleApplies := curStack == nil && toAdd[0].Rank == basic.King ||
		//В стак с картами можно положить стак, начинающийся с карты на 1 ниже рангом и с такой-же мастью
		curStack != nil && curStackLastCard.Rank.Value-1 == toAdd[0].Rank.Value && curStackLastCard.Suit == toAdd[0].Suit

	if ruleApplies {
		newD := d.Copy()
		newD.Stacks[stackIndex] = append(d.Stacks[stackIndex], toAdd...)
		return newD, true
	}
	return Diagonal{}, false
}

func (d Diagonal) TryAddToReserveStack(toAdd basic.Stack) (Diagonal, bool) {
	//Пропускаем пустой добавляемый стак
	if toAdd == nil {
		return Diagonal{}, false
	}

	tmpStack := append(d.ReserveStack, toAdd...)
	if tmpStack.IsPerfectlySorted() {
		newD := d.Copy()
		newD.ReserveStack = tmpStack
		return newD, true
	}
	return Diagonal{}, false
}

func (d Diagonal) FindSteps() []Diagonal {
	var result []Diagonal
	//Для каждой открытой карты пробуем переместить ее в другой стак
	for fromStackIndex, stack := range append(d.Stacks, d.ReserveStack) {
		for cardIndex, card := range stack {
			if !card.Open {
				continue
			}
			curStack := stack[cardIndex:]
			for toStackIndex := range d.Stacks {
				if fromStackIndex == toStackIndex {
					continue
				}
				newDiagonal, ok := d.TryAddToMainStack(toStackIndex, curStack)
				if ok {
					result = append(result, newDiagonal)
				}
			}
			if fromStackIndex == len(d.Stacks) {
				continue
			}
			newDiagonal, ok := d.TryAddToReserveStack(curStack)
			if ok {
				result = append(result, newDiagonal)
			}
		}
	}
	return result
}
