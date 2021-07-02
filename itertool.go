package gcollections

import (
	"sort"
)

type IterTool struct{}

func (it *IterTool) PermuteInt(nums []int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	if length == 0 {
		return append(res, make([]int, 0))
	}
	for i := 0; i < length; i++ {
		removeIndexI := append(append(make([]int, 0), nums[0:i]...), nums[i+1:]...)
		for _, line := range it.PermuteInt(removeIndexI) {
			res = append(res, append([]int{nums[i]}, line...))
		}
	}
	return res
}

func (it *IterTool) PermuteString(nums []string) [][]string {
	res := make([][]string, 0)
	length := len(nums)
	if length == 0 {
		return append(res, make([]string, 0))
	}
	for i := 0; i < length; i++ {
		removeIndexI := append(append(make([]string, 0), nums[0:i]...), nums[i+1:]...)
		for _, line := range it.PermuteString(removeIndexI) {
			res = append(res, append([]string{nums[i]}, line...))
		}
	}
	return res
}

func (it *IterTool) PermuteIntWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	var perm []int
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

func (it *IterTool) PermuteStringWithDup(nums []string) (ans [][]string) {
	sort.Strings(nums)
	n := len(nums)
	var perm []string
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]string(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

func (it *IterTool) SubsetsInt(nums []int) [][]int {
	var res [][]int

	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)
		for j := i; j < len(nums); j++ {
			list = append(list, nums[j])
			dfs(j+1, list)
			list = list[:len(list)-1]
		}
	}

	dfs(0, []int{})
	return res
}

func (it *IterTool) SubsetsString(nums []string) [][]string {
	var res [][]string

	var dfs func(i int, list []string)
	dfs = func(i int, list []string) {
		tmp := make([]string, len(list))
		copy(tmp, list)
		res = append(res, tmp)
		for j := i; j < len(nums); j++ {
			list = append(list, nums[j])
			dfs(j+1, list)
			list = list[:len(list)-1]
		}
	}

	dfs(0, []string{})
	return res
}

func (it *IterTool) SubsetsIntWithDup(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	tmp := make([]int, 0)
	ans := make([][]int, 0)
	var dfs func(x int)
	dfs = func(u int) {
		if u >= n {
			cp := make([]int, len(tmp))
			copy(cp, tmp)
			ans = append(ans, cp)
			return
		}
		k := 0
		for u+k < n && nums[u] == nums[u+k] {
			k++
		}
		for i := 0; i <= k; i++ {
			dfs(u + k)
			tmp = append(tmp, nums[u])
		}

		tmp = tmp[0 : len(tmp)-k-1]
	}
	dfs(0)
	return ans
}

func (it *IterTool) SubsetsStingWithDup(nums []string) [][]string {
	n := len(nums)
	sort.Strings(nums)
	tmp := make([]string, 0)
	ans := make([][]string, 0)
	var dfs func(x int)
	dfs = func(u int) {
		if u >= n {
			cp := make([]string, len(tmp))
			copy(cp, tmp)
			ans = append(ans, cp)
			return
		}
		k := 0
		for u+k < n && nums[u] == nums[u+k] {
			k++
		}
		for i := 0; i <= k; i++ {
			dfs(u + k)
			tmp = append(tmp, nums[u])
		}

		tmp = tmp[0 : len(tmp)-k-1]
	}
	dfs(0)
	return ans
}
