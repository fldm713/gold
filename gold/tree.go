package gold

import (
	"strings"
)

type trieNode struct {
	name     string
	children []*trieNode
	pattern  string
}

// Insert path: /user/get/:id
func (t *trieNode) Insert(path string) {
	root := t
	strs := strings.Split(path, "/")
	for index, name := range strs {
		if index == 0 {
			continue
		}
		children := t.children
		isMatched := false
		for _, node := range children {
			if node.name == name {
				isMatched = true
				t = node
				break
			}
		}
		if !isMatched {
			node := &trieNode{
				name:     name,
				children: make([]*trieNode, 0),
				pattern:    t.pattern + "/" + name,
			}
			children = append(children, node)
			t.children = children
			t = node
		}
	}
	t = root
}

func (t *trieNode) Find(path string) (*trieNode, string, map[string]string) {
	var bestMatched *trieNode
	var params map[string]string = make(map[string]string)
	strs := strings.Split(path, "/")
	for index, name := range strs {
		if index == 0 {
			continue
		}
		children := t.children
		bestMatched = nil
		for _, node := range children {
			if node.name == name  {
				bestMatched = node
				break
			} else if strings.HasPrefix(node.name, ":") {
				bestMatched = node
			} else if node.name == "*" {
				if bestMatched == nil {
					bestMatched = node
				}
			}
		}
		if bestMatched == nil {
			break
		}
		t = bestMatched
		if strings.HasPrefix(bestMatched.name, ":") {
			key := bestMatched.name[1:]
			params[key] = name
		}
		if index == len(strs)-1 {
			return t, t.pattern, params
		}
	}
	return nil, "", params
}
