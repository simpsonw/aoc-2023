package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"sort"
)
var cardValue = map[rune]int {
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveofAKind
)

type Wager struct {
	cards string
	bid int
	hand HandType
}

func (w *Wager) less(w2 *Wager) bool  {
	if w.hand == w2.hand {
		for k, v := range w.cards {
			w2Card := rune(w2.cards[k])
			if v != w2Card {
				return cardValue[v] < cardValue[w2Card]
			}
		}
		return false
	} else {
		return w.hand < w2.hand
	}
}

func (w Wager) String() string {
	return fmt.Sprintf("%s (type: %d) %d", w.cards, w.hand, w.bid)
}

func (w *Wager) init()  {
	repetitions := make(map[rune]int, len(w.cards))
	for _, card := range w.cards {
		repetitions[card]++
	}
	groupings := make(map[int]int, len(w.cards))
	for _, v :=  range repetitions {
		groupings[v]++
	}
	keys := []int{5,4,3,2,1}
	for _, key := range keys {
		_, ok := groupings[key]
		if ok {
			switch key {
			case 5:
				w.hand = FiveofAKind
				return
			case 4:
				w.hand = FourOfAKind
				return
			case 3:
				if groupings[2] == 1 {
					w.hand = FullHouse
					return
				} else {
					w.hand = ThreeOfAKind
					return
				}
			case 2:
				if groupings[2] == 2 {
					w.hand = TwoPair
					return
				} else {
					w.hand = OnePair
					return
				}
			}
		}
		w.hand = HighCard
	}
}

func NewWager(cards string, bid int)  *Wager{
	wager := &Wager{
		cards: cards,
		bid:   bid,
	}
	wager.init()
	return wager
}

func main() {
	if len(os.Args) < 1 {
		log.Fatalf("Usage: %s <filename> \n", os.Args[0])
	}
	filename := os.Args[1]
	var content, err = os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	var wagers []*Wager
	for _, line := range lines {
		if line == "" {
			continue
		}
		components := strings.Split(line, " ")
		bid, _ := strconv.Atoi(components[1])
		wagers = append(wagers, NewWager(components[0], bid))
	}
	sort.Slice(wagers, func(i, j int) bool {
		return wagers[i].less(wagers[j])
	})
	totalWinnings := 0
	for k, v := range wagers {
		winnings := v.bid*(k+1)
		totalWinnings += winnings
	}
	fmt.Printf("Total winnings: %d\n", totalWinnings)
}