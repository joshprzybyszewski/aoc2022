package fetch

import (
	"fmt"
	"os"
	"strings"
)

func Submit(
	day int,
	part1, part2 func(string) (string, error),
) (_, __ func(string) (string, error)) {
	return submissionWrapper(day, 1, part1),
		submissionWrapper(day, 2, part2)
}

func submissionWrapper(
	day, level int,
	in func(string) (string, error),
) func(string) (string, error) {
	return func(input string) (string, error) {
		answer, err := in(input)
		if err != nil {
			return ``, err
		}
		submitAnswer(day, level, answer)
		return answer, nil
	}
}

func submitAnswer(
	day, level int,
	answer string,
) {
	if hasCorrectAnswer(day, level) {
		return
	}
	if level == 2 && !hasCorrectAnswer(day, 1) {
		return
	}

	var resp string
	fmt.Printf("Submit %q as answer for day %d part %d? (Y/n)\n", answer, day, level)
	fmt.Scanf("%s", &resp)
	if len(resp) == 0 || (resp != `y` && resp != `Y`) {
		return
	}
	resp, err := postAnswerToWebsite(day, level, answer)
	if err != nil {
		fmt.Printf("error while submitting: %v\n", err)
		return
	}

	if isCorrect(resp) {
		fmt.Printf("Successfully submitted correct answer: %q\n", resp)
		recordCorrectAnswer(day, level, answer)
	} else {
		fmt.Printf("Submitted WRONG answer: %q\n", resp)
	}
}

func isCorrect(
	resp string,
) bool {
	return strings.Contains(resp, `That's the right answer!`) || strings.Contains(resp, `Did you already complete it?`)
}

func hasCorrectAnswer(
	day, level int,
) bool {
	filename := fmt.Sprintf("answers/day%d-level%d.txt", day, level)
	data, err := os.ReadFile(filename)
	if err != nil {
		return false
	}
	return data != nil && string(data) != ``
}

func recordCorrectAnswer(
	day, level int,
	answer string,
) {
	filename := fmt.Sprintf("answers/day%d-level%d.txt", day, level)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("error while creating: %v\n", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(answer)
	if err != nil {
		fmt.Printf("error while writing: %v\n", err)
	}
}
