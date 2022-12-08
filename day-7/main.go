package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	// filepath := "./problem/input-simple.txt"
	filepath := "./problem/input.txt"
	input, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// initialize root node
	root := createNode("dir root")

	// initialize current node pointer
	currNode := root

	for scanner.Scan() {
		// parse file line by line
		currLine := scanner.Text()
		if isCommand(currLine) {
			command := extractCommand(currLine)
			switch command[0] {
			// execute cd
			case "cd":
				nextNode := currNode.cd(command[1])
				log.Println("Navigating from", currNode.Metadata.Name, "to", nextNode.Metadata.Name)
				currNode = nextNode
				continue
			// ls will just log, we'll handle the actual logic in else loop below since we can consider
			// non-commands (files and `dir` to be our base input)
			case "ls":
				log.Println("Showing contents of", currNode.Metadata.Name)
			}
		} else {
			// create a new node to store the current line's data in and then link it to its parent
			newNode := createNode(currLine)
			newNode.Parent = currNode
			log.Println("Found:", newNode)
			currNode.Children = append(currNode.Children, newNode)
		}
	}

	// for solving problem: list of all directory sizes and list of names
	sizes := []int{}
	names := []string{}

	root.getAllFolderNamesAndSizes(&sizes, &names)
	// log.Println("SIZES:", sizes)
	// log.Println("NAMES:", names)
	log.Println(sizes[findClosestLargerOrEqualNumberIdx(3636666, sizes)])
	log.Println(names[findClosestLargerOrEqualNumberIdx(3636666, sizes)])

	// printAllSizes(root)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func printAllSizes(root *node) {
	log.Println(root.Metadata.Size, root.Metadata.Name)
	if len(root.Children) != 0 {
		for _, c := range root.Children {
			printAllSizes(c)
		}
	}
}

// sumSlice is a helper func to sum slice up for 1st star
func sumSlice(is []int) int {
	var total int
	for _, num := range is {
		total += num
	}
	return total
}

// findClosestLargerOrEqualNumberIdx finds the smallest number in a slice that's >= target
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

// getAllFolderNamesAndSizes attempts to update 2 input arrays with all the folder
// names and sizes of folders
func (n *node) getAllFolderNamesAndSizes(sizes *[]int, names *[]string) {
	n.Metadata.Size = n.calculateSize()
	if n.Metadata.DataType != FILE {
		*sizes = append(*sizes, n.Metadata.Size)
		*names = append(*names, n.Metadata.Name)
	}
	for _, child := range n.Children {
		child.getAllFolderNamesAndSizes(sizes, names)
	}
}

// getAllFolderNamesAndSizesUnderX attempts to update 2 input arrays with all the folder
// names and sizes of folders < size x
func (n *node) getAllFolderNamesAndSizesUnderX(x int, sizes *[]int, names *[]string) {
	n.Metadata.Size = n.calculateSize()
	if n.Metadata.Size <= x && n.Metadata.DataType != FILE {
		*sizes = append(*sizes, n.Metadata.Size)
		*names = append(*names, n.Metadata.Name)
	}
	for _, child := range n.Children {
		child.getAllFolderNamesAndSizesUnderX(x, sizes, names)
	}
}

// getAllFolderNamesAndSizesOverX attempts to update 2 input arrays with all the folder
// names and sizes of folders > size x
func (n *node) getAllFolderNamesAndSizesOverX(x int, sizes *[]int, names *[]string) {
	n.Metadata.Size = n.calculateSize()
	if n.Metadata.Size >= x && n.Metadata.DataType != FILE {
		*sizes = append(*sizes, n.Metadata.Size)
		*names = append(*names, n.Metadata.Name)
	}
	for _, child := range n.Children {
		child.getAllFolderNamesAndSizesOverX(x, sizes, names)
	}
}

// calculateSize recursively calculates the size of a folder or file
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

// isCommand checks whether a file line is a command, since commands start with '$'
func isCommand(s string) bool {
	return s[0] == DOLLAR_SIGN
}

// extractCommand removes the '$' from the start of a command and returns a slice
// containing the command followed by all its arguments
func extractCommand(s string) []string {
	split := strings.Split(s, " ")
	return split[1:]
}

// createNode initializes a node
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

// cd is a method on the node type that navigates from the node
// to the destination and returns a pointer to it
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

// findChild is a method on the node type that returns a pointer to
// a child specified by name
func (n *node) findChild(name string) *node {
	for _, child := range n.Children {
		if child.Metadata.Name == name {
			return child
		}
	}
	return nil
}
