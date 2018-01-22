package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type slice []string

func (s slice) Len() int {
	return len(s)
}
func (s slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s slice) Less(i, j int) bool {
	sa := strings.Split(s[i], "/")
	sb := strings.Split(s[j], "/")
	a, _ := strconv.Atoi(sa[1])
	b, _ := strconv.Atoi(sb[1])
	return a < b
}

func main() {
	var xxx = []string{"192.168.34.40/42", "192.168.34.40/32"}
	fmt.Println(xxx[0])
	sort.Sort(slice(xxx))
	fmt.Println(xxx[0])
}
