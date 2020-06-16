package main

import (
	"fmt"
	"github.com/JohnGeorge47/cardgame/cmd"
	"github.com/JohnGeorge47/cardgame/pkg/card"
	"testing"
)

func TestFindScore(t *testing.T) {
	card1:=card.Card{
		Suit: 1,
		Rank: 1,
	}
	card2:=card.Card{
		Suit: 1,
		Rank: 1,
	}
	card3:=card.Card{
		Suit: 1,
		Rank: 1,
	}
	cardSlice:=make([]card.Card,3)
	cardSlice=append(cardSlice, card1,card2,card3)
	fmt.Println(card1,card2,card3)
	player1:=cmd.Player{
		Maxval:   0,
		PlayerId: 1,
		Cards:    cardSlice,
		Score:    0,
	}
	score:=FindScore(player1,0)
	if len(score)!=3{
		t.Errorf("got %d, want %d", len(score), 3)
	}
}

