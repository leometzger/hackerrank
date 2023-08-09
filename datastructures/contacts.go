package datastructures

// https://www.hackerrank.com/challenges/contacts/problem?isFullScreen=false

type ContactsTrieNode struct {
	isEndOfWord bool
	char        string
	counter     int32
	children    map[rune]*ContactsTrieNode
}

func (node *ContactsTrieNode) countEndOfWords() int32 {
	return node.counter
}

func (node *ContactsTrieNode) increment() {
	node.counter++
}

type ContactsTrie struct {
	rootNode *ContactsTrieNode
}

func (trie *ContactsTrie) add(item string) {
	currentNode := trie.rootNode
	currentNode.increment()

	for index, c := range item {
		if currentNode.children[c] != nil {
			currentNode = currentNode.children[c]
			currentNode.increment()
		} else {
			node := &ContactsTrieNode{
				isEndOfWord: index == len(item)-1,
				children:    make(map[rune]*ContactsTrieNode),
				char:        string(c),
				counter:     1,
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

	return node.countEndOfWords()
}

func contacts(queries [][]string) []int32 {
	var result []int32
	trie := &ContactsTrie{
		rootNode: &ContactsTrieNode{
			isEndOfWord: false,
			counter:     0,
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
