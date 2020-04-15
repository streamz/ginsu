package ginsu

import (
	"errors"
	"reflect"
)

// F typeclass container for a function
type F struct {
	I interface{}
}

// T typeclass container for a type
type T struct {
	I interface{}
}

type _R struct {
	I reflect.Kind
	O reflect.Kind
}

var rstst = _R{reflect.Struct, reflect.Struct}
var rstbool = _R{reflect.Struct, reflect.Bool}

// Filter T Generic
func Filter(t T, f F) (T, error) {
	return t.filter(f)
}

// ForAll T Generic
func ForAll(t T, fn F) (bool, error) {
	return t.foranyall(fn, true)
}

// ForAny T Generic
func ForAny(t T, fn F) (bool, error) {
	return t.foranyall(fn, false)
}

// ForEach T Generic
func ForEach(t T, fn F) {
	t.foreach(fn)
}

// Map T Generic
func Map(t T, fn F) (T, error) {
	return t.fmap(fn)
}

func (fn F) assertArity(t reflect.Type, k []reflect.Kind) bool {
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

func (fn F) assertIn(k reflect.Kind) bool {
	val := reflect.ValueOf(fn.I)
	t := val.Type()
	return (val.Kind() == reflect.Func &&
		t.NumIn() == 1 &&
		t.In(0).Kind() == k)
}

func (fn F) assert(r _R) bool {
	val := reflect.ValueOf(fn.I)
	t := val.Type()
	return (val.Kind() == reflect.Func &&
		t.NumOut() == 1 &&
		t.In(0).Kind() == r.I &&
		t.Out(0).Kind() == r.O)
}

func (t T) filter(fn F) (T, error) {
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
		if f.Call(p[:])[0].Bool() {
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
	if this.Kind() == reflect.Slice && fn.assertIn(reflect.Struct) {
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

