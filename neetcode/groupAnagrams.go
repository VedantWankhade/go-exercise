// https://leetcode.com/problems/group-anagrams/description/
package neetcode

func groupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}
	hash := make(map[[26]int][]string)

	for _, s := range strs {
		v := [26]int{}
		for i := 0; i < len(s); i++ {
			v[s[i]-'a']++
		}
		hash[v] = append(hash[v], s)
	}
	res := [][]string{}
	for _, v := range hash {
		res = append(res, v)
	}
	return res
}
