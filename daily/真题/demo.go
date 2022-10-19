package main

import "fmt"

//去除重复字符串
//ffdcvf
//利用map的key唯一性
func removedup(str string) string {
	//定义一个map
	targer := make(map[string]int)

	//循环原字符串
	var counts = 1
	for _, i := range str {
		//fmt.Println(string(i))
		//判断i在不在map里，不在的话，添加，在的话，continue循环
		if _, ok := targer[string(i)]; !ok {
			targer[string(i)] = counts
		} else {
			counts += 1
		}
	}

	//最后循环map,输出key,连接起来，map是无序的
	result := ""
	for k, _ := range targer {
		result = result + k
	}
	return result
}

func main() {
	res := removedup("ffdcvf")
	fmt.Println(res)
}





package main

// "954891"，17 => "89"
// 548 加起来也等于 17，符合条件的最短子串
import (
"fmt"
"strconv"
)

func demo(str string, target int) string {
	//i,j 下标，循环出新的切片
	minlen := 10
	minstr := " "

	for i := 0; i < len(str); i++ {
		//注意 j=i
		for j := i; j < len(str); j++ {
			//计算 slice 切片的和
			newstr := str[i:j]
			number := 0
			for _, str2 := range newstr {
				a, _ := strconv.Atoi(string(str2))
				number = number + a
			}
			//fmt.Printf("number is %d",number)
			//满足条件的话，存最短切片的长度，继续循环
			if number == target {
				nowlen := len(newstr)
				//比较长度，如果更短，替换；如果没有，不变
				if nowlen < minlen {
					minlen = nowlen
					minstr = newstr
				}
			}
		}
	}
	//fmt.Printf("minstr is %s ---",minstr)
	return minstr
	//最后输出最短
}

func main() {
	s := demo("954891", 17)
	fmt.Println(strconv.Quote(s))
}

// 输入数字、字母和特殊符号
// 1、只保留字母并且去重
// 2、只保留字母，并且统计字母中重复字母及其个数并输出
package main

import (
"fmt"
"os"
"unicode"
)

func demo() target string {
//循环命令行参数，输出 v
//新建map
m := map[s string] int
var counts = 1
for k, v := range os.Args{
fmt.Printf("args[%v]=[%v]\n", k, v)

//判断输入的类型，是字母，加入map的key, VALUE为字母出现的次数
if unicode.IsLetter(v){
//如果没出现过，加入
if _, ok := range m; !ok{
map[v] = counts
}
//如果出现过，加1
if _, ok := range m; ok{
map[v]++
}
}
}

//遍历map
s := ""
for k, v := range m{
s = s+k
if v >= 2{
fmt.Printf("重复字母为%s,重复个数为%d", k, v)
}

}
//返回拼接的去重的字母串
return s
}

// 第1天，1个硬币，第2，3天，2个，第 4,5,6 天,3个
// 输入n, 输出第 n 天硬币总数

// 递归

//找出数组中的某个元素出现次数超过长度二分之一
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

//hash
func majorityElement(nums []int) int {
	m := make(map[int]int)
	t := len(nums) / 2
	for _, v := range nums {
		if m[v] >= t {
			return v
		}
		m[v]++
	}
	return -1
}

//快排排序
func quicksort(nums []int, p int r int) {
	//递归终止条件
	if p > r {
		return
	}

	//获取分区位置
	q := partition(nums, p, r)
	quicksort(nums, p, q-1)
	quicksort(nums, q+1, r)
}

func partition(nums []int, p int r int) int {
	//初始化分区位置,一般取最后一个元素
	pivot := nums[r]
	i := p
	for j := p; i < j; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j]
			nums[i]
			i++
		}
	}

	nums[i], nums[r] = pivot, nums[i]
	return
}

func demo(nums []int) int {
	//新建一个map,key存数字，value存出现的次数
	m := make(map(int) int)
	var counts = 1

	for i, v := range nums {
		//判断是不是空数组
		if nums == nil {
			fmt.Printf("空数组")
		}

		//如果没出现过，加入map
		if _, ok := m[v]; !ok {
			m[v] := counts
		} else {
			m[v]++
		}
	}
}







最长不重复子串
窗口中没有重复元素，改变 “右窗棱”的值
窗口中出现重复元素，改变 “右窗棱”的值 ，同时改变 “左窗棱”的值，
hash map 用来做重复判断

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



相交链表的交叉节点

解题思路
双指针
a + b = b + a

for 内循环在两个 链表（a+b和b+a）都结束后就自动结束了。

如果链表有相交，那么会在中途相等，返回相交节点；
如果链表不相交，那么最后会 nil == nil，返回 nil；

代码

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

作者：panda8z
链接：https: //daily-cn.com/problems/intersection-of-two-linked-lists/solution/go-shuang-zhi-zhen-wu-nao-xie-by-panda8z-rs6q/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。




现有一个文件file.txt，请统计文件中各字符出现的次数，并根据次数进行从大到小排序，最后输出字符和次数 如：
file.txt 内容
aaaaa
234
111
***
1
输出 a 5 1 4 * 3 2 1 3 1 4 1

cat words.txt |tr -s ' ' '\n' |sort|uniq -c|sort -r|awk '{print $2,$1}'

1、首先cat命令查看words.txt
2、tr -s ' ' '\n'将空格都替换为换行 实现分词
3、sort排序 将分好的词按照顺序排序
4、uniq -c 统计重复次数（此步骤与上一步息息相关，-c原理是字符串相同则加一，如果不进行先排序的话将无法统计数目）
5、sort -r 将数目倒序排列
6、awk '{print $2,$1}' 将词频和词语调换位置打印出来




第一个只出现一次的字符
在第一次遍历时，我们使用哈希映射统计出字符串中每个字符出现的次数。在第二次遍历时，我们只要遍历到了一个只出现一次的字符，那么就返回该字符，否则在遍历结束后返回空格。

func firstUniqChar(s string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}



我:
二叉树后序遍历
我:
func postorderTraversal(root *TreeNode) (res []int) {
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return
}



B 树跟 B+ 树的不同点主要集中在这几个地方：
B+ 树中的节点不存储数据，只是索引，而 B 树中的节点存储数据；
B 树中的叶子节点并不需要链表来串联。




SELECT "最高分", t.*
FROM t, (SELECT MAX(score) AS score, `subject` FROM t GROUP BY `subject`)b
WHERE t.`score` = b.score
AND t.`subject` = b.subject
UNION
SELECT "最低分", t.*
FROM t, (SELECT MIN(score) AS score, `subject` FROM t GROUP BY `subject`)b
WHERE t.`score` = b.score
AND t.`subject` = b.subject


二分查找也称折半查找，要求线性表必须采用顺序存储结构，而且表中元素按关键字有序排列。
时间复杂度: O(logn)
func binarySearch(arr []int, k int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := (l + r) / 2
		if k == arr[mid] {
			return mid
		}
		if k < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}



python
def deal_string(sourceStr):
while(1):
descStr = sourceStr.replace("AB", "")
if sourceStr == descStr:
return descStr
sourceStr = descStr


if __name__ == '__main__':
# sourceStr = "ABAABCDABC"
sourceStr = "AAABBBBCDAABBABA"
str= deal_string(sourceStr)
print(str)




判断是完全二叉树
1.如果当前是空节点，满足完全二叉树，那剩下遍历的都是空节点，做标记
2.如果已经做了标记，还是遍历到非空节点，那就不是完全二叉树，返回


func isCompleteTree1(root *TreeNode) bool {
	flag := false
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			e := queue[0]
			queue = queue[1:]
			// 如果当前是空节点，满足完全二叉树，那剩下遍历的都是空节点，做标记
			if e == nil {
				flag = true
				continue
				　　　　　　 // 如果已经做了标记，还是遍历到非空节点，那就不是完全二叉树，返回
			} else if flag && e != nil {
				return false
			}
			queue = append(queue, e.Left)
			queue = append(queue, e.Right)
		}
	}
	return true
}

作者：Flare-
链接：https: //daily-cn.com/problems/check-completeness-of-a-binary-tree/solution/go-jian-dan-de-ceng-xu-bian-li-jie-fa-by-tbt4/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



如何在一次遍历中找到单个链表的中间节点？
从示例中可以看出，如果链表长度为奇数，要求的是中间节点；否则，要求的是下中点。也就是说如果链表长度是偶数（有两个中间节点），求第二个中间节点。
采用两根（快慢）指针，fast/slow 刚开始均指向链表头节点，然后每次快节点走两步，慢指针走一步，直至快指针指向 null，此时慢节点刚好来到链表的下中节点。
func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	slow, fast := head.Next, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}

	return slow
}



查询学生的总成绩并进行排名（不重点）
SELECT s_id, SUM(s_score)'总成绩' FROM score
GROUP BY s_id
ORDER BY SUM(s_score) DESC





package main

import "fmt"

//删除有序数组中的重复项
//1，2，2，3，5，5
//相对位置保持一致，需要用同向双指针

func remove(nums []int) int {
	//第一步，初始化双指针
	//0，i-1 是需要保留的，i到j-1 是处理了不需要保留的，j到 len(nums) 是未处理的
	i, j := 0, 0

	//第二步，判断条件
	//如果i为0，第一个左边没值，不需要操作
	//如果 i-1 和 j 比较不等，则需要把 j的值赋值给i, 并且往前挪

	//注意，这里用 for，而不是 if,for 类似while
	//为什么用 if 会报错？
	for j < len(nums) {
		if i == 0 || nums[i-1] != nums[j] {
			nums[i] = nums[j]
			i++
			j++
		} else {
			//否则i不动，j+1
			j++
		}
		// 0 到 i-1
	}
	return i
}

func main() {
	nums := []int{1, 2, 2, 3, 5, 5}
	result := remove(nums)
	fmt.Println(result)
}

package main

import "fmt"

//定义一个 node
type ListNode struct {
	data int
	next *ListNode
}

//linked list 找中间节点
//快慢指针，一个走1，一个走2，当j走到最后，i就是中间节点
func middlelink(head *ListNode) *ListNode {
	i, j := head, head
	for j != nil && j.next != nil {
		i = i.next
		j = j.next.next
	}
	return i
}

//倒数第k个节点
//j先走k步，i j间隔k
//当j超出时，i就是倒数第k个节点
func lastk(head *ListNode, k int) *ListNode {
	i, j := head, head

	//j 先往前 k 步,间隔k
	for c := 0; c < k; c++ {
		j = j.next
	}
	//指针前进的速度，一步
	for j != nil {
		i = i.next
		j = j.next
	}
	return i
}

//反转
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}

作者：ji-chang-_joey
链接：https: //daily-cn.com/problems/reverse-linked-list/solution/jian-lue-xie-fa-li-yong-go-de-yu-fa-tang-by-ji-cha/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


func main() {
	//定义链表
	f := ListNode{data: 5, next: nil}
	e := ListNode{data: 4, next: &f}
	d := ListNode{data: 3, next: &e}
	c := ListNode{data: 2, next: &d}
	b := ListNode{data: 1, next: &c}
	a := ListNode{data: 0, next: &b}
	head := &a

	//first := &ListNode{data: 1}
	//second := &ListNode{data: 2}
	//third := &ListNode{data: 3}
	//first.next = second
	//second.next = third
	//head := first

	result := middlelink(head)
	lastk1 := lastk(head, 2)
	fmt.Printf("cur %v  , data %d \n", result, result.data)
	fmt.Printf("cur %v  , data %d \n", lastk1, lastk1.data)
}
