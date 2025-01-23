package main

import "fmt"

type menu struct {
	val int
}

func (m *menu) mul() {
	for i := 1; i <= m.val; i++ {
		for j := 1; j <= m.val; j++ {
			fmt.Printf("%d*%d = %d ", i, j, i*j)
		}
		fmt.Println()
	}
}
func main() {
	me := menu{val: 9}
	me.mul()
}
