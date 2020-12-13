/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */
package leetgo

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	resMap := make(map[[26]int][]string)
	for _, s := range strs {
		key := [26]int{}
		for _, item := range s {
			key[item-'a'] += 1
		}
		resMap[key] = append(resMap[key], s)
	}
	res := make([][]string, 0, len(resMap))
	for _, itemList := range resMap {
		res = append(res, itemList)
	}
	return res
}

// @lc code=end
