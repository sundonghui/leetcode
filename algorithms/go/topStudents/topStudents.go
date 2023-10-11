package topstudents

import (
	"fmt"
	"sort"
	"strings"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	words := make(map[string]int, len(positive_feedback)+len(negative_feedback))
	for _, v := range positive_feedback {
		words[v] = 3
	}
	for _, v := range negative_feedback {
		words[v] = -1
	}
	fmt.Println(words)
	type pair struct {
		id    int
		score int
	}
	list := make([]pair, len(report))
	for i, v := range report {
		score := 0
		for _, w := range strings.Split(v, " ") {
			score += words[w]
		}
		list[i] = pair{
			score: score,
			id:    student_id[i],
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].score > list[j].score || list[i].score == list[j].score && list[i].id < list[j].id
	})
	res := make([]int, k)
	for i, p := range list[:k] {
		res[i] = p.id
	}
	return res
}
