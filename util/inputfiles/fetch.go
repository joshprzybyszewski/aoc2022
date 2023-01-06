package inputfiles

import (
	_ "embed"
	"fmt"
	"os"
)

const (
	year = 2022
)

var (
	//go:embed day1.txt
	day1Input string

	//go:embed day2.txt
	day2Input string

	//go:embed day3.txt
	day3Input string

	//go:embed day4.txt
	day4Input string

	//go:embed day5.txt
	day5Input string

	//go:embed day6.txt
	day6Input string

	//go:embed day7.txt
	day7Input string

	//go:embed day8.txt
	day8Input string

	//go:embed day9.txt
	day9Input string

	//go:embed day10.txt
	day10Input string

	//go:embed day11.txt
	day11Input string

	//go:embed day12.txt
	day12Input string

	//go:embed day13.txt
	day13Input string

	//go:embed day14.txt
	day14Input string

	//go:embed day15.txt
	day15Input string

	//go:embed day16.txt
	day16Input string

	//go:embed day17.txt
	day17Input string

	//go:embed day18.txt
	day18Input string

	//go:embed day19.txt
	day19Input string

	//go:embed day20.txt
	day20Input string

	//go:embed day21.txt
	day21Input string

	//go:embed day22.txt
	day22Input string

	//go:embed day23.txt
	day23Input string

	//go:embed day24.txt
	day24Input string

	//go:embed day25.txt
	day25Input string
)

func Fetch(
	day int,
) (string, error) {
	input := getFromMemory(day)
	if input != `` {
		return input, nil
	}
	return getInputFromLocalFile(day)
}

func getFromMemory(
	day int,
) string {
	switch day {
	case 1:
		return day1Input
	case 2:
		return day2Input
	case 3:
		return day3Input
	case 4:
		return day4Input
	case 5:
		return day5Input
	case 6:
		return day6Input
	case 7:
		return day7Input
	case 8:
		return day8Input
	case 9:
		return day9Input
	case 10:
		return day10Input
	case 11:
		return day11Input
	case 12:
		return day12Input
	case 13:
		return day13Input
	case 14:
		return day14Input
	case 15:
		return day15Input
	case 16:
		return day16Input
	case 17:
		return day17Input
	case 18:
		return day18Input
	case 19:
		return day19Input
	case 20:
		return day20Input
	case 21:
		return day21Input
	case 22:
		return day22Input
	case 23:
		return day23Input
	case 24:
		return day24Input
	case 25:
		return day25Input
	}
	return ``
}

func getInputFromLocalFile(
	day int,
) (string, error) {
	data, err := os.ReadFile(getInputFilname(day))
	if err != nil {
		return ``, err
	}
	return string(data), nil
}

func Store(
	day int,
	input string,
) error {
	f, err := os.Create(getInputFilname(day))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(input)
	if err != nil {
		return err
	}
	return nil
}

func getInputFilname(day int) string {
	return fmt.Sprintf("util/inputfiles/day%d.txt", day)
}
