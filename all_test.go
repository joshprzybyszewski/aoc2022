package main

import (
	"testing"

	"github.com/joshprzybyszewski/aoc2022/fetch"
	"github.com/joshprzybyszewski/aoc2022/five"
	"github.com/joshprzybyszewski/aoc2022/four"
	"github.com/joshprzybyszewski/aoc2022/one"
	"github.com/joshprzybyszewski/aoc2022/seven"
	"github.com/joshprzybyszewski/aoc2022/six"
	"github.com/joshprzybyszewski/aoc2022/three"
	"github.com/joshprzybyszewski/aoc2022/two"
)

func BenchmarkAll(b *testing.B) {
	/*
		goos: linux
		goarch: amd64
		pkg: github.com/joshprzybyszewski/aoc2022
		cpu: Intel(R) Core(TM) i5-3570 CPU @ 3.40GHz
		BenchmarkAll/Day_1/Part_One-4         	   17386	     70449 ns/op	   40984 B/op	       3 allocs/op
		BenchmarkAll/Day_1/Part_Two-4         	   13556	     88780 ns/op	   45096 B/op	      13 allocs/op
		BenchmarkAll/Day_2/Part_One-4         	    2935	    376330 ns/op	  120965 B/op	    2502 allocs/op
		BenchmarkAll/Day_2/Part_Two-4         	    2919	    386575 ns/op	  120965 B/op	    2502 allocs/op
		BenchmarkAll/Day_3/Part_One-4         	   21768	     53668 ns/op	    4868 B/op	       2 allocs/op
		BenchmarkAll/Day_3/Part_Two-4         	   18825	     59562 ns/op	   13780 B/op	      55 allocs/op
		BenchmarkAll/Day_4/Part_One-4         	    2728	    434128 ns/op	  145197 B/op	    3004 allocs/op
		BenchmarkAll/Day_4/Part_Two-4         	    3025	    454100 ns/op	  145195 B/op	    3004 allocs/op
		BenchmarkAll/Day_5/Part_One-4         	     681	   1726078 ns/op	  485602 B/op	    7139 allocs/op
		BenchmarkAll/Day_5/Part_Two-4         	    1003	   1134753 ns/op	  189452 B/op	    3469 allocs/op
		BenchmarkAll/Day_6/Part_One-4         	   81715	     16304 ns/op	   16388 B/op	       2 allocs/op
		BenchmarkAll/Day_6/Part_Two-4         	  341350	      2967 ns/op	       4 B/op	       1 allocs/op
		BenchmarkAll/Day_7/Part_One-4         	    4124	    276696 ns/op	   91818 B/op	    1877 allocs/op
		BenchmarkAll/Day_7/Part_Two-4         	    4582	    263500 ns/op	   92339 B/op	    1871 allocs/op
		PASS
		ok  	github.com/joshprzybyszewski/aoc2022	22.697s
	*/

	b.Run(`Day 1`, func(b *testing.B) {
		input, err := fetch.Input(1)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = one.One(input)
				if answer != `68787` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = one.Two(input)
				if answer != `198041` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})

	b.Run(`Day 2`, func(b *testing.B) {
		input, err := fetch.Input(2)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = two.One(input)
				if answer != `11873` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = two.Two(input)
				if answer != `12014` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})

	b.Run(`Day 3`, func(b *testing.B) {
		input, err := fetch.Input(3)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = three.One(input)
				if answer != `7863` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = three.Two(input)
				if answer != `2488` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})

	b.Run(`Day 4`, func(b *testing.B) {
		input, err := fetch.Input(4)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = four.One(input)
				if answer != `500` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = four.Two(input)
				if answer != `815` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})

	b.Run(`Day 5`, func(b *testing.B) {
		input, err := fetch.Input(5)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = five.One(input)
				if answer != `TWSGQHNHL` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = five.Two(input)
				if answer != `JNRSCDWPP` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})

	b.Run(`Day 6`, func(b *testing.B) {
		input, err := fetch.Input(6)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = six.One(input)
				if answer != `1300` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = six.Two(input)
				if answer != `3986` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})

	b.Run(`Day 7`, func(b *testing.B) {
		input, err := fetch.Input(7)
		if err != nil {
			b.Fail()
		}

		b.Run(`Part One`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = seven.One(input)
				if answer != `1915606` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})

		b.Run(`Part Two`, func(b *testing.B) {
			var answer string
			var err error
			for n := 0; n < b.N; n++ {
				answer, err = seven.Two(input)
				if answer != `5025657` {
					b.Fail()
				}
				if err != nil {
					b.Fail()
				}
			}
		})
	})
}
