package main

import (
	"fmt"
	"go-redis/src"
	"log"
	"time"
)

func main() {
	//Iterator := src.NewStringOperation().MGet("name", "age", "ab").Iterator()
	//for Iterator.HasNext() {
	//	fmt.Println(Iterator.Next())
	//}

	//fmt.Println(src.NewStringOperation().Set("addr1","beijing",
	//	src.WithExpire(time.Second*5),
	//	src.WithNx(),
	//))

	newscache := src.NewSimpleCache(src.NewStringOperation(), time.Second*17)
	newscache.DBGetter = func() string {
		log.Println("get from db")
		return "NewsData"
	}
	fmt.Println(newscache.GetCache("123"))
}
