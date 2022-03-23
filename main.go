package main

import "lem-in/pkg"

var g pkg.Graph

func main() {
	g.ReadFile()
	g.SearchForAllPaths()
	g.CombineAndSort()
	g.GetBestComb()
	g.SendAnts(g.AntNumber, g.BestComb)
}
