package main

import (
	"log"

	"github.com/petar/GoLLRB/llrb"
)

type Num struct {
	num int
}

func NewNum(num int) *Num {
	return &Num{num: num}
}

func (n *Num) Less(than llrb.Item) bool {
	return n.num < than.(*Num).num
}

func Iterator(item llrb.Item) bool {
	log.Println("item: ", item)
	return true
}

func main() {
	tree := llrb.New()

	tree.InsertNoReplace(NewNum(1))
	tree.InsertNoReplace(NewNum(1))
	tree.InsertNoReplace(NewNum(3))
	tree.InsertNoReplace(NewNum(3))
	tree.InsertNoReplace(NewNum(2))
	log.Printf("len: %v", tree.Len())
	tree.AscendGreaterOrEqual(NewNum(1), Iterator)
	deleted := tree.Delete(NewNum(3))
	log.Println("deleted: ", deleted)
	len := tree.Len()
	for i := 0; i < len; i++ {
		itemMin := tree.Min()
		itemMax := tree.Max()
		log.Printf("min %v max %v", itemMin, itemMax)
		tree.DeleteMin()
	}

}
