package main

import (
	"hash/maphash"
)

type Color byte

const (
	AIR    Color = iota
	VIOLET Color = iota
	PINK   Color = iota
	RED    Color = iota
	ORANGE Color = iota
	YELLOW Color = iota
	BROWN  Color = iota
	GRAY   Color = iota
	DGREEN Color = iota
	LGREEN Color = iota
	BGREEN Color = iota
	DBLUE  Color = iota
	LBLUE  Color = iota
)

type Vial [4]Color

type Level struct {
	Vials []Vial
}

func (v *Vial) Valid() bool {
	onAir := v[0] == AIR
	for i := 0; i < 4; i++ {
		if v[i] == AIR {
			if !onAir {
				return false
			}
		} else {
			onAir = false
		}
	}
	return true
}

func (v *Vial) Finished() bool {
	c0 := v[0]
	if c0 == AIR {
		return false
	}
	for i := 1; i < 4; i++ {
		if v[i] != c0 {
			return false
		}
	}
	return true
}

func (v *Vial) TopColor() Color {
	for i := 0; i < 4; i++ {
		if v[i] != AIR {
			return v[i]
		}
	}
	return AIR
}

func (v *Vial) TopQty() (qty int) {
	c := v.TopColor()
	if c == AIR {
		return 0
	}
	i := 0
	for ; i < 4 && v[i] != c; i++ {
	}
	qty = 0
	for ; i < 4 && v[i] == c; i++ {
		qty++
	}
	return
}

func (v *Vial) SpaceLeft() (left int) {
	left = 0
	for i := 0; i < 4; i++ {
		if v[i] == AIR {
			left++
		} else {
			return
		}
	}
	return
}

func (v *Vial) CanPourInto(o *Vial) bool {
	if o.Finished() {
		return false
	}
	vColor := v.TopColor()
	if vColor == AIR {
		return false
	}
	oColor := o.TopColor()
	if vColor != oColor && oColor != AIR {
		return false
	}
	return v.TopQty() <= o.SpaceLeft()
}

func (v *Vial) PourInto(o *Vial) {
	vColor := v.TopColor()
	vQty := v.TopQty()
	for i := o.SpaceLeft() - 1; vQty > 0; i-- {
		o[i] = vColor
		vQty--
	}
	vQty = v.TopQty()
	vLeft := v.SpaceLeft()
	for i := vLeft; i < vLeft+vQty; i++ {
		v[i] = AIR
	}
}

var levelHashSeed = maphash.MakeSeed()

func (l *Level) HashCode() uint64 {
	var h maphash.Hash
	h.SetSeed(levelHashSeed)
	for _, vial := range l.Vials {
		for i := 0; i < len(vial); i++ {
			if h.WriteByte(byte(vial[i])) != nil {
				panic("Couldn't calculate hash")
			}
		}
	}
	return h.Sum64()
}

func (l *Level) DeepCopy() (copy *Level) {
	copy = &Level{
		Vials: make([]Vial, len(l.Vials)),
	}
	for i, vial := range l.Vials {
		copy.Vials[i] = vial
	}
	return
}
