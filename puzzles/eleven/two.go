package eleven

const (
	oneMillion = 999999
)

func Two(
	input string,
) (int, error) {
	answer := solveForExpansion(input, oneMillion)
	return answer, nil
}
