# ginsu
The ginsu library provides higher order functions for slices.

Unfortunately golang does not support generics :(

Any type of slice can be wrapped in type T

Unfortunately higher order functions require the use of reflection, so they are not recommended for performance sensitve code.

*supported higer order functions*:

```golang
Compare(t0 T, t1 T, fn F) (bool, error)

Filter(t T, fn F) (T, error)

FilterNot(t T, fn F) (T, error)

ForAll(t T, fn F) (bool, error)

ForAny(t T, fn F) (bool, error)

ForEach(t T, fn F)

Map(t T, fn F) (T, error)

Reduce(initial T, t T, fn F) (T, error)
```

*Usage:*

```golang
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
    ok, err := Compare(T{td.in}, T{td.expect}, F{func(p0 point, p1 point) bool {
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
    res := make([]point, 0, len(td.expect))
    
    ForEach(T{td.in}, F{func(p point) {
        res = append(res, point{p.x*2, p.y*2})
    }})
  
    result := res.I.([]point)


    // Map
    res, _ := Map(T{in}, F{func(p point) line {
        return line{p, point{p.x*10, p.y*10}}
    }})
    
    result := res.I.([]line)


    // Reduce
    res, _ := Reduce(T{point{0,0}}, T{td}, F{func(acc point, p point) point {
        return point{acc.x + p.x, acc.y + p.y}
    }})
    
    result := res.I.(point)
```  
