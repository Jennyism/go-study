package gee

import (
	"strings"
)

type geeNode interface {
	matchChild(part string) *node
	matchChildren(part string) []*node
	insert(pattern string, parts []string, height int)
	search(parts []string, height int) *node
}

type node struct {
	pattern  string // 待匹配路由，如果为空，表示该节点是中间节点，而非匹配路由
	part     string
	children []*node
	isWild   bool
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.pattern == part || child.isWild {
			return child
		}
	}

	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			pattern: pattern,
			part:    part,
			isWild:  part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		res := child.search(parts, height+1)
		if res != nil {
			return res
		}
	}

	return nil
}
