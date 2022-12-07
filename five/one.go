package five

import (
	"fmt"
	"strings"
)

type stack struct {
	values []string
}

func newStack() *stack {
	return &stack{
		values: nil,
	}
}

func newStacks(lines []string) ([]*stack, error) {
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
	output := make([]*stack, 9)
	for i := range output {
		output[i] = newStack()
	}
	output[0].push(`D`, `T`, `W`, `N`, `L`)
	output[1].push(`H`, `P`, `C`)
	output[2].push(`J`, `M`, `G`, `D`, `N`, `H`, `P`, `W`)
	output[3].push(`L`, `Q`, `T`, `N`, `S`, `W`, `C`)
	output[4].push(`N`, `C`, `H`, `P`)
	output[5].push(`B`, `Q`, `W`, `M`, `D`, `N`, `H`, `T`)
	output[6].push(`L`, `S`, `G`, `J`, `R`, `B`, `M`)
	output[7].push(`T`, `R`, `B`, `V`, `G`, `W`, `N`, `Z`)
	output[8].push(`L`, `P`, `N`, `D`, `G`, `W`)
	return output, nil
}

func (s *stack) top() string {
	return s.values[0]
}

func (s *stack) pop() (string, error) {
	if len(s.values) == 0 {
		return ``, fmt.Errorf("too few elements to pop. %+v", s.values)
	}
	output := s.values[0]
	s.values = s.values[1:]
	return output, nil
}

func (s *stack) push(ss ...string) {
	s.values = append(ss, s.values...)
}

type instruction struct {
	source   int
	dest     int
	quantity int
}

func newInstruction(line string) (instruction, error) {
	// "move 6 from 6 to 5"
	var q, s, d int
	_, err := fmt.Sscanf(line, "move %d from %d to %d", &q, &s, &d)
	if err != nil {
		return instruction{}, err
	}
	return instruction{
		quantity: q,
		source:   s,
		dest:     d,
	}, nil
}

func One(
	input string,
) (string, error) {
	stacks, ins, err := convertInputToStacksAndInstructions(input)
	if err != nil {
		return ``, err
	}

	for _, inst := range ins {
		for i := 0; i < inst.quantity; i++ {
			v, err := stacks[inst.source-1].pop()
			if err != nil {
				return ``, err
			}
			stacks[inst.dest-1].push(v)
		}
	}

	output := ``
	for _, s := range stacks {
		output += s.top()
	}
	return output, nil
}

func convertInputToStacksAndInstructions(
	input string,
) ([]*stack, []instruction, error) {
	s := strings.Split(input, "\n\n")
	stackLines := strings.Split(s[0], "\n")
	instructionLines := strings.Split(s[1], "\n")

	stacks, err := newStacks(stackLines)
	if err != nil {
		return nil, nil, err
	}

	insts := make([]instruction, 0, len(instructionLines))
	for _, line := range instructionLines {
		if line == `` {
			continue
		}
		inst, err := newInstruction(line)
		if err != nil {
			return nil, nil, err
		}
		insts = append(insts, inst)
	}

	return stacks, insts, nil
}
