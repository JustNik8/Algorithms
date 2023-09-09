package main

func main() {

}

type set map[byte]*TrieNode

type TrieNode struct {
	data     byte
	isKey    bool
	children set
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	root := &TrieNode{
		data:     ' ',
		isKey:    false,
		children: make(set),
	}

	return Trie{
		root: root,
	}
}

func (this *Trie) Insert(word string) {
	var dfs func(node *TrieNode, index int)
	dfs = func(node *TrieNode, index int) {
		letter := word[index]
		child, hasChild := node.children[letter]

		if !hasChild {
			child = &TrieNode{
				data:     letter,
				isKey:    false,
				children: make(set),
			}
			node.children[letter] = child
		}

		isKey := len(word) == index+1

		if !isKey {
			dfs(child, index+1)
		} else {
			child.isKey = isKey
		}
	}

	dfs(this.root, 0)

}

func (this *Trie) Search(word string) bool {
	var dfs func(node *TrieNode, index int) bool
	dfs = func(node *TrieNode, index int) bool {
		if index == len(word) {
			return node.isKey
		}

		letter := word[index]
		child, hasChild := node.children[letter]

		if !hasChild {
			return false
		}

		return dfs(child, index+1)

	}

	return dfs(this.root, 0)
}

func (this *Trie) StartsWith(prefix string) bool {
	var dfs func(node *TrieNode, index int) bool
	dfs = func(node *TrieNode, index int) bool {
		if index == len(prefix) {
			return true
		}

		letter := prefix[index]
		child, hasChild := node.children[letter]

		if !hasChild {
			return false
		}

		return dfs(child, index+1)

	}

	return dfs(this.root, 0)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
