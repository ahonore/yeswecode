package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type node struct {
	value    rune
	children map[rune]*node
}

func newNode(value rune) *node {
	return &node{
		value:    value,
		children: map[rune]*node{},
	}
}

func (n *node) countNodes() int {
	sum := 1
	for _, v := range n.children {
		sum += v.countNodes()
	}

	return sum
}

func (n *node) addChild(child *node) {
	if child == nil {
		return
	}

	n.children[child.value] = child
}

func (n *node) addEntry(person string) {
	parentNode := n
	for _, r := range person {
		nextNode, ok := parentNode.children[r]
		if !ok {
			nextNode = newNode(r)
			parentNode.addChild(nextNode)
		}
		parentNode = nextNode
	}
}

const (
	inputFile = "input.txt"
	flagFile  = "flag.txt"
)

func main() {
	in, _ := os.Open(inputFile)
	defer in.Close()
	out, _ := os.Create(flagFile)
	defer out.Close()

	sc := bufio.NewScanner(in)
	rmDiac := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	t := newNode(' ')
	for sc.Scan() {
		s, _, _ := transform.String(rmDiac, strings.ToLower(sc.Text()))
		t.addEntry(s)
	}
	out.WriteString(strconv.Itoa(t.countNodes() - 1))
}
