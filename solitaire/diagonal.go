package solitaire

import (
	"Solitaire/basic"
	"slices"
)

const (
	diagonalSize = 9
)

type Diagonal struct {
	Deck    basic.Deck
	Stacks  []basic.Stack
	History []Diagonal
}

func NewDiagonal() Diagonal {
	var result Diagonal
	//Колода
	result.Deck = append(basic.New52Deck(), basic.New52Deck()...)
	result.Deck.Shuffle()
	//Создаем основные стаки
	mainStacksRule := func(curStack basic.Stack, addList basic.CardsList) bool {
		//В пустой стак можно положить стак, начинающийся с любого короля
		return curStack.Cards == nil &&
			addList[0].Rank == basic.King ||
			//В стак с картами можно положить стак, начинающийся с карты на 1 ниже рангом и с такой-же мастью
			curStack.Cards != nil &&
				curStack.Cards[len(curStack.Cards)-1].Rank.Value-1 == addList[0].Rank.Value &&
				curStack.Cards[len(curStack.Cards)-1].Suit == addList[0].Suit

	}
	for i := 0; i < diagonalSize; i++ {
		result.Stacks = append(result.Stacks, basic.NewStack(mainStacksRule))
	}
	//Заполняем основные стаки
	openFrom := diagonalSize - 1
	for i := 0; i < diagonalSize; i++ {
		for j := 0; j < diagonalSize; j++ {
			curCard := result.Deck[0]
			if j >= openFrom {
				curCard.Open = true
			}
			result.Stacks[j].Cards = append(result.Stacks[j].Cards, curCard)
			result.Deck = result.Deck[1:]
		}
		openFrom--
	}
	//Резервный стак
	reserveStackRule := func(curStack basic.Stack, addList basic.CardsList) bool {
		tmpStack := append(curStack.Cards, addList...)
		return tmpStack.IsPerfectlySorted()
	}
	result.Stacks = append(result.Stacks, basic.NewStack(reserveStackRule))
	return result
}

func (d Diagonal) Copy() Diagonal {
	result := Diagonal{
		Deck:    make(basic.Deck, len(d.Deck)),
		History: make([]Diagonal, len(d.History)),
	}
	copy(result.Deck, d.Deck)
	for _, diagonal := range d.History {
		result.History = append(result.History, diagonal.Copy())
	}
	for _, stack := range d.Stacks {
		result.Stacks = append(result.Stacks, stack.Copy())
	}
	return result
}

func (d Diagonal) Print() error {
	printer := basic.NewPrinter()
	//Пустая строка
	printer.Repeat(diagonalSize + 4).Blank().Ln()
	//Основное поле со стаками
	var index int
	for {
		printer.Blank()
		stacksPrinted := false
		for i, stack := range d.Stacks {
			if i == len(d.Stacks)-1 {
				printer.Blank()
			}
			stacksPrinted = printer.Stack(stack, index) || stacksPrinted
		}
		printer.Blank().Ln()
		index++
		if !stacksPrinted {
			break
		}
	}
	//Строка с колодой
	printer.Repeat(diagonalSize + 2).Blank().Deck(d.Deck).Blank().Ln()
	//Пустая строка
	return printer.Repeat(diagonalSize + 4).Blank().Ln().Error()
}

func (d Diagonal) TryMoveBetweenStacks(fromStackIndex int, toStackIndex int, fromCardIndex int) (Diagonal, bool) {
	//Пропускаем пустой неправильные индексы и количество взятых карт
	if toStackIndex < 0 || toStackIndex >= len(d.Stacks) ||
		fromStackIndex < 0 || fromStackIndex >= len(d.Stacks) ||
		fromCardIndex < 0 || fromCardIndex >= len(d.Stacks[fromStackIndex].Cards) {
		return Diagonal{}, false
	}

	add := d.Stacks[fromStackIndex].Cards[fromCardIndex:]
	if d.Stacks[toStackIndex].CanBeAdded(add) {
		newD := d.Copy()
		newD.Stacks[toStackIndex].Cards = append(d.Stacks[toStackIndex].Cards, add...)
		newD.Stacks[fromStackIndex].Cards = newD.Stacks[fromStackIndex].Cards[:fromCardIndex]
		newD.History = append(newD.History, d)
		return newD, true
	}
	return Diagonal{}, false
}

func (d Diagonal) FindSteps() []Diagonal {
	var result []Diagonal
	//Для каждой открытой карты пробуем переместить ее в другой стак
	for fromStackIndex, stack := range d.Stacks {
		for cardIndex, card := range stack.Cards {
			if !card.Open {
				continue
			}
			for toStackIndex := range d.Stacks {
				if fromStackIndex == toStackIndex {
					continue
				}
				newDiagonal, ok := d.TryMoveBetweenStacks(fromStackIndex, toStackIndex, cardIndex)
				if ok {
					result = append(result, newDiagonal)
				}
			}
		}
	}
	//Раздача из колоды
	newDiagonal, ok := d.LayOutCardsFromDeck()
	if ok {
		result = append(result, newDiagonal)
	}
	return result
}

func (d Diagonal) LayOutCardsFromDeck() (Diagonal, bool) {
	if d.Deck == nil {
		return Diagonal{}, false
	}
	newD := d.Copy()
	for i := range newD.Stacks[:len(newD.Stacks)-1] {
		cards, ok := newD.Deck.Take(1, true)
		if !ok {
			break
		}
		newD.Stacks[i].Cards = append(newD.Stacks[i].Cards, cards...)
	}
	return newD, true
}

func (d Diagonal) Solved() bool {
	if d.Deck != nil || d.Stacks[len(d.Stacks)-1].Cards != nil {
		return false
	}
	for _, stack := range d.Stacks {
		if stack.Cards != nil &&
			(stack.Cards[0].Rank != basic.King || stack.Cards[len(stack.Cards)-1].Rank != basic.Ace || !stack.Cards.IsPerfectlySorted()) {
			return false
		}
	}
	return true
}

func (d Diagonal) Equal(another Diagonal) bool {
	if !slices.Equal(d.Deck, another.Deck) ||
		len(d.Stacks) != len(another.Stacks) {
		return false
	}
	for i := range d.Stacks {
		if !slices.Equal(d.Stacks[i].Cards, another.Stacks[i].Cards) {
			return false
		}
	}
	return true
}

func (d Diagonal) FindAllSolutions() []Diagonal {
	curBranches := d.FindSteps()
	//TODO
	return curBranches
}
