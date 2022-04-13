package fool

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/tests/deck"
)

const (
	UNDEFINED_VALUE = iota
	GAMESTATE_SETUP
	GAMESTATE_END_TURN
	GAMESTATE_END_GAME
	GAMESTATE_ATTACK_PLAYER_1
	GAMESTATE_ATTACK_PLAYER_2
	WRONG_VALUE
)

type Pool struct {
	name  string
	cards []*deck.Card
}

func NewPool(name string) *Pool {
	return &Pool{name: name}
}

type Game struct {
	deck      *deck.Deck
	player    []*Player
	graveyard *deck.Deck
	gamestate int
}

func NewGame(players int) *Game {
	g := Game{}
	g.deck = deck.StandardDeck()
	for i := 0; i < players; i++ {
		g.player = append(g.player, NewPlayer(fmt.Sprintf("Player %v", i+1)))
	}
	g.graveyard = deck.NewDeck("Graveyard")
	g.gamestate = GAMESTATE_SETUP
	return &g
}

func (g *Game) Play() error {
	loop := 0
	for g.gamestate != GAMESTATE_END_GAME {
		fmt.Println("Loop", loop)
		if loop > 10 {
			break
		}
	}
	return nil
}

type Player struct {
	name           string
	inHand         []*deck.Card          //карты в руке - информацию о них знает только игрок
	knowledge_Card map[*deck.Card]string //каждый игрок может знать доступную ему часть информации о картах
}

func NewPlayer(name string) *Player {
	pl := Player{}
	pl.name = name
	pl.knowledge_Card = make(map[*deck.Card]string)
	return &pl
}

/*
Gameflow:
1. раздать карты
2. определить козырь
3. определить агрессора
4. агрессор атакует
5. битва
6. обновление рук и колоды
7. если игроков больше 1 - вернуться к 3
8. объявление дурака

*/
