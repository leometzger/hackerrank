package datastructures

// https://www.hackerrank.com/challenges/contacts/problem?isFullScreen=false

type ContactsTrieNode struct {
	isEndOfWord bool
	char        string
	children    map[rune]*ContactsTrieNode
}

func (node *ContactsTrieNode) countEndOfWords(currentCount int32) int32 {
	if node.isEndOfWord {
		currentCount += 1
	}

	for _, subnode := range node.children {
		currentCount += subnode.countEndOfWords(0)
	}

	return currentCount
}

type ContactsTrie struct {
	rootNode *ContactsTrieNode
}

func (trie *ContactsTrie) add(item string) {
	currentNode := trie.rootNode

	for index, c := range item {
		if currentNode.children[c] != nil {
			currentNode = currentNode.children[c]
		} else {
			node := &ContactsTrieNode{
				isEndOfWord: index == len(item)-1,
				children:    make(map[rune]*ContactsTrieNode),
				char:        string(c),
			}
			currentNode.children[c] = node
			currentNode = node
		}
	}
}

func (trie *ContactsTrie) countPrefix(prefix string) int32 {
	var node *ContactsTrieNode = trie.rootNode

	for _, c := range prefix {
		if node.children[c] == nil {
			return 0
		}

		node = node.children[c]
	}

	return node.countEndOfWords(0)
}

func contacts(queries [][]string) []int32 {
	var result []int32
	trie := &ContactsTrie{
		rootNode: &ContactsTrieNode{
			isEndOfWord: false,
			children:    make(map[rune]*ContactsTrieNode),
		},
	}

	for _, query := range queries {
		operation := query[0]
		value := query[1]

		if operation == "add" {
			trie.add(value)
		}

		if operation == "find" {
			result = append(result, trie.countPrefix(value))
		}
	}

	return result
}
