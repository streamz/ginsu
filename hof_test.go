
// Copyright 2020 streamz
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ginsu

import (
	"errors"
	"sync"
	"testing"
)

type point struct {
	x, y int
}

type line struct {
	x, y point
}

func TestAsyncRepeat(t *testing.T) {
	e := errors.New("TestAsyncRepeat Failed")
	count := 0
	invoked := false
	
	var wg0 sync.WaitGroup
	wg0.Add(1)

	defered := func() {
		defer wg0.Done()
		invoked = true
	}

	var wg1 sync.WaitGroup
	wg1.Add(5)
	
	stop := AsyncRepeat(F{func() {
		if count < 5 {
			count++
			wg1.Done()
		} 
	}}, defered)

	wg1.Wait()

	stop()

	wg0.Wait()

	if count != 5 || !invoked {
		t.Errorf(e.Error())
	}
}

func TestApply(t *testing.T) {
	e := errors.New("TestApply Failed")
	fn, err := Apply(F{func(a, b int) int {
		return a + b
	}}, T{10}, T{10})

	if err != nil {
		t.Errorf(err.Error())
	}

	res := fn()

	if res.I.(int) != 20 {
		t.Errorf(e.Error())
	}
}

func TestApplyBadFn(t *testing.T) {
	e := errors.New("TestApply Failed")
	fn, err := Apply(F{func(a, b int) int {
		return a + b
	}}, T{10}, T{"10"})

	if err == nil {
		fn()
		t.Errorf(e.Error())
	}
}

func TestCompare(t *testing.T) {
	td := struct {
		in, expect []point
	}{
		[]point{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
		[]point{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
	}

	ok, err := Compare(T{td.in}, T{td.expect}, F{func(p0, p1 point) bool {
		return p0 == p1
	}})

	if !ok || err != nil {
		t.Errorf(err.Error())
	}
}

func TestCompareNot(t *testing.T) {
	td := struct {
		in, expect []point
	}{
		[]point{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
		[]point{{2, 2}, {4, 4}, {6, 6}, {8, 8}},
	}

	ok, err := Compare(T{td.in}, T{td.expect}, F{func(p0, p1 point) bool {
		return p0 == p1
	}})

	if ok || err != nil {
		t.Errorf(err.Error())
	}
}

func TestCompareNotLength(t *testing.T) {
	td := struct {
		in, expect []point
	}{
		[]point{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
		[]point{{1, 1}, {2, 2}, {3, 3}},
	}

	ok, err := Compare(T{td.in}, T{td.expect}, F{func(p0, p1 point) bool {
		return p0 == p1
	}})

	if ok || err != nil {
		t.Errorf(err.Error())
	}
}

func TestFilter(t *testing.T) {
	td := struct {
		in, expect []point
	}{
		[]point{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
		[]point{{2, 2}, {4, 4}},
	}

	res, _ := Filter(T{td.in}, F{func(p point) bool {
		return (p.x%2 == 0)
	}})

	rs := res.I.([]point)
	ok, err := Compare(T{rs}, T{td.expect}, F{func(p0, p1 point) bool {
		return p0 == p1
	}})

	if !ok || err != nil {
		t.Errorf(err.Error())
	}
}

func TestFilterNot(t *testing.T) {
	td := struct {
		in, expect []point
	}{
		[]point{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
		[]point{{1, 1}, {3, 3}},
	}

	res, _ := FilterNot(T{td.in}, F{func(p point) bool {
		return (p.x%2 == 0)
	}})

	rs := res.I.([]point)
	ok, err := Compare(T{rs}, T{td.expect}, F{func(p0, p1 point) bool {
		return p0 == p1
	}})

	if !ok || err != nil {
		t.Errorf(err.Error())
	}
}

func TestForAll(t *testing.T) {
	td := []point{{1, 1}, {2, 2}, {3, 3}, {4, 4}}

	ok, err := ForAll(T{td}, F{func(p point) bool {
		return (p.x == p.y)
	}})

	if !ok || err != nil {
		t.Errorf(err.Error())
	}
}

func TestFailForAll(t *testing.T) {
	td := []point{{1, 2}, {2, 2}, {3, 3}, {4, 4}}

	ok, err := ForAll(T{td}, F{func(p point) bool {
		return (p.x == p.y)
	}})

	if ok || err != nil {
		t.Errorf(err.Error())
	}
}

func TestForAny(t *testing.T) {
	td := []point{{1, 1}, {2, 2}, {3, 3}, {4, 4}}

	ok, err := ForAny(T{td}, F{func(p point) bool {
		return (p == point{2, 2})
	}})

	if !ok || err != nil {
		t.Errorf(err.Error())
	}
}

func TestForEach(t *testing.T) {
	td := struct {
		in, expect []point
	}{
		[]point{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
		[]point{{2, 2}, {4, 4}, {6, 6}, {8, 8}},
	}

	res := make([]point, len(td.expect))
	idx := 0
	err := ForEach(T{td.in}, F{func(p point) {
		res[idx] = point{p.x * 2, p.y * 2}
		idx++
	}})

	ok, err := Compare(T{res}, T{td.expect}, F{func(p0, p1 point) bool {
		return p0 == p1
	}})

	if !ok || err != nil {
		t.Errorf(err.Error())
	}
}

func TestForEachPtr(t *testing.T) {
	in := make([]*point, 4)
	for i := 0; i < cap(in); i++ {
		xy := i+1
		in[i] = &point{x: xy, y: xy}
	}

	expect := make([]*point, 4)
	for i := 0; i < cap(expect); i++ {
		xy := (i+1)*2
		expect[i] = &point{x: xy, y: xy}
	}

	td := struct {
		in, expect []*point
	}{
		in, 
		expect,
	}

	res := make([]*point, len(td.expect))
	idx := 0
	err := ForEach(T{td.in}, F{func(p *point) {
		res[idx] = &point{p.x * 2, p.y * 2}
		idx++
	}})

	ok, err := Compare(T{res}, T{td.expect}, F{func(p0, p1 *point) bool {
		return *p0 == *p1
	}})

	if !ok || err != nil {
		t.Errorf(err.Error())
	}
}

func TestMap(t *testing.T) {
	e := errors.New("TestMap Failed")
	in := []point{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	expect := []line{
		{in[0], point{10, 10}},
		{in[1], point{20, 20}},
		{in[2], point{30, 30}},
		{in[3], point{40, 40}},
	}

	r, err := Map(T{in}, F{func(p point) line {
		return line{p, point{p.x * 10, p.y * 10}}
	}})

	res := r.I.([]line)

	if len(res) != 4 {
		t.Errorf(e.Error())
	}

	ok, err := Compare(T{res}, T{expect}, F{func(l0, l1 line) bool {
		return l0 == l1
	}})

	if !ok || err != nil {
		t.Errorf(err.Error())
	}
}

func TestMapX(t *testing.T) {
	e := errors.New("TestMap Failed")
	in := []point{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	expect := []int{1, 2, 3, 4}

	r, err := Map(T{in}, F{func(p point) int {
		return p.x
	}})

	res := r.I.([]int)

	if len(res) != 4 {
		t.Errorf(e.Error())
	}

	ok, err := Compare(T{res}, T{expect}, F{func(i0, i1 int) bool {
		return i0 == i1
	}})

	if !ok || err != nil {
		t.Errorf(err.Error())
	}
}

func TestReduce(t *testing.T) {
	td := []point{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	expect := point{10, 10}

	res, err := Reduce(T{point{0, 0}}, T{td}, F{func(acc, p point) point {
		return point{acc.x + p.x, acc.y + p.y}
	}})

	if err != nil {
		t.Errorf(err.Error())
	}

	p := res.I.(point)

	if p != expect {
		t.Errorf("TestReduce Failed")
	}

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAllBadSlice(t *testing.T) {
	td := point{1, 1}

	ok, err := Compare(T{td}, T{td}, F{func(p0, p1 point) bool {
		return p0 == p1
	}})

	if ok {
		t.Errorf(err.Error())
	}

	ok, err = ForAll(T{td}, F{func(p point) bool {
		return (p.x == p.y)
	}})

	if ok {
		t.Errorf(err.Error())
	}

	okT, errT := Filter(T{td}, F{func(p point) bool {
		return (p.x%2 == 0)
	}})

	if okT.I != nil {
		t.Errorf(errT.Error())
	}

	okT, errT = Map(T{td}, F{func(p point) line {
		return line{p, point{p.x * 10, p.y * 10}}
	}})

	if okT.I != nil {
		t.Errorf(errT.Error())
	}

	okT, errT = Reduce(T{point{0, 0}}, T{td}, F{func(acc, p point) point {
		return point{acc.x + p.x, acc.y + p.y}
	}})

	if okT.I != nil {
		t.Errorf(errT.Error())
	}
}

func TestAllBadFn(t *testing.T) {
	e := errors.New("TestAllBadFn Failed")
	td := struct {
		in, expect []point
	}{
		[]point{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
		[]point{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
	}

	ok, err := Compare(T{td.in}, T{td.expect}, F{func(i0, i1 int) bool {
		return true
	}})

	if ok {
		t.Errorf(err.Error())
	}

	ok, err = Compare(T{td.in}, T{td.expect}, F{func(i0 int) bool {
		return true
	}})

	if ok {
		t.Errorf(err.Error())
	}

	ok, err = ForAll(T{td.in}, F{1})

	if ok {
		t.Errorf(err.Error())
	}

	ok, err = ForAll(T{td.in}, F{func(p point) int {
		return 0
	}})

	if ok {
		t.Errorf(err.Error())
	}

	err = ForEach(T{td.in}, F{func(p point) int {
		return 0
	}})

	if err == nil {
		t.Errorf(err.Error())
	}

	err = ForEach(T{0}, F{func(p point) int {
		return 0
	}})

	if err == nil {
		t.Errorf(err.Error())
	}

	_, errT := Filter(T{td.in}, F{func() bool {
		return true
	}})

	if errT == nil {
		t.Errorf(e.Error())
	}

	_, errT = Map(T{td.in}, F{func(b bool) bool {
		return b
	}})

	if errT == nil {
		t.Errorf(e.Error())
	}

	_, errT = Reduce(T{point{0, 0}}, T{td.in}, F{func(acc point, i int) bool {
		return false
	}})

	if errT == nil {
		t.Errorf(e.Error())
	}

	_, errT = Reduce(T{0}, T{td.in}, F{func(acc point, i int) bool {
		return false
	}})

	if errT == nil {
		t.Errorf(e.Error())
	}
}
