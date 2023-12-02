package two

import (
	"fmt"
	"strings"
)

func One(
	input string,
) (int, error) {
	max := handful{
		red: 12,
		green: 13,
		blue: 14,
	}

	lines := strings.Split(input, "\n")

	sum := 0
	for i, line := range lines {

	handfuls := strings.Split(line, ";")
	for _, handfulStr := range handfuls {
		handful :=interpretSeen(handfulStr)
		if !isPossible(handful, max) {
			sum += (i+1)
		}
	}

		
	}

	return sum, nil
}

func interpretSeen(handfulString string) handful {

}

type handful struct{
	red int
	blue int
	green int
}

func isPossible(
	seen, max handful,
) bool {

}