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
