package main

import "fmt"

func StudentScores() map[string]int {
	// TODO: создайте map, добавьте студентов, измените баллы и удалите Веру.
	scores := make(map[string]int)
	scores["Анна"] = 3
	scores["Борис"] = 0
	scores["Вера"] = 5
	scores["Анна"] += 2
	delete(scores, "Вера")
	return scores
}

func main() {
	scores := StudentScores()
	fmt.Println(scores["Анна"]) // 5
	for _, name := range []string{"Борис", "Глеб"} {
		score, ok := scores[name]
		fmt.Println(score, ok)
	}
	fmt.Println(len(scores)) // 2
}
