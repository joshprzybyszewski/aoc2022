package twentyone

import "fmt"

const (
	validateTwo = false
)

func Two(
	input string,
) (int, error) {

	monkeys, nameToIndex, err := convertToMonkeys(input)
	if err != nil {
		return 0, err
	}
	i, ok := nameToIndex[`humn`]
	if !ok {
		return 0, fmt.Errorf(`could not find humn`)
	}
	humn := monkeys[i]
	if humn == nil {
		return 0, fmt.Errorf(`could not find humn`)
	}

	i, ok = nameToIndex[`root`]
	if !ok {
		return 0, fmt.Errorf(`could not find root`)
	}
	root := monkeys[i]
	if root == nil {
		return 0, fmt.Errorf(`could not find root`)
	}
	if root.left == nil || root.right == nil {
		return 0, fmt.Errorf(`root's left and right are both nil`)
	}

	dep, known := root.left, root.right

	if root.right.dependsOn(humn) {
		if root.left.dependsOn(humn) {
			// I don't know how to solve it if both sides depend on humn
			return 0, fmt.Errorf(`oof this is a harder problem to solve than I'm cut out for.`)
		}
		known = root.left
		dep = root.right
	}

	v := known.eval()
	found, ok := dep.reverseEval(v, humn)
	if !ok {
		return 0, fmt.Errorf(`reverse eval is not ok`)
	}

	if validateTwo {
		humn.value = found
		if r := root.right.eval(); r != v {
			return 0, fmt.Errorf(
				`found the wrong answer: %d => %d instead of %d`,
				found,
				r, v,
			)
		}
	}

	return int(found), nil
}
