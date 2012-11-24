package oceansim

import "fmt"

var debug = false

func trace(x Space, methodName string) string {
	if debug {
		m := fmt.Sprintf("%v::%s", x, methodName)
		fmt.Println("entering ", m)
		return m
	}
	return ""
}

func un(s string) {
	if debug {
		fmt.Println("leaving ", s)
	}
}

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
	defer un(trace(x, "Act"))
	// do nothing
}

func (x Empty) React(Space) (blocks, reacts bool) {
	defer un(trace(x, "React"))
	// do not block or react
	return false, false
}

func (x Empty) String() string {
	return "Empty"
}

////////////////////////////////////////////////////////////////
// Rock Space
type Rock struct {
}

func (x Rock) Act() {
	defer un(trace(x, "Act"))
	// do nothing
}

func (x Rock) React(Space) (blocks, reacts bool) {
	defer un(trace(x, "React"))
	// block
	return true, false
}

func (x Rock) String() string {
	return "Rock"
}

// Critter interface -- is a space that moves and reproduces
type Critter interface {
	// include Space interface
	Space
	Move() Space
	Reproduce() Space

	// private methods
	reproduce() bool
	move() bool
}


// Plankton is simplest form of Critter -
//  no size, but moves and reproduces at regular intervals
type Plankton struct {
	Critter

	// private data members
	timeSinceReproduced_ uint
	hunger_ uint
}

// Public interface
func (x Plankton) Act() {
	defer un(trace(x, "Act"))
	if x.isTimeToReproduce() {
		x.reproduce()
	} else {
		x.move()
	}
}

// Public
//  React to another invading my space
func (x Plankton) React(other Space) (blocks, reacts bool) {
	defer un(trace(x, "React"))
	return true, true
}

func (x Plankton) String() string {
	return "Plankton"
}

// Private
//  is it time to reproduce?
func (x Plankton) isTimeToReproduce() bool {
	return x.timeSinceReproduced_ >= 3
}

func (x Plankton) reproduce() {
	x.timeSinceReproduced_ = 0
}

func (x Plankton) move() {
	defer un(trace(x, "move"))
}


// Fish is a Critter with size that eats lesser Critters
//  including Plankton and smaller Fish.
// It reproduces at intervals depending on size.
// Size is dependent on age and diet
type Fish struct {
	// include Critter interface
	Critter

	// private data members
	size_ uint
	timeSinceReproduced_ uint
	hunger_ uint
}

// Public interface
func (x Fish) Act() {
	defer un(trace(x, "Act"))
	if x.isHungry() {
		x.eat()
		if x.isStarving() {
			//x.die()
		}
	} else if x.isTimeToReproduce() {
		if x.reproduce() {
			x.timeSinceReproduced_ = 0
		}
	} else {
		x.move()
	}
	x.timeSinceReproduced_ += 1
	x.hunger_ += 1
	x.size_ += 1
}

//  React to another invading my space
func (x Fish) React(other Space) (blocks, reacts bool) {
	defer un(trace(x, "React"))
	return true, true
}

func (x Fish) String() string {
	return "Fish"
}

// Private methods
//  is it time to reproduce?
func (x Fish) isTimeToReproduce() bool {
	return x.timeSinceReproduced_ >= 3
}

func (x Fish) isHungry() bool {
	return x.hunger_ > 10
}

func (x Fish) isStarving() bool {
	return x.hunger_ > 20
}

func (x Fish) eat() {
	// TODO: find something nearby to eat
	x.hunger_ = 0
}

func (x Fish) reproduce() bool {
	return false
}

func (x Fish) move() bool {
	return false
}



// Shark is a Fish. It will eat any size Fish and smaller
// Sharks.  It will attack a larger Shark if starving.
type Shark struct {
	// include Fish interface
	Fish

	// private data members
	size_ uint
	timeSinceReproduced_ uint
	hunger_ uint
}

// Public interface
func (x Shark) Act() {
	defer un(trace(x, "Act"))
	if x.isHungry() {
		x.eat()
		if x.isStarving() {
			//x.die()
		}
	} else if x.isTimeToReproduce() {
		if x.reproduce() {
			x.timeSinceReproduced_ = 0
		}
	} else {
		x.move()
	}
	x.timeSinceReproduced_ += 1
	x.hunger_ += 1
	x.size_ += 1
}

//  React to another invading my space
func (x Shark) React(other Space) (blocks, reacts bool) {
	defer un(trace(x, "React"))
	return true, true
}

func (x Shark) String() string {
	return "Shark"
}

// Private methods
//  is it time to reproduce?
func (x Shark) isTimeToReproduce() bool {
	return x.timeSinceReproduced_ >= 3
}

func (x Shark) isHungry() bool {
	return x.hunger_ > 10
}

func (x Shark) isStarving() bool {
	return x.hunger_ > 20
}

func (x Shark) eat() {
	// TODO: find something nearby to eat
	x.hunger_ = 0
}

func (x Shark) reproduce() bool {
	return false
}

func (x Shark) move() bool {
	return false
}

