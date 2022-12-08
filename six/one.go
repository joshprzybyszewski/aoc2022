package six

func One(
	input string,
) (string, error) {
	return getMarkerOfUniqueWindow(
		input,
		4,
	)
}
