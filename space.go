package oceansim

// Space interface -- each type of space must implement Act() and React()
type Space interface {
	Act()
	React(Space) (blocks, reacts bool)
}

////////////////////////////////////////////////////////////////
// Empty Space
type Empty struct {
}

func (x Empty) Act() {
	// do nothing
}

func (x Empty) React(Space) (blocks, reacts bool) {
	// do not block or react
	return false, false
}

////////////////////////////////////////////////////////////////
// Rock Space
type Rock struct {
}

func (x Rock) Act() {
	// do nothing
}

func (x Rock) React(Space) (blocks, reacts bool) {
	// block
	return true, false
}


// Critter interface -- is a space that moves and reproduces
type Critter interface {
	// include Space interface
	Act()
	React(Space) (blocks, reacts bool)
	Move() Space
	Reproduce() Space

	// private methods
	endCycle() bool
	reproduce() bool
	move() bool
}


// Plankton is simplest form of Critter -
//  no size, but reproduces at regular intervals
type Plankton struct {
	// private data members
	age uint
	timeSinceReproduced uint
	timeSinceAte uint

}

// Public interface
func (x Plankton) Act() {
	if x.endCycle() {
		x.reproduce()
	} else {
		x.move()
	}
}

func (x Plankton) reproduce() bool {
	return false
}

func (x Plankton) move() bool {
	return false
}

// Public
//  React to another invading my space
func (x Plankton) React(other Space) (blocks, reacts bool) {
	return true, true
}

// Private
//  is it time to reproduce?
func (x Plankton) endCycle() bool {
	return x.age > 0 && x.age % x.timeSinceReproduced == 0
}

// Fish is a Critter with size that eats lesser Critters
//  including Plankton and smaller Fish.
// It reproduces at intervals depending on size.
// Size is dependent on age and diet
type Fish struct {
	// include Critter interface
	Critter
}

// Shark is a Fish. It will eat any size Fish and smaller
// Sharks.  It will attack a larger Shark if starving.
type Shark struct {
	Fish
}
