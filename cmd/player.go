package cmd

import (
"github.com/JohnGeorge47/cardgame/pkg/card"
)

type Player struct {
	Maxval   int
	PlayerId uint64
	Cards    []card.Card
	Score    int
}

func NewPlayers(playerCount int) []Player {
	playersArray := make([]Player, playerCount)
	for i := 0; i < len(playersArray); i++ {
		playersArray[i].PlayerId = uint64(i + 1)
	}
	return playersArray
}

