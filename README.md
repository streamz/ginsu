# ginsu

[![CircleCI](https://circleci.com/gh/streamz/ginsu.svg?style=svg)](https://circleci.com/gh/streamz/ginsu)
[![GoDoc](https://godoc.org/github.com/streamz/ginsu?status.svg)](https://godoc.org/github.com/streamz/ginsu)
[![Go Report Card](https://goreportcard.com/badge/github.com/streamz/ginsu)](https://goreportcard.com/report/github.com/streamz/ginsu)

The ginsu library provides higher order functions mainly for slices.

Unfortunately golang does not support generics :(

Any type of slice can be wrapped in type T

Unfortunately higher order functions require the use of reflection, so they are not recommended for performance sensitve code.

*supported higer order functions*:

```golang
Apply(fn F, args ...T) (func()T, error)

AsyncRepeat(fn F, defered func()) func()

Compare(t0, t1 T, fn F) (bool, error)

Filter(t T, fn F) (T, error)

FilterNot(t T, fn F) (T, error)

ForAll(t T, fn F) (bool, error)

ForAny(t T, fn F) (bool, error)

ForEach(t T, fn F)

Map(t T, fn F) (T, error)

Reduce(initial, t T, fn F) (T, error)
```

*Usage:*

```golang
    // simple hof(s)

    // Apply
    fn, _ := Apply(F{func(a, b int) int {
	return a + b
    }}, T{1}, T{1})
    
    // res == 2
    res := fn().I.(int)

    // AsyncRepeat
    stop := AsyncRepeat(F{func() {
	// do something
    }}, func() {
        // do something defered
    })
    
    // stop doing something
    stop()


    // slice hof(s)

    type point struct {
        x, y int
    }
    
    type line struct {
        x, y point
    }
    
    td := struct {
        in, expect []point
    }{
        []point{{1, 1}, {2, 2}, {3, 3},{4, 4}},
        []point{{1, 1}, {2, 2}, {3, 3},{4, 4}},
    }
    

    // Compare
    ok, err := Compare(T{td.in}, T{td.expect}, F{func(p0, p1 point) bool {
        return p0 == p1
    }})


    // Filter
    res, err := Filter(T{td.in}, F{func(p point) bool {
        return (p.x % 2 == 0)
    }})

    result := res.I.([]point)


    // FilterNot
    res, _ := FilterNot(T{td.in}, F{func(p point) bool {
        return (p.x % 2 == 0)
    }})
    
    result := res.I.([]point)
    

    // ForAll
    ok, err := ForAll(T{td}, F{func(p point) bool {
        return (p.x == p.y)
    }})
  
  
    // ForAny
    ok, err := ForAny(T{td}, F{func(p point) bool {
        return (p == point{2, 2})
    }})
  

    // ForEach
    res := make([]point, len(td.expect))
    idx := 0
    ForEach(T{td.in}, F{func(p point) {
        res[idx] = point{p.x*2, p.y*2}
        idx++
    }})
  
    result := res.I.([]point)


    // Map
    res, _ := Map(T{in}, F{func(p point) line {
        return line{p, point{p.x*10, p.y*10}}
    }})
    
    result := res.I.([]line)


    // Reduce
    res, _ := Reduce(T{point{0,0}}, T{td}, F{func(acc, p point) point {
        return point{acc.x + p.x, acc.y + p.y}
    }})
    
    result := res.I.(point)
```  
