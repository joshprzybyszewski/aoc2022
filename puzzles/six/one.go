package six

import "sort"

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
) (int, error) {
	mult := uint(1)
	for _, r := range races {
		mult *= getPossibilities(r)
	}
	return int(mult), nil
}

func getPossibilities(r race) uint {
	tmp := sort.Search(int(r.duration/2), func(h int) bool {
		return uint(h)*(r.duration-uint(h)) > r.distance
	})
	minHold := uint(tmp)

	// maxHold = r.duration - minHold
	return r.duration - minHold - minHold + 1
}
