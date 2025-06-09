
# https://leetcode.com/problems/longest-common-subsequence/

class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        t1Len, t2Len = len(text1), len(text2)
        if t1Len < t2Len:
            t1Len, t2Len = t2Len, t1Len
            text1, text2 = text2, text1

        dp = [0 for i in range(t2Len+1)]

        for i in range(t1Len - 1, -1, -1):
            prev = 0
            for j in range(t2Len - 1, -1, -1):
                tmp = dp[j]
                if text1[i] == text2[j]:
                    dp[j] = prev + 1
                else:
                    dp[j] = max(dp[j], dp[j + 1])
                prev = tmp

        return dp[0]