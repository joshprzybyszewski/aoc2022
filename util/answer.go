package util

import "fmt"

func Answers(day int) (string, string, error) {
	if !hasCorrectAnswer(day, 1) {
		return ``, ``, fmt.Errorf("Does not have answer for day %d level 1", day)
	}
	if !hasCorrectAnswer(day, 2) {
		return ``, ``, fmt.Errorf("Does not have answer for day %d level 1", day)
	}

	a1, err := getCorrectAnswer(day, 1)
	if err != nil {
		return ``, ``, err
	}
	a2, err := getCorrectAnswer(day, 2)
	if err != nil {
		return ``, ``, err
	}
	return a1, a2, nil
}
