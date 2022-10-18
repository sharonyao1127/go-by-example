package main

import "reflect"

//88.合并两个有序数组
//https://leetcode.cn/problems/merge-sorted-array/
//expected 'nums1' to have 1 <= size <= 1 but got 0 ？？？
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if m == 0 && n != 0 {
		nums1 = nums2
	}

	i, j, p := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] <= nums2[j] {
			nums1[p] = nums2[j]
			j--
			p--
		} else {
			nums1[p] = nums1[i]
			i--
			p--
		}
	}
	return nums1
}

//9.回文数（stack)
//为什么只能通过部分用例？
func isAnagram(s string, t string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	if len(s) != len(t) {
		return false
	}
	for i := range s {
		if _, ok := m1[s[i]]; ok {
			m1[s[i]] = m1[s[i]] + 1
		}
		m1[s[i]] = 1
	}
	for j := range t {
		if _, ok := m2[t[j]]; ok {
			m2[t[j]] = m2[t[j]] + 1
		}
		m2[t[j]] = 1
	}
	return reflect.DeepEqual(m1, m2)
}
