package oceansim

import (
	"testing"
)

func TestSpaces(t *testing.T) {
	s := Empty{}
	r := Rock{}
	s.Act()
	// space does not block
	if s.React(r) {
		t.Fail()
	}
	// rock blocks
	if ! r.React(s) {
		t.Fail()
	}
}

