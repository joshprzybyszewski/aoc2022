package six

func Two(
	input string,
) (int, error) {
	r := race{
		duration: 42686985,
		distance: 284100511221341,
	}

	possiblites := getPossibilities(r)

	return int(possiblites), nil
}
