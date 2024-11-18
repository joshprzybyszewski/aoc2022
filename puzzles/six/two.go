package six

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

	var tmp int
	start := -1
	end := window - 1
START:
	tmp = end
	seen = 1 << (input[tmp] % 32)
	tmp--
	for start < tmp {
		bit = 1 << (input[tmp] % 32)
		if seen&bit != 0 {
			start = tmp
			end = tmp + window
			goto START
		}

		seen |= bit
		tmp--
	}
	return end + 1, nil
}
