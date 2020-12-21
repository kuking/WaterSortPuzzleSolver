package main

import (
	"fmt"
	"runtime"
	"time"
)

var NO_SOLUTION = [][2]int{}

type SolvingState struct {
	t0       time.Time
	shortest bool
	verbose  bool
	explored map[uint64]bool
	moves    uint64
	depth    int
}

func (l *Level) exploreMove(ss *SolvingState, i, j int) (solution *[][2]int) {
	vialI, vialJ := l.Vials[i], l.Vials[j]
	if !vialI.CanPourInto(&vialJ) {
		return nil
	}
	innocuous := vialJ.SpaceLeft() == 4 && vialI.TopQty()+vialI.SpaceLeft() == 4
	if innocuous {
		return nil
	}
	work := l.BufferedDeepCopy()
	defer work.ReturnBuffer()
	thisSolution := [][2]int{{i, j}}
	work.Vials[i].PourInto(&work.Vials[j])
	if work.Solved() {
		return &thisSolution
	}
	ss.depth++
	tailSolution := work.solveRecurse(ss)
	ss.depth--
	if len(tailSolution) > 0 {
		sol := append(thisSolution, tailSolution...)
		return &sol
	}
	return nil
}

func (l *Level) solveRecurse(ss *SolvingState) (solution [][2]int) {
	hash := l.HashCode()
	if ss.explored[hash] {
		return NO_SOLUTION
	}
	ss.explored[hash] = true
	ss.moves++

	if ss.depth > len(l.Vials)*5 {
		return NO_SOLUTION
	}
	if ss.moves%1_000_000 == 0 {
		if ss.verbose {
			d := time.Now().Sub(ss.t0)
			fmt.Printf("mega moves: %vm, t: %v, p: %v/m, d: %v, ts: %v\n",
				ss.moves/1_000_000,
				d.Truncate(time.Second),
				(d / time.Duration(ss.moves)).Truncate(time.Nanosecond),
				ss.depth,
				l)
		}
		runtime.GC()
	}
	var best *[][2]int

	// optimisations
	var finished = []bool{}
	var topColor = []Color{}
	for i := 0; i < l.Size; i++ {
		finished = append(finished, l.Vials[i].Finished())
		topColor = append(topColor, l.Vials[i].TopColor())
	}

	for i := 0; i < l.Size; i++ {
		for j := i + 1; j < l.Size; j++ {
			// optimisation
			if (topColor[i] == topColor[j] || topColor[i] == AIR || topColor[j] == AIR) && !finished[i] && !finished[j] {
				// left->right: i->j
				sol := l.exploreMove(ss, i, j)
				if sol != nil && (best == nil || len(*best) > len(*sol)) {
					best = sol
				}
				// right->left: j->i
				sol = l.exploreMove(ss, j, i)
				if sol != nil && (best == nil || len(*best) > len(*sol)) {
					best = sol
				}
				// shall continue or is good enough?
				if best != nil && (len(*best) == 1 || !ss.shortest) {
					return *best
				}
			}
		}
	}
	if best == nil {
		return NO_SOLUTION
	}
	return *best
}

func (l *Level) Solve(shortest bool, verbose bool) (solution [][2]int) {
	if !l.Valid() {
		return NO_SOLUTION
	}
	ss := SolvingState{
		t0:       time.Now(),
		shortest: shortest,
		verbose:  verbose,
		explored: map[uint64]bool{},
		moves:    0,
		depth:    0,
	}
	work := l.BufferedDeepCopy()
	defer work.ReturnBuffer()
	solution = work.solveRecurse(&ss)
	duration := time.Now().Sub(ss.t0)
	if verbose {
		fmt.Printf("Solution took: %v, exploring %v moves, or %v/move, %v moves-per-second\n",
			duration, ss.moves, duration/time.Duration(ss.moves), ss.moves/uint64(duration.Seconds()+1))

	}
	return
}
