package main

// https://leetcode.com/problems/longest-common-subsequence/

func longestCommonSubsequence(text1 string, text2 string) int {
	t1Len, t2Len := len(text1), len(text2)
	dp := make(map[[2]int]int, t1Len*2)

	for i := t1Len - 1; i >= 0; i-- {
		for j := t2Len - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[[2]int{i, j}] = 1 + dp[[2]int{i + 1, j + 1}]
			} else {
				dp[[2]int{i, j}] = max(dp[[2]int{i, j + 1}], dp[[2]int{i + 1, j}])
			}
		}
	}

	return dp[[2]int{0, 0}]
}

// FastLongestCommonSubsequence the fastest solution because of using array instead of hash map
func FastLongestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}

	dp := make([]int, len(text2)+1)
	for i := len(text1) - 1; i >= 0; i-- {
		prev := 0
		for j := len(text2) - 1; j >= 0; j-- {
			tmp := dp[j]
			if text1[i] == text2[j] {
				dp[j] = prev + 1
			} else {
				dp[j] = max(dp[j], dp[j+1])
			}
			prev = tmp
		}
	}

	return dp[0]
}

// SlowLongestCommonSubsequence slower solution
func SlowLongestCommonSubsequence(text1 string, text2 string) int {
	t1Len, t2Len := len(text1), len(text2)
	dp := make(map[[2]int]int, t1Len*2)

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == t1Len || j == t2Len {
			return 0
		}
		point := [2]int{i, j}
		if v, ok := dp[point]; ok {
			return v
		}

		var res int
		if text1[i] == text2[j] {
			res = 1 + dfs(i+1, j+1)
		} else {
			res = max(dfs(i+1, j), dfs(i, j+1))
		}

		dp[point] = res
		return res
	}

	return dfs(0, 0)
}
