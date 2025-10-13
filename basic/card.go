package basic

const (
	Red = 1 + iota
	Black
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

	Hearts   = Suit{Symbol: '♥', Color: Red}
	Diamonds = Suit{Symbol: '♦', Color: Red}
	Clubs    = Suit{Symbol: '♣', Color: Black}
	Spades   = Suit{Symbol: '♠', Color: Black}

	RedJoker   = Suit{Symbol: 'J', Color: Red}
	BlackJoker = Suit{Symbol: 'J', Color: Black}
)

type Color int

type Rank struct {
	Symbol rune
	Value  int
}

type Suit struct {
	Symbol rune
	Color  Color
}

type Card struct {
	Suit Suit
	Rank Rank
	Open bool
}
