package main

import "fmt"

var Level105 = Level{Vials: []Vial{
	{DBLUE, DGREEN, LBLUE, LBLUE},
	{VIOLET, PINK, LGREEN, GRAY},
	{ORANGE, VIOLET, RED, BROWN},
	{ORANGE, PINK, RED, ORANGE},
	{DGREEN, RED, YELLOW, DBLUE},
	{YELLOW, DGREEN, BROWN, DGREEN},
	{BROWN, VIOLET, RED, BGREEN},
	{BGREEN, VIOLET, PINK, BGREEN},
	{LGREEN, GRAY, LBLUE, DBLUE},
	{BROWN, YELLOW, GRAY, LGREEN},
	{GRAY, YELLOW, BGREEN, DBLUE},
	{LGREEN, LBLUE, PINK, ORANGE},
	{AIR, AIR, AIR, AIR},
	{AIR, AIR, AIR, AIR},
}}

func main() {

	fmt.Println("Solving Level 105")
	sol := Level105.Solve()
	for i, s := range sol {
		fmt.Printf("Move %2d: %2d -> %2d\n", i+1, s[0]+1, s[1]+1)
	}

}