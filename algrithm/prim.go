package main

import (
	"fmt"
	"strconv"
)

type point struct {
	name string
}

type edge struct {
	lp    *point
	rp    *point
	value int
}

type points []point
var Ps points
type edges []edge
var Es edges

func prepare() {
	Ps = make([]point, 6)
	for i := 0; i < len(Ps); i++ {
		Ps[i].name = "city" + strconv.Itoa(i)
	}
	Es = make([]edge, 10)
	for i := 0; i < 6; i++ {
		Es[i].lp = &Ps[i%6]
		Es[i].rp = &Ps[(i+1)%6]
		Es[i].value = i + 14%20
	}
	for i := 6; i < 10; i++ {
		Es[i].lp = &Ps[i%6]
		Es[i].rp = &Ps[(i+2)%6]
		Es[i].value = i + 14%20
	}
}

func (this points) contain(p *point) bool {
	for _, v := range this {
		if v.name == p.name {
			return true
		}
	}
	return false
}

func (e *edge) compare(p edge) {

}

func (e *edge)belong(p *point) bool{

}

func (this points)Equal(p points ) bool{
	if len(Ps)<len(this) {
		return false	
	}
	return true
}

func prim() {
	fmt.Println(Ps, Es)
	newPoints := make([]point, 6)
	newPoints[0] = Ps[3]
	
}
