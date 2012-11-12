package oceansim

import (
	"testing"
)

func TestSpaces(t *testing.T) {
	s0,s1 := Empty{}, Empty{}
	r0,r1 := Rock{}, Rock{}

	////////////////////
	// Empty Space
	s0.Act()
	blocks, reacts := s0.React(s1)

	// does not block
	if blocks {
		t.Fail()
	}
	// does not react
	if reacts {
		t.Fail()
	}

	////////////////////
	// Rock
	blocks, reacts = r0.React(r1)
	// blocks
	if ! blocks {
		t.Fail()
	}
	// does not react
	if reacts {
		t.Fail()
	}
}

