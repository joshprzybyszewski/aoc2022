package twentytwo

type heading uint8

const (
	right heading = 0
	down  heading = 1
	left  heading = 2
	up    heading = 3
)

type direction struct {
	dist    uint
	heading heading
}

type distanceAndRotation struct {
	dist      uint
	clockwise bool
}
