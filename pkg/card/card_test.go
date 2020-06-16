package card

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{
		Suit: Heart,
		Rank: Ace,
	})
	fmt.Println(Card{
		Suit: Spade,
		Rank: Two,
	})
	//Output:
	//Ace of Hearts
	//Two of Spades

}

func TestSort(t *testing.T) {
	cards := New(DefaultSort)
	expected := Card{
		Suit: Spade,
		Rank: Ace,
	}
	if cards[0] != expected {
		t.Error("Expected ace of spades but got:", cards[0])
	}
}
func TestNewSort(t *testing.T) {
	cards := New(Sort(Less))
	expected := Card{
		Suit: Spade,
		Rank: Ace,
	}
	if cards[0] != expected {
		t.Error("Expected ace of spades but got:", cards[0])
	}
}

func TestAddJoker(t *testing.T) {
	cards := New(AddJoker(4))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 4 {
		t.Error("Expected 4 Jokers, received:", count)

	}
}

func TestRemoveCards(t *testing.T) {
	cards := New(RemoveCards(Two))
	fmt.Println(cards)
	for _, c := range cards {
		if c.Rank == Two {
			t.Error("Expected no 2's but...")
		}
	}
}

func TestMultipleDecks(t *testing.T) {
	MultipleDecks(3)
}

