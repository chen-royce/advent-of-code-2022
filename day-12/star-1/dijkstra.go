package main

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

type adjacencyList map[int][][]int

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
// valid edge after calculating the weight via the elevationsMap
func addEdgeToAdjacencyList(currLocation, neighborLocation int, currNodeVal, neighboringVal uint8, elevationsMap map[uint8]int, adjList adjacencyList) {
	weight := elevationsMap[neighboringVal] - elevationsMap[currNodeVal]
	if weight <= 1 {
		edgeWeightArray := []int{neighborLocation, weight}
		adjList[currLocation] = append(adjList[currLocation], edgeWeightArray)
	}
}

func shortestPath() {
	// intialize list of visited notes
	// start from src
	// go to vertex with smallest-known cost
}
