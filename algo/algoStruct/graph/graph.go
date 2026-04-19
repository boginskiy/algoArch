package main

/*
	Breadth-First Search - алгорит поиска в ширину

	Поиск в ширину позволяет найти кратчайшее расстояние между двумя объ­ектами.

	Алгоритм поиска, который работает для графа
	Отвечает на вопросы двух типов.
		тип 1: существует ли путь от узла А к узлу В?
		тип 2: как выглядит кратчайший путь от узла А к узлу В?
*/

/*
	Depth First Search - поиск в глубину

*/

type Graph struct {
	ajList map[int][]int //Список смежности вершин
}

func NewGraph(verticesCount int) *Graph {
	return &Graph{
		ajList: make(map[int][]int),
	}
}

// Метод добавления ребра
func (g *Graph) AddEdge(u, v int) {
	g.ajList[u] = append(g.ajList[u], v)
	g.ajList[v] = append(g.ajList[v], u)
}

func main() {
	// Graph List
	graph := map[string][]string{}
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
}
