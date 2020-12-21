package main

import (
	"fmt"
	"runtime"
	"time"
)

var NO_SOLUTION = [][2]int{}

func (l *Level) exploreMove(explored map[uint64]bool, shortest bool, verbose bool, deep int, i int, j int) (solution *[][2]int) {
	vialI := l.Vials[i]
	vialJ := l.Vials[j]
	if !vialI.CanPourInto(&vialJ) {
		return nil
	}
	innocuous := vialI.TopQty()+vialI.SpaceLeft() == 4 && vialJ.SpaceLeft() == 4
	if innocuous {
		return nil
	}
	work := l.DeepCopy()
	work.Vials[i].PourInto(&work.Vials[j])
	thisSolution := [][2]int{{i, j}}
	if work.Solved() {
		return &thisSolution
	}
	tailSolution := work.solveRecurse(explored, shortest, verbose, deep+1)
	if len(tailSolution) > 0 {
		sol := append(thisSolution, tailSolution...)
		return &sol
	}
	return nil
}

func (l *Level) solveRecurse(explored map[uint64]bool, shortest bool, verbose bool, deep int) (solution [][2]int) {
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
		if verbose {
			fmt.Println("Explored playouts", len(explored), "deep", deep, "snapshot:", *l)
		}
		runtime.GC()
	}

	var best *[][2]int
	for i, _ := range l.Vials {
		for j, _ := range l.Vials {
			if i < j {
				// left->right: i->j
				sol := l.exploreMove(explored, shortest, verbose, deep, i, j)
				if sol != nil && (best == nil || len(*best) > len(*sol)) {
					best = sol
				}
				// right->left: j->i
				sol = l.exploreMove(explored, shortest, verbose, deep, j, i)
				if sol != nil && (best == nil || len(*best) > len(*sol)) {
					best = sol
				}
			}
			if best != nil && (len(*best) == 1 || !shortest) {
				return *best
			}
		}
	}

	if best == nil {
		return NO_SOLUTION
	} else {
		return *best
	}
}

func (l *Level) Solve(shortest bool, verbose bool) (solution [][2]int) {
	t0 := time.Now()
	if !l.Valid() {
		return NO_SOLUTION
	}
	work := l.DeepCopy()
	var explored = map[uint64]bool{}
	solution = work.solveRecurse(explored, shortest, verbose, 0)
	duration := time.Now().Sub(t0)
	if verbose {
		fmt.Printf("Solution took: %v, exploring %v moves, or %v/move\n", duration, len(explored), duration/time.Duration(len(explored)))
	}
	return
}
