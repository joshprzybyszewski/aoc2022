package twentyone

import (
	"os"
	"strings"
)

const (
	mathFileName = `puzzles/twentyone/math.go`
)

var (
	oneAnswer int
)

func One(
	input string,
) (int, error) {
	exists, err := checkMathExists()
	if err != nil {
		return 0, err
	}
	if !exists {
		err = writeMathFile(input)
		if err != nil {
			return 0, err
		}
	}
	return oneAnswer, nil
}

func checkMathExists() (bool, error) {
	data, err := os.ReadFile(mathFileName)
	if err != nil {
		if strings.Contains(err.Error(), `no such file or directory`) {
			return false, nil
		}
		return false, err
	}
	return len(data) > 0, nil
}

func writeMathFile(
	input string,
) error {
	f, err := os.Create(mathFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(`package twentyone

func init() {
	oneAnswer = root
}

var (`)
	if err != nil {
		return err
	}

	_, err = f.WriteString(strings.ReplaceAll(input, `:`, `=`))
	if err != nil {
		return err
	}

	_, err = f.WriteString(`)`)
	if err != nil {
		return err
	}
	return nil
}
