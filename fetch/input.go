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
