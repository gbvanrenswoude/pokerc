package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Suit int
type Rank int

const (
	BS Suit = iota
	RH
	BD
	RC
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

type Card struct {
	suit Suit
	rank Rank
}

func parseCard(s string) Card {
	suit := s[:2]
	rank := s[2:]

	var cardSuit Suit
	switch suit {
	case "BS":
		cardSuit = BS
	case "RH":
		cardSuit = RH
	case "BD":
		cardSuit = BD
	case "RC":
		cardSuit = RC
	default:
		fmt.Println("Invalid suit:", suit)
		os.Exit(1)
	}

	cardRank, err := strconv.Atoi(rank)
	if err != nil {
		switch rank {
		case "A":
			cardRank = int(Ace)
		case "J":
			cardRank = int(Jack)
		case "Q":
			cardRank = int(Queen)
		case "K":
			cardRank = int(King)
		default:
			fmt.Println("Invalid rank:", rank)
			os.Exit(1)
		}
	} else {
		if cardRank < int(Two) || cardRank > int(King) {
			fmt.Println("Invalid rank:", rank)
			os.Exit(1)
		}
	}

	return Card{suit: cardSuit, rank: Rank(cardRank)}
}


func deal(deck *[]Card) Card {
	i := rand.Intn(len(*deck))
	card := (*deck)[i]
	*deck = append((*deck)[:i], (*deck)[i+1:]...)
	return card
}

func simulate(players int, hand []Card, community []Card, trials int) float64 {
	wins := 0

	for t := 0; t < trials; t++ {
		deck := createDeck(hand, community)

		comm := make([]Card, len(community))
		copy(comm, community)
		for i := len(comm); i < 5; i++ {
			comm = append(comm, deal(&deck))
		}

		bestHand := handValue(hand, comm)
		winner := true

		for i := 0; i < players-1; i++ {
			opponentHand := []Card{deal(&deck), deal(&deck)}
			opponentValue := handValue(opponentHand, comm)
			if opponentValue >= bestHand {
				winner = false
				break
			}
		}

		if winner {
			wins++
		}
	}

	return float64(wins) / float64(trials)
}

func hasFourOfAKind(cards []Card) (bool, Rank) {
    rankCounts := make(map[Rank]int)
    for _, card := range cards {
        rankCounts[card.rank]++
        if rankCounts[card.rank] == 4 {
            return true, card.rank
        }
    }
    return false, 0
}

func hasFullHouse(cards []Card) (bool, Rank, Rank) {
    rankCounts := make(map[Rank]int)
    for _, card := range cards {
        rankCounts[card.rank]++
    }

    threeOfAKind := Rank(0)
    pair := Rank(0)

    for r, count := range rankCounts {
        if count == 3 {
            threeOfAKind = r
        } else if count == 2 {
            pair = r
        }
    }

    if threeOfAKind > 0 && pair > 0 {
        return true, threeOfAKind, pair
    }
    return false, 0, 0
}

func hasThreeOfAKind(cards []Card) (bool, Rank) {
    rankCounts := make(map[Rank]int)
    for _, card := range cards {
        rankCounts[card.rank]++
        if rankCounts[card.rank] == 3 {
            return true, card.rank
        }
    }
    return false, 0
}

func hasTwoPair(cards []Card) (bool, Rank, Rank) {
    rankCounts := make(map[Rank]int)
    for _, card := range cards {
        rankCounts[card.rank]++
    }

    firstPair := Rank(0)
    secondPair := Rank(0)

    for r, count := range rankCounts {
        if count == 2 {
            if firstPair == 0 {
                firstPair = r
            } else {
                secondPair = r
                break
            }
        }
    }

    if firstPair > 0 && secondPair > 0 {
        return true, firstPair, secondPair
    }
    return false, 0, 0
}

func hasOnePair(cards []Card) (bool, Rank) {
    rankCounts := make(map[Rank]int)
    for _, card := range cards {
        rankCounts[card.rank]++
        if rankCounts[card.rank] == 2 {
            return true, card.rank
        }
    }
    return false, 0
}

func highCard(cards []Card) Rank {
    highest := Rank(0)
    for _, card := range cards {
        if card.rank > highest {
            highest = card.rank
        }
    }
    return highest
}


func hasFlush(cards []Card) bool {
	suitCounts := make(map[Suit]int)
	for _, card := range cards {
		suitCounts[card.suit]++
		if suitCounts[card.suit] >= 5 {
			return true
		}
	}
	return false
}

func hasStraight(cards []Card) bool {
	rankCounts := make(map[Rank]int)
	for _, card := range cards {
		rankCounts[card.rank] = 1
	}

	if rankCounts[Ace] > 0 {
		rankCounts[King+1] = 1 // Ace can also be considered as the highest card
	}

	consecutive := 0
	for r := Ace; r <= King+1; r++ {
		if rankCounts[r] > 0 {
			consecutive++
			if consecutive >= 5 {
				return true
			}
		} else {
			consecutive = 0
		}
	}

	return false
}


func handValue(hand, community []Card) int {
	cards := append(hand, community...)

	// Check for Straight Flush
	flush := hasFlush(cards)
	straight := hasStraight(cards)
	if flush && straight {
		return 9
	}

	// Check for Four of a Kind
	fourOfAKind, _ := hasFourOfAKind(cards)
	if fourOfAKind {
		return 8
	}

	// Check for Full House
	fullHouse, _, _ := hasFullHouse(cards)
	if fullHouse {
		return 7
	}

	// Check for Flush
	if flush {
		return 6
	}

	// Check for Straight
	if straight {
		return 5
	}

	// Check for Three of a Kind
	threeOfAKind, _ := hasThreeOfAKind(cards)
	if threeOfAKind {
		return 4
	}

	// Check for Two Pair
	twoPair, _, _ := hasTwoPair(cards)
	if twoPair {
		return 3
	}

	// Check for One Pair
	onePair, _ := hasOnePair(cards)
	if onePair {
		return 2
	}

	// return the High Card sadness
	return 1
}


func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: pokerc <num_players> <hand> [community_cards]")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	players, err := strconv.Atoi(os.Args[1])
	if err != nil || players < 2 || players > 8 {
		fmt.Println("Invalid number of players, 2 to 8 players allowed.")
		os.Exit(1)
	}

	handStr := strings.Split(os.Args[2], " ")
	hand := make([]Card, 0, 2)
	for _, cardStr := range handStr {
		hand = append(hand, parseCard(cardStr))
	}

	community := make([]Card, 0, 5)
	if len(os.Args) > 4 {
		for i := 3; i < len(os.Args); i++ {
			community = append(community, parseCard(os.Args[i]))
		}
	}

	communityLen := len(community)
	if communityLen != 0 && communityLen != 4 && communityLen != 5 && communityLen != 6 {
		fmt.Println("Invalid number of community cards. Must be 0, 3, 4, or 5.")
		os.Exit(1)
	}

	trials := 100000
	winProb := simulate(players, hand, community, trials)
	fmt.Printf("Winning probability: %.2f%%\n", winProb*100)
}



func createDeck(exclude ...[]Card) []Card {
	deck := make([]Card, 0, 52)
	excludeMap := make(map[Card]bool)

	for _, excludeCards := range exclude {
		for _, excludeCard := range excludeCards {
			if excludeMap[excludeCard] {
				fmt.Printf("Error: Duplicate card found (%s%d)\n", excludeCard.suit, excludeCard.rank)
				os.Exit(1)
			}
			excludeMap[excludeCard] = true
		}
	}

	for _, s := range []Suit{BS, RH, BD, RC} {
		for r := Ace; r <= King; r++ {
			card := Card{suit: s, rank: r}
			if !excludeMap[card] {
				deck = append(deck, card)
			}
		}
	}
	return deck
}
