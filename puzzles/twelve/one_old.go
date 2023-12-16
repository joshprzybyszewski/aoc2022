package twelve

// type part uint8

// const (
// 	safe    part = 0
// 	broken  part = 1
// 	unknown part = 2
// )

// func (p part) toString() byte {
// 	switch p {
// 	case safe:
// 		return '.'
// 	case broken:
// 		return '#'
// 	case unknown:
// 		return '?'
// 	}
// 	return 'X'
// }

// type row struct {
// 	parts    [20]part
// 	numParts int
// }

// func (r row) String() string {
// 	var sb strings.Builder
// 	for i := 0; i < r.numParts; i++ {
// 		sb.WriteByte(r.parts[i].toString())
// 	}
// 	return sb.String()
// }

// func (r row) isSolution(indexes []int) bool {
// 	ii := 0
// 	cur := 0
// 	for i := 0; i < len(r.parts); i++ {
// 		if r.parts[i] == broken {
// 			cur++
// 		} else if cur > 0 {
// 			if ii >= len(indexes) || cur != indexes[ii] {
// 				return false
// 			}
// 			ii++
// 			cur = 0
// 		}
// 	}
// 	if ii < len(indexes) {
// 		if cur != indexes[ii] {
// 			return false
// 		}
// 		cur = 0
// 		ii++
// 	}
// 	return ii == len(indexes) && cur == 0
// }

// func (r row) getPossibilities(indexes []int) int {
// 	// fmt.Printf("-- %s %v\n", r, indexes)
// 	total := solveNext(r, 0, indexes)
// 	// fmt.Printf("   %d\n", total)
// 	return total
// }

// func solveNext(
// 	r row,
// 	i int,
// 	indexes []int,
// ) int {
// 	for i < len(r.parts) && r.parts[i] != unknown {
// 		i++
// 	}

// 	if i >= len(r.parts) {
// 		if r.isSolution(indexes) {
// 			// fmt.Printf("   %s\n", r)
// 			return 1
// 		}
// 		return 0
// 	}

// 	r1 := r
// 	r1.parts[i] = broken
// 	r.parts[i] = safe
// 	i++

// 	return solveNext(r1, i, indexes) + solveNext(r, i, indexes)
// }

// func getNumConfigurations(line string) int {
// 	var r row
// 	var i int
// 	for i = 0; i < len(line); i++ {
// 		if line[i] == ' ' {
// 			r.numParts = i
// 			i++
// 			break
// 		}
// 		switch line[i] {
// 		case '?':
// 			r.parts[i] = unknown
// 		case '#':
// 			r.parts[i] = broken
// 		}
// 	}

// 	curNum := 0
// 	var indexes []int
// 	for ; i < len(line); i++ {
// 		if line[i] == ',' {
// 			indexes = append(indexes, curNum)
// 			curNum = 0
// 			continue
// 		}
// 		curNum *= 10
// 		curNum += int(line[i] - '0')
// 	}
// 	indexes = append(indexes, curNum)

// 	return r.getPossibilitiesV2(indexes)
// }

// func OneOld(
// 	input string,
// ) (int, error) {
// 	total := 0
// 	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
// 		total += getNumConfigurations(input[:nli])
// 		input = input[nli+1:]
// 	}
// 	return total, nil
// }

// func (r row) getPossibilitiesV2(
// 	busted []int,
// ) int {

// 	maxToSkip := r.numParts + 1 - len(busted)
// 	for bi := range busted {
// 		maxToSkip -= busted[bi]
// 	}

// 	total := 0

// 	for toSkip := 0; toSkip <= maxToSkip; toSkip++ {
// 		total += r.getPossibilitiesV2_recursive(0, toSkip, busted)
// 	}

// 	return total
// }

// func (r row) getPossibilitiesV2_recursive(
// 	minUncheckedIndex int,
// 	toSkip int,
// 	remainingBusted []int,
// ) int {
// 	if !r.isPossibleV2(minUncheckedIndex, toSkip, remainingBusted) {
// 		return 0
// 	}

// 	minUncheckedIndex += toSkip + remainingBusted[0]
// 	remainingBusted = remainingBusted[1:]

// 	if len(remainingBusted) == 0 {
// 		for ; minUncheckedIndex < r.numParts; minUncheckedIndex++ {
// 			if r.parts[minUncheckedIndex] == broken {
// 				return 0
// 			}
// 		}
// 		return 1
// 	}

// 	numPossible := 0
// 	maxToSkip := r.numParts - minUncheckedIndex - len(remainingBusted)
// 	for toSkip = 1; toSkip <= maxToSkip; toSkip++ {
// 		numPossible += r.getPossibilitiesV2_recursive(
// 			minUncheckedIndex,
// 			toSkip,
// 			remainingBusted,
// 		)
// 	}
// 	return numPossible
// }

// func (r row) isPossibleV2(
// 	minUncheckedIndex int,
// 	toSkip int,
// 	remainingBusted []int,
// ) bool {

// 	var n int

// 	i := minUncheckedIndex

// 	if i+toSkip+remainingBusted[0] > r.numParts {
// 		return false
// 	}
// 	for n = 0; n < toSkip; n++ {
// 		if r.parts[i] == broken {
// 			return false
// 		}
// 		i++
// 	}

// 	for n = 0; n < remainingBusted[0]; n++ {
// 		if r.parts[i] == safe {
// 			return false
// 		}
// 		i++
// 	}

// 	for n = 1; n < len(remainingBusted); n++ {
// 		i += 1 + remainingBusted[n]
// 	}

// 	return i <= r.numParts
// }
