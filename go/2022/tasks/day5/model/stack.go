package model

type Stack struct {
	Crates []Crate
}

func (s *Stack) Add(crate Crate) {
	s.Crates = append(s.Crates, crate)
}

func (s *Stack) Remove() (crate Crate) {
	crate = s.Crates[len(s.Crates)-1]
	s.Crates = s.Crates[0 : len(s.Crates)-1]
	return
}

func (s *Stack) Top() (crate Crate) {
	return s.Crates[len(s.Crates)-1]
}
