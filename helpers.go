package main

import (
	Fr "practice/fileReaders"
)

func Find_Node_Ids(src_name, dst_name string) (srcID, dstID int) {
	// Retrieve array of all n
	nodes := Fr.Get_All_Nodes("data.txt")
	// Initialise the IDs
	srcID, dstID = -1, -1
	// Calculate the node ids
	for i:=0; i<len(nodes); i++ {
		if nodes[i].Name == src_name {
			srcID = i
		} else if nodes[i].Name == dst_name {
			dstID = i
		}
		if srcID!=-1 && dstID!=-1 {
			break
		}
	}
	return
}

func Find_Node_Name(node_id int) string {
	// Retrieve array of all n
	nodes := Fr.Get_All_Nodes("data.txt")
	// Calculate the node ids
	for i:=0; i<len(nodes); i++ {
		if nodes[i].Id == node_id {
			return nodes[i].Name
		}
	}
	return ""
}
