package main

// https://leetcode.com/problems/number-of-1-bits/

func hammingWeight(n int) int {
    c := 0
    for n > 0 {
        if n & 1 == 1 {
            c++
        }
        n >>= 1
    }
    return c
}