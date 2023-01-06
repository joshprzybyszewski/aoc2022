package twentythree

const (
	// min: {x:-14 y:-13}
	// max: {x:127 y:126}
	maxDir = 144
	// offset by 15, because that's as far negative as it goes in p2
	offset = uint8(15)
)

// [x][y]
type space [maxDir][maxDir]bool

func convertInputToElfLocations(input string) (space, []coord) {
	var output space
	coords := make([]coord, 0, 74*74/2)

	x, y := offset, offset
	for _, ch := range input {
		switch ch {
		case '#':
			output[x][y] = true
			coords = append(coords, coord{
				x: x,
				y: y,
			})
			x++
		case '.':
			x++
		case '\n':
			y++
			x = offset
		}
	}

	return output, coords
}
