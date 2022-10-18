package main

import "fmt"

func main(){
	s1 := []byte{'h','e','l','l','o'}
	fmt.Println(reverseString(s1))
}

//反转字符串
//https://leetcode.cn/problems/reverse-string/
//? append out of memory?
//func reverseString(s []byte) []byte {
//	ans := []byte{}
//	for len(s) > 0 {
//		x := pop(s)
//		ans = append(ans, x)
//	}
//	return ans
//}
//
//func pop(s2 []byte) byte {
//	ans := s2[len(s2)-1]
//	return ans
//}
func reverseString(s []byte) []byte{
	i, j := 0,len(s)-1
	for i<=j{
		s[i],s[j] = s[j],s[i]
		i++
		j--
	}
	return s

}

//有序数组的平方？
//https://leetcode.cn/problems/squares-of-a-sorted-array/submissions/
func sortedSquares(nums []int) []int {
	ans := []int{}
	for len(nums) > 0 {
		x := nums[0] * nums[0]
		ans = append(ans, x)
	}

	return ans
}

// 也是 append out of memory

//206. 反转链表
type ListNode struct {
	Val  int
	Next *ListNode
}

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

//这样算反转成功了吗？如何打印出链表的一长串内容？
//题目是什么意思？需要再次遍历出val?

//206. 反转链表
//type ListNode struct {
//	val  int
//	Next *ListNode
//}

func reverse(list *ListNode) *ListNode {
	//链表可以循环吗？最后一个结点指向 nil
	//定义一个空结点
	prehead := &ListNode{}
	//当前头结点
	head := list
	//把初始链表存一份
	//temp := list //为什么如果这么写， temp.Next 变成了 指向0，而不是2？提前存了一份，为什么会变化
	temp := list.Next
	//list
	//怎么让 temp 也在往后走
	for temp != nil {
		head.Next = prehead
		prehead = head
		head = temp
		temp = temp.Next
	}
	return head
}

//判断字符是否唯一（go 实现 set）
//https://leetcode.cn/problems/is-unique-lcci/
//定义 set, 不太理解
type Set map[interface{}]struct{}

//

func (s *Set) Add(k interface{}) {
	(*s)[k] = struct{}{}
}

func (s *Set) Remove(k interface{}) {
	delete((*s), k)
}

func (s *Set) Has(k interface{}) bool {
	_, ok := (*s)[k]
	return ok
}

func isUnique(astr string) bool {
	var s Set = Set{}
	//s := []int{}
	for _, i := range astr {
		if s.Has(i) {
			return false
		}
		s.Add(i)
	}
	return true
}

// 直接把一个数组变成 set？python可以直接 set（astr）
// 比较前后的长度是否一致，一致则返回 true
// 不同的语言，处理起来效率不同，这题是否 python 更快？
