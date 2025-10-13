package basic

import (
	"math/rand"
)

type Stack []Card

type Deck []Card

func (d *Deck) Shuffle() {
	for i := range *d {
		j := rand.Intn(len(*d))
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
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
	deck = append(deck, Card{Suit: RedJoker, Rank: Joker})
	deck = append(deck, Card{Suit: BlackJoker, Rank: Joker})
	return deck
}
