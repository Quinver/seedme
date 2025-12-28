package search

import (
	"seedme/internal/model"
	"sort"
)

func All(query []string) ([]model.Torrent, error) {
	var results []model.Torrent

	nyaa, _ := SearchNyaa(query)
	uIndex, _ := SearchUIndex(query)

	results = append(results, nyaa...)
	results = append(results, uIndex...)

	sort.Slice(results, func(i, j int) bool {
		return results[i].Seeds > results[j].Seeds
	})

	return results, nil
}
