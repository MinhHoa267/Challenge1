package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0)
	for i := 1; i <= 5; i++ {
		s = append(s, i)
		fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))
	}
	t := []int{5, 10, 15, 20, 25}
	tslice := t[1:4]
	fmt.Println("slice:", t)
	fmt.Println("slice sau khi cat:", tslice)

	s1 := []string{"a", "b", "c"}
	s2 := append([]string(nil), s1...)
	s1[0] = "Hòa"
	s2[1] = "Hòa"
	fmt.Println("s1:", s1, "s2:", s2)
}