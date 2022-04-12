package deck

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/dice"
)

func Suits() []string {
	return []string{
		"\u2663", //'♣'
		"\u2666", //'♦'
		"\u2665", //'♥'
		"\u2660", //'♠'
	}
}

func Values() []string {
	return []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
}

type Card struct {
	suite    string
	val      string
	location string
	reverse  bool
	shirt    bool
}

func NewCard(suit, value string) Card {
	c := Card{}
	c.suite = suit
	c.val = value
	if c.val == "Joker" {
		c.suite = ""
	}
	c.location = "Undistributed"
	return c
}

func (c *Card) String() string {
	return fmt.Sprintf("[%v%v]", c.val, c.suite)
}

type Deck struct {
	cards []Card
	name  string
}

func NewDeck(name string) *Deck {
	d := Deck{}
	d.name = name
	return &d
}

func StandardDeck() *Deck {
	d := Deck{}
	for _, s := range Suits() {
		for _, v := range Values() {
			card := NewCard(s, v)
			d.Add(card)
		}
	}
	return &d
}

func (d *Deck) String() string {
	str := "Deck: "
	for _, v := range d.cards {
		str += v.String()
	}
	return str
}

func (d *Deck) Shuffle() {
	random := dice.New()
	totalCards := len(d.cards)
	for lap := 0; lap < 1000; lap++ {
		r1 := random.Roll("1d" + fmt.Sprintf("%v", totalCards)).DM(-1).Sum()
		r2 := random.Roll("1d" + fmt.Sprintf("%v", totalCards)).DM(-1).Sum()
		for r1 == r2 {
			r2 = random.Roll("1d" + fmt.Sprintf("%v", totalCards)).DM(-1).Sum()
		}
		d.cards[r1], d.cards[r2] = d.cards[r2], d.cards[r1]
	}
}

func (d *Deck) Add(c Card) {
	d.cards = append(d.cards, c)
}

func (d *Deck) Remove(c Card) {
	for i, card := range d.cards {
		if card.String() == c.String() {
			d.cards = append(d.cards[0:i], d.cards[i+1:]...)
			return
		}
	}
}

func MoveCardFromTo(card Card, deck1, deck2 *Deck) error {
	for _, crd := range deck1.cards {
		if crd.String() != card.String() {
			continue
		}
		card.location = deck2.name
		deck1.Remove(card)
		deck2.Add(card)
	}
	return nil
}
