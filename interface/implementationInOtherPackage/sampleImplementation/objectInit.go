package sampleImplementation

type InterfaceCopy interface {
	Execute()
}

type ST struct {
	inter InterfaceCopy
}

func NewST(inter InterfaceCopy) *ST {
	return &ST{inter: inter}
}

func (s *ST) Process() {
	s.inter.Execute()
}