package find

func sort(ans []int) []int {
	ans_len := len(ans)
	for i := 0; i < ans_len-1; i++ {

		for j := i + 1; j < ans_len; j++ {
			if ans[i] > ans[j] {
				ans[i], ans[j] = ans[j], ans[i]
			}
		}
	}
	return ans
}

func Find(x int, nums ...int) []int {
	ans := []int{}
	for _, num := range nums {
		if x > num {
			ans = append(ans, num)
		}
	}
	return sort(ans)
}
