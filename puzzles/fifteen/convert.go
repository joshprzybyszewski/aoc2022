package fifteen

import (
	"sort"
	"strconv"
	"strings"
)

func getReports(
	input string,
) ([numReports]report, error) {
	// Sensor at x=2, y=18: closest beacon is at x=-2, y=15

	slice := make([]report, 0, numReports)

	var i1, i2,
		sx, sy,
		bx, by int
	var err error
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}
		i1 = strings.Index(input, `x=`) + 2
		i2 = i1 + strings.Index(input[i1:], `,`)
		sx, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return [numReports]report{}, err
		}

		i1 = strings.Index(input, `y=`) + 2
		i2 = i1 + strings.Index(input[i1:], `:`)
		sy, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return [numReports]report{}, err
		}

		i1 = i2 + strings.Index(input[i2:], `x=`) + 2
		i2 = i1 + strings.Index(input[i1:], `,`)
		bx, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return [numReports]report{}, err
		}

		i1 = i2 + strings.Index(input[i2:], `y=`) + 2
		by, err = strconv.Atoi(input[i1:nli])
		if err != nil {
			return [numReports]report{}, err
		}

		slice = append(slice, newReport(
			sx, sy,
			bx, by,
		))
		input = input[nli+1:]
	}

	sort.Slice(slice, func(i, j int) bool {
		if slice[i].sx == slice[j].sx {
			return slice[i].sy < slice[j].sy
		}
		return slice[i].sx < slice[j].sx
	})

	var output [numReports]report
	for i := range slice {
		output[i] = slice[i]
	}
	return output, nil
}
