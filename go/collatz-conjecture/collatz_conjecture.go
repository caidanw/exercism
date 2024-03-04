package collatzconjecture

import "errors"

func CollatzConjecture(num int) (int, error) {
	if num < 1 {
		return 0, errors.New("invalid number")
	}

	steps := 0
	for num > 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num*3 + 1
		}
		steps++
	}
	return steps, nil
}
