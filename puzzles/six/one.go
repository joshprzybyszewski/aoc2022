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
	halfDur := r.duration / 2

	var minHold, maxHold uint

	// TODO
	tmp := sort.Search(int(halfDur), func(h int) bool {
		return uint(h)*(r.duration-uint(h)) > r.distance
	})
	minHold = uint(tmp)

	tmp = sort.Search(int(halfDur)+1, func(h int) bool {
		return (uint(h)+halfDur)*(r.duration-(uint(h)+halfDur)) < r.distance
	})
	maxHold = halfDur + uint(tmp) - 1

	// for r.distance < maxHold*(r.duration-maxHold) {
	// 	maxHold++
	// }
	// for r.distance > maxHold*(r.duration-maxHold) {
	// 	maxHold--
	// }

	return maxHold - minHold + 1
}
