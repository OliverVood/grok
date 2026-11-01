package grok

import (
	"fmt"

	"github.com/OliverVood/grok/internal/search"
)

func main() {
	value, ok := search.Binary([]int{1, 3}, -3)

	fmt.Println(value, ok)
}

func Binary(items []int, item int) (int, error) {
	return search.Binary(items, item)
}
