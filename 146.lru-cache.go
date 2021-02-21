/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 *
 * https://leetcode.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (35.44%)
 * Likes:    7669
 * Dislikes: 315
 * Total Accepted:    691.9K
 * Total Submissions: 2M
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * Design a data structure that follows the constraints of a Least Recently
 * Used (LRU) cache.
 *
 * Implement the LRUCache class:
 *
 *
 * LRUCache(int capacity) Initialize the LRU cache with positive size
 * capacity.
 * int get(int key) Return the value of the key if the key exists, otherwise
 * return -1.
 * void put(int key, int value) Update the value of the key if the key exists.
 * Otherwise, add the key-value pair to the cache. If the number of keys
 * exceeds the capacity from this operation, evict the least recently used
 * key.
 *
 *
 * Follow up:
 * Could you do get and put in O(1) time complexity?
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * Output
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 *
 * Explanation
 * LRUCache lRUCache = new LRUCache(2);
 * lRUCache.put(1, 1); // cache is {1=1}
 * lRUCache.put(2, 2); // cache is {1=1, 2=2}
 * lRUCache.get(1);    // return 1
 * lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
 * lRUCache.get(2);    // returns -1 (not found)
 * lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
 * lRUCache.get(1);    // return -1 (not found)
 * lRUCache.get(3);    // return 3
 * lRUCache.get(4);    // return 4
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= capacity <= 3000
 * 0 <= key <= 3000
 * 0 <= value <= 10^4
 * At most 3 * 10^4 calls will be made to get and put.
 *
 *
*/
package leetgo

import "sync"

// @lc code=start
type DLinkedNode struct {
	Key   int
	Value int
	Pre   *DLinkedNode
	Post  *DLinkedNode
}

type LRUCache struct {
	sync.Mutex

	Cache    map[int]*DLinkedNode
	Head     *DLinkedNode
	Tail     *DLinkedNode
	Count    int
	Capacity int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Cache:    make(map[int]*DLinkedNode),
		Head:     &DLinkedNode{},
		Tail:     &DLinkedNode{},
		Capacity: capacity,
	}
	lru.Head.Post = lru.Tail
	lru.Tail.Pre = lru.Head
	return lru
}

func (lru *LRUCache) Get(key int) int {
	lru.Lock()
	defer lru.Unlock()
	if node, ok := lru.Cache[key]; ok {
		lru.moveToFirst(node)
		return node.Value
	} else {
		return -1
	}
}

func (lru *LRUCache) Put(key int, value int) {
	lru.Lock()
	defer lru.Unlock()
	if node, ok := lru.Cache[key]; ok {
		node.Value = value
		lru.moveToFirst(node)
	} else {
		node := &DLinkedNode{
			Key:   key,
			Value: value,
		}

		lru.Cache[key] = node
		lru.addNode(node)
		lru.Count++
		if lru.Count > lru.Capacity {
			last := lru.popLast()
			delete(lru.Cache, last.Key)
			lru.Count--
		}
	}
}

func (lru *LRUCache) addNode(node *DLinkedNode) {
	node.Pre = lru.Head
	node.Post = lru.Head.Post
	lru.Head.Post.Pre = node
	lru.Head.Post = node
}

func (lru *LRUCache) delNode(node *DLinkedNode) {
	pre := node.Pre
	post := node.Post

	pre.Post = post
	post.Pre = pre

	node.Pre = nil
	node.Post = nil
}

func (lru *LRUCache) moveToFirst(node *DLinkedNode) {
	lru.delNode(node)
	lru.addNode(node)
}

func (lru *LRUCache) popLast() *DLinkedNode {
	last := lru.Tail.Pre
	lru.delNode(last)
	return last
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
