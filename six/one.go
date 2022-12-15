package six

func One(
	input string,
) (int, error) {
	return getMarkerOfUniqueWindow(
		input,
		4,
	)
}
