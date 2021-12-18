package main

import (
	"fmt"
	"strings"
)

func main() {
	recurse(`package main

import (
	"fmt"
	"strings"
)

func main() {`, `
}`, `
func recurse(A, B, R string) {
	fmt.Println(A)
`+"	fmt.Println(\"	recurse(`%s`, `%s`, `%s`), A, B, R)\")"+`
	fmt.Println(B)
	fmt.Println(R)
}`)
}

func recurse(A, B, R string) {
	fmt.Println(A)
	repl := strings.ReplaceAll(R, "\"", "\\\"")
	fmt.Printf("	recurse(`%s`, `%s`, `%s`)", A, B, repl)
	fmt.Println(B)
	fmt.Println(R)
}
