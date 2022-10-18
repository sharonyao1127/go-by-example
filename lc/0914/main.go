package main

//相交链表
//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
//循环链表i,j,if 指向的 val 相等，return
type ListNode struct {
	Val  int
	Next *ListNode
}

//我这是什么思路？不理解相交是什么？
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	for headA != nil {
		for headB != nil {
			if headA.Next == headB.Next {
				return headA.Next
			}
			headB.Next = headB
		}
		headA.Next = headA
	}
	return nil
}

//链表的中间结点
//快慢指针
func middleNode(headA *ListNode) *ListNode {
	i, j := 0, 0
	for b != nil {
		a := headA.Next
		i++

		b := headA.Next.Next
		j = j + 2
	}
}

//只出现一次的数字

//三数之和

//LRU 缓存

func lengthOfLongestSubstring(s string) int {
	res := 0
	left, right := 0, 0
	for right < len(s) {
		for i := left; i < right; i++ {
			if s[i] == s[right] {
				res = max(res, right-left)
				left = i + 1
				break
			}
		}
		right++
	}
	res = max(res, right-left)
	return res
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

作者：wangzhuooooo
链接：https: //leetcode.cn/problems/longest-substring-without-repeating-characters/solution/shuang-zhi-zhen-goshuang-bai-shi-xian-ji-583l/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。





下面的理解，上面的还不理解
func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	if strLen == 0 {
		return 0
	}

	left, right, ans := 0, 0, 0
	m := map[byte]int{}
	for right < strLen {
		if _, ok := m[s[right]]; !ok {
			m[s[right]] = right
		} else {
			if m[s[right]]+1 >= left {
				left = m[s[right]] + 1
			}
			m[s[right]] = right
		}
		ans = max(right-left+1, ans)
		right++
	}

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

作者：xiao-hao-8o
链接：https: //leetcode.cn/problems/longest-substring-without-repeating-characters/solution/gohua-dong-chuang-kou-xun-zhao-zi-fu-chu-eis5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
