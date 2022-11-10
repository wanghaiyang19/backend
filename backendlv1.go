package main

import "fmt"

var (
	s             []int
	m, a, l, r, t int
)

func main() {
	fmt.Scan(&a)
	fmt.Scan(&l)
	fmt.Scanln(&r)
	for i := 0; i < a; i++ {
		fmt.Scan(&m)
		s = append(s, m)
	}
	for t := l; t < r; t++ {
		if s[l] > s[l+1] {
			s[l], s[l+1] = s[l+1], s[l]
		}
	}
	fmt.Println(s)
}
