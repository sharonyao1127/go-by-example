package main

import (
	"fmt"
	"sort"
)

func binarysearch(nums []int,num int,low int,high int)int{
	//递归终止条件
	if low>high{
		return -1
	}

	//通过中间元素进行二分查找
	mid := (low + high) / 2
	//递归查找
	if num > nums[mid]{
		return binarysearch(nums,num,mid+1,high)
	}else if num < nums[mid]{
		return binarysearch(nums,num,low,mid-1)
	}else{
		//返回索引
		return mid
	}
}

//查找出现第一次的位置：
//如果此时 mid 位置已经到了序列的最左边，不能再往左了，或者序列中索引小于 mid 的上一个元素值不等于待查找元素值，那么此时 mid 就是第一个等于待查找元素值的位置

//查找最后一次出现的位置：
//mid 位置到了序列的最右边，不能再往后了，或者索引大于 mid 的后一个元素值不等于待查找元素，才返回 mid


//针对有序数据序列，适用于变动不是很频繁的静态序列集
//时间复杂度是 O(logn)
func main(){
	nums := []int{4,6,5,3,1,8,2,7}
	//先排序
	sort.Ints(nums)
	fmt.Println(nums)

	num := 5

	index := binarysearch(nums,num,0,len(nums)-1)
	if index != -1{
		fmt.Printf("Find num %d at index %d\n",num,index)
	} else {
		fmt.Printf("Num %d not exists in nums\n",num)
	}
}