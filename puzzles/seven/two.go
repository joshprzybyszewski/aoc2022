package seven

import (
	"sort"
	"strconv"
	"strings"
)

const (
	wildIndex = 0 // convertCardWildsToOrder('J')
)

func convertCardWildsToOrder(b byte) int {
	if b >= '2' && b <= '9' {
		return int(b - '1')
	}
	switch b {
	case 'T':
		return 9
	case 'J':
		return 0
	case 'Q':
		return 10
	case 'K':
		return 11
	case 'A':
		return 12
	}
	panic(`unsupported char: ` + string(b))
}

type cardWilds uint16

func newCardWilds(b byte) cardWilds {
	return 1 << convertCardWildsToOrder(b)
}

func (c cardWilds) String() string {
	switch c {
	case 1 << 9:
		return `T`
	case 1 << 0:
		return `J`
	case 1 << 10:
		return `Q`
	case 1 << 11:
		return `K`
	case 1 << 12:
		return `A`
	}
	b := '2'
	c >>= 1
	for c != 0 {
		if c&1 == 1 {
			return string(b)
		}
		c >>= 1
		b++
	}
	panic(`ahh`)
}

func newHandWildsType(cards string) handType {
	countByCard := [16]uint8{}
	index := [5]int{}
	var i int
	for i = 0; i < 5; i++ {
		index[i] = convertCardWildsToOrder(cards[i])
		countByCard[index[i]]++
	}

	numWilds := countByCard[wildIndex]
	if numWilds >= 4 {
		return 1 << 6 // 5 of a kind
	}
	numTwos := 0
	seenThree := false

	for i = 0; i < 5; i++ {
		if index[i] == wildIndex {
			continue
		}

		switch countByCard[index[i]] + numWilds {
		case 5:
			return 1 << 6
		case 4:
			return 1 << 5
		}

		switch countByCard[index[i]] {
		case 3:
			if numTwos > 0 {
				return 1 << 4 // full house
			}
			seenThree = true
		case 2:
			if seenThree || numWilds == 2 {
				return 1 << 4 // full house
			}
			numTwos++
		}
	}

	if numTwos == 4 && numWilds == 1 ||
		seenThree && numWilds > 1 {
		return 1 << 4 // full house
	}

	if seenThree ||
		numWilds == 2 ||
		numWilds == 1 && numTwos == 2 {
		return 1 << 3 // 3 of a kind
	}
	if numTwos == 4 { // they get seen twice
		return 1 << 2 // two pair
	}
	if numTwos == 2 || numWilds == 1 {
		return 1 << 1 // one pair
	}
	return 1
}

type handWilds struct {
	cards [5]cardWilds

	handType handType
	bid      uint
}

func newHandWilds(line string) handWilds {
	h := handWilds{}
	h.cards[0] = newCardWilds(line[0])
	h.cards[1] = newCardWilds(line[1])
	h.cards[2] = newCardWilds(line[2])
	h.cards[3] = newCardWilds(line[3])
	h.cards[4] = newCardWilds(line[4])

	h.bid = uint(line[6] - '0')
	for i := 7; i < len(line); i++ {
		h.bid *= 10
		h.bid += uint(line[i] - '0')
	}

	h.handType = newHandWildsType(line[:5])

	return h
}

func (h handWilds) String() string {
	return h.cards[0].String() +
		h.cards[1].String() +
		h.cards[2].String() +
		h.cards[3].String() +
		h.cards[4].String() +
		` ` + strconv.Itoa(int(h.bid)) +
		` (` + h.handType.String() + `)`
}

func Two(
	input string,
) (int, error) {

	hands := make([]handWilds, 0, 1000)

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		hands = append(hands, newHandWilds(input[:nli]))

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

	return total, nil
}
