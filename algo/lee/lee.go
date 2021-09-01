package Lee

import (
	"fmt"
	"strconv"
)

/**
27块钱
手里有若干面值5元 7元 2元若干，求最少硬币
f(27) = min{f(27-2)+1,f(27-5)+1,f(27-7)+1}
*/
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func Recursion(x int) int {
	var total = MaxInt
	if x == 0 {
		return 0
	}

	if x == 1 {
		fmt.Println("errr")
		return MaxInt
	}

	//fmt.Println(x)
	if x >= 2 {
		total = Min(Recursion(x-2)+1, total)
	}
	if x >= 5 {
		total = Min(Recursion(x-5)+1, total)
	}
	if x >= 7 {
		total = Min(Recursion(x-7)+1, total)
	}
	return total
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
322. Coin Change
https://leetcode.com/problems/coin-change/
*/
func CoinChange(coins []int, amount int) int {
	var minCoinArr = make([]int, amount+1)
	if amount == 0 {
		return 0
	}
	for i := 0; i < amount+1; i++ {
		x := 0
		minCoinArr[i] = MaxInt
		minCoinArr[0] = 0
		for _, c := range coins {
			if i-c >= 0 {
				ret := 1 + minCoinArr[i-c]
				if ret < 0 {
					continue
				}
				if ret > minCoinArr[i] {
					continue
				}
				x = c
				minCoinArr[i] = ret
			}
		}
		fmt.Println(i, "投入：", x)
	}

	if minCoinArr[amount] == MaxInt {
		return -1
	}
	return minCoinArr[amount]
}

/**
有一个国家发现了5座金矿，每座金矿的黄金储量不同，需要参与挖掘的工人数也不同。参与挖矿工人的总数是10人。每座金矿要么全挖，要么不挖，不能派出一半人挖取一半金矿。要求用程序求解出，要想得到尽可能多的黄金，应该选择挖取哪几座金矿？
1、400金/5人   2、500金/5人   3、200金/3人    4、300金/4人    5、350金/3人
*/
func Gold(source map[int]int, people int) int {
	var maxCoinArr = make([]int, people+1)
	var storage = make(map[int]map[int]int)
	for i := 0; i <= people; i++ {
		maxCoinArr[i] = 0
		storage[i] = map[int]int{}
		var x int
		for g, p := range source {
			if i >= p {
				c := maxCoinArr[i-p] + g
				if c > maxCoinArr[i] {
					maxCoinArr[i] = c
					x = g
				}
			}
		}
		if x > 0 {
			fmt.Println(i, x)
			fmt.Println(source[x])
			storage[i] = storage[i-source[x]]
			storage[i][x] = storage[i][x] + 1
			fmt.Println("zhiq", storage[i-source[x]])
			fmt.Println(i, "个工人", "wa", x, maxCoinArr[i])

		}
	}
	for i, i2 := range storage {
		fmt.Println(i, "=>", i2)
	}
	return maxCoinArr[people]
}

/**
62. Unique Paths
https://leetcode.com/problems/unique-paths/
*/
func UniquePaths(m int, n int) int {
	var ret = make(map[int]map[int]int)
	for i := 0; i < m; i++ { //left to right
		ret[i] = make(map[int]int)
		for j := 0; j < n; j++ { // up to down
			if i == 0 || j == 0 {
				ret[i][j] = 1
			} else {
				ret[i][j] = ret[i-1][j] + ret[i][j-1]
			}
		}
	}
	return ret[m-1][n-1]
}

/**
https://leetcode.com/problems/jump-game/
55. Jump Game
*/
func CanJump(nums []int) bool {
	var ret = make(map[int]bool, len(nums))
	ret[0] = true
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if ret[j] && j+nums[j] >= i {
				ret[i] = true
				break
			}
		}
	}
	return ret[len(nums)-1]
}

/**
No. 1
*/
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == tmp {
				return []int{i, j}
			}
		}
	}
	return nil
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nums []int) (ret *ListNode) {
	var tmp *ListNode
	for _, v := range nums {
		if ret == nil {
			ret = &ListNode{Val: v}
		} else {
			if tmp == nil {
				ret.Next = &ListNode{Val: v}
				tmp = ret.Next
			} else {
				tmp.Next = &ListNode{Val: v}
				tmp = tmp.Next
			}
		}
	}
	return
}

func AWN(a []int, b []int) int {
	l1 := NewListNode(a)
	l2 := NewListNode(b)
	ls := addTwoNumbers(l1, l2)
	fmt.Println(ls)
	i := 1
	sum := 0
	for {
		if ls == nil {
			break
		}
		sum += ls.Val * i
		i = 10 * i
		ls = ls.Next
	}
	return sum
}

func MWN(a []int, b []int) int {
	l1 := NewListNode(a)
	l2 := NewListNode(b)
	ls := mergeTwoLists(l1, l2)
	i := 1
	sum := 0
	for {
		if ls == nil {
			break
		}
		sum += ls.Val * i
		i = 10 * i
		ls = ls.Next
	}
	return sum
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	var head = new(ListNode)
	var root = head
	for l1 != nil || l2 != nil || sum != 0 {
		head.Next = new(ListNode)
		head = head.Next
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		head.Val = sum % 10
		sum = sum / 10
	}
	return root.Next
}

/**

 */
func convertToBase7(num int) string {
	var s string
	flag := false
	if num < 0 {
		flag = true
		num = 0 - num
	}
	for {
		x := num % 7
		s = strconv.Itoa(x) + s
		num = num / 7
		if num == 0 {
			break
		}
	}
	if flag {
		s = "-" + s
	}
	return s
}

/**
给定一个 每个结点的值互不相同 的二叉树，和一个目标值 k，找出树中与目标值 k 最近的叶结点。

这里，与叶结点 最近 表示在二叉树中到达该叶节点需要行进的边数与到达其它叶结点相比最少。而且，当一个结点没有孩子结点时称其为叶结点。

在下面的例子中，输入的树以逐行的平铺形式表示。实际上的有根树 root 将以TreeNode对象的形式给出。

示例 1：

输入：
root = [1, 3, 2], k = 1
二叉树图示：
          1
         / \
        3   2

输出： 2 (或 3)

解释： 2 和 3 都是距离目标 1 最近的叶节点。


示例 2：

输入：
root = [1], k = 1
输出：1

解释： 最近的叶节点是根结点自身。


示例 3：

输入：
root = [1,2,3,4,null,null,null,5,null,6], k = 2
二叉树图示：
             1
            / \
           2   3
          /
         4
        /
       5
      /
     6

输出：3
解释： 值为 3（而不是值为 6）的叶节点是距离结点 2 的最近结点。


注：

root 表示的二叉树最少有 1 个结点且最多有 1000 个结点。
每个结点都有一个唯一的 node.val ，范围为 [1, 1000]。
给定的二叉树中有某个结点使得 node.val == k。

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(arr []*int) *TreeNode {
	var root *TreeNode
	for _, i2 := range arr {
		if i2 != nil {
			if root == nil {
				root = &TreeNode{Val: *i2}
			} else {
				root.Left = &TreeNode{Val: *i2}
			}
		}
	}
	return root
}

func findClosestLeaf(root *TreeNode, k int) int {
	return 0
}

func findMinStep(board string, hand string) int {
	var hands = make(map[int32]int)
	for _, i2 := range hand {
		hands[i2] ++
	}
	return 0
}

func DFS(board string, hand map[int32]int) {
	//var count int
	var c uint8
	i := 0
	for i < len(board) {
		for c == board[i] {
			i++
		}
	}
}

func remove(board string) string {
	var before int32
	max := 0
	for i, i2 := range board {
		if i2 != before {
			before = i2
			if max >= 3 {
				return remove(board[:i-1] + board[i+1:])
			}
			max = 1
		} else {
			max += 1
		}
	}
	return board
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var h *ListNode
	for {
		if head == nil {
			break
		}
		h = &ListNode{Val: head.Val, Next: h}
		head = head.Next
	}
	return h
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = new(ListNode)
	var root = head
	for l1 != nil || l2 != nil {
		for l2 != nil {
			if l1 != nil {
				if l2.Val > l1.Val {
					break
				}
			}
			head.Next = &ListNode{Val: l2.Val}
			head = head.Next
			l2 = l2.Next
		}
		if l1 != nil {
			head.Next = &ListNode{Val: l1.Val}
			head = head.Next
			l1 = l1.Next
		}
	}
	return root.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
}

const full = 50

/**
500 km  50L 邮箱
1. 每50KM保存50L
2. 25km依次递增
*/
func FiveKM(dist int) (int, int) {
	var oil = make(map[int]int)
	var times = make(map[int]int)
	times[0] = 0
	times[1] = 1
	oil[0] = 0
	oil[1] = full / 2
	for i := 1; i < dist/full*2; i++ {
		times[i] = times[i-1]*2 + 1
		oil[i] = oil[i-1]*2 + full/2
	}
	fmt.Println(oil)
	fmt.Println(times)
	return oil[dist/full*2-1], times[dist/full*2-1]
}

/**
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
*/
func lengthOfLongestSubstring(s string) int {
	longest := make(map[int]int)
	longest[0] = 1
	var sub string
	for i := 0; i < len(s); i++ {
		if i == 0 {
			sub = string(s[i])
			longest[i] = 1
		} else {
			sub = existsChar(sub, s[i])
			longest[i] = Max(longest[i-1], len(sub))
		}
	}
	return longest[len(s)-1]
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func existsChar(s string, c uint8) string {
	for j := 0; j < len(s); j++ {
		if s[j] == c {
			return s[j+1:] + string(c)
		}
	}
	return s + string(c)
}

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestPalindrome(s string) string {
	var drome = make(map[int]string)
	var sub string
	for i := 0; i < len(s); i++ {
		if i == 0 {
			drome[i] = string(s[i])
			sub = string(s[i])
		} else {
			var d string
			sub += string(s[i])
			d = checkPalindrome(sub)
			if d != "" && len(drome[i-1]) < len(d) {
				drome[i] = d
			} else {
				drome[i] = drome[i-1]
			}
		}
	}
	return drome[len(s)-1]
}

func isPalindrome(s string) bool {
	mid := len(s) / 2
	if len(s)%2 != 0 {
		mid = len(s)/2 - 1
	}
	for i := 0; i <= mid; i++ {
		if i == len(s)-i {
			break
		}
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func checkPalindrome(s string) (b string) {
	if len(s) < 2 {
		return s
	}
	for len(s) > 1 {
		fmt.Println(s)
		c := s[len(s)-1]
		if s[0] == c {
			if isPalindrome(s[1 : len(s)-1]) {
				return s
			} else {
				s = s[1:]
			}
		} else {
			s = s[1:]
		}
	}
	return s
}

func lastPalindrome(s string) string {
	n := len(s)
	var start, end, len, gStart, i int
	for i < n {
		start = i
		end = i
		for end+1 < n && s[end] == s[end+1] { //same char
			end++
		}
		i = end + 1
		for start-1 >= 0 && end+1 < n && s[start-1] == s[end+1] {
			start--
			end++
		}
		if end-start+1 > len {
			len = end - start + 1
			gStart = start
		}
		fmt.Println("[", start, end, "]", gStart, len, "<------>", s[gStart:gStart+len])
	}
	return s[gStart : gStart+len]
}

func GetFirstChar(s string) uint8 {
	first := make(map[uint8]int)
	var arr []uint8
	for i := 0; i < len(s); i++ {
		c := s[i]
		first[c]++
		if first[c] == 1 {
			arr = append(arr, c)
		}
	}
	for _, c := range arr {
		if first[c] == 1 {
			return c
		}
	}
	return ' '
}

func Hanoi(n int, from, to, space string) (sum int) {
	if n == 0 {
		return sum
	}
	if n > 0 {
		sum += Hanoi(n-1, from, space, to) //  除了最后一个，将其他的移动到中转
		sum++                              //  移动最后一个盘到目标盘
		sum += Hanoi(n-1, space, to, from) //  将中转盘移动到目标盘
	}
	return sum
}

func Sqrt(x int) {
	//FixPoint := func() {}
	//
	//AverageDamp := func() {}
	//avg:=func(y int) { x / y }
	//FixPoint(AverageDamp(avg(0),1))

}

/**
(define (fact-tail n)
  (fact-rec n n))

(define (fact-rec n p)
  (if (= n 1)
      p
      (let ((m (- n 1)))
    (fact-rec m (* p m)))))
*/

func FactTail(n int) int {
	return FactRec(n, n)
}
func FactRec(m, n int) int {
	if m == 1 {
		return n
	}
	p := m - 1
	return FactRec(p, p*n)
}


func Queen() {

}