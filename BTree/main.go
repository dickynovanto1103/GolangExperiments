package main

import (
	"fmt"
	"github.com/dickynovanto1103/GolangExperiments/BTree/model"
	"github.com/google/btree"
	"log"
)

func Iterator(item btree.Item) bool {
	student := item.(*model.Student)
	fmt.Println("student: ", student)
	return true
}

func main() {
	tree := btree.New(2)
	tree.ReplaceOrInsert(model.NewStudent(1, "dicky"))
	tree.ReplaceOrInsert(model.NewStudent(4, "andre"))
	tree.ReplaceOrInsert(model.NewStudent(3, "andre"))

	itemSmallest := model.NewStudent(1, "andre")
	itemBiggest := model.NewStudent(4, "dicky")
	fmt.Println("ascend range")
	tree.AscendRange(itemSmallest, itemBiggest, Iterator)

	log.Println("tree len: ", tree.Len())
	log.Println("tree: ", *tree)
	log.Println("new tree len: ", tree.Len())
	tree.Ascend(Iterator)
}
