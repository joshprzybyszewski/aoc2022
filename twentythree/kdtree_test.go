package twentythree

import (
	"fmt"
	"sort"
	"testing"
)

func TestKDTreeSearch(t *testing.T) {

	// +012345
	// 0X-X-X-
	// 1-X-X-X
	// 2X-X--X
	// 3--X---
	// 4------
	// 5-----X
	input := []coord{{
		x: 0, y: 0,
	}, {
		x: 2, y: 0,
	}, {
		x: 4, y: 0,
	}, {
		x: 1, y: 1,
	}, {
		x: 3, y: 1,
	}, {
		x: 5, y: 1,
	}, {
		x: 0, y: 2,
	}, {
		x: 2, y: 2,
	}, {
		x: 5, y: 2,
	}, {
		x: 2, y: 3,
	}, {
		x: 5, y: 5,
	}}

	kdt := newKDTree(input)

	testCases := []struct {
		cr  coordRange
		exp []coord
	}{{
		cr: coordRange{
			x0: 0,
			x1: 2,
			y0: 0,
			y1: 2,
		},
		exp: []coord{{
			x: 0, y: 0,
		}, {
			x: 2, y: 0,
		}, {
			x: 1, y: 1,
		}, {
			x: 0, y: 2,
		}, {
			x: 2, y: 2,
		}},
	}, {
		cr: coordRange{
			x0: 1,
			x1: 3,
			y0: 1,
			y1: 3,
		},
		exp: []coord{{
			x: 1, y: 1,
		}, {
			x: 3, y: 1,
		}, {
			x: 2, y: 2,
		}, {
			x: 2, y: 3,
		}},
	}, {
		cr: coordRange{
			x0: 3,
			x1: 5,
			y0: 3,
			y1: 5,
		},
		exp: []coord{{
			x: 5, y: 5,
		}},
	}}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%+v", tc.cr), func(t *testing.T) {
			act := kdt.search(tc.cr)
			if len(act) != len(tc.exp) {
				t.Logf("Expected %d elements, found %d\n", len(tc.exp), len(act))
				t.Fail()
				return
			}

			sort.Slice(act, func(i, j int) bool {
				if act[i].x == act[j].x {
					return act[i].y < act[j].y
				}
				return act[i].x < act[j].x
			})
			sort.Slice(tc.exp, func(i, j int) bool {
				if tc.exp[i].x == tc.exp[j].x {
					return tc.exp[i].y < tc.exp[j].y
				}
				return tc.exp[i].x < tc.exp[j].x
			})

			for i := range tc.exp {
				if tc.exp[i] != act[i] {
					t.Logf("Expected act[%d] to be %v, but was %v\n", i, tc.exp[i], act[i])
					t.Fail()
				}
			}
		})
	}
}
