package model

import "github.com/google/btree"

type Student struct {
	Id   int
	Name string
}

func NewStudent(id int, name string) *Student {
	return &Student{
		Id:   id,
		Name: name,
	}
}

func (s *Student) Less(than btree.Item) bool {
	return s.Id < than.(*Student).Id
}