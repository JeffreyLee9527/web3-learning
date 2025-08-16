package go_base

import "slices"

/*
*
最长公共前缀
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]

}

/*
*
有效的括号
*/
func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}
	mp := map[rune]rune{')': '(', ']': '[', '}': '{'}
	st := []rune{}

	for _, c := range s {
		if mp[c] == 0 {
			st = append(st, c)
		} else {
			if len(st) == 0 || st[len(st)-1] != mp[c] {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}

/*
*加一
 */
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

/*
* 删除有序数组中的重复项
 */
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	i := 1
	for j := 1; j < n; j++ {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

/*
*合并区间
 */
func merge(intervals [][]int) (ret [][]int) {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] })

	for _, p := range intervals {
		n := len(ret)
		if n == 0 {
			ret = append(ret, p)
			continue
		}

		if ret[n-1][1] >= p[0] {
			ret[n-1][1] = max(ret[n-1][1], p[1])
		} else {
			ret = append(ret, p)
		}
	}
	return
}

/*
* 两数之和
 */
func twoSum(nums []int, target int) []int {
	var ret = make([]int, 2)
	if len(nums) < 2 {
		return ret
	}
	mp := make(map[int]int)
	for i, v := range nums {
		if mv, ok := mp[target-v]; ok {
			ret[0] = mv
			ret[1] = i
			return ret
		}
		mp[v] = i
	}
	return ret
}
