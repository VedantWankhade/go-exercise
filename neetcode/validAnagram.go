// https://leetcode.com/problems/valid-anagram/description/

package neetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var freq [26]int
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}
	// return !slices.ContainsFunc(freq[:], func(e int) bool {
	// 	return e != 0
	// })
	for i := range freq {
		if i != 0 {
			return false
		}
	}
	return true
}
