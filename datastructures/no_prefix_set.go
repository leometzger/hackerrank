package datastructures

import "errors"

// Trie node
type NoPrefixSetNode struct {
	isEndOfWord bool
	children    map[byte]*NoPrefixSetNode
}

func (node *NoPrefixSetNode) add(word string) error {
	currentNode := node
	createdNew := false

	for i := 0; i < len(word); i++ {
		if currentNode.children[word[i]] != nil {
			currentNode = currentNode.children[word[i]]
		} else {
			newNode := &NoPrefixSetNode{
				isEndOfWord: false,
				children:    make(map[byte]*NoPrefixSetNode),
			}
			currentNode.children[word[i]] = newNode
			currentNode = newNode
			createdNew = true
		}

		if currentNode.isEndOfWord {
			return errors.New("has prefix")
		}
	}
	currentNode.isEndOfWord = true

	if !createdNew {
		return errors.New("is prefix")
	}

	return nil
}

func noPrefix(words []string) string {
	root := &NoPrefixSetNode{children: make(map[byte]*NoPrefixSetNode), isEndOfWord: false}

	for _, word := range words {
		err := root.add(word)
		if err != nil {
			return word
		}
	}

	return ""
}
