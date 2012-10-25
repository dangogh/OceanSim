package oceansim

// Space interface -- each type of space must implement Act() and React()
type Space interface {
  Act()
  React(Space) bool
}

////////////////////////////////////////////////////////////////
// Empty Space
type Empty struct {
}

func (x Empty) Act() {
	// do nothing
}

func (x Empty) React(Space) bool {
	// do not block
	return false
}

////////////////////////////////////////////////////////////////
// Rock Space
type Rock struct {
}

func (x Rock) Act() {
	// do nothing
}

func (x Rock) React(Space) bool {
	// block
	return true
}
