package strutil

func TrimSpaces(input string) string {
	i := 0
	for i = 0; i < len(input) && input[i] == ' '; i++ {
	}
	input = input[i:]
	for i = len(input) - 1; i >= 0 && input[i] == ' '; i-- {
	}
	input = input[:i+1]

	return input
}
