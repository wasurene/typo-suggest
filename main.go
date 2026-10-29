package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

type Result struct {
	Dist int
	Word string
}

func main() {
	input := flag.String("input", "data/words.txt", "Файл с словарем")

	output := flag.String("output", "out.txt", "Файл с результатами")

	max := flag.Int("max", 2, "Максимальное расстояние Левенштейна")

	query := flag.String("query", "", "Слово для поиска")
	flag.Parse()

	words, err := readWords(*input)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}

	if *query == "" {
		fmt.Println("Ошибка: не указан поисковый запрос")
		os.Exit(1)
	}

	var result []Result

	for i := range words {
		dist := levenshtein(*query, words[i])
		if dist <= *max {
			result = append(result, Result{
				Dist: dist,
				Word: words[i],
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Dist == result[j].Dist {
			return result[i].Word < result[j].Word
		}
		return result[i].Dist < result[j].Dist
	})

	for i := range result {
		fmt.Printf("%s (Расстояние: %d)\n", result[i].Word, result[i].Dist)
	}

	err = writeResult(result, *output)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}
}
