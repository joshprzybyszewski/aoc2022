package util

import (
	"fmt"
	"os"
)

const (
	year = 2022
)

func Input(
	day int,
) (string, error) {
	input, err := getInputFromLocalFile(day)
	if err != nil {
		fmt.Printf("Input file not found at %q\n", getInputFilname(day))
		fmt.Printf("Attempting to fetch input file from website...\n")
		err = writeInputToLocalFile(day)
		if err != nil {
			return ``, err
		}
		return getInputFromLocalFile(day)
	}

	return input, nil
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

func writeInputToLocalFile(
	day int,
) error {
	input, err := getInputFromWebsite(day)
	if err != nil {
		return err
	}
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
	return fmt.Sprintf("input/day%d.txt", day)
}
