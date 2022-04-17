package main

import (
	Fr "practice/fileReaders"
)

func BellmanFord(srcNode string, dstNode string) (totalWeight int, path []string) {

	// Retrive an array of nodes and an array of edges
	nodes := Fr.Get_All_Nodes("data.txt")
	edges := Fr.Get_All_Edges("data.txt")
	// Create distances and predecessors arrays
	var distances []int
	var predecessors []int
	// Find ids of srcNode and dstNode in the overall nodes array
	srcId, dstId := Find_Node_Ids(srcNode, dstNode)
	// Predefine values in distances and predecessors
	for i := 0; i < len(nodes); i++ {
		if i == srcId {
			distances[i] = 0
		} else {
			// Make sure the value is greater than the sum of all weights - needs to be larger than the maximum possible length path
			distances[i] = 100000000
		}
		predecessors[i] = -1
	}

	// Perform the body of the algorithm n-1 times where n is the number of vertices
	// Note: This can be optimised
	for x := 0; x < len(nodes); x++ {
		// Loop through all of the edges
		for j:=0; j<len(edges); j++ {

			// Extract necessary data from the edges
			startNode := edges[j].U
			endNode := edges[j].V
			weight := edges[j].W
			// Get the ids of the start and end nodes
			startId, endId := Find_Node_Ids(startNode, endNode)

			// Set up variable to identify negative-weight loops
			negLoop := false

			// For the given start node (for the current edge) calculate the current shortest path
			for currentId:=startId; currentId!=srcId; {
				nextNodeId := predecessors[currentId]
				// If there is a null value in the shortest path then there is NO shortest path currently. This would cause an infinite loop
				if nextNodeId==-1 {
					break
				}
				// If one of the cities in the shortest path matches the id for the end node then considering the current edge would cause a negative-weight cycle
				if nextNodeId==endId {
					// Notice that this sets neg_loop=true
					negLoop = true
					break
				}
				currentId = nextNodeId
			}

			// If neg_loop=true this calculation of the shortest path gets skipped
			if negLoop==false {
				// If the weight of the current shortest path for the end node is greater than the path to the start node along with the current edge then we have a new shortest path to the end node via the start node
				if distances[startId] + weight < distances[endId] {
					distances[endId] = distances[startId] + weight
					predecessors[endId] = startId
				}
			}
		}
	}

	// Find total weight of path to specified destination
	totalWeight = distances[dstId]
	// Calculate the path from the predecessors array
	// Keeping track of n just allows us to keep track of the length of the path
	n := 1
	for currentIndex := dstId; currentIndex!=srcId; {
		nextNodeId := predecessors[currentIndex]
		path = append(path, nodes[nextNodeId].Name)
		currentIndex = nextNodeId
		n++
	}

	return
}
