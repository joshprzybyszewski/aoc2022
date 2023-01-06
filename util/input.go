package util

import (
	_ "embed"
	"fmt"

	"github.com/joshprzybyszewski/aoc2022/util/inputfiles"
)

const (
	year = 2022
)

func Input(
	day int,
) (string, error) {
	input, err := inputfiles.Fetch(day)
	if err != nil {
		fmt.Printf("inputfiles.Fetch errored: %q\n", err.Error())
		fmt.Printf("Attempting to fetch input file from website...\n")
		err = writeInputToLocalFile(day)
		if err != nil {
			return ``, err
		}
		return inputfiles.Fetch(day)
	}

	return input, nil
}

func writeInputToLocalFile(
	day int,
) error {
	input, err := getInputFromWebsite(day)
	if err != nil {
		return err
	}
	return inputfiles.Store(day, input)
}
