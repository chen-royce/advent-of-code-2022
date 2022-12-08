package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
	Size     int
	DataType dataType
}

type dataType string

const (
	FOLDER dataType = "folder"
	FILE   dataType = "file"

	DOLLAR_SIGN uint8 = 36
)

func main() {
	filepath := "./problem/input-simple.txt"
	// filepath := "./problem/input.txt"
	input, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	root := createNode("dir root")
	currNode := root

	for scanner.Scan() {
		currLine := scanner.Text()
		if isCommand(currLine) {
			command := extractCommand(currLine)
			switch command[0] {
			case "cd":
				nextNode := currNode.cd(command[1])
				log.Println("Navigating from", currNode.Metadata.Name, "to", nextNode.Metadata.Name)
				currNode = nextNode
				continue
			case "ls":
				log.Println("Showing contents of", currNode.Metadata.Name)
			}
		} else {
			newNode := createNode(currLine)
			newNode.Parent = currNode
			log.Println("Found:", newNode)
			currNode.Children = append(currNode.Children, newNode)
		}
	}

	sizes := []int{}
	names := []string{}

	root.getAllFolderNamesAndSizesOverX(0, &sizes, &names)
	log.Println("SIZES:", sizes)
	log.Println("NAMES:", names)
	log.Println(sumIntSlice(sizes))
	log.Println(sizes[findClosestLargerOrEqualNumberIdx(8381165, sizes)])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func sumIntSlice(is []int) int {
	var total int
	for _, num := range is {
		total += num
	}
	return total
}

func findClosestLargerOrEqualNumberIdx(target int, input []int) int {
	var currClosest int
	var currClosestIdx int
	var largerOrEqualNumFound bool

	for idx, num := range input {
		if num >= target {
			if !largerOrEqualNumFound {
				currClosest = num
				currClosestIdx = idx
				largerOrEqualNumFound = true
			} else if num < currClosest {
				currClosest = num
				currClosestIdx = idx
			}
		}
	}

	return currClosestIdx
}

func (n *node) getAllFolderNamesAndSizesUnderX(x int, sizes *[]int, names *[]string) {
	n.Metadata.Size = n.calculateSize()
	if n.Metadata.Size < x && n.Metadata.DataType != FILE {
		*sizes = append(*sizes, n.Metadata.Size)
		*names = append(*names, n.Metadata.Name)
	}
	for _, child := range n.Children {
		child.getAllFolderNamesAndSizesUnderX(x, sizes, names)
	}
}

func (n *node) getAllFolderNamesAndSizesOverX(x int, sizes *[]int, names *[]string) {
	n.Metadata.Size = n.calculateSize()
	if n.Metadata.Size > x && n.Metadata.DataType != FILE {
		*sizes = append(*sizes, n.Metadata.Size)
		*names = append(*names, n.Metadata.Name)
	}
	for _, child := range n.Children {
		child.getAllFolderNamesAndSizesOverX(x, sizes, names)
	}
}

func (n *node) calculateSize() int {
	if len(n.Children) == 0 {
		return n.Metadata.Size
	}
	var sumOfChildrenSizes int
	for _, child := range n.Children {
		sumOfChildrenSizes += child.calculateSize()
	}
	n.Metadata.Size = sumOfChildrenSizes
	return sumOfChildrenSizes
}

func isCommand(s string) bool {
	return s[0] == DOLLAR_SIGN
}

func extractCommand(s string) []string {
	split := strings.Split(s, " ")
	return split[1:]
}

func createNode(s string) *node {
	nodeData := strings.Split(s, " ")

	var name string = nodeData[1]
	var size int
	var dataType dataType

	if nodeData[0] == "dir" {
		dataType = FOLDER
	} else {
		var err error
		size, err = strconv.Atoi(nodeData[0])
		if err != nil {
			log.Fatal("Convert file size to int")
		}
		dataType = FILE
	}
	return &node{
		Metadata: metadata{
			DataType: dataType,
			Name:     name,
			Size:     size,
		},
	}
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
