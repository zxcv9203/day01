package deepcopy

import "errors"

func DeepCopy(origin []int, st, ed int) ([]int, error) {
	if st > ed {
		return nil, errors.New("the start index is larger than the end index")
	}
	if st < 0 || ed > len(origin) {
		return nil, errors.New("the index has gone beyond the scope of the slice")
	}
	ans := make([]int, ed-st)
	for i, value := range origin[st:ed] {
		ans[i] = value
	}
	return ans, nil
}
