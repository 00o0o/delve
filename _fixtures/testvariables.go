package main

import "fmt"

type FooBar struct {
	Baz int
	Bur string
}

func foobar(baz string) {
	var (
		a1 = "foo"
		a2 = 6
		a3 = 7.23
		a4 = [2]int{1, 2}
		a5 = []int{1, 2, 3, 4, 5}
		a6 = FooBar{Baz: 8, Bur: "word"}
		a7 = &FooBar{Baz: 5, Bur: "strum"}
	)

	fmt.Println(a1, a2, a3, a4, a5, a6, a7, baz)
}

func main() {
	foobar("bazburzum")
}
