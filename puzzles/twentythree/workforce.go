package twentythree

import (
	"runtime"
	"sync"
)

type workforce struct {
	space      *space
	elves      []coord
	roundIndex uint8

	proposals     map[coord]uint16
	proposalsLock sync.Mutex

	wg   sync.WaitGroup
	work chan uint16
}

func newWorkforce(
	space *space,
	elves []coord,
) workforce {
	return workforce{
		space:     space,
		elves:     elves,
		proposals: make(map[coord]uint16, len(elves)),
		work:      make(chan uint16, len(elves)),
	}
}

func (w *workforce) start() {
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			var c, p coord
			var cl clears

			checkElf := func(ci uint16) {
				c = w.elves[ci]
				cl = allClear

				if w.space[c.x-1][c.y-1] {
					cl &= southEast
				}
				if w.space[c.x+1][c.y+1] {
					cl &= northWest
				}

				if (cl&northEast != 0) && w.space[c.x+1][c.y-1] {
					cl &= southWest
				}
				if (cl&southWest != 0) && w.space[c.x-1][c.y+1] {
					cl &= northEast
				}

				if cl == noneClear {
					// already eliminated all directions. do nothing
					return
				}

				if cl&north == north && w.space[c.x][c.y-1] {
					cl &= notNorth
				}
				if cl&south == south && w.space[c.x][c.y+1] {
					cl &= notSouth
				}
				if cl&east == east && w.space[c.x+1][c.y] {
					cl &= notEast
				}
				if cl&west == west && w.space[c.x-1][c.y] {
					cl &= notWest
				}

				if cl == allClear || cl == noneClear {
					// do nothing
					return
				}

				p = c
				switch w.roundIndex {
				case 0:
					if cl&north == north {
						p.y--
					} else if cl&south == south {
						p.y++
					} else if cl&west == west {
						p.x--
					} else if cl&east == east {
						p.x++
					}
				case 1:
					if cl&south == south {
						p.y++
					} else if cl&west == west {
						p.x--
					} else if cl&east == east {
						p.x++
					} else if cl&north == north {
						p.y--
					}
				case 2:
					if cl&west == west {
						p.x--
					} else if cl&east == east {
						p.x++
					} else if cl&north == north {
						p.y--
					} else if cl&south == south {
						p.y++
					}
				case 3:
					if cl&east == east {
						p.x++
					} else if cl&north == north {
						p.y--
					} else if cl&south == south {
						p.y++
					} else if cl&west == west {
						p.x--
					}
				}

				w.proposalsLock.Lock()
				defer w.proposalsLock.Unlock()
				if _, ok := w.proposals[p]; ok {
					w.proposals[p] = uint16(len(w.elves))
				} else {
					w.proposals[p] = ci
				}
			}

			for i := range w.work {
				checkElf(i)
				w.wg.Done()
			}
		}()
	}
}

func (w *workforce) stop() {
	close(w.work)
}

func (w *workforce) run(
	roundIndex uint8,
) map[coord]uint16 {
	w.roundIndex = roundIndex
	// it requires fewer allocs to delete all entries rather than
	// re-make-ing the map.
	for k := range w.proposals {
		delete(w.proposals, k)
	}

	for i := range w.elves {
		w.wg.Add(1)
		w.work <- uint16(i)
	}

	w.wg.Wait()

	return w.proposals
}
