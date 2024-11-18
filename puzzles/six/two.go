package six

import (
	"fmt"
)

const (
	lastSeenSize = int(byte('z')) + 1
)

// This has been finessed and optimized a lot.
func Two(
	input string,
) (int, error) {
	return getMarkerOfUniqueWindow(
		input,
		14,
	)
}

func getMarkerOfUniqueWindow(
	input string,
	window int,
) (int, error) {

	// we know that there's only 26 bytes that could be seen
	var seen uint32
	var bit uint32

	var min, j int
	i := window
	max := len(input)
START:
	for max > i {
		seen = 0
		for j = i - 1; min <= j; j-- {
			bit = 1 << (input[j] % 32)
			if seen&bit != 0 { // have we already seen this character in this window?
				min = j + 1
				i = min + window // move the end of the window forward to after this known duplicate
				goto START
			}

			seen |= bit
		}
		return i, nil
	}

	return 0, fmt.Errorf("didn't find a window of %d unique characters\n", window)
}
