package pkg

import (
	"fmt"
	"os"
)

// making combinations of different non-intersecting paths
func (g *Graph) CombineAndSort() {
	if len(g.SolvedPaths) == 0 {
		fmt.Println("no paths leading to end Room")
		os.Exit(1)
	}
	g.SortPaths() // we need it to combine our groups as short as possible
	var GroupsMix [][][]string
	allPaths := g.SolvedPaths
	for i, onePath := range allPaths {
		var Group [][]string
		Group = append(Group, onePath)       // adding it to group [[path]]
		GroupsMix = append(GroupsMix, Group) // adding it to groupmix [[[path]]]
		for j := i + 1; j < len(allPaths); j++ {
			pathToCompare := allPaths[j]        // taking next path to compare
			if !compare(Group, pathToCompare) { // if it is not the same, we add it to group, and later group to groupMix
				Group = append(Group, pathToCompare)
				GroupsMix = append(GroupsMix, Group)
			}
		}
	}
	g.sortCo(GroupsMix)
}

// comparing rooms for possible non-intersecting path
func compare(grup [][]string, path []string) bool {
	for _, pathInGrup := range grup {
		for _, roomX := range pathInGrup[:len(pathInGrup)-1] { // trimming end room
			for _, roomY := range path[:len(path)-1] {
				if roomX == roomY {
					return true
				}
			}
		}
	}
	return false
}

// with sorting combinations function we
// fill up the map with only possible shortest combinations of n paths in there
func (g *Graph) sortCo(combs [][][]string) {
	g.PathGroups = make(map[int][][]string) // paths combinations under int(number of paths in combination)
	for _, comb := range combs {
		category := len(comb)                    // 1, 2, 3.. number of paths in combination
		currentCombLenght := getCombLenght(comb) // lenght of combination paths
		_, ok := g.PathGroups[category]          // checking if something lays under given category
		if ok {                                  // if smth is in the map, we compare it`s lenght with new one, if new one is shorter, the assign shortest path to be the right one
			valueInMap := g.PathGroups[category]
			valueCombLenght := getCombLenght(valueInMap)
			if currentCombLenght < valueCombLenght {
				g.PathGroups[category] = comb
			}
		} else { // if nothing is in the map, we assign first value(our combination) for the map
			g.PathGroups[category] = comb
		}
	}
}

func getCombLenght(comb [][]string) int {
	counter := 0
	for _, path := range comb {
		counter += len(path)
	}
	return counter
}

func (g *Graph) SortPaths() {
	for i := 0; i < len(g.SolvedPaths)-1; i++ {
		for j := 0; j < len(g.SolvedPaths)-i-1; j++ {
			if len(g.SolvedPaths[j]) > len(g.SolvedPaths[j+1]) {
				g.SolvedPaths[j+1], g.SolvedPaths[j] = g.SolvedPaths[j], g.SolvedPaths[j+1]
			}
		}
	}
}

// we need to achieve
// [ [[h n e end]]  [[h n e end] [t E a m end]]  [[t E a m end] [0 o n e end] [h A c k end]] ]
