package main

const ALPHABET_SIZE = 26

type TrieNode struct {
	children    [ALPHABET_SIZE]*TrieNode
	isEndOfWord bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{isEndOfWord: false},
	}
}

func (t *TrieNode) Insert(word string) {
	node := t
	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			node.children[index] = &TrieNode{isEndOfWord: false}
		}
		node = node.children[index]
	}
	node.isEndOfWord = true
}

func (t *TrieNode) Search(word string) bool {
	node := t
	for i, char := range word {
		if char == '.' {
			// 对 .，递归检查所有分支
			for _, child := range node.children {
				if child != nil && child.Search(word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			index := char - 'a'
			if node.children[index] == nil {
				return false
			}
			node = node.children[index]
		}
	}
	return node.isEndOfWord
}

func (this *WordDictionary) AddWord(word string) {
	this.root.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.root.Search(word)
}
