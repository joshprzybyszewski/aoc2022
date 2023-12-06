package six

func Two(
	input string,
) (int, error) {
	r := race{
		duration: 42686985,
		distance: 284100511221341,
	}

	var hold, myDistance uint
	possiblites := 0
	for hold = 1; hold < r.duration; hold++ {
		myDistance = hold * (r.duration - hold)
		if myDistance > r.distance {
			possiblites++
		} else if possiblites > 0 {
			return int(possiblites), nil
		}
	}

	return 0, nil
}
