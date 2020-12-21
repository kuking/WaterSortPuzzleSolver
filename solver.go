package main

import (
	"fmt"
	"runtime"
)

var NO_SOLUTION = [][2]int{}

func (l *Level) solveRecurse(explored map[uint64]bool, deep int) (solution [][2]int) {
	if explored[l.HashCode()] {
		return NO_SOLUTION
	} else {
		explored[l.HashCode()] = true
	}
	if deep > 150 {
		//fmt.Println("abort, too-deep", deep, l)
		return NO_SOLUTION
	}

	if len(explored)%1000000 == 0 {
		fmt.Println("Explored playouts", len(explored), "deep", deep, "snapshot:", *l)
		runtime.GC()
	}

	for i, vialI := range l.Vials {
		for j, vialJ := range l.Vials {
			if i < j {
				// left->right: i->j
				// avoid pouring an full vial into an empty one
				if !vialI.Finished() && vialI.CanPourInto(&vialJ) {
					work := l.DeepCopy()
					work.Vials[i].PourInto(&work.Vials[j])
					thisSolution := [][2]int{{i, j}}
					if work.Solved() {
						return thisSolution
					}
					tailSolution := work.solveRecurse(explored, deep+1)
					if len(tailSolution) > 0 {
						return append(thisSolution, tailSolution...)
					}
				}
// dont move full into empty
				// right->left: j->i
				if !vialJ.Finished() && vialJ.CanPourInto(&vialI) {
					work := l.DeepCopy()
					work.Vials[j].PourInto(&work.Vials[i])
					thisSolution := [][2]int{{j, i}}
					if work.Solved() {
						return thisSolution
					}
					tailSolution := work.solveRecurse(explored, deep+1)
					if len(tailSolution) > 0 {
						return append(thisSolution, tailSolution...)
					}
				}
			}
		}
	}
	return NO_SOLUTION
}

func (l *Level) Solve() (solution [][2]int) {
	if !l.Valid() {
		return NO_SOLUTION
	}
	work := l.DeepCopy()
	var explored = map[uint64]bool{}
	return work.solveRecurse(explored, 0)
}
