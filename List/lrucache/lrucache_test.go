package lrucache

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//创建lru缓存
	lru := Constructor(2)
	lru.Put(1, 1)           //{1=1}
	lru.Put(2, 2)           //{1=1, 2=2}
	fmt.Println(lru.Get(1)) //返回1
	lru.Put(3, 3)           // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	fmt.Println(lru.Get(2)) // 返回 -1 (未找到)
	lru.Put(4, 4)           // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(lru.Get(1)) // 返回 -1 (未找到)
	fmt.Println(lru.Get(3)) // 返回 3
	fmt.Println(lru.Get(4)) // 返回 4
}
