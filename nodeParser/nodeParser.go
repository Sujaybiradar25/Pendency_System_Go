package nodeParser

import "awesomeProject2/src/Pendency_System_Go"

type childNode struct {
	nodeName   string
	count      int
	childNodes map[string]*childNode
}
type node struct {
	values     map[string][]string
	childNodes *childNode
}

func New() Pendency_System_Go.NodeParser {
	n := node{}
	n.values = make(map[string][]string)
	n.childNodes = &childNode{}
	return &n
}

/*
AddNode(path []string, value string)
DeleteNode(value string)
*/

func (n *node) AddNode(path []string, value string) {
	n.values[value] = path
	root := n.childNodes
	for i := 0; i < len(path); i++ {
		if root.childNodes == nil {
			root.childNodes = make(map[string]*childNode)
		}
		if _, exists := root.childNodes[path[i]]; !exists {
			t := new(childNode)
			root.childNodes[path[i]] = t
			root.childNodes[path[i]].nodeName = path[i]
			root.childNodes[path[i]].count++
			root = t
		} else {
			root.childNodes[path[i]].count++
			root = root.childNodes[path[i]]
		}
	}
}
func (n *node) DeleteNode(value string) {
	path := n.values[value]
	delete(n.values, value)
	root := n.childNodes
	for i := 0; i < len(path); i++ {
		if root.childNodes[path[i]] == nil {
			return
		}
		root.childNodes[path[i]].count--
		root = root.childNodes[path[i]]
	}
	return
}

func (n *node) GetCount(path []string) int {
	root := n.childNodes
	exists := false
	for i := 0; i < len(path); i++ {
		if root, exists = root.childNodes[path[i]]; !exists {
			return 0
		}
	}
	return root.count
}
