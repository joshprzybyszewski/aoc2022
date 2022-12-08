package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	year = 2022
)

var (
	secret = ``
)

func init() {
	secret = os.Getenv(`AOC_2022_SECRET`)
}

func Input(
	day int,
) (string, error) {
	input, err := getInputFromLocalFile(day)
	if err != nil {
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
	filename := fmt.Sprintf("input/day%d.txt", day)
	data, err := os.ReadFile(filename)
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
	filename := fmt.Sprintf("input/day%d.txt", day)
	f, err := os.Create(filename)
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

func getInputFromWebsite(
	day int,
) (string, error) {
	url := fmt.Sprintf(`https://adventofcode.com/%d/day/%d/input`, year, day)

	req, err := http.NewRequest(`GET`, url, nil)
	if err != nil {
		return ``, err
	}

	req.Header.Add(`cookie`, secret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ``, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ``, err
	}

	return string(body), nil
}
