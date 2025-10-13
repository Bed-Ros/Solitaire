package basic

import (
	"github.com/fatih/color"
	"math/rand"
	"strconv"
)

type Deck []Card

func (d *Deck) Shuffle() {
	for i := range *d {
		j := rand.Intn(len(*d))
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

func (d *Deck) Print() error {
	l := len(*d)
	if l == 0 {
		_, err := color.New(DescColor).Print(BlankCell)
		if err != nil {
			return err
		}
		return nil
	}
	lStr := strconv.Itoa(l)
	_, err := color.New(BackColor, BackFontColor).Print(lStr[len(lStr)-2:])
	if err != nil {
		return err
	}
	return nil
}

type Stack []Card

func (s Stack) Print(index int) (bool, error) {
	if index == 0 && len(s) == 0 {
		_, err := color.New(EmptyColor).Print(BlankCell)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	if index < len(s) {
		_, err := color.New(DescColor).Print(s[index].String())
		if err != nil {
			return false, err
		}
		return true, nil
	}
	_, err := color.New(DescColor).Print(BlankCell)
	if err != nil {
		return false, err
	}
	return false, nil
}

func NewAllCardsOfSuit(suit Suit) []Card {
	var result []Card
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	for _, rank := range ranks {
		result = append(result, Card{Suit: suit, Rank: rank})
	}
	return result
}

func New52Deck() Deck {
	var result Deck
	suits := []Suit{Clubs, Diamonds, Hearts, Spades}
	for _, suit := range suits {
		result = append(result, NewAllCardsOfSuit(suit)...)
	}
	return result
}

func New54Deck() Deck {
	deck := New52Deck()
	deck = append(deck, Card{Suit: Suit{Symbol: 'J', Color: color.FgBlack}, Rank: Joker})
	deck = append(deck, Card{Suit: Suit{Symbol: 'J', Color: color.FgRed}, Rank: Joker})
	return deck
}
