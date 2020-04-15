package ginsu 

import (
	"testing"
)

func TTestForEach(t *testing.T) {
	data := struct { 
		in, expect []int 
	}{
		[]int{1,2,3,4,5,6,7,8,9,0},
		[]int{10,20,30,40,50,60,70,80,90,0},
	}

	is := IntSlice{data.in}
	res := IntSlice{make([]int, 0, len(is.T))}
	is.ForEach(func(i int) {
		res.T = append(res.T, i*10)
	})

	if !res.Eq(data.expect) {
		t.Errorf("IntSlice.ForEach failed")
	}
}

func TTestFilter(t *testing.T) {
	data := struct { 
		in, expect []int 
	}{
		[]int{1,2,3,4,5,6,7,8,9,0},
		[]int{2,4,6,8,0},
	}
	
	is := IntSlice{data.in}
	res := is.Filter(func(i int) bool {
		return i % 2 == 0
	})
	
	if !res.Eq(data.expect) {
		t.Errorf("IntSlice.Filter failed")
	}
}










