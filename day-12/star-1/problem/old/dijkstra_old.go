package main

import "sort"

const (
	S uint8 = 83
	Z uint8 = 90
)

var (
	elevations = "abcdefghijklmnopqrstuvwxyz"
)

func generateElevationsMap() map[uint8]int {
	// initialize map
	elevationsMap := make(map[uint8]int)

	// generate mapping from letters -> elevations
	for i := range elevations {
		elevationsMap[elevations[i]] = i
	}
	// add special characters to map
	elevationsMap[S] = 0
	elevationsMap[Z] = 25

	return elevationsMap
}

type adjacencyList map[int](map[int]int)

// buildAdjacencyList maps the position of a node from a 2D input
// array to arrays containing references to: 1) another node that
// the node is connected to, and 2) edge weights;
func buildAdjacencyList(input [][]uint8) adjacencyList {
	elevationsMap := generateElevationsMap()
	adjacencyList := adjacencyList{}

	rowLength := len(input[0])
	for rowNum, row := range input {
		for colNum, curr := range row {
			currLocation := (rowNum * rowLength) + colNum
			// find adjacent vertices
			if rowNum != 0 {
				topNeighborLocation := ((rowNum - 1) * rowLength) + colNum
				topNeighborVal := input[rowNum-1][colNum]
				addEdgeToAdjacencyList(currLocation, topNeighborLocation, curr, topNeighborVal, elevationsMap, adjacencyList)
			}
			if rowNum != len(input)-1 {
				botNeighborLocation := ((rowNum + 1) * rowLength) + colNum
				botNeighborVal := input[rowNum+1][colNum]
				addEdgeToAdjacencyList(currLocation, botNeighborLocation, curr, botNeighborVal, elevationsMap, adjacencyList)
			}
			if colNum != 0 {
				leftNeighborLocation := (rowNum * rowLength) + (colNum - 1)
				leftNeighbor := input[rowNum][colNum-1]
				addEdgeToAdjacencyList(currLocation, leftNeighborLocation, curr, leftNeighbor, elevationsMap, adjacencyList)
			}
			if colNum != len(input[0])-1 {
				rightNeighborLocation := (rowNum * rowLength) + (colNum + 1)
				rightNeighbor := input[rowNum][colNum+1]
				addEdgeToAdjacencyList(currLocation, rightNeighborLocation, curr, rightNeighbor, elevationsMap, adjacencyList)
			}
		}
	}
	return adjacencyList
}

// addEdgeToAdjacencyList takes a node and its neighbor, an elevationsMap,
// and an adjacencyList and updates the adjacencyList if it can find a
// valid edge after calculating the elevation difference via the elevationsMap
func addEdgeToAdjacencyList(currLocation, neighborLocation int, currNodeVal, neighboringVal uint8, elevationsMap map[uint8]int, adjList adjacencyList) {
	elevationDiff := elevationsMap[neighboringVal] - elevationsMap[currNodeVal]
	if elevationDiff <= 1 {
		adjList[currLocation][neighborLocation] = 1
	}
}

// findNodeLocation finds the first instance of a node, delimited
// by its uint8/byte character
func findNodeLocation(needle uint8, haystack [][]uint8) int {
	inputWidth := len(haystack[0])
	for row := range haystack {
		for col := range haystack[0] {
			if haystack[row][col] == needle {
				return (row * inputWidth) + col
			}
		}
	}
	return -1
}

func shortestPath(startingLoc, endingLoc, gridLength, gridHeight int, adjList adjacencyList) int {
	// intialize map of visited nodes and the cost to get there
	shortestPaths := make(map[int]int)
	shortestPaths[startingLoc] = 0 // seed value

	startRow := startingLoc % gridLength
	startCol := startingLoc - (startRow * gridLength)

	// TODO: recursively traverse neighbors until ending location is found

	return shortestPaths[endingLoc]
}

func checkNeighbors(currRow, currCol, endingLoc, gridLength, gridHeight int, adjList adjacencyList, shortestPaths map[int]int) {
	var neighbors []int
	currLoc := currRow * gridLength + currCol
	if currRow != 0 {
		topNeighborLocation := ((currRow - 1) * gridLength) + currCol
		// check if a path exists to top neighbor
		edges := adjList[currLoc]
		if 
		neighbors = append(neighbors, topNeighborLocation)
		if shortest, ok := shortestPaths[topNeighborLocation]; !ok || adjList[]
	}
	if currRow != gridHeight-1 {
		botNeighborLocation := ((currRow + 1) * gridLength) + currCol
		neighbors = append(neighbors, botNeighborLocation)
	}
	if currCol != 0 {
		leftNeighborLocation := (currRow * gridLength) + (currCol - 1)
		neighbors = append(neighbors, leftNeighborLocation)
	}
	if currCol != gridLength-1 {
		rightNeighborLocation := (currRow * gridLength) + (currCol + 1)
		neighbors = append(neighbors, rightNeighborLocation)
	}

	sort.Ints(neighbors)

	for _, neighbor := range neighbors {

	}
}
