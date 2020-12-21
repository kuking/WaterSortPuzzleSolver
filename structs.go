package main

import (
	"fmt"
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
	Size  int
	Vials [20]Vial
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
	for i := 3; i >= 0; i-- {
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

func (v *Vial) TopQty() (qty int) { // perf
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
	oSpaceLeft := o.SpaceLeft()
	if oSpaceLeft == 0 {
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
	vTopQty := v.TopQty()
	return vTopQty <= oSpaceLeft
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

func (l *Level) Valid() bool {
	var counts = map[Color]int{}
	for i := 0; i < l.Size; i++ {
		if !l.Vials[i].Valid() {
			return false
		}
		for j := 0; j < 4; j++ {
			counts[l.Vials[i][j]]++
		}
	}
	for _, v := range counts {
		if v%4 != 0 {
			return false
		}
	}
	return true
}

var levelHashSeed = maphash.MakeSeed()

func (l *Level) HashCode() uint64 {
	var h maphash.Hash
	h.SetSeed(levelHashSeed)
	for i := 0; i < l.Size; i++ {
		for j := 0; j < len(l.Vials[i]); j++ {
			if h.WriteByte(byte(l.Vials[i][j])) != nil {
				panic("Couldn't calculate hash")
			}
		}
	}
	return h.Sum64()
}

func BuildLevel(vials []Vial) (l Level) {
	l = Level{
		Size:  len(vials),
		Vials: [20]Vial{},
	}
	for i, vial := range vials {
		l.Vials[i] = vial
	}
	return
}

var levelBuffers []*Level

func InitialiseLevelBuffers(size int) {
	levelBuffers = []*Level{}
	for i := 0; i < size; i++ {
		level := BuildLevel([]Vial{})
		levelBuffers = append(levelBuffers, &level)
	}
}

func (l *Level) BufferedDeepCopy() (copy *Level) {
	if len(levelBuffers) == 0 {
		panic("I ran out of Level buffers")
	}
	copy = levelBuffers[0]
	levelBuffers = levelBuffers[1:]
	*copy = *l
	return
}

func (l *Level) ReturnBuffer() {
	l.Size = 0
	levelBuffers = append(levelBuffers, l)
}

func (l *Level) Solved() bool {
	for i := 0; i < l.Size; i++ {
		if !l.Vials[i].Finished() && l.Vials[i].SpaceLeft() != 4 {
			return false
		}
	}
	return true
}

var hexes = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func (l *Level) String() string {
	res := fmt.Sprintf("{s: %2d", l.Size)
	for i := 0; i < l.Size; i++ {
		res += " "
		for j := 0; j < len(l.Vials[i]); j++ {
			res += string(hexes[int(l.Vials[i][j])])
		}
	}
	res += "}"
	return res
}
