package five

import "fmt"

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
