package fourteen

import (
	"strings"
)

func Two(
	input string,
) (int, error) {
	lines := strings.Split(input, "\n")
	g, err := getGrid(lines)
	if err != nil {
		return 0, err
	}

	// fmt.Printf("got grid\n%s\n", g)

	g.addFloor()

	// fmt.Printf("got grid\n%s\n", g)

	units := 0
	for g.addSand(500, 0) {
		units++
		if g.check(500, 0) == sand {
			break
		}
	}

	// fmt.Printf("%s\n", g)

	return units, nil
}
