package four

import (
	"fmt"
	"strconv"
	"strings"
)

type assignment struct {
	start int
	end   int
}

func newAssignment(r string) (assignment, error) {
	vals := strings.Split(r, `-`)
	if len(vals) != 2 {
		return assignment{}, fmt.Errorf("should provide two values: %q", r)
	}

	start, err := strconv.Atoi(vals[0])
	if err != nil {
		return assignment{}, err
	}

	end, err := strconv.Atoi(vals[1])
	if err != nil {
		return assignment{}, err
	}

	return assignment{
		start: start,
		end:   end,
	}, nil
}
