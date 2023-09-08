package gbank

import "strings"

type Node struct {
	part     string
	pattern  string
	isWild   bool
	children []*Node
}

func (n *Node) Insert(parts []string, pattern string, height int) {
	if len(parts) == height {
		// 若所有的part都成功插入，设置当前节点的pattern，设置可以匹配到完整路径
		n.pattern = pattern
		return
	}

	// 查找匹配的子节点，若存在则递归向下继续查找，若没有则新建节点
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		// 新建节点，若part以 : 或 * 开头，则是通配节点
		child = &Node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.Insert(parts, pattern, height+1)
}

func (n *Node) Search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// 若所有的part都匹配到了对应的节点 或 当前节点的part是 * 说明通配part后面全部的字段
		if n.pattern == "" {
			// 这里再做一次校验，若当前节点的pattern不是空，说明匹配到了完整路径，则匹配成功
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		// 递归查询子节点
		res := child.Search(parts, height+1)
		if res != nil {
			return res
		}
	}
	return nil
}

// 匹配当前节点的子节点中是否有指定part的节点
func (n *Node) matchChild(part string) *Node {
	for _, child := range n.children {
		if child.pattern == part || child.isWild {
			return child
		}
	}
	return nil
}

// 匹配所有part等于传入参数相等的子节点
func (n *Node) matchChildren(part string) []*Node {
	var res []*Node
	for _, child := range n.children {
		if child.part == part || child.isWild {
			res = append(res, child)
		}
	}
	return res
}
