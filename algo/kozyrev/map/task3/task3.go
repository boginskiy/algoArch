package main

import "fmt"

// Фильтрация по условию - работа с множествами и пересечениями

func main() {
	set1 := map[int]struct{}{
		1: struct{}{},
		2: struct{}{},
		3: struct{}{},
		4: struct{}{},
		5: struct{}{},
	}

	set2 := map[int]struct{}{
		3: struct{}{},
		4: struct{}{},
		6: struct{}{},
		8: struct{}{},
	}

	res := SetIntersection(set1, set2)
	fmt.Println(res)

}

func SetIntersection(set1, set2 map[int]struct{}) map[int]struct{} {
	for k := range set2 {
		if _, ok := set1[k]; !ok {
			delete(set2, k)
		}
	}
	return set2
}
