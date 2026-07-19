package main

import "fmt"

func main() {
	// 1,取重复数结果
	fmt.Println(findSingle([]int{2, 2, 1}))       // [1]
	fmt.Println(findSingle([]int{2, 2, 1, 2}))    // [1]
	fmt.Println(findSingle([]int{2, 2, 1, 2, 3})) // [1 3]
	// 2,回文数结果
	fmt.Println(isPalindrome(121))   // true
	fmt.Println(isPalindrome(-121))  // false
	fmt.Println(isPalindrome(12321)) // true
	fmt.Println(isPalindrome(123))   // false
	// 3,字符判断结果
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([])"))   // true
	fmt.Println(isValid("([)]"))   // false
	// 4,公共前缀结果
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
	fmt.Println(longestCommonPrefix([]string{"ab", "xb"}))                 // ""
	// 5,加一结果
	fmt.Println(plusOne([]int{1, 2, 3}))    // [1,2,4]
	fmt.Println(plusOne([]int{9}))          // [1,0]
	fmt.Println(plusOne([]int{4, 1, 2, 2})) // [4,1,2,3]
	fmt.Println(plusOne([]int{9, 9, 9}))    // [1，0,0,0]
	fmt.Println(plusOne([]int{1, 9, 9}))    // [1，0,0,0]
	// 6,删除重复项结果
	var nums = []int{1, 1, 2}
	fmt.Println(removeDuplicates(&nums)) // 2
	fmt.Println(nums)                    // [1,2]
	// 7,合并区间
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1,6],[8,10],[15,18]]
	fmt.Println(merge([][]int{{4, 7}, {1, 4}}))                    // [[1,6],[8,10],[15,18]]
	fmt.Println(merge([][]int{{1, 5}, {2, 3}}))                    // [[1,5]]
	fmt.Println(merge([][]int{{0, 0}, {1, 4}}))                    // [[0,0],[1,4]]
	// 8,两数之合
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1,2]
}
