package pkg

func (g *Graph) GetBestComb() {
	minLvl := 999999
	var bestComb [][]string
	for _, comb := range g.PathGroups {
		freeSpace := 0
		totalCombLen := 0
		level := 0
		LongestPath := len(comb[len(comb)-1]) // as we got it sorted by lenght , we take longest one from the end of combination
		for i := 0; i < len(comb); i++ {
			freeSpace = freeSpace + LongestPath - len(comb[i])
			totalCombLen = totalCombLen + len(comb[i])
		}
		level = (g.AntNumber-freeSpace)/len(comb) + LongestPath // (antsnumber - freespace) / amountofpathsincomb + longhestpath
		if level < minLvl {
			minLvl = level
			bestComb = comb
		}
		// fmt.Println(comb, "free space :", freeSpace, "rooms/steps in combin :", totalCombLen)
		// fmt.Println("start level for", antsNbr, "ants is level", level)
		// fmt.Println()
	}
	g.BestComb = bestComb
}
