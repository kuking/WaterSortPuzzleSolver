package main

import "testing"

// I am not expecting to implement complex solving tests, but a couple just to verify it returns data in correct format
func TestLevel_Solve(t *testing.T) {

	InitialiseLevelBuffers(10)
	var level = BuildLevel([]Vial{
		{AIR, AIR, BROWN, BROWN},
		{AIR, AIR, BROWN, BROWN},
	})

	sol := level.Solve(false, false)
	if len(sol) != 1 ||
		sol[0][0] != 0 || sol[0][1] != 1 { // pour vial 0 into vial 1
		t.Fatal()
	}
}

var forBenchmark = BuildLevel([]Vial{
	{VIOLET, RED, BGREEN, PINK},
	{PINK, LGREEN, DBLUE, DBLUE},
	{DBLUE, LBLUE, LBLUE, GRAY},
	{PINK, LGREEN, BGREEN, LGREEN},
	{GRAY, VIOLET, ORANGE, ORANGE},
	{PINK, RED, ORANGE, LGREEN},
	{RED, BGREEN, GRAY, GRAY},
	{ORANGE, VIOLET, RED, BGREEN},
	{VIOLET, DBLUE, LBLUE, LBLUE},
	{AIR, AIR, AIR, AIR},
	{AIR, AIR, AIR, AIR},
})

func BenchmarkLevel_Solve(b *testing.B) {
	InitialiseLevelBuffers(100)
	forBenchmark.Solve(true, false)

}
