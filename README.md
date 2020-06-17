This is a simple card game with the following rules <br/>
- A trail (three cards of the same number) is the highest possible combination.
- The next highest is a sequence (numbers in order, e.g., 4,5,6. A is considered to have a value of 1).
- The next highest is a pair of cards (e.g.: two Kings or two 10s).
- If all else fails, the top card (by number value wins).
- If the top card has the same value, each of the tied players draws a single card from the deck until a winner is found.
- Only the newly drawn cards are compared to decide a tie. The top card wins a tie.

If you do not have go installed  theres a binary built for linux
you can run that or you can go run game.go

There are testcases written for the card deck which is in pkg/card.
pkg/card is a pkg which can be used to shuffle cards add jokers and multiple decks
