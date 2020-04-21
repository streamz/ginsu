package ginsu

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

// hof.go Higher Order Functions for golang

import (
	"errors"
	"fmt"
	"reflect"
)

// T container for a type
//
// Used to wrap any builtin or user type:
// 		T{[]int{1, 2, 3}}
// To extract the value out of the container apply the appropriate type assertion:
// 		t := T{[]int{1, 2, 3}}
// 		i := t.I.([]int)
type T struct {
	I interface{}
}

// F container for a function, type alias for T
//
// Used to wrap a func variable or literal:
// 		F{func(p point) bool {
//			return (p.x == p.y)
//		}}
type F = T


type _K = []reflect.Kind

type _R struct {
	I _K
	O _K
}

var none = T{}
var rstst = _R{_K{reflect.Struct}, _K{reflect.Struct}}
var rstbool = _R{_K{reflect.Struct}, _K{reflect.Bool}}
var rstunit = _R{_K{reflect.Struct}, _K{}}
var structstruct = _R{_K{reflect.Struct}, _K{reflect.Struct}}
var struct2 = _K{reflect.Struct, reflect.Struct}

var invalidfn = "fn is of type %T it is not a function"
var invalidnin = "invalid arity expected %d in params, received %d"
var invalidnout = "invalid arity expected %d params, received %d"
var invalidkin = "invalid input kind param num %d expected %d, received %d"
var invalidkout = "invalid output kind param num %d expected %d, received %d"
var invalidslice = "invalid slice received %T"

// Compare T Generic
// Compare two type T by applying a binary function:
// 		Compare(T{[]int{1, 2, 3}}, T{[]int{1, 2, 3}}, F{func(i0 int, i1 int) bool {
//				return i0 == i1
//		}})
// returns (true || false) || error
func Compare(this T, that T, fn F) (bool, error) {
	return this.compare(that, fn)
}

// Filter T Generic
// Filter a slice by items that match the given predicate:
// 		Filter(T{[]int{1, 2, 3}}, F{func(i int) bool {
//				return i % 2 == 0
//		}})
// returns T{[]int} || error
func Filter(t T, fn F) (T, error) {
	return t.filter(fn, false)
}

// FilterNot T Generic
// Filter a slice by items that do NOT match the given predicate:
// 		FilterNot(T{[]int{1, 2, 3}}, F{func(i int) bool {
//				return i % 2 == 0
//		}})
// returns T{[]int} || error
func FilterNot(t T, fn F) (T, error) {
	return t.filter(fn, true)
}

// ForAll T Generic
// Return true if ALL items match the given predicate:
// 		ForAll(T{[]int{1, 2, 3}}, F{func(i int) bool {
//				return i < 10
//		}})
// returns (true || false) || error
func ForAll(t T, fn F) (bool, error) {
	return t.foranyall(fn, true)
}

// ForAny T Generic
// Return true if ANY item match the given predicate:
// 		ForAny(T{[]int{1, 2, 3}}, F{func(i int) bool {
//				return i == 2
//		}})
// returns (true || false) || error
func ForAny(t T, fn F) (bool, error) {
	return t.foranyall(fn, false)
}

// ForEach T Generic
// Apply the fn to T:
// 		ForEach(T{[]int{1, 2, 3}}, F{func(i int) {
//			fmt.PrintF("%d\n", i)
//		}})
// returns error
func ForEach(t T, fn F) error {
	return t.foreach(fn)
}

// Map T Generic
// Apply the transform fn to A and return a new slice:
// 		Map(A{[]int{1, 2, 3}}, F{Itoa})
// returns B || error
func Map(t T, fn F) (T, error) {
	return t.fmap(fn)
}

// Reduce T Generic
// Reduces the elements of T by applying fn as an associative binary operator
//		Reduce(T{0}}, T{[]int{1, 2, 3}}, F{func(acc int, i int) int {
//			return acc + i
//		}})
// returns a single reduced element of wrapped slice in T{} || error
func Reduce(initial T, t T, fn F) (T, error) {
	return t.reduce(initial)(fn)
}

func assertslice(t reflect.Type) error {
	if t.Kind() != reflect.Slice {
		return fmt.Errorf(invalidslice, t)
	}
	return nil
}

func (fn F) assert(r _R) error {
	t := reflect.ValueOf(fn.I).Type()
	if t.Kind() != reflect.Func {
		return fmt.Errorf(invalidfn, t)
	}

	nin := t.NumIn()
	enin := len(r.I)
	if nin != enin {
		return fmt.Errorf(invalidnin, enin, nin)
	}

	nout := t.NumOut()
	enout := len(r.O)
	if nout != enout {
		return fmt.Errorf(invalidnout, enout, nout)
	}

	for i := 0; i < nin; i++ {
		in := t.In(i).Kind()
		expect := r.I[i]
		if in != expect {
			return fmt.Errorf(invalidkin, i, expect, in)
		}
	}

	for i := 0; i < nout; i++ {
		out := t.Out(i).Kind()
		expect := r.O[i]
		if out != expect {
			return fmt.Errorf(invalidkout, i, expect, out)
		}
	}

	return nil
}

func (t T) compare(other T, fn F) (bool, error) {
	this := reflect.ValueOf(t.I)
	that := reflect.ValueOf(other.I)

	if err := assertslice(this.Type()); err != nil {
		return false, err
	}

	// assert function arity
	if err := fn.assert(_R{struct2, _K{reflect.Bool}}); err != nil {
		return false, err
	}

	// compare len of slices
	if this.Len() != that.Len() {
		return false, nil
	}

	f := reflect.ValueOf(fn.I)
	var p [2]reflect.Value
	for i := 0; i < this.Len(); i++ {
		p[0] = this.Index(i)
		p[1] = that.Index(i)
		if !f.Call(p[:])[0].Bool() {
			return false, nil
		}
	}

	return true, nil
}

func (t T) filter(fn F, not bool) (T, error) {
	this := reflect.ValueOf(t.I)

	if err := assertslice(this.Type()); err != nil {
		return none, err
	}

	if err := fn.assert(rstbool); err != nil {
		return none, err
	}

	f := reflect.ValueOf(fn.I)
	o := reflect.MakeSlice(this.Type(), 0, this.Len())

	var p [1]reflect.Value
	for i := 0; i < this.Len(); i++ {
		p[0] = this.Index(i)
		r := f.Call(p[:])[0].Bool()
		if (r && !not) || (!r && not) {
			o = reflect.Append(o, this.Index(i))
		}
	}

	return T{o.Interface()}, nil
}

func (t T) foranyall(fn F, all bool) (bool, error) {
	this := reflect.ValueOf(t.I)

	if err := assertslice(this.Type()); err != nil {
		return false, err
	}

	if err := fn.assert(rstbool); err != nil {
		return false, err
	}

	f := reflect.ValueOf(fn.I)

	var p [1]reflect.Value
	for i := 0; i < this.Len(); i++ {
		p[0] = this.Index(i)
		call := f.Call(p[:])[0].Bool()
		if all && !call {
			return false, nil
		} else if !all && call {
			return true, nil
		}
	}

	return all, nil
}

func (t T) foreach(fn F) error {
	this := reflect.ValueOf(t.I)

	if err := assertslice(this.Type()); err != nil {
		return err
	}

	if err := fn.assert(rstunit); err != nil {
		return err
	}

	f := reflect.ValueOf(fn.I)
	var p [1]reflect.Value
	for i := 0; i < this.Len(); i++ {
		p[0] = this.Index(i)
		f.Call(p[:])
	}
	
	return nil
}

func (t T) fmap(fn F) (T, error) {
	this := reflect.ValueOf(t.I)

	if err := assertslice(this.Type()); err != nil {
		return none, err
	}

	if err := fn.assert(structstruct); err != nil {
		return none, err
	}

	f := reflect.ValueOf(fn.I)
	v := f.Type().Out(0)
	vs := reflect.SliceOf(v)
	o := reflect.MakeSlice(vs, 0, this.Len())

	var p [1]reflect.Value
	for i := 0; i < this.Len(); i++ {
		p[0] = this.Index(i)
		o = reflect.Append(o, f.Call(p[:])[0])
	}

	return T{o.Interface()}, nil
}

func (t T) reduce(initial T) func(fn F) (T, error) {
	return func(fn F) (T, error) {
		this := reflect.ValueOf(t.I)
		out := reflect.ValueOf(initial.I)

		if err := assertslice(this.Type()); err != nil {
			return none, err
		}
	
		it := this.Type().Elem().Kind()
		ik := out.Kind()
		if it != ik {
			return T{}, errors.New("Type mismatch")
		}

		if err := fn.assert(_R{struct2, _K{it}}); err != nil {
			return none, err
		}

		var p [2]reflect.Value
		f := reflect.ValueOf(fn.I)

		for i := 0; i < this.Len(); i++ {
			p[0] = out
			p[1] = this.Index(i)
			out = f.Call(p[:])[0]
		}

		return T{out.Interface()}, nil
	}
}
