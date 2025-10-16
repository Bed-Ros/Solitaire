package basic

import (
	"math/rand"
)

type CardsList []Card

func (l CardsList) IsPerfectlySorted() bool {
	if len(l) <= 1 {
		return true
	}
	curSuit := l[0].Suit
	for i := 1; i < len(l); i++ {
		if curSuit != l[i].Suit || l[i-1].Rank.Value-1 != l[i].Rank.Value {
			return false
		}
	}
	return true
}

type Stack struct {
	Cards    CardsList
	addRules []func(Stack, CardsList) bool
}

func NewStack(rulesToAdd ...func(Stack, CardsList) bool) Stack {
	var result Stack
	result.addRules = rulesToAdd
	return result
}

func (s Stack) Copy() Stack {
	result := Stack{
		Cards:    make(CardsList, len(s.Cards)),
		addRules: s.addRules,
	}
	copy(result.Cards, s.Cards)
	return result
}

func (s Stack) CanBeAdded(list CardsList) bool {
	for _, rule := range s.addRules {
		if !rule(s, list) {
			return false
		}
	}
	return true
}

type Deck CardsList

func (d *Deck) Take(n int, open bool) (CardsList, bool) {
	if n < 0 || n >= len(*d) {
		return nil, false
	}
	result := (*d)[len(*d)-n:]
	if open {
		for i := range result {
			result[i].Open = true
		}
	}
	*d = (*d)[:len(*d)-n]
	return CardsList(result), true
}

func (d *Deck) Shuffle() {
	for i := range *d {
		j := rand.Intn(len(*d))
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

func NewAllCardsOfSuit(suit Suit) CardsList {
	var result CardsList
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
