package twelve

type coord struct {
	row, col int
}

// [row][col]
type grid [41][161]uint8

func newGrid(input string) (grid, coord, coord) {
	var s, e coord
	var g grid
	var i int
	var b byte
	for r := range g {
		for c := range g[r] {
			b = input[i]
			if b == 'E' {
				e = coord{
					row: r,
					col: c,
				}
				g[r][c] = 25
			} else if b == 'S' {
				s = coord{
					row: r,
					col: c,
				}
				g[r][c] = 0
			} else {
				g[r][c] = b - 'a'
			}
			i++
		}
		// skip the newline char
		i++
	}
	return g, s, e
}

func One(
	input string,
) (int, error) {
	g, s, e := newGrid(input)
	steps := paint(&g, e, s)
	return steps[s.row][s.col], nil
}

func paint(
	g *grid,
	zero coord,
	target coord,
) [41][161]int {
	max := len(g) * len(g[0])
	var output [41][161]int
	for i := range output {
		for j := range output[i] {
			output[i][j] = max
		}
	}
	output[zero.row][zero.col] = 0

	pending := make([]coord, 0, 8196)
	pending = append(pending, zero)

	var dest, s coord
	var dv int
	var msv uint8
	for len(pending) > 0 {
		dest = pending[0]
		if dest == target {
			break
		}

		dv = output[dest.row][dest.col] + 1

		msv = g[dest.row][dest.col] - 1
		if dest.row > 0 &&
			g[dest.row-1][dest.col] >= msv &&
			output[dest.row-1][dest.col] > dv {

			s = coord{
				row: dest.row - 1,
				col: dest.col,
			}
			output[s.row][s.col] = dv
			pending = append(pending, s)
		}

		if dest.row < len(g)-1 &&
			g[dest.row+1][dest.col] >= msv &&
			output[dest.row+1][dest.col] > dv {

			s = coord{
				row: dest.row + 1,
				col: dest.col,
			}
			output[s.row][s.col] = dv
			pending = append(pending, s)
		}

		if dest.col > 0 &&
			g[dest.row][dest.col-1] >= msv &&
			output[dest.row][dest.col-1] > dv {

			s = coord{
				row: dest.row,
				col: dest.col - 1,
			}
			output[s.row][s.col] = dv
			pending = append(pending, s)
		}

		if dest.col < len(g[dest.row])-1 &&
			g[dest.row][dest.col+1] >= msv &&
			output[dest.row][dest.col+1] > dv {

			s = coord{
				row: dest.row,
				col: dest.col + 1,
			}
			output[s.row][s.col] = dv
			pending = append(pending, s)
		}

		pending = pending[1:]
	}

	return output
}
