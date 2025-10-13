package basic

import (
	"github.com/fatih/color"
)

const (
	DescColor     = color.BgGreen
	BackColor     = color.BgWhite
	BackFontColor = color.FgBlue
	EmptyColor    = color.BgBlack

	BlankCell = "  "
)

var (
	Ace   = Rank{Symbol: 'A', Value: 1}
	Two   = Rank{Symbol: '2', Value: 2}
	Three = Rank{Symbol: '3', Value: 3}
	Four  = Rank{Symbol: '4', Value: 4}
	Five  = Rank{Symbol: '5', Value: 5}
	Six   = Rank{Symbol: '6', Value: 6}
	Seven = Rank{Symbol: '7', Value: 7}
	Eight = Rank{Symbol: '8', Value: 8}
	Nine  = Rank{Symbol: '9', Value: 9}
	Ten   = Rank{Symbol: '❿', Value: 10}
	Jack  = Rank{Symbol: 'J', Value: 11}
	Queen = Rank{Symbol: 'Q', Value: 12}
	King  = Rank{Symbol: 'K', Value: 13}
	Joker = Rank{Symbol: 'J', Value: 14}

	Hearts   = Suit{Symbol: '♥', Color: color.FgRed}
	Diamonds = Suit{Symbol: '♦', Color: color.FgRed}
	Clubs    = Suit{Symbol: '♣', Color: color.FgBlack}
	Spades   = Suit{Symbol: '♠', Color: color.FgBlack}
)

type Rank struct {
	Symbol rune
	Value  int
}

type Suit struct {
	Symbol rune
	Color  color.Attribute
}

type Card struct {
	Suit Suit
	Rank Rank
	Open bool
}

func (c Card) String() string {
	if c.Open {
		decorator := color.New(c.Suit.Color)
		return decorator.Sprintf("%c%c", c.Suit.Symbol, c.Rank.Symbol)
	} else {
		decorator := color.New(BackColor)
		return decorator.Sprint("  ")
	}
}
