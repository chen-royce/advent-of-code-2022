package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// steps
// 1. build tree
//		1a. extract commands and run them
//			- ls populates children
//			- cd navigates
// 2. populate tree with values (do i need this step?)
// 3. search tree

type node struct {
	Parent   *node
	Children []*node
	Metadata metadata
}

type metadata struct {
	Name     string
	Size     string
	DataType dataType
}

type dataType string

const (
	FOLDER dataType = "folder"
	FILE   dataType = "file"

	DOLLAR_SIGN uint8 = 36
)

func main() {
	filepath := "./problem/input.txt"
	input, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var root *node
	currNode := root

	for scanner.Scan() {
		currLine := scanner.Text()
		if isCommand(currLine) {
			command := extractCommand(currLine)
			switch command[0] {
			case "cd":
				nextNode := currNode.cd(command[1])
				log.Println("Navigating from %s to %s", currNode.Metadata.Name, nextNode.Metadata.Name)
				currNode = nextNode
				continue
			case "ls":
				log.Println("Showing contents of")
			}
		} else {
			newNode
			currNode.Children = append(currNode.Children)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// func to extract node data from current line if not command

func isCommand(s string) bool {
	return s[0] == DOLLAR_SIGN
}

func extractCommand(s string) []string {
	split := strings.Split(s, " ")
	return split[1:]
}

func (n *node) ls(s *Scanner) err {

}

func (n *node) cd(dst string) *node {
	switch dst {
	case "/":
		if n.Parent == nil {
			return n
		} else {
			return n.Parent.cd("/")
		}
	case "..":
		return n.Parent
	default:
		return n.findChild(dst)
	}
}

func (n *node) findChild(name string) *node {
	for _, child := range n.Children {
		if child.Metadata.Name == name {
			return child
		}
	}
	return nil
}
