package main

import (
	"fmt"
	"go-redis/src"
)

func main() {
	Iterator := src.NewStringOperation().MGet("name", "age", "ab").Iterator()
	for Iterator.HasNext() {
		fmt.Println(Iterator.Next())
	}

}
