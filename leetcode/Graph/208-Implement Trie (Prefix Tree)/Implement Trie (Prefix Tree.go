package main

// const ALPHABET_SIZE = 26

// type Trie struct {
// 	children  [ALPHABET_SIZE]*Trie
// 	isEndWord bool
// }

// func Constructor() Trie {
// 	return Trie{}
// }

// func (this *Trie) Insert(word string) {
// 	node := this
// 	for _, char := range word {
// 		index := char - 'a'
// 		if node.children[index] == nil {
// 			node.children[index] = &Trie{}
// 		}
// 		node = node.children[index]
// 	}
// 	node.isEndWord = true
// }

// func (this *Trie) Search(word string) bool {
// 	node := this
// 	for _, char := range word {
// 		index := char - 'a'
// 		if node.children[index] == nil {
// 			return false
// 		}
// 		node = node.children[index]
// 	}
// 	return node.isEndWord
// }

// func (this *Trie) StartsWith(prefix string) bool {
// 	node := this
// 	for _, char := range prefix {
// 		index := char - 'a'
// 		if node.children[index] == nil {
// 			return false
// 		}
// 		node = node.children[index]
// 	}
// 	return true
// }

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			node.children[index] = &Trie{}
		}
		node = node.children[index]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, char := range word {
		index := char - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, char := range prefix {
		index := char - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true
}
