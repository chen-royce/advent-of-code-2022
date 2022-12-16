package main

import (
	"log"
	"strconv"
)

const (
	S uint8 = 83
	E uint8 = 69
)

var (
	elevations = "abcdefghijklmnopqrstuvwxyz"
)

// generateElevationsMap maps elevations a-z and S/E to their elevations
func generateElevationsMap() map[uint8]int {
	// initialize map
	elevationsMap := make(map[uint8]int)

	// generate mapping from letters -> elevations
	for i := range elevations {
		elevationsMap[elevations[i]] = i
	}
	// add special characters to map
	elevationsMap[S] = 0
	elevationsMap[E] = 25

	return elevationsMap
}

// findNodeLocation finds the first instance of a node, delimited
// by its uint8/byte character
func findNodeLocation(needle uint8, haystack [][]uint8) string {
	for row := range haystack {
		for col := range haystack[0] {
			if haystack[row][col] == needle {
				return coordinatesToString(row, col)
			}
		}
	}
	return ""
}

// coordinatesToString takes a row # and col # and creates a string to use
// as a UUID for the node at given coordinates
func coordinatesToString(row, col int) string {
	return strconv.Itoa(row) + "-" + strconv.Itoa(col)
}

// adjacencyList is a map of node positions (represented as a string "row-col")
// to a map from the neighbor's position (also represented "row-col") to its
// weight
type adjacencyList map[string]map[string]int

// buildAdjacencyList returns an adjacencyList from an input
func buildAdjacencyList(input [][]uint8) adjacencyList {
	elevationsMap := generateElevationsMap()
	adjList := map[string]map[string]int{}

	for rowNum, row := range input {
		for colNum, curr := range row {
			if rowNum != 0 {
				topNeighbor := input[rowNum-1][colNum]
				if elevationsMap[curr]-elevationsMap[topNeighbor] >= -1 {
					_, ok := adjList[coordinatesToString(rowNum, colNum)]
					if !ok {
						adjList[coordinatesToString(rowNum, colNum)] = make(map[string]int)
					}
					adjList[coordinatesToString(rowNum, colNum)][coordinatesToString(rowNum-1, colNum)] = 1
				}
			}
			if rowNum != len(input)-1 {
				botNeighbor := input[rowNum+1][colNum]
				if elevationsMap[curr]-elevationsMap[botNeighbor] >= -1 {
					_, ok := adjList[coordinatesToString(rowNum, colNum)]
					if !ok {
						adjList[coordinatesToString(rowNum, colNum)] = make(map[string]int)
					}
					adjList[coordinatesToString(rowNum, colNum)][coordinatesToString(rowNum+1, colNum)] = 1
				}
			}
			if colNum != 0 {
				leftNeighbor := input[rowNum][colNum-1]
				if elevationsMap[curr]-elevationsMap[leftNeighbor] >= -1 {
					_, ok := adjList[coordinatesToString(rowNum, colNum)]
					if !ok {
						adjList[coordinatesToString(rowNum, colNum)] = make(map[string]int)
					}
					adjList[coordinatesToString(rowNum, colNum)][coordinatesToString(rowNum, colNum-1)] = 1
				}
			}
			if colNum != len(input[0])-1 {
				rightNeighbor := input[rowNum][colNum+1]
				if elevationsMap[curr]-elevationsMap[rightNeighbor] >= -1 {
					_, ok := adjList[coordinatesToString(rowNum, colNum)]
					if !ok {
						adjList[coordinatesToString(rowNum, colNum)] = make(map[string]int)
					}
					adjList[coordinatesToString(rowNum, colNum)][coordinatesToString(rowNum, colNum+1)] = 1
				}
			}
		}
	}

	return adjList
}

// findShortestPath utilizes Dijkstra's algorithm in order to create a map of
// the shortest path from a designated start node
func findShortestPaths(start, finish string, adjList adjacencyList) int {
	shortestPaths := make(map[string]int)
	shortestPaths[start] = 0
	checkNeighbors(start, adjList, shortestPaths)
	return shortestPaths[finish]
}

func checkNeighbors(node string, adjList adjacencyList, shortestPathsMap map[string]int) {
	log.Println("START NODE:", node)
	// Check node's distance from its neighbors
	for neighbor := range adjList[node] {
		log.Println("NEIGHBOR:", neighbor)
		// If we've discovered a new neighbor, add to shortestPathsMap
		if _, ok := shortestPathsMap[neighbor]; !ok {
			shortestPathsMap[neighbor] = shortestPathsMap[node] + adjList[node][neighbor]
		}
		// If we've found a shorter path to a neighbor, update the shortest paths map
		if shortestPathsMap[node]+adjList[node][neighbor] < shortestPathsMap[neighbor] {
			shortestPathsMap[neighbor] = shortestPathsMap[node] + adjList[node][neighbor]
		}
	}
	// Remove current node from adjacency list to mark it as visited
	delete(adjList, node)
	// Repeat process on smallest unvisited node in shortestPathsMap
	var toVisit []string
	for node := range shortestPathsMap {
		if _, unvisited := adjList[node]; unvisited {
			toVisit = append(toVisit, node)
		}
	}
	if len(toVisit) > 0 {
		checkNeighbors(toVisit[0], adjList, shortestPathsMap)
	}
}
