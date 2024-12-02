package fifteen

import (
	"fmt"
	"slices"
	"strings"
)

type lens struct {
	label string
	value int
}

type boxes [256][]lens

func (b *boxes) String() string {
	var sb strings.Builder
	for i := range b {
		if len(b[i]) == 0 {
			continue
		}
		sb.WriteString(fmt.Sprintf("Box %d: %v\n", i+1, b[i]))
	}
	return sb.String()
}

func (b *boxes) remove(
	hash uint8,
	label string,
) {
	i := slices.IndexFunc(b[hash], func(l lens) bool { return l.label == label })
	if i < 0 {
		// not found
		return
	}
	b[hash] = slices.Delete(b[hash], i, i+1)
}

func (b *boxes) add(
	hash uint8,
	label string,
	val int,
) {
	i := slices.IndexFunc(b[hash], func(l lens) bool { return l.label == label })
	if i < 0 {
		b[hash] = append(b[hash], lens{
			label: label,
			value: val,
		})
	} else {
		b[hash][i].value = val
	}
}

func (b *boxes) total() int {
	out := 0

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			out += (i + 1) * (j + 1) * b[i][j].value
		}
	}

	return out
}

func Two(
	input string,
) (int, error) {

	var b boxes

	var cur uint8
	// var labelLen int
	var label string

	for len(input) > 0 {
		switch input[0] {
		case ',':
			cur = 0
			label = ``
			// labelLen = 0
			// label = input[1:]
		case '=':
			// TODO get the value following this.
			val := int(input[1] - '0')
			b.add(cur, label, val)
			// b.add(cur, label[:labelLen], val)
			// labelLen = 0
		case '-':
			b.remove(cur, label)
			// b.remove(cur, label[:labelLen])
			// labelLen = 0
		case '\n':
			// do nothing
		default:
			cur += uint8(input[0])
			cur *= 17
			label += string(input[0])
			// labelLen++
		}
		input = input[1:]
	}

	return b.total(), nil
}
