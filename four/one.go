package four

import (
	"fmt"
	"strconv"
	"strings"
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	var s1, e1, s2, e2 int
	var err error
	total := 0
	var ranges, nums []string

	for _, line := range lines {
		if line == `` {
			continue
		}

		// Sscanf takes 5000 allocs, but only 87000 B
		// strings.Split takes 3000 allocs, but 112000 B
		// TODO replace with an "index access"?
		// _, err = fmt.Sscanf(line, "%d-%d,%d-%d", &s1, &e1, &s2, &e2)
		ranges = strings.Split(line, `,`)
		if len(ranges) != 2 {
			return ``, fmt.Errorf("unexpected line: %q", line)
		}
		nums = strings.Split(ranges[0], `-`)
		if len(nums) != 2 {
			return ``, fmt.Errorf("unexpected line: %q", line)
		}
		s1, err = strconv.Atoi(nums[0])
		if err != nil {
			return ``, err
		}
		e1, err = strconv.Atoi(nums[1])
		if err != nil {
			return ``, err
		}
		nums = strings.Split(ranges[1], `-`)
		if len(nums) != 2 {
			return ``, fmt.Errorf("unexpected line: %q", line)
		}
		s2, err = strconv.Atoi(nums[0])
		if err != nil {
			return ``, err
		}
		e2, err = strconv.Atoi(nums[1])
		if err != nil {
			return ``, err
		}

		if (s1 >= s2 && e1 <= e2) ||
			(s2 >= s1 && e2 <= e1) {
			// the ranges are overlapping
			total++
		}
	}

	return strconv.Itoa(total), nil
}
