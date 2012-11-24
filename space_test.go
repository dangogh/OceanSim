package oceansim

import (
	"testing"
	"fmt"
)

var testDebug = false
func printDebug(f string, x... interface {}) {
	if testDebug {
		fmt.Printf(f, x...)
	}
}

func TestSpaces(t *testing.T) {
	var objs = [...]Space{
		Plankton{},
		Empty{},
		Rock{},
		Fish{},
		Shark{} }
	for ii, obj0 := range objs {
		obj0.Act()
		printDebug("Testing %v::React\n", obj0)
		for jj, obj1 := range objs {
			printDebug("%v(%v).Reacts(%v(%v)) ?\n", obj0, ii, obj1, jj)
			blocks, reacts := obj0.React(obj1)
			printDebug("   %v, %v\n", blocks, reacts)
		}
	}
}
