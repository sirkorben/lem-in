package pkg

import (
	"fmt"
	"runtime"
	"time"
)

var possiblePath = []string{}
var isVisited = make(map[string]bool)

func (g *Graph) SearchForAllPaths() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	start := time.Now()
	fmt.Println("Searching paths using DFS...")
	g.DFSrecursive(g.StartRoom)
	fmt.Printf("Done. Found #%v paths in %v.\n", len(g.SolvedPaths), time.Since(start))
	fmt.Println()

}

func (g *Graph) DFSrecursive(room string) {
	isVisited[room] = true
	possiblePath = append(possiblePath, room)
	if room == g.EndRoom {
		// cannot add possiblePath straight to g.SolvedPaths
		// without making copy I loose some room names in finded possiblePath. for example [start 8 5 end] laying down in g.SolvedPath as [[start 8]]
		var copy []string
		copy = append(copy, possiblePath[1:]...)    // trim off start
		g.SolvedPaths = append(g.SolvedPaths, copy) // can trim off finish here [:len(copy)-1]
	} else {
		var neighbours []string = g.AdjacencyList[room]
		for _, val := range neighbours {
			if !isVisited[val] {
				g.DFSrecursive(val)
			}
		}
	}
	possiblePath = possiblePath[:len(possiblePath)-1]
	isVisited[room] = false
}
