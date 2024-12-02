package eighteen

func Two(
	input string,
) (int, error) {
	l := newLagoon(input, true)

	return l.numDug(), nil
}

func newLineV2(
	cur coord,
	input string,
) (line, coord, string) {

	input = input[2:]
	for input[0] != ' ' {
		input = input[1:]
	}
	input = input[2:]

	dir := input[6]

	num := getDistanceFromHexCode(input[:7])

	var l line
	switch dir {
	case '0': // east / Right
		l.isHorizontal = true
		l.val = cur.row
		l.low = cur.col
		cur.col += num
		l.high = cur.col
	case '2': // west / Left
		l.isHorizontal = true
		l.val = cur.row
		l.high = cur.col
		cur.col -= num
		l.low = cur.col
	case '3': // north / Up
		l.isHorizontal = false
		l.val = cur.col
		l.high = cur.row
		cur.row -= num
		l.low = cur.row
	case '1': // south / Down
		l.isHorizontal = false
		l.val = cur.col
		l.low = cur.row
		cur.row += num
		l.high = cur.row
	default:
		panic(`unexpected heading ` + string(dir))
	}

	return l, cur, input[8:]
}

func getDistanceFromHexCode(
	hexCode string,
) int {
	num := 0
	for i := 1; i < 6; i++ {
		// bit shifting 4 places is multiplying by 16.
		num <<= 4
		if hexCode[i] <= '9' {
			num += int(hexCode[i] - '0')
		} else {
			num += int(hexCode[i]-'a') + 10

		}
	}
	return num
}
