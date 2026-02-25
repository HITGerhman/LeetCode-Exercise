package main

import (
	"fmt"
)

type Investigator struct {
	ID int
}

func main() {
	users := []Investigator{{ID: 1}, {ID: 2}, {ID: 3}}
	var usePtrs []*Investigator
	for _, u := range users {
		usePtrs = append(usePtrs, &u)
	}
	fmt.Println("%d,%d,%d", usePtrs[0].ID, usePtrs[1].ID, usePtrs[2].ID)
}
