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

	g.addFloor()

	units := 0
	for g.addSand(500, 0) {
		units++
		if g.check(500, 0) == sand {
			break
		}
	}

	return units, nil
}
