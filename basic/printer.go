package basic

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

const (
	blankCell = "  "
	emptyCell = " ."
	cardBack  = "<>"

	descColor = color.BgGreen
)

var (
	desc  = color.New(descColor, color.FgWhite)
	back  = color.New(color.BgWhite, color.FgBlue)
	empty = color.New(color.BgBlack, color.FgWhite)
	red   = color.New(descColor, color.FgRed)
	black = color.New(descColor, color.FgBlack)
)

type Printer struct {
	err    error
	repeat int
}

func NewPrinter() *Printer {
	return &Printer{repeat: 1}
}

func (p *Printer) addError(err error) *Printer {
	p.err = errors.Join(p.err, err)
	return p
}

func (p *Printer) Error() error {
	return p.err
}

func (p *Printer) Repeat(n int) *Printer {
	if n >= 0 {
		p.repeat = n
	}
	return p
}

func (p *Printer) Ln() *Printer {
	for i := 0; i < p.repeat; i++ {
		_, err := fmt.Println()
		p.addError(err)
	}
	return p
}

func (p *Printer) genericPrint(decorator *color.Color, defStr string, str ...string) *Printer {
	text := defStr
	if str != nil {
		text = str[0][len(str[0])-2:]
	}
	for i := 0; i < p.repeat; i++ {
		_, err := decorator.Print(text)
		p.addError(err)
	}
	p.repeat = 1
	return p
}

func (p *Printer) Blank(str ...string) *Printer {
	return p.genericPrint(desc, blankCell, str...)
}

func (p *Printer) Empty(str ...string) *Printer {
	return p.genericPrint(empty, emptyCell, str...)
}

func (p *Printer) Back(str ...string) *Printer {
	return p.genericPrint(back, cardBack, str...)
}

func (p *Printer) Card(card Card) *Printer {
	if !card.Open {
		return p.Back()
	}
	var decorator *color.Color
	switch card.Suit.Color {
	case Red:
		decorator = red
		break
	case Black:
		decorator = black
		break
	default:
		return p.addError(errors.New("color mismatch"))
	}
	return p.genericPrint(decorator, fmt.Sprintf("%c%c", card.Suit.Symbol, card.Rank.Symbol))
}

func (p *Printer) Deck(deck Deck) *Printer {
	l := len(deck)
	if l == 0 {
		return p.Back()
	}
	return p.Back(strconv.Itoa(l))
}

func (p *Printer) Stack(stack Stack, index int) bool {
	if index == 0 && len(stack) == 0 {
		p.Empty()
		return true
	}
	if index < len(stack) {
		p.Card(stack[index])
		return true
	}
	p.Blank()
	return false
}
