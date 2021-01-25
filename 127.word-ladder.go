/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 *
 * https://leetcode.com/problems/word-ladder/description/
 *
 * algorithms
 * Hard (31.43%)
 * Likes:    4502
 * Dislikes: 1380
 * Total Accepted:    530.3K
 * Total Submissions: 1.7M
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * Given two words beginWord and endWord, and a dictionary wordList, return the
 * length of the shortest transformation sequence from beginWord to endWord,
 * such that:
 *
 *
 * Only one letter can be changed at a time.
 * Each transformed word must exist in the word list.
 *
 *
 * Return 0 if there is no such transformation sequence.
 *
 *
 * Example 1:
 *
 *
 * Input: beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log","cog"]
 * Output: 5
 * Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" ->
 * "dog" -> "cog", return its length 5.
 *
 *
 * Example 2:
 *
 *
 * Input: beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log"]
 * Output: 0
 * Explanation: The endWord "cog" is not in wordList, therefore no possible
 * transformation.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= beginWord.length <= 100
 * endWord.length == beginWord.length
 * 1 <= wordList.length <= 5000
 * wordList[i].length == beginWord.length
 * beginWord, endWord, and wordList[i] consist of lowercase English
 * letters.
 * beginWord != endWord
 * All the strings in wordList are unique.
 *
 *
 */
package leetgo

// BFS
// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]bool, len(wordList))
	for _, item := range wordList {
		wordMap[item] = true
	}
	if !wordMap[endWord] {
		return 0
	}
	wordQueue := []string{beginWord}

	pathCnt := map[string]int{beginWord: 1}

	for len(wordQueue) > 0 {
		t := wordQueue[0]
		wordQueue = wordQueue[1:]
		for i := 0; i < len(t); i++ {
			w := []rune(t)
			for c := 'a'; c <= 'z'; c++ {
				w[i] = c
				s := string(w)
				if wordMap[s] && s == endWord {
					return pathCnt[t] + 1
				}
				if wordMap[s] && pathCnt[s] == 0 {
					pathCnt[s] = pathCnt[t] + 1
					wordQueue = append(wordQueue, s)
				}
			}
		}
	}
	return 0
}

// @lc code=end
