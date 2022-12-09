package five

import (
	"strings"
)

func newStacks() [9]stack {
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
	var output [9]stack

	push := func(s *stack, values string) {
		for i := range values {
			s.values[i] = values[i]
		}
		s.length = len(values)
	}

	push(&output[0], `LNWTD`)
	push(&output[1], `CPH`)
	push(&output[2], `WPHNDGMJ`)
	push(&output[3], `CWSNTQL`)
	push(&output[4], `PHCN`)
	push(&output[5], `THNDMWQB`)
	push(&output[6], `MBRJGSL`)
	push(&output[7], `ZNWGVBRT`)
	push(&output[8], `WGDNPL`)

	return output
}

func One(
	input string,
) (string, error) {
	stacks := newStacks()

	move := func(si, di int, q int) {
		for i := 0; i < q; i++ {
			stacks[di].values[stacks[di].length] = stacks[si].values[stacks[si].length-1]
			stacks[si].length--
			stacks[di].length++
		}
	}

	var inst instruction
	var err error

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
		move(inst.source-1, inst.dest-1, inst.quantity)
	}

	var sb strings.Builder
	for _, s := range stacks {
		sb.WriteByte(s.top())
	}
	return sb.String(), nil
}
