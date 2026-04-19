package graph

import (
	"fmt"
	"strings"
)

// Освоив основы, переходите к алгоритмам решения базовых задач на графах:
// Алгоритм Дейкстры для нахождения кратчайших путей
// Поиск в глубину (DFS) и поиск в ширину (BFS)
// Алгоритмы Краскала и Прима для построения минимального остовного дерева
// Матрица смежности и список ребер для представления графа

// Примеры задач, которые полезно решить на начальном этапе:
// Нахождение всех простых циклов в графе
// Проверка связности графа
// Определение диаметра графа

/*
	Граф на map
	интерфейс:
*/

// Vertex представляет вершину графа
type Vertex struct {
	Name string
}

func NewVertex(name string) *Vertex {
	return &Vertex{Name: name}
}

// Graph представляет собой ориентированный граф
type Graph struct {
	// Карта вершин и смежных вершин
	vertices map[*Vertex][]*Vertex
}

// AddVertex добавляет новую вершину в граф
func (g *Graph) AddVertex(v *Vertex) {
	if g.vertices == nil {
		g.vertices = make(map[*Vertex][]*Vertex)
	}
	// Изначально добавляем пустую коллекцию соседей
	g.vertices[v] = []*Vertex{}
}

// AddEdge добавляет ребро между двумя вершинами
func (g *Graph) AddEdge(from, to *Vertex) {

	// Если вершины ещё нет, добавляем её
	if _, ok := g.vertices[from]; !ok {
		g.AddVertex(from)
	}

	// Проверяем вторую вершину тоже
	if _, ok := g.vertices[to]; !ok {
		g.AddVertex(to)
	}

	// Добавляем соседнюю вершину
	g.vertices[from] = append(g.vertices[from], to)
}

// String возвращает строковое представление графа
func (g *Graph) String() string {
	var sb strings.Builder

	for v, neighbors := range g.vertices {
		sb.WriteString(fmt.Sprintf("%v -> ", v.Name))
		for i, n := range neighbors {
			sb.WriteString(n.Name)
			if i != len(neighbors)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
