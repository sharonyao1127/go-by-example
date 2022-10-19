package main

import "fmt"

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Println(reverselist(list))

}

//206. 反转链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//反转并遍历打印出 val
func reverselist(list *ListNode) *ListNode {
	temp := list
	ret := &ListNode{}
	fmt.Println("ret", ret)
	for list != nil {
		temp = list.Next
		fmt.Println("temp:", temp)
		list.Next = ret
		if temp == nil {
			return list
		}
		list = temp
		fmt.Println("list:", list)
	}
	return nil
}

//遍历
//for list != nil {
//		fmt.Println(*list)
//		list = list.Next //移动指针
//	}

//215. 数组中的第K个最大元素
//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

//示例 1:
//
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
//示例2:
//
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4

//直接排序，找出索引对应的元素

//map,map不行，map是无序的，循环 map,比较大小

//无重复字串的最长子串
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
//示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
