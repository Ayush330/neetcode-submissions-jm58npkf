type PrefixTree struct {
	root *Node
}

type Node struct{
	children [26]*Node
	isEnd bool
}

func Constructor() PrefixTree {
    return PrefixTree{
		&Node{},
	}
}

func (this *PrefixTree) Insert(word string) {
	curr := this.root
	for i:=0; i<len(word); i++{
		idx := word[i] - 'a'
		if curr.children[idx] == nil{
			curr.children[idx] = &Node{}
		}
		curr = curr.children[idx] 
	}
	curr.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	curr := this.root
	for i:=0; i<len(word); i++{
		idx := word[i] - 'a'
		if curr.children[idx] == nil{
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this.root
	for i:=0; i<len(prefix); i++{
		idx := prefix[i] - 'a'
		if curr.children[idx] == nil{
			return false
		}
		curr = curr.children[idx]
	}
	return true
}
