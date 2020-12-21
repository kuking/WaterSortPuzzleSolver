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

var Level106 = Level{Vials: []Vial{
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
}}

func main() {

	fmt.Println("Solving Level 105")
	sol := Level105.Solve(false)
	for i, s := range sol {
		fmt.Printf("Move %2d: %2d -> %2d\n", i+1, s[0]+1, s[1]+1)
	}

}
