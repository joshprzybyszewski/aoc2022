package util

import "fmt"

func Part1Answer(day int) (string, error) {
	if !hasCorrectAnswer(day, 1) {
		return ``, fmt.Errorf("Does not have answer for day %d level 1", day)
	}

	a1, err := getCorrectAnswer(day, 1)
	if err != nil {
		return ``, err
	}
	return a1, nil
}

func Part2Answer(day int) (string, error) {
	if !hasCorrectAnswer(day, 2) {
		return ``, fmt.Errorf("Does not have answer for day %d level 2", day)
	}

	a2, err := getCorrectAnswer(day, 2)
	if err != nil {
		return ``, err
	}
	return a2, nil
}
