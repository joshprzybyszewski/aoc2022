package util

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var (
	secret = ``
)

func init() {
	secret = os.Getenv(`AOC_2022_SECRET`)
}

func getInputFromWebsite(
	day int,
) (string, error) {
	inputURL := fmt.Sprintf(`https://adventofcode.com/%d/day/%d/input`, year, day)

	req, err := http.NewRequest(`GET`, inputURL, nil)
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

func postAnswerToWebsite(
	day int,
	level int,
	answer string,
) (string, error) {

	submitURL := fmt.Sprintf(`https://adventofcode.com/%d/day/%d/answer`, year, day)

	formData := url.Values{}
	formData.Set("answer", answer)
	formData.Set("level", strconv.Itoa(level))

	req, err := http.NewRequest(`POST`, submitURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return ``, err
	}

	req.Header.Add(`cookie`, secret)
	req.Header.Set(`Content-Type`, `application/x-www-form-urlencoded`)

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
