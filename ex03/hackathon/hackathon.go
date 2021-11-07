package hackathon

import "sort"

var (
	cnt    int
	team   [][]int
	visit  []bool
	circle []bool
	not    []int
)

func dfs(idx int, want []int) {
	next := want[idx]
	visit[idx] = true
	if !visit[next] {
		dfs(next, want)
	} else {
		if !circle[next] {
			team = append(team, []int{})
			for i := next; i != idx; i = want[i] {
				team[cnt] = append(team[cnt], i)
			}
			team[cnt] = append(team[cnt], idx)
			cnt++
		} else {
			not = append(not, idx)
		}
	}
	circle[idx] = true
}

func Match(want []int) [][]int {
	cnt = 0
	not = nil
	team = nil
	size := len(want)
	visit = make([]bool, size)
	circle = make([]bool, size)
	for i := 0; i < size; i++ {
		if !visit[i] {
			dfs(i, want)
		}
	}
	if not != nil {
		team = append(team, []int{})
		team[cnt] = append(team[cnt], not...)
	}
	for i := range team {
		sort.Ints(team[i])
	}
	return team
}
