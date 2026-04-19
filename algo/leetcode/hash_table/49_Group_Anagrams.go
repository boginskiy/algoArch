package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]uint][]string, len(strs))

	for _, sss := range strs {
		// Собираем маску строки
		var mask [26]uint
		for _, s := range sss {
			mask[s-'a']++
		}
		//
		if _, ok := groups[mask]; !ok {
			groups[mask] = make([]string, 0, 10)
		}
		groups[mask] = append(groups[mask], sss)
	}

	// Подготовка результата
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func main() {
	arr := []string{"abbbbbbbbbbb", "aaaaaaaaaaab"}

	res := groupAnagrams(arr)
	fmt.Println(res)

}

// Решение с leetcode -2
// func mapToKey(s string) [26]int8 {
// 	var k [26]int8
// 	for _, c := range s {
// 		k[c-'a']++
// 	}
// 	return k
// }
// func groupAnagrams(strs []string) [][]string {
// 	anagrams := make(map[[26]int8][]string, len(strs))

// 	for _, s := range strs {
// 		k := mapToKey(s)
// 		if _, ok := anagrams[k]; !ok {
// 			anagrams[k] = make([]string, 0, 5)
// 		}
// 		anagrams[k] = append(anagrams[k], s)
// 	}

// 	result := make([][]string, 0, len(strs))
// 	for _, v := range anagrams {
// 		result = append(result, v)
// 	}
// 	return result
// }

// Решение с leetcode.
// func groupAnagrams(strs []string) [][]string {
// 	mapper := make(map[[26]int][]string)

// 	for _, s := range strs {
// 		var count [26]int
// 		for _, ch := range s {
// 			count[ch-'a']++
// 		}
// 		mapper[count] = append(mapper[count], s)
// 	}

// 	result := make([][]string, 0, len(mapper))
// 	for _, group := range mapper {
// 		result = append(result, group)
// 	}

// 	return result
// }

// Улучшенное решение 2
// func groupAnagrams(strs []string) [][]string {
// 	result := make([][]string, 0, len(strs)/5+1)
// 	idxs := make(map[string]int)

// 	for _, v := range strs {
// 		delta := [26]int{}

// 		//
// 		for _, b := range v {
// 			i := int(b - 'a')
// 			delta[i]++
// 		}

// 		// Формируем компактный ключ
// 		k := make([]string, 0, 24)
// 		for i, count := range delta {
// 			if count != 0 {
// 				k = append(k, strconv.Itoa(i), strconv.Itoa(count))
// 			}
// 		}
// 		key := strings.Join(k, "|")

// 		//
// 		if _, ok := idxs[key]; !ok {
// 			idxs[key] = len(idxs)
// 			result = append(result, nil)
// 		}
// 		result[idxs[key]] = append(result[idxs[key]], v)
// 	}
// 	return result
// }

// Улучшенное решение
// func groupAnagrams(strs []string) [][]string {
// 	result := make([][]string, 0, len(strs)/5+1) // Предполагаем среднее распределение групп
// 	idxs := make(map[string]int)

// 	for _, v := range strs {
// 		delta := [26]int{}

// 		for _, b := range v {
// 			i := int(b - 'a')
// 			delta[i]++
// 		}

// 		key := "" // Формируем компактный ключ
// 		for _, count := range delta {
// 			key += strconv.Itoa(count) + "#"
// 		}

// 		if _, exists := idxs[key]; !exists {
// 			idxs[key] = len(result)
// 			result = append(result, nil) // Оставляем резерв под новую группу
// 		}
// 		result[idxs[key]] = append(result[idxs[key]], v)
// 	}

// 	return result
// }

// Мое решение. Неоптимальное ППЦ
// func groupAnagrams(strs []string) [][]string {
// 	result := make([][]string, 0, 10)
// 	idxs := make(map[[26]int]int)

// 	for _, v := range strs {
// 		delta := [26]int{}

// 		for _, b := range v {
// 			i := int(b - 'a')
// 			delta[i] = delta[i] + 1
// 		}

// 		if i, ok := idxs[delta]; !ok {
// 			idxs[delta] = len(idxs)
// 			result = append(result, []string{v})

// 		} else {
// 			result[i] = append(result[i], v)
// 		}
// 	}

// 	return result
// }
