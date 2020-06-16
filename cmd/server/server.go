package main

import (
	"fmt"
	"github.com/JohnGeorge47/cardgame/pkg/card"
	"github.com/JohnGeorge47/cardgame/cmd"
	"sort"
)

const (
	triplets = 100
	doubles  = 50
	sequenceofthree=75
)

type Game struct {
	players     []cmd.Player
	currentDeck []card.Card
}

//This function takes in number of cards and deals them to players
//this is then put in the player struct and inserts it in the cmd struct
func (g *Game) DealCards(cardsToDealperPlayer int) {
	for i := 0; i < len(g.players); i++ {
		g.players[i].Cards, g.currentDeck = g.currentDeck[0:cardsToDealperPlayer], g.currentDeck[cardsToDealperPlayer:len(g.currentDeck)]
		sort.Slice(g.players[i].Cards, func(j, k int) bool {
			return int(g.players[i].Cards[j].Rank) < int(g.players[i].Cards[k].Rank)
		})
	}
}

//Creates a new instance of the game
func NewBlackJackGame(playerCount int) *Game {
	return &Game{
		players:     cmd.NewPlayers(playerCount),
		currentDeck: card.New(card.Shuffle),
	}
}

func main() {
	start := NewBlackJackGame(4)
	roundno := 1
	for len(start.currentDeck) >= 3 {
		if roundno == 1 {
			start.DealCards(3)
		} else {
			start.DealCards(1)
		}
		for _, player := range start.players{
			fmt.Println("Player hands ",player.PlayerId,player.Cards)
		}
		start.logic()
		fmt.Println(start.players)
		if len(start.players) == 1 {
			fmt.Println("We have a winner Player number:", start.players[0].PlayerId)
			break
		}
		roundno++
	}
}

//Badly named but this is the core logic of the cmd
//It takes the cmd struct and performs the operations on top of it
//on the players array field present in it
func (g *Game) logic() {
	for i, player := range g.players {
		//Check if they are triplets
		scorearr:=FindScore(player,0)
		//Check if its a sequence of 3 consequetive numbers
		if len(scorearr)==3{
			g.players[i].Score = triplets
		}else {
			seqencescore:=FindScore(player,1)
			if len(seqencescore)==3{
				g.players[i].Score=sequenceofthree
			}else if len(scorearr)==2 {
				g.players[i].Score=doubles
			}else {
				g.players[i].Score=0
			}
		}
		//Check if the first value is an ace since it has highest value
		if int(g.players[i].Cards[0].Rank) == 1 {
			tmp := g.players[i].Cards[0]
			g.players[i].Cards[0] = g.players[i].Cards[len(player.Cards)-1]
			g.players[i].Cards[len(player.Cards)-1] = tmp
		}
		g.players[i].Maxval = CardValue(player.Cards[len(player.Cards)-1])
	}
	var victor []cmd.Player
	victor = append(victor,g.players[0])
	for i := 1; i < len(g.players); i++ {
		if g.players[i].Score > victor[len(victor)-1].Score {
			victor[0] = g.players[i]
			victor = victor[0:1]
		} else if g.players[i].Score == victor[len(victor)-1].Score {
			if g.players[i].Maxval == victor[len(victor)-1].Maxval {
				victor = append(victor, g.players[i])
			} else if g.players[i].Maxval > victor[len(victor)-1].Maxval {
				victor[0] = g.players[i]
				victor = victor[0:1]
			}
		}
	}
	g.players = victor
}

//This function takes a count parameter this count is basically to
//check if there are a sequence of 3 or pairs are present
func FindScore(player cmd.Player,count int) []card.Card {
	var scorearr []card.Card
	var curr_arr []card.Card
	curr_arr = append(curr_arr, player.Cards[0])
	scorearr = append(scorearr, player.Cards[0])
	for i := 1; i < len(player.Cards); i++ {
		if int(player.Cards[i].Rank) == int(curr_arr[len(curr_arr)-1].Rank)+count {
			curr_arr = append(curr_arr, player.Cards[i])
		} else {
			curr_arr[0] = player.Cards[i]
			curr_arr = curr_arr[0:1]
		}
		if len(curr_arr) >= len(scorearr) {
			scorearr = curr_arr
		}
	}
	return scorearr
}


func CardValue(c card.Card) int {
	if int(c.Rank) == 1 {
		return int(c.Rank) + 13
	}
	return int(c.Rank)
}

