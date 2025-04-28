package leetcode

func LengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if pos, exists := charMap[s[i]]; exists {
			start = max(start, pos+1)
		}
		charMap[s[i]] = i
		maxLength = max(maxLength, i-start+1)
	}
	return maxLength
}
