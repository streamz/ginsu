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
	"reflect"
)

// F container for a function
//
// Used to wrap a func variable or literal:
// 		F{func(p point) bool {
//			return (p.x == p.y)
//		}}
type F struct {
	I interface{}
}

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

type _R struct {
	I reflect.Kind
	O reflect.Kind
}

var rstst = _R{reflect.Struct, reflect.Struct}
var rstbool = _R{reflect.Struct, reflect.Bool}
var struct2 = []reflect.Kind{reflect.Struct, reflect.Struct}

// Compare T Generic
// Compare two type T by applying a binary function:
// 		Compare(T{[]int{1, 2, 3}}, T{[]int{1, 2, 3}}, F{func(i0 int, i1 int) bool {
//				return i0 == i1
//		}})
// returns true|false|runtime error
func Compare(this T, that T, fn F) (bool, error) {
	return this.compare(that, fn)
}

// Filter T Generic
// Filter a slice by items that match the given predicate:
// 		Filter(T{[]int{1, 2, 3}}, F{func(i int) bool {
//				return i % 2 == 0
//		}})
// returns T{[]int} | runtime error
func Filter(t T, fn F) (T, error) {
	return t.filter(fn, false)
}

// FilterNot T Generic
// Filter a slice by items that do NOT match the given predicate:
// 		FilterNot(T{[]int{1, 2, 3}}, F{func(i int) bool {
//				return i % 2 == 0
//		}})
// returns T{[]int} | runtime error
func FilterNot(t T, fn F) (T, error) {
	return t.filter(fn, true)
}

// ForAll T Generic
// Return true if ALL items match the given predicate:
// 		ForAll(T{[]int{1, 2, 3}}, F{func(i int) bool {
//				return i < 10
//		}})
// returns true|false|runtime error
func ForAll(t T, fn F) (bool, error) {
	return t.foranyall(fn, true)
}

// ForAny T Generic
// Return true if ANY item match the given predicate:
// 		ForAny(T{[]int{1, 2, 3}}, F{func(i int) bool {
//				return i == 2
//		}})
// returns true|false|runtime error
func ForAny(t T, fn F) (bool, error) {
	return t.foranyall(fn, false)
}

// ForEach T Generic
// Apply the fn to T:
// 		ForEach(T{[]int{1, 2, 3}}, F{func(i int) {
//			fmt.PrintF("%d\n", i)
//		}})
func ForEach(t T, fn F) {
	t.foreach(fn)
}

// Map T Generic
// Apply the transform fn to T and return a new slice:
// 		Map(T{[]int{1, 2, 3}}, F{Itoa})
// returns T{} | error
func Map(t T, fn F) (T, error) {
	return t.fmap(fn)
}

// Reduce T Generic
// Reduces the elements of T by applying fn as an associative binary operator
//		Reduce(T{0}}, T{[]int{1, 2, 3}}, F{func(acc int, i int) int {
//			return acc + i
//		}})
// returns a single reduced element of wrapped slice in T{} | error
func Reduce(initial T, t T, fn F) (T, error) {
	return t.reduce(initial)(fn)
}

func assertArity(t reflect.Type, k []reflect.Kind) bool {
	if t.NumIn() != len(k) {
		return false
	}

	for i := 0; i < len(k); i++ {
		if t.In(i).Kind() != k[i] {
			return false
		}
	}
	return true
}

func assertIn(t reflect.Type, k reflect.Kind) bool {
	return (t.NumIn() == 1 && t.In(0).Kind() == k)
}

func assertOut(t reflect.Type, k reflect.Kind) bool {
	return (t.NumOut() == 1 && t.Out(0).Kind() == k)
}

func assertfn(t reflect.Type) bool {
	return (t.Kind() == reflect.Func)
}

func (fn F) assert(r _R) bool {
	ft := reflect.ValueOf(fn.I).Type()
	return (assertfn(ft) &&
		assertIn(ft, r.I) &&
		assertOut(ft, r.O))
}

func (t T) compare(other T, fn F) (bool, error) {
	this := reflect.ValueOf(t.I)
	that := reflect.ValueOf(other.I)

	if this.Kind() != reflect.Slice {
		return false, errors.New("Kind is not a slice")
	}

	val := reflect.ValueOf(fn.I)
	ft := val.Type()
	if !assertArity(ft, struct2) || !assertOut(ft, reflect.Bool) {
		return false, errors.New("Invalid params on predicate")
	}

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

	if this.Kind() != reflect.Slice {
		return T{}, errors.New("Kind is not a slice")
	}

	if !fn.assert(rstbool) {
		return T{}, errors.New("Invalid params on predicate")
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

	if this.Kind() != reflect.Slice {
		return false, errors.New("Kind is not a slice")
	}

	if !fn.assert(rstbool) {
		return false, errors.New("Invalid params on predicate")
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

func (t T) foreach(fn F) {
	this := reflect.ValueOf(t.I)
	rt := reflect.ValueOf(fn.I).Type()
	if this.Kind() == reflect.Slice && assertfn(rt) && assertIn(rt, reflect.Struct) {
		f := reflect.ValueOf(fn.I)
		var p [1]reflect.Value
		for i := 0; i < this.Len(); i++ {
			p[0] = this.Index(i)
			f.Call(p[:])
		}
	}
}

func (t T) fmap(fn F) (T, error) {
	this := reflect.ValueOf(t.I)

	if this.Kind() != reflect.Slice {
		return T{}, errors.New("Kind is not a slice")
	}

	if !fn.assert(rstst) {
		return T{}, errors.New("Invalid params on predicate")
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

		if this.Kind() != reflect.Slice {
			return T{}, errors.New("Kind is not a slice")
		}

		it := this.Type().Elem().Kind()
		ik := out.Kind()
		if it != ik {
			return T{}, errors.New("Type mismatch")
		}

		val := reflect.ValueOf(fn.I)
		ft := val.Type()
		if !assertArity(ft, struct2) || !assertOut(ft, ik) {
			return T{}, errors.New("Invalid params on predicate")
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
