package deck

import (
	"fmt"
	"testing"
)

func TestSuitsValues(y *testing.T) {
	fmt.Println(Suits())
	fmt.Println(Values())
}

func TestCard(t *testing.T) {
	for _, s := range Suits() {
		for _, v := range Values() {
			card := NewCard(s, v)
			if card.val == "Joker" && card.suite != "" {
				t.Errorf("Card %v is impossible (Joker's suite must be ``)", card)
			}
		}
	}
}

func TestDeck(t *testing.T) {
	sd := StandardDeck()
	fmt.Println(sd)
	sd.Add(Card{val: "Joker"})
	sd.Add(Card{val: "Joker"})
	sd.Add(Card{val: "Joker"})
	sd.Remove(Card{val: "Joker"})
	fmt.Println(sd)
	if len(sd.cards) > 54 {
		t.Errorf("to many cards!")
	}
}

func TestShuffle(t *testing.T) {
	dk := StandardDeck()
	initialOrder := []string{}
	newOrder := []string{}
	for _, v := range dk.cards {
		initialOrder = append(initialOrder, v.String())
	}
	fmt.Println(dk)
	fmt.Println("SHUFFLE")
	dk.Shuffle()
	match := 0
	for _, v := range dk.cards {
		newOrder = append(newOrder, v.String())
	}
	fmt.Println(dk)
	for i := 0; i < len(dk.cards); i++ {
		if initialOrder[i] != newOrder[i] {
			match++
		}
	}
	if match < 1 {
		t.Errorf("deck was not shuffled (%v)", match)
	}
}
