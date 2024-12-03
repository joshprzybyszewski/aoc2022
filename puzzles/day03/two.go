package day03

func Two(
	inputString string,
) (int, error) {
	sum := 0

	var l1, l2 int
	var ok bool

	isEnabled := true
	input := []byte(inputString)
	maxI := len(input) - 8 // need space at the end for mul(x,y)

	getL1 := func(i *int) bool {
		l1 = 0
		for *i < maxI {
			if input[*i] == ',' {
				*i++
				return true
			}
			if input[*i] < '0' || input[*i] > '9' {
				return false
			}
			l1 *= 10
			l1 += int(input[*i] - '0')
			*i++
		}
		return false
	}

	getL2 := func(i *int) bool {
		l2 = 0
		for *i < maxI {
			if input[*i] == ')' {
				*i++
				return true
			}
			if input[*i] < '0' || input[*i] > '9' {
				return false
			}
			l2 *= 10
			l2 += int(input[*i] - '0')
			*i++
		}
		return false
	}

	for i := 0; i < maxI; {
		if input[i] == 'd' {
			i++
			if input[i] != 'o' {
				continue
			}
			i++
			if isEnabled {
				if input[i] != 'n' {
					continue
				}
				i++
				if input[i] != '\'' {
					continue
				}
				i++
				if input[i] != 't' {
					continue
				}
				i++
				if input[i] != '(' {
					continue
				}
				i++
				if input[i] != ')' {
					continue
				}
				isEnabled = false
			} else {
				if input[i] != '(' {
					continue
				}
				i++
				if input[i] != ')' {
					continue
				}
				isEnabled = true
			}

			continue
		}
		if input[i] != 'm' {
			i++
			continue
		}
		if !isEnabled {
			// doesn't matter
			i++
			continue
		}
		if input[i+1] != 'u' {
			i += 1
			continue
		}
		if input[i+2] != 'l' {
			i += 2
			continue
		}
		if input[i+3] != '(' {
			i += 3
			continue
		}
		i += 4
		ok = getL1(&i)
		if !ok {
			continue
		}

		ok = getL2(&i)
		if !ok {
			continue
		}
		sum += (l1 * l2)
	}

	// 153469856 is too high

	return sum, nil
}
