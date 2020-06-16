//go:generate stringer -type=Suit,Rank
package card

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint
type Rank uint

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

var suits = []Suit{Spade, Diamond, Club, Heart}

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, s := range suits {
		for i := minRank; i <= maxRank; i++ {
			cards = append(cards, Card{
				Suit: s,
				Rank: i,
			})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return CardRank(cards[i]) < CardRank(cards[j])
	}
}

func Sort(less func(cards []Card) func(i, j int) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	shuffled := make([]Card, len(cards))
	perm := r.Perm(len(shuffled))
	for i, shuff := range perm {
		shuffled[i] = cards[shuff]
	}
	return shuffled
}

func CardRank(c Card) int {
	r := int(c.Suit)*int(maxRank) + int(c.Rank)
	return r
}

func AddJoker(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			j := Card{
				Suit: Joker,
				Rank: Rank(i),
			}
			cards = append(cards, j)
		}
		return cards
	}
}

func RemoveCards(cardTypes ...Rank) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for _, cardtype := range cardTypes {
			cards = Remove(cardtype, cards)
		}
		return cards
	}
}

func Remove(cardType Rank, cards []Card) []Card {
	newDeck := make([]Card, len(cards))
	j := 0
	for i := 0; i < len(cards); i++ {
		if cardType != cards[i].Rank {
			newDeck[j] = cards[i]
			j++
		}
	}
	return newDeck
}

func MultipleDecks(deckCount int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		if deckCount == 1 {
			return cards
		}
		for i := 0; i < deckCount-1; i++ {
			cards = append(cards, cards...)
		}
		return cards
	}
}

