package main

import (
	"fmt"
	"strings"
)

type trieTree struct {
	root *node
}

type node struct {
	char      rune
	childrens map[rune]*node
	exist     []int
	fail      *node
	count     int
}

func (t *trieTree) Insert(rowStr string) {
	itorater := t.root
	for _, char := range rowStr {
		if _, ok := itorater.childrens[char]; !ok {
			itorater.childrens[char] = newNode(char, 0)
		}
		itorater = itorater.childrens[char]
		itorater.count++
	}
	itorater.exist = []int{len(rowStr)}
}

func (t *trieTree) Remove(rowStr string) {
	itorater := t.root
	for _, char := range rowStr {
		child, ok := itorater.childrens[char]
		if !ok {
			break
		}
		if child.count == 1 {
			delete(itorater.childrens, char)
			break
		} else {
			child.count--
			itorater = child
		}
	}
}

func (t *trieTree) BuildAc() {
	var queue = []*node{t.root}
	for len(queue) != 0 {
		var newQueue []*node
		for _, treeNode := range queue {
			for char, child := range treeNode.childrens {
				newQueue = append(newQueue, child)
				if treeNode != t.root {
					p := treeNode.fail
					if p != nil {
						if _, ok := p.childrens[char]; ok {
							child.fail = p.childrens[char]
							if len(child.fail.exist) != 0 {
								child.exist = append(child.exist, child.fail.exist...)
							}
							break
						}
					}
				}
				child.fail = t.root
			}
		}
		queue = newQueue
	}
}

func (t *trieTree) Print() {
	var queue = []*node{t.root}
	for len(queue) != 0 {
		var newQueue []*node
		for _, char := range queue {
			fmt.Printf("[char:%c,exsit:%v] ", char.char, char.exist)
			for _, child := range char.childrens {
				newQueue = append(newQueue, child)
			}
		}
		fmt.Println()
		queue = newQueue
	}
}

func (t *trieTree) Query(str string) [][]int {
	str = strings.ToLower(str)
	var res [][]int = make([][]int, len(str))
	itorater := t.root
	for idx, char := range str {
		if ignoreChar(byte(char)) {
			continue
		}
		if _, ok := itorater.childrens[char]; ok {
			itorater = itorater.childrens[char]
			if len(itorater.exist) != 0 { //找到单词
				res[idx] = itorater.exist
			}
			continue
		}
		for itorater.fail != nil {
			itorater = itorater.fail
			if _, ok := itorater.childrens[char]; ok {
				itorater = itorater.childrens[char]
				if len(itorater.exist) != 0 { //找到单词
					res[idx] = itorater.exist
				}
				break
			}
		}
	}
	return res
}

func Fillter(str string, res [][]int) string {
	var newStr = make([]rune, len(str))
	var times int
	for i := len(str) - 1; i >= 0; i-- {
		if !ignoreChar(str[i]) {
			if times != 0 {
				times--
				newStr[i] = '*'
				continue
			}
			if len(res[i]) > 0 {
				times = res[i][0] - 1
				newStr[i] = '*'
				continue
			}
		}
		newStr[i] = rune(str[i])
	}
	return string(newStr)
}

func newNode(char rune, count int) *node {
	n := new(node)
	n.char = char
	n.childrens = make(map[rune]*node)
	n.count = count
	return n
}

func ignoreChar(char byte) bool {
	if char == byte(' ') || char == byte('*') {
		return true
	}
	return false
}

func main() {
	Tree := new(trieTree)
	Tree.root = newNode(rune(-1), 1)
	Tree.Insert("fuck")
	Tree.Insert("shit")
	Tree.Insert("gay")
	Tree.Insert("bitch")
	Tree.Insert("funcyou")
	Tree.BuildAc()
	Tree.Print()
	Tree.Remove("fuck")
	Tree.BuildAc()
	Tree.Print()
	// str := "i f uC k this b*i t* Ch feel like fuck shit you are a gay"
	// res := Tree.Query(str)
	// fmt.Println("res", res)
	// Fillter(str,res)
	// fmt.Println("newStr:", str)
}
