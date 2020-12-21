package main

import "testing"

// I am not expecting to implement complex solving tests, but a couple just to verify it returns data in correct format
func TestLevel_Solve(t *testing.T) {

	var level = Level{Vials: []Vial{
		{AIR, AIR, BROWN, BROWN},
		{AIR, AIR, BROWN, BROWN},
	}}

	sol := level.Solve(false)
	if len(sol) != 1 ||
		sol[0][0] != 0 || sol[0][1] != 1 { // pour vial 0 into vial 1
		t.Fatal()
	}
}
