package main

import (
	"fmt"
	"strings"
)

func main() {

	root := NewTrie()

	var n int
	fmt.Scanf("%d", &n)

	words := make([]string, 0)
	for i := 0; i < n; i++ {
		var word string
		fmt.Scanf("%s", &word)
		words = append(words, word)
		root.Insert(word)
	}

	for _, v := range words {
		fmt.Printf("%v %s\n", v, root.FindPrefix(v))
	}
}

type Trie struct {
	C      byte
	IsWord bool
	Word   string
	Next   map[byte]*Trie
	Parent *Trie
}

func NewTrie() *Trie {
	t := &Trie{}
	t.Next = make(map[byte]*Trie)
	return t
}

func (t *Trie) Insert(word string) {
	insert(t, word, word)
}

func insert(t *Trie, word, nochangeWord string) {
	if len(word) == 0 {
		t.IsWord = true
		t.Word = nochangeWord
		return
	}

	c := byte(word[0])
	n, ok := t.Next[c]
	if !ok {
		newNode := NewTrie()
		newNode.C = c
		newNode.Parent = t
		t.Next[c] = newNode
		n = newNode
	}
	insert(n, word[1:], nochangeWord)
}

func (t *Trie) Index(word string) *Trie {
	return indexWord(t, word)
}

func indexWord(t *Trie, word string) *Trie {
	if len(word) == 0 {
		return t
	}

	c := byte(word[0])
	n := t.Next[c]
	return indexWord(n, word[1:])
}

func (t *Trie) FindPrefix(word string) string {
	if len(word) == 1 {
		return word
	}

	node := t.Index(word)
	if len(node.Next) > 0 {
		return word
	}

	p1 := node.Parent
	p2 := node

	for p1 != nil && p1.Parent != nil {
		if p1.IsWord {
			return p1.Word + string(p2.C)
		}

		if len(p1.Next) > 1 {
			return getWord(p1) + string(p2.C)
		}

		p1 = p1.Parent
		p2 = p2.Parent
	}

	return string(p2.C)
}

func getWord(t *Trie) string {

	wordSlice := make([]byte, 0)
	wordSlice = innerGetWord(t, wordSlice)
	newByte := make([]byte, 0)
	i := len(wordSlice) - 1
	for i >= 0 {
		newByte = append(newByte, wordSlice[i])
		i--
	}
	return strings.TrimSpace(string(newByte))
}

func innerGetWord(t *Trie, s []byte) []byte {
	if t == nil {
		return s
	}

	if t.Parent != nil {
		s = append(s, t.C)
	}
	return innerGetWord(t.Parent, s)
}
