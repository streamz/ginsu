# ginsu
higher order functions for working with slices in golang

Unfortunately golang does not support generics.

The ginsu library provides higher order functions for the following types of slices:

*build in types* such as []int, []string etc...

*user types* such as []T where T is a struct

Unfortunately higher order functions for user types require the use of reflection, so they are not recommended for performance sensitve code.

*supported higer order functions*:

```golang
Filter(t T, f F) (T, error)

ForAll(t T, fn F) (bool, error)

ForEach(t T, fn F)

Map(t T, fn F) (T, error)
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
      []point{{2, 2}, {4, 4}, {6, 6},{8, 8}},
  }

  // Filter
  fn := func(p point) bool {
      return (p.x % 2 == 0)
  }

  res, _ := Filter(T{td.in}, F{fn})
  rs := res.I.([]point)


  // ForEach
  res := make([]point, 0, len(td.expect))
  fn := func(p point) {
      res = append(res, point{p.x*2, p.y*2})
  }
    
  ForEach(T{td.in}, F{fn})


  // ForAll
  fn := func(p point) bool {
      return (p.x == p.y)
  }
    
  res, _ := ForAll(T{td.in}, F{fn})
  
  
  // ForAny
  fn := func(p point) bool {
      return (p == point{2, 2})
  }
    
  res, _ := ForAny(T{td.in}, F{fn})
  
  
  // Map
  fn := func(p point) line {
      return line{p, point{p.x*10, p.y*10}}
  }

  r, _ := Map(T{td.in}, F{fn})
  res := r.I.([]line)
```  
