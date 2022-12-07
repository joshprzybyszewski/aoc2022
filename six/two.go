package six

import (
	"fmt"
	"strconv"
)

const (
	window       = 14
	halfWindow   = window / 2
	lastSeenSize = int(byte('z')) + 1
)

// This has been finessed and optimized a lot.
func Two(
	input string,
) (string, error) {
	// we know that the whole input is in the range A-Za-z, so this array acts as a map lookup
	// of the last time a given byte (aka character) was seen.
	lastSeen := [lastSeenSize]int{}

	var oi, j, min int
	for i := window; i <= len(input); {
		min = i - window // only do the subtraction once, and enable a simple comparison to know if the window is legit
		for j = i - 1; j >= min; j-- {
			oi = lastSeen[input[j]]
			if oi > j { // have we already seen this character in this window?
				i = j + window + 1 // move the end of the window forward to after this known duplicate
				break
			} else {
				lastSeen[input[j]] = j // write that we saw this character at this index
			}
		}
		if j < min { // the inner loop iterated the whole window
			return strconv.Itoa(i), nil
		}
	}

	return ``, fmt.Errorf("didn't find a window of %d unique characters\n", window)
}
