package main

import "fmt"

//https://geekr.dev/posts/go-sorting-algorithm-quick

//快速排序
//思路：从 p 到 r 的数组，q 为分区点，一般为最后一个元素
//通过两个变量遍历，和分区点做比较，j比分区小的,交换ij，i+1;否则不换，j一直往后
 func quicksort(nums []int,p int,r int){
 	//递归终止条件
 	if p>r{
 		return
	}

	//获取分区位置
	q:= partition(nums,p,r)
	//递归分区
	quicksort(nums,p,q-1)
	quicksort(nums,q+1,r)
 }


 func partition(nums []int, p int, r int) int{
 	//初始分区点
 	pivot := nums[r]
 	//初始化下标
 	i := p
 	//后移j下标的遍历过程
 	for j:=p;j<r;j++{
 		//将比pivot小的数丢到[p...i-1]里，剩下的[i...j]都是比pivot大的
 		if nums[j]<pivot{
 			// 互换 i，j下标
 			nums[i],nums[j] = nums[j],nums[i]
 			// 将i下标后移一位
 			i++
		}
	}

	//最后将 pivot 与 i 下标对应数据值互换
	//pivot 位于数据中间，i也就是pivot的下标
	nums[i],nums[r] = pivot,nums[i]
	return i
 }

 func main(){
 	nums := []int{4,5,6,7,8,3,2,1}
 	quicksort(nums,0,len(nums)-1)
 	fmt.Println(nums)
 }