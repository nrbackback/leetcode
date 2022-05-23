package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string]int, len(strs))
	var r [][]string
	for _, s := range strs {
		sortedResult := sortString(s)
		index, ok := m[sortedResult]
		if ok {
			r[index] = append(r[index], s)
			continue
		}
		r = append(r, []string{s})
		m[sortedResult] = len(r) - 1
	}
	sort.Slice(r, func(i, j int) bool {
		return len(r[i]) < len(r[j])
	})
	return r
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
