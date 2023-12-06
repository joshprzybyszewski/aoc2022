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
	possiblites := [4]int{}

	var hold, myDistance uint

	for i, r := range races {
		for hold = 1; hold < r.duration; hold++ {
			myDistance = hold * (r.duration - hold)
			if myDistance > r.distance {
				possiblites[i]++
			} else if possiblites[i] > 0 {
				break
			}
		}
	}

	mult := int64(1)
	for _, p := range possiblites {
		mult *= int64(p)
	}

	return mult, nil
}
