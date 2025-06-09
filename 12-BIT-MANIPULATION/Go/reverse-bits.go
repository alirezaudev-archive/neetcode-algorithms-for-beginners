package main

// https://leetcode.com/problems/reverse-bits/

func reverseBits(num uint32) uint32 {
    var res uint32
    for i := 0; i < 32; i++ {
        bit := (num >> i) & 1
        res |= bit << (31 - i)
    }

    return res
}