package twenty

import "fmt"

func Two(
	input string,
) (int, error) {
	m := newMachine(input)
	// fmt.Printf("machine: %+v\n", m)

	p := newProcessor(&m)
	p.addRx()

	for n := 0; ; n++ {
		p.pushButton()
		if p.numLowPulsesReceived > 0 {
			fmt.Printf("n = %d, rx %d\n", n, p.numLowPulsesReceived)
		}
		if n%10000 == 0 {
			fmt.Printf("n = %10d\n", n)
			fmt.Printf("  low  = %10d\n", p.numLowPulsesSent)
			fmt.Printf("  high = %10d\n", p.numHighPulsesSent)
		}
		if p.numLowPulsesReceived == 1 {
			// it's currently higher than 48810000
			// we need a way to work backwards to detect what the answer is.
			return n, nil
		}
	}
}
