package pkg

type Graph struct {
	AntNumber     int
	Ants          []Ant
	StartRoom     string
	EndRoom       string
	AdjacencyList map[string][]string
	SolvedPaths   [][]string
	PathGroups    map[int][][]string
	BestComb      [][]string
}

type Ant struct {
	Id       int
	Skip     bool
	Path     []string
	RoomId   int
	Previous string
}

var Task string

func (g *Graph) GetAdjList(rooms []string, links [][]string) {
	if g.AdjacencyList == nil {
		g.AdjacencyList = make(map[string][]string)
	}
	for _, v := range rooms {
		g.AdjacencyList[v] = []string{}
	}
	for _, value := range links {
		first := value[0]
		second := value[1]
		g.AdjacencyList[first] = append(g.AdjacencyList[first], second)
		g.AdjacencyList[second] = append(g.AdjacencyList[second], first)
	}
}
