package ginsu 

import (
	"testing"
)

type point struct {
	x, y int
}

type line struct {
	x, y point
}

func TestFilter(t *testing.T) {
	td := struct {
		in, expect []point
	}{
		[]point{{1, 1}, {2, 2}, {3, 3},{4, 4}},
		[]point{{2, 2}, {4, 4}},
	}

	fn := func(p point) bool {
		return (p.x % 2 == 0)
	}

	res, _ := Filter(T{td.in}, F{fn})
	rs := res.I.([]point)

	if len(rs) != 2 {
		t.Errorf("filter failed")
	}
}

func TestForEach(t *testing.T) {
	td := struct {
		in, expect []point
	}{
		[]point{{1, 1}, {2, 2}, {3, 3},{4, 4}},
		[]point{{2, 2}, {4, 4}, {6, 6},{8, 8}},
	}

	res := make([]point, 0, len(td.expect))
	fn := func(p point) {
		res = append(res, point{p.x*2, p.y*2})
	}
	
	ForEach(T{td.in}, F{fn})
	
	if len(res) != len(td.expect) {
		t.Errorf("results are not equal")
	}

	for i, v := range res {
		if v != res[i] {
			t.Errorf("results are not equal")
		}
	}	
}

func TestForAll(t *testing.T) {
	td := []point{{1, 1}, {2, 2}, {3, 3},{4, 4}}

	fn := func(p point) bool {
		return (p.x == p.y)
	}
	
	res, _ := ForAll(T{td}, F{fn})
	
	if !res {
		t.Errorf("all values are not the same")
	}
}

func TestForAny(t *testing.T) {
	td := []point{{1, 1}, {2, 2}, {3, 3},{4, 4}}

	fn := func(p point) bool {
		return (p == point{2, 2})
	}
	
	res, _ := ForAny(T{td}, F{fn})
	
	if !res {
		t.Errorf("could not find value")
	}
}

func TestMap(t *testing.T) {
	in := []point{{1, 1}, {2, 2}, {3, 3},{4, 4}}
	expect := []line{
		{in[0], point{10, 10}},
		{in[1], point{20, 20}},
		{in[2], point{30, 30}},
		{in[3], point{40, 40}},
	}

	fn := func(p point) line {
		return line{p, point{p.x*10, p.y*10}}
	}

	r, _ := Map(T{in}, F{fn})
	res := r.I.([]line)

	if len(res) != 4 {
		t.Errorf("Map failed")
	}

	for i, v := range res {
		if v != expect[i] {
			t.Errorf("results are not equal")
		}
	}
}










