package main

type Trie struct {
	child [26]*Trie
	isEnd bool
}

func ConstructorXIII() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	trie := this

	for _, ch := range word {
		idx := ch - 'a'
		if trie.child[idx] == nil {
			trie.child[idx] = &Trie{}
		}
		trie = trie.child[idx]
	}
	trie.isEnd = true
}

func (this *Trie) Search(word string) bool {
	trie := this

	for _, ch := range word {
		idx := ch - 'a'
		if trie.child[idx] == nil {
			return false
		}
		trie = trie.child[idx]
	}
	return trie.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	trie := this

	for _, ch := range prefix {
		idx := ch - 'a'
		if trie.child[idx] == nil {
			return false
		}
		trie = trie.child[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
