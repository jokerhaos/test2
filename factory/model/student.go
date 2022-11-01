package model

type student struct {
	name  string
	Score float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		n,
		s,
	}
}

func (s *student) GetName() string {
	// return (*s).name
	return s.name
}
