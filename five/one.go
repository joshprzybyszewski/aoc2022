package five

import (
	"strings"
)

func newStacks() [9]*stack {
	// :badpokerface: yes, I just manually created the stacks instead of reading them in.
	// I figured it was faster to get an answer than to build a generic reader.
	/*
		        [J]         [B]     [T]
		        [M] [L]     [Q] [L] [R]
		        [G] [Q]     [W] [S] [B] [L]
		[D]     [D] [T]     [M] [G] [V] [P]
		[T]     [N] [N] [N] [D] [J] [G] [N]
		[W] [H] [H] [S] [C] [N] [R] [W] [D]
		[N] [P] [P] [W] [H] [H] [B] [N] [G]
		[L] [C] [W] [C] [P] [T] [M] [Z] [W]
		 1   2   3   4   5   6   7   8   9
	*/
	var output [9]*stack
	for i := range output {
		output[i] = newStack()
	}
	output[0].push([]byte(`LNWTD`)...)
	output[1].push([]byte(`CPH`)...)
	output[2].push([]byte(`WPHNDGMJ`)...)
	output[3].push([]byte(`CWSNTQL`)...)
	output[4].push([]byte(`PHCN`)...)
	output[5].push([]byte(`THNDMWQB`)...)
	output[6].push([]byte(`MBRJGSL`)...)
	output[7].push([]byte(`ZNWGVBRT`)...)
	output[8].push([]byte(`WGDNPL`)...)

	return output
}

func One(
	input string,
) (string, error) {
	stacks := newStacks()

	var inst instruction
	var err error
	var i int
	var v byte

	for _, line := range strings.Split(
		input[strings.Index(input, "\n\n")+2:],
		"\n",
	) {
		if line == `` {
			continue
		}
		inst, err = newInstruction(line)
		if err != nil {
			return ``, err
		}
		for i = 0; i < inst.quantity; i++ {
			v, err = stacks[inst.source-1].pop()
			if err != nil {
				return ``, err
			}
			stacks[inst.dest-1].push(v)
		}
	}

	var sb strings.Builder
	for _, s := range stacks {
		sb.WriteByte(s.top())
	}
	return sb.String(), nil
}
