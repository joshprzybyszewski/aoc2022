package six

type race struct {
	duration uint
	distance uint
}

var (
	races = [4]race{{
		duration: 42,
		distance: 284,
	}, {
		duration: 68,
		distance: 1005,
	}, {
		duration: 69,
		distance: 1122,
	}, {
		duration: 85,
		distance: 1341,
	}}
)

func One(
	input string,
) (int64, error) {
	possiblites := [4]uint{}

	for i, r := range races {
		possiblites[i] = getPossibilities(r)
	}

	mult := int64(1)
	for _, p := range possiblites {
		mult *= int64(p)
	}

	return mult, nil
}

func getPossibilities(r race) uint {
	var hold, myDistance uint
	possiblites := uint(0)

	for hold = 1; hold < r.duration; hold++ {
		myDistance = hold * (r.duration - hold)
		if myDistance > r.distance {
			possiblites++
		} else if possiblites > 0 {
			return possiblites
		}
	}
	return 0
}
