package main

func main() {
	//fmt.Println(insertsort([]int{6, 5, 3, 2}))
	// 直接print 打印的是地址？
	//print(sortxx([]int{1, 5, 3, 2}))
	//提交仍然不通过是为什么？
}

//func insertsort(nums []int) []int {
//	//插入排序
//	for i := 0; i < len(nums); i++ {
//		j := i
//		//如果用 j-1 >0,j=1 就进不去循环，为了防止超出 index 想设置 j-1,也应该是 >=0
//		for j > 0 {
//			if nums[j] < nums[j-1] {
//				nums[j], nums[j-1] = nums[j-1], nums[j]
//			}
//			j--
//		}
//	}
//	return nums
//}

//插入排序？
//https://leetcode.cn/problems/sort-an-array/
func sortArray(nums []int) []int {
	//算法：
	//1、取第二个元素，跟前面的元素一一对比，如果比前面的小，放入，并把前面的值往后挪
	//2、循环对比，直到前面没有元素
	//3、再取第三个元素，重复前面的步骤，直到遍历到最后一个元素，排序完成

	//不知道如何初始化 j 比较好？
	i, j := 1, 1

	for i < len(nums) {
		j = i
		for j > 0 {
			if nums[j] < nums[j-1] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
				j = j - 1
			}
			j = j - 1 //为什么还是超出时间范围？
		}
		i++
	}
	return nums
}

// 这么写是哪里有问题呢？
