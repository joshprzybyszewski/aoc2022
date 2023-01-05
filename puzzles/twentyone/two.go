package twentyone

import "fmt"

func Two(
	input string,
) (string, error) {

	// 	input = `root: sjmn + pppw
	// dbpl: 5
	// cczh: sllz + lgvd
	// zczc: 2
	// ptdq: humn - dvpt
	// dvpt: 3
	// lfqf: 4
	// humn: 5
	// ljgn: 2
	// sjmn: drzm * dbpl
	// sllz: 4
	// pppw: cczh / lfqf
	// lgvd: ljgn * ptdq
	// drzm: hmdt - zczc
	// hmdt: 32
	// `

	monkeys, nameToIndex, err := convertToMonkeys(input)
	if err != nil {
		return ``, err
	}
	i, ok := nameToIndex[`humn`]
	if !ok {
		return ``, fmt.Errorf(`could not find humn`)
	}
	humn := monkeys[i]
	if humn == nil {
		return ``, fmt.Errorf(`could not find humn`)
	}

	i, ok = nameToIndex[`root`]
	if !ok {
		return ``, fmt.Errorf(`could not find root`)
	}
	root := monkeys[i]
	if root == nil {
		return ``, fmt.Errorf(`could not find root`)
	}
	if root.left == nil || root.right == nil {
		return ``, fmt.Errorf(`root's left and right are both nil`)
	}

	fmt.Println("root:\n")
	root.Print()
	// 60656487515872 is too high
	// -1002, -987, 759013 isn't right
	// 11065578908532 is wrong.
	// 11065578908535.363281 wrong

	if root.left.dependsOn(humn) {
		if root.right.dependsOn(humn) {
			// I don't know how to solve it if both sides depend on humn
			return ``, fmt.Errorf(`oof this is a harder problem to solve than I'm cut out for.`)
		}

		v := root.right.eval()
		out, ok := root.left.reverseEval(uint64(v), humn)
		if !ok {
			return ``, fmt.Errorf(`left's reverse eval is not ok`)
		}
		humn.value = int(out)
		if l := root.left.eval(); l != v {
			return ``, fmt.Errorf(`found the wrong answer: %d => %d instead of %d`, out, l, v)
		}
		return fmt.Sprintf("%d", out), nil
	}

	v := root.left.eval()
	out, ok := root.right.reverseEval(uint64(v), humn)
	if !ok {
		return ``, fmt.Errorf(`right's reverse eval is not ok`)
	}
	humn.value = int(out)
	if r := root.right.eval(); r != v {
		return ``, fmt.Errorf(`found the wrong answer: %d => %d instead of %d`, out, r, v)
	}
	return fmt.Sprintf("%d", out), nil
}
