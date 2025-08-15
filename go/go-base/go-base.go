package go_base

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
