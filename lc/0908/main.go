package main

import "fmt"

func main() {
	//prices := []int{7, 1, 5, 3, 6, 4}
	//fmt.Println(maxProfit(prices))
	//
	//怎么定义相交，通过val可以吗？
	headA := &ListNode{2, &ListNode{3, &ListNode{8, &ListNode{10, &ListNode{9, nil}}}}}
	headB := &ListNode{5, &ListNode{4, &ListNode{8, &ListNode{10, &ListNode{9, nil}}}}}
	fmt.Println(getIntersectionNode(headA, headB))
}

//121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	//算法：双指针
	// i 买入
	// j 卖出,默认 i+1
	// 后面一天比之前的买入价小，在后一天买入 若j<=i, i++, j=i+1
	// 后面一天比之前买入价高，若j>i，记录此时的利润，同时和之前的最大利润进行对比，存下最大利润,同时 j++
	// 直到 j < len(prices),退出循环
	// 最后返回最大利润
	ans, maxans := 0, 0
	i := 0
	j := 1
	for j < len(prices) {
		if prices[j] > prices[i] {
			ans = prices[j] - prices[i]
			maxans = max(maxans, ans)
			j++
		} else {
			i++
			j = i + 1
		}
		////不满足和满足都会走下面的？不能这么直接写吗？走完上面的if判断还会走下面的？
		//i++
		//j = i + 1
		//上面的判断完，j已经被加1了，此时下面就会报错，应该用else
		// 循环，逻辑分支不熟悉
		//if prices[j] <= prices[i] {
		//	i++
		//	j = i + 1
		//}
	}
	return maxans
}

//教一下 debug

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
