package dataStructure

import "testing"

func TestBuildTree(t *testing.T) {
	//nodesStr := []string{"10","5","15","null","null","6","20"}
	nodesStr := []string{"3", "1", "5", "0", "2", "4", "6", "null", "null", "null", "3"}
	//BuildTree(nodesStr)
	root := BuildTree(nodesStr)
	root.PrintTree()
}
