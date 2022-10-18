package main

import "fmt"

func main() {
	//nums := []int{2, 5, 9, 10, 3, 7}
	//quicksort(nums, 0, 5)
	//fmt.Println(nums)
	//res := search(nums, 3, 0, 5)
	//fmt.Println(res)

	fmt.Println(str("abcabcbb"))

	////定义一个链表
	//list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	//fmt.Println(reverse(list))
	////怎么打印出链表后面连着的所有的值

}

////补充题4. 手撕快速排序
//func quicksort(nums []int, p int, r int) {
//	//递归需要终止条件
//	if p >= r {
//		return
//	}
//	//获取分区位置
//	pivotindex := partition(nums, p, r)
//	quicksort(nums, p, pivotindex-1)
//	quicksort(nums, pivotindex+1, r)
//}

//func partition(nums []int, p int, r int) int {
//	//入参需要哪几个变化的参数，需要想清楚
//
//	pivot := nums[r]
//	i := p
//	//j 需要小于r,而不是固定的长度
//	for j := p; j < r; j++ {
//		if nums[j] < pivot {
//			nums[i], nums[j] = nums[j], nums[i]
//			i++
//		}
//	}
//
//	// i 在循环里面变化也会影响外面吗？这里是变量的作用范围理解不够
//	// 最后将 pivot 与 i 下标对应数据值互换
//	// 这样一来，pivot 就位于当前数据序列中间，i 也就是 pivot 值对应的下标
//	nums[i], nums[r] = pivot, nums[i]
//	// 返回 i 作为 pivot 分区位置
//	return i
//}

//215. 数组中的第K个最大元素
//快排的思路，放在正确的位置
//如果中间的分区点 q+1 = K,则 nums[q] 等于该元素
func search(nums []int, k int, p int, r int) int {
	//如果分到最后找不到，就返回-1
	if p >= r {
		return -1
	}
	q := partition(nums, p, r)
	if q+1 == k {
		return nums[q]
	}
	if q+1 < k {
		search(nums, k, q, r)
	}
	if q+1 > k {
		search(nums, k, p, q-1)
	}
	//为什么这里还要return
	return -1
}

func partition(nums []int, p int, r int) int {
	//
	i := p
	pivot := nums[r]
	for j := p; j < r; j++ {
		if nums[j] > pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	//互换的时候，怎么判断用哪个变量，nums[i],nums[r] = pivot, nums[i]
	//nums[i],nums[r] = nums[r], nums[i]
	//pivot = nums[r]
	//跟下面的是一个意思吗？

	nums[i], nums[r] = pivot, nums[i]
	return i
}

//3. 无重复字符的最长子串
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//滑动窗口
func str(s string) string {
	i := 0
	//写在外面会影响取值吗？
	//j := i + 1
	//ans := s[i:j]
	m := make(map[int]string)
	maxcount := 0

	for i+1 < len(s) {
		j := i + 1
		for j+1 < len(s) {
			if s[j] != s[i] {
				j++
			}
			//只判断了开头和结尾相等，缺了一部分
			//遍历此时和i之间的，是否重复
			//判断重复用 hash，思考如何进行判断
			if s[j] == s[i] {
				ans := s[i:j]
				//存进map
				m[len(ans)] = ans
				//存一下最大的 count 数
				maxcount = max(maxcount, len(ans))
				i++
				break
			}
		}
	}
	//有多个怎么处理？
	return m[maxcount]
}

//'abcadcdbb'
//l=a
//r=l+1
func strv2(s string) string {
	l := 0
	r := 0
	m := make(map[byte]int)
	for r < len(s) {
		//存字符和出现的索引值
		m[s[r]] = r //a,0
		//如果右侧出现过，移动左侧的位置,同时更新索引的位置
		if _, ok := m[s[r+1]]; ok {
			l = r
			m[s[r]] = r //a,3
		} else {
			r++ //没出现过，右侧往右挪 2，3
		}

	}

}

//变量的作用范围
//loop
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
