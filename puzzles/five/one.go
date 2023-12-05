package five

import (
	"strings"
)

type mapping struct {
	dest   int
	source int
	length int
}

func newMapping(line string) mapping {
	m := mapping{}
	tmp := 0
	numSpaces := 0
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			if numSpaces == 0 {
				m.dest = tmp
			} else {
				m.source = tmp
			}
			numSpaces++
			tmp = 0
			continue
		}
		tmp *= 10
		tmp += int(line[i] - '0')
	}
	m.length = tmp
	return m
}

func (m mapping) transform(src int) (int, bool) {
	if src < m.source || src > m.source+m.length {
		return 0, false
	}
	return m.dest + (src - m.source), true
}

type allMappings []mapping

func (am allMappings) add(m mapping) allMappings {
	return append(am, m)
}

func (am allMappings) transform(src int) int {
	var dest int
	var ok bool
	for _, m := range am {
		dest, ok = m.transform(src)
		if ok {
			return dest
		}
	}
	return src
}

type multiMaps []allMappings

func (mm multiMaps) add(am allMappings) multiMaps {
	if am == nil {
		return mm
	}
	return append(mm, am)
}

func (mm multiMaps) transform(src int) int {
	for _, am := range mm {
		src = am.transform(src)
	}
	return src
}

func One(
	input string,
) (int, error) {

	si, tmp := 0, 0
	seeds := [20]int{}
	nli := strings.Index(input, "\n")
	if input[:7] != `seeds: ` {
		panic(`dev error`)
	}
	for i := 7; i < len(input[:nli]); i++ {
		if input[i] == ' ' {
			seeds[si] = tmp
			si++
			tmp = 0
			continue
		}
		tmp *= 10
		tmp += int(input[i] - '0')
	}
	input = input[nli+1:]

	var mm multiMaps
	var cur allMappings

	for nli = strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[1:]
			continue
		}

		if input[nli-1] == ':' {
			// this line is dictating that we have the start of a new mapping.
			// Add the current set to the multi.
			mm = mm.add(cur)
			// clear out the current
			cur = nil
		} else {
			cur = cur.add(newMapping(input[:nli]))
		}

		input = input[nli+1:]
	}

	lowest := mm.transform(seeds[0])

	for _, s := range seeds {
		tmp = mm.transform(s)
		if tmp < lowest {
			lowest = tmp
		}
	}

	// 694734096 is too high
	return lowest, nil
}
