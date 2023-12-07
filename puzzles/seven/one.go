package seven

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func convertCardToOrder(b byte) int {
	if b >= '2' && b <= '9' {
		return int(b - '2')
	}
	switch b {
	case 'T':
		return 8
	case 'J':
		return 9
	case 'Q':
		return 10
	case 'K':
		return 11
	case 'A':
		return 12
	}
	panic(`unsupported char: ` + string(b))
}

type card uint8

func newCard(b byte) card {
	return card(convertCardToOrder(b))
}

func (c card) String() string {
	if c <= 7 {
		return string('2' + c)
	}
	switch c {
	case 8:
		return `T`
	case 9:
		return `J`
	case 10:
		return `Q`
	case 11:
		return `K`
	case 12:
		return `A`
	}
	panic(`ahh`)
}

type handType uint8

const (
	kind5     handType = 6
	kind4     handType = 5
	fullHouse handType = 4
	kind3     handType = 3
	pair2     handType = 2
	pair1     handType = 1
	highCard  handType = 0
)

func newHandType(cards string) handType {
	countByCard := [16]uint8{}
	index := [5]int{}
	var i int
	for i = 0; i < 5; i++ {
		index[i] = convertCardToOrder(cards[i])
		countByCard[index[i]]++
	}

	numTwos := 0
	seenThree := false

	for i = 0; i < 5; i++ {
		switch countByCard[index[i]] {
		case 5:
			return kind5
		case 4:
			return kind4
		case 3:
			if numTwos > 0 {
				return fullHouse
			}
			seenThree = true
		case 2:
			if seenThree {
				return fullHouse
			}
			numTwos++
		}
	}

	if seenThree {
		return kind3
	}
	if numTwos == 4 { // they get seen twice
		return pair2
	}
	if numTwos == 2 {
		return pair1
	}
	return highCard
}

func (ht handType) String() string {
	switch ht {
	case kind5:
		return `5 of a kind`
	case kind4:
		return `4 of a kind`
	case fullHouse:
		return `full house`
	case kind3:
		return `3 of a kind`
	case pair2:
		return `two pair`
	case pair1:
		return `one pair`
	case highCard:
		return `high card`
	}
	return fmt.Sprintf("%d", ht)
}

type hand struct {
	cards [5]card

	handType handType
	bid      uint16
}

func newHand(line string) hand {
	h := hand{}
	h.cards[0] = newCard(line[0])
	h.cards[1] = newCard(line[1])
	h.cards[2] = newCard(line[2])
	h.cards[3] = newCard(line[3])
	h.cards[4] = newCard(line[4])

	h.bid = uint16(line[6] - '0')
	for i := 7; i < len(line); i++ {
		h.bid *= 10
		h.bid += uint16(line[i] - '0')
	}

	h.handType = newHandType(line[:5])

	return h
}

func (h hand) String() string {
	return h.cards[0].String() +
		h.cards[1].String() +
		h.cards[2].String() +
		h.cards[3].String() +
		h.cards[4].String() +
		` ` + strconv.Itoa(int(h.bid)) +
		` (` + h.handType.String() + `)`
}

func One(
	input string,
) (int, error) {

	hands := make([]hand, 0, 1000)

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		hands = append(hands, newHand(input[:nli]))

		input = input[nli+1:]
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType != hands[j].handType {
			return hands[i].handType < hands[j].handType
		}
		for ci := range hands[i].cards {
			if hands[i].cards[ci] == hands[j].cards[ci] {
				continue
			}
			return hands[i].cards[ci] < hands[j].cards[ci]
		}
		panic(`ahh`)
	})

	total := 0
	for i := range hands {
		total += (i + 1) * int(hands[i].bid)
	}

	// 252785659 is too low
	// 253603890
	return total, nil
}
