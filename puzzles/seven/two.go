package seven

import (
	"slices"
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

type cardWilds uint8

func newCardWilds(b byte) cardWilds {
	return cardWilds(convertCardWildsToOrder(b))
}

func (c cardWilds) String() string {
	switch c {
	case 0:
		return `J`
	case 9:
		return `T`
	case 10:
		return `Q`
	case 11:
		return `K`
	case 12:
		return `A`
	}
	if c >= 1 && c <= 8 {
		return string('1' + c)
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
		return kind5
	}
	numTwos := 0
	seenThree := false

	for i = 0; i < 5; i++ {
		if index[i] == wildIndex {
			continue
		}

		switch countByCard[index[i]] + numWilds {
		case 5:
			return kind5
		case 4:
			return kind4
		}

		switch countByCard[index[i]] {
		case 3:
			if numTwos > 0 {
				return fullHouse
			}
			seenThree = true
		case 2:
			if seenThree || numWilds == 2 {
				return fullHouse
			}
			numTwos++
		}
	}

	if numTwos == 4 && numWilds == 1 ||
		seenThree && numWilds > 1 {
		return fullHouse
	}

	if seenThree ||
		numWilds == 2 ||
		numWilds == 1 && numTwos == 2 {
		return kind3
	}
	if numTwos == 4 { // they get seen twice
		return pair2
	}
	if numTwos == 2 || numWilds == 1 {
		return pair1
	}
	return highCard
}

type handWilds struct {
	cards [5]cardWilds

	handType handType
	bid      uint16
}

func (h handWilds) toInt() int {
	// assumes we're running on 64bit architecture:#
	return int(h.handType)<<56 |
		int(h.cards[0])<<48 |
		int(h.cards[1])<<40 |
		int(h.cards[2])<<32 |
		int(h.cards[3])<<24 |
		int(h.cards[4])<<16 |
		int(h.bid) // need 16 bits
}

func newHandWilds(line string) handWilds {
	h := handWilds{}
	h.cards[0] = newCardWilds(line[0])
	h.cards[1] = newCardWilds(line[1])
	h.cards[2] = newCardWilds(line[2])
	h.cards[3] = newCardWilds(line[3])
	h.cards[4] = newCardWilds(line[4])

	h.bid = uint16(line[6] - '0')
	for i := 7; i < len(line); i++ {
		h.bid *= 10
		h.bid += uint16(line[i] - '0')
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

	hi := 0
	handInts := make([]int, 1000)

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		handInts[hi] = newHandWilds(input[:nli]).toInt()
		hi++

		input = input[nli+1:]
	}

	slices.Sort(handInts)

	total := 0
	for i := range handInts {
		total += (i + 1) * getBidFromInt(handInts[i])
	}

	return total, nil
}
