package main

import (
	"testing"
)

var EMPTY_VIAL = Vial{AIR, AIR, AIR, AIR}
var HALF_RED_VIAL = Vial{AIR, AIR, RED, RED}
var ONE_GRAY_VIAL = Vial{AIR, AIR, AIR, GRAY}
var ONE_RED_VIAL = Vial{AIR, AIR, AIR, RED}
var FULL_ORANGE_VIAL = Vial{ORANGE, ORANGE, ORANGE, ORANGE}
var FULL_MIX_VIAL = Vial{GRAY, GRAY, ORANGE, RED}
var FULL_MIXDUPE_VIAL = Vial{RED, GRAY, ORANGE, RED}
var INVALID_AIR_VIAL = Vial{RED, AIR, ORANGE, DBLUE}

func TestVial_Valid(t *testing.T) {
	if !EMPTY_VIAL.Valid() || !HALF_RED_VIAL.Valid() || !FULL_MIX_VIAL.Valid() || !FULL_MIXDUPE_VIAL.Valid() {
		t.Fatal()
	}
	if INVALID_AIR_VIAL.Valid() {
		t.Fatal()
	}
}

func TestVial_Finished(t *testing.T) {
	if EMPTY_VIAL.Finished() || HALF_RED_VIAL.Finished() || FULL_MIXDUPE_VIAL.Finished() {
		t.Fatal()
	}
	if !FULL_ORANGE_VIAL.Finished() {
		t.Fatal()
	}
}

func TestVial_TopColor(t *testing.T) {
	if EMPTY_VIAL.TopColor() != AIR || HALF_RED_VIAL.TopColor() != RED || FULL_MIXDUPE_VIAL.TopColor() != RED {
		t.Fatal()
	}
}

func TestVial_TopQty(t *testing.T) {
	if EMPTY_VIAL.TopQty() != 0 {
		t.Fatal()
	}
	if HALF_RED_VIAL.TopQty() != 2 || FULL_ORANGE_VIAL.TopQty() != 4 || FULL_MIX_VIAL.TopQty() != 2 {
		t.Fatal()
	}
}

func TestVial_SpaceLeft(t *testing.T) {
	if EMPTY_VIAL.SpaceLeft() != 4 || HALF_RED_VIAL.SpaceLeft() != 2 || FULL_ORANGE_VIAL.SpaceLeft() != 0 {
		t.Fatal()
	}
}

func TestVial_Empty(t *testing.T) {
	if !EMPTY_VIAL.Empty() || HALF_RED_VIAL.Empty() || FULL_ORANGE_VIAL.Empty() || FULL_MIXDUPE_VIAL.Empty() {
		t.Fail()
	}
}

func TestVial_Full(t *testing.T) {
	if EMPTY_VIAL.Full() || HALF_RED_VIAL.Full() || !FULL_ORANGE_VIAL.Full() || !FULL_MIXDUPE_VIAL.Full() {
		t.Fail()
	}
}

func TestVial_CanPourInto(t *testing.T) {
	if EMPTY_VIAL.CanPourInto(&EMPTY_VIAL) || EMPTY_VIAL.CanPourInto(&HALF_RED_VIAL) {
		t.Fatal()
	}
	if !HALF_RED_VIAL.CanPourInto(&EMPTY_VIAL) ||
		!HALF_RED_VIAL.CanPourInto(&HALF_RED_VIAL) ||
		!ONE_RED_VIAL.CanPourInto(&HALF_RED_VIAL) ||
		!HALF_RED_VIAL.CanPourInto(&ONE_RED_VIAL) {
		t.Fatal()
	}
	if ONE_GRAY_VIAL.CanPourInto(&ONE_RED_VIAL) {
		t.Fatal()
	}
	if ONE_GRAY_VIAL.CanPourInto(&FULL_MIX_VIAL) {
		t.Fatal()
	}
	if !FULL_ORANGE_VIAL.CanPourInto(&EMPTY_VIAL) ||
		!FULL_MIX_VIAL.CanPourInto(&EMPTY_VIAL) ||
		!ONE_RED_VIAL.CanPourInto(&EMPTY_VIAL) {
		t.Fatal()
	}
}

func TestVial_PourInto(t *testing.T) {
	HRV := HALF_RED_VIAL
	EV := EMPTY_VIAL
	HRV.PourInto(&EV)
	if HRV != EMPTY_VIAL || EV != HALF_RED_VIAL {
		t.Fatal()
	}

	OV := FULL_ORANGE_VIAL
	EV = EMPTY_VIAL
	OV.PourInto(&EV)
	if OV != EMPTY_VIAL || EV != FULL_ORANGE_VIAL {
		t.Fatal()
	}

	V1 := FULL_MIXDUPE_VIAL
	V2 := HALF_RED_VIAL
	V1.PourInto(&V2)
	V1Exp := Vial{AIR, GRAY, ORANGE, RED}
	V2Exp := Vial{AIR, RED, RED, RED}
	if V1 != V1Exp || V2 != V2Exp {
		t.Fatal()
	}

	V1 = HALF_RED_VIAL
	V2 = HALF_RED_VIAL
	V1.PourInto(&V2)
	V2Exp = Vial{RED, RED, RED, RED}
	if V1 != EMPTY_VIAL || V2 != V2Exp {
		t.Fatal()
	}
}

func TestLevel_HashCode_DeepCopy(t *testing.T) {
	InitialiseLevelBuffers(10)

	var level = BuildLevel([]Vial{
		{DBLUE, DGREEN, LBLUE, LBLUE},
		{VIOLET, PINK, LGREEN, GRAY},
		{ORANGE, VIOLET, RED, BROWN},
	})
	levelHash := level.HashCode()
	level.Vials[2][0] = RED
	if level.HashCode() == levelHash {
		t.Fatal()
	}
	levelHash = level.HashCode()
	levelCopy := *level.BufferedDeepCopy()
	if levelHash != levelCopy.HashCode() {
		t.Fatal()
	}
	levelCopy.Vials[2][0] = VIOLET
	if levelHash == levelCopy.HashCode() {
		t.Fatal()
	}
}

func TestLevel_Solved(t *testing.T) {
	var level = BuildLevel([]Vial{
		{DBLUE, DGREEN, LBLUE, LBLUE},
		{VIOLET, PINK, LGREEN, GRAY},
		{AIR, AIR, RED, BROWN},
	})
	if level.Solved() {
		t.Fatal()
	}

	level = BuildLevel([]Vial{
		{DBLUE, DBLUE, DBLUE, DBLUE},
		{AIR, AIR, AIR, VIOLET},
	})
	if level.Solved() {
		t.Fatal()
	}

	level = BuildLevel([]Vial{
		{DBLUE, DBLUE, DBLUE, DBLUE},
		{VIOLET, VIOLET, VIOLET, VIOLET},
		{AIR, AIR, AIR, AIR},
	})
	if !level.Solved() {
		t.Fatal()
	}

}
