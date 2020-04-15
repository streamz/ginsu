package ginsu

// Would be nice if Go had generics :(

// BoolSlice wraps []bool
type BoolSlice struct {
	T []bool
}

// Filter BoolSlice
func (A BoolSlice) Filter(fn func(bool) bool) BoolSlice {
	var r []bool
	for _, v := range A.T {
		if fn(v) {
			r = append(r, v)
		}
	}
	return BoolSlice{r}
}

// ForEach BoolSlice
func (A BoolSlice) ForEach(fn func(bool)) {
	for _, v := range A.T {
		fn(v)
	}
}

// ForAll BoolSlice
func (A BoolSlice) ForAll(fn func(bool) bool) bool {
	for _, v := range A.T {
		if !fn(v) {
			return false
		}
	}
	return true
}

// ForAny BoolSlice
func (A BoolSlice) ForAny(fn func(bool) bool) bool {
	for _, v := range A.T {
		if fn(v) {
			return true
		}
	}
	return false
}

// Eq BoolSlice
func (A BoolSlice) Eq(other []bool) bool {
	if len(A.T) == len(other) {
		for i, v := range A.T {
			if other[i] != v {
				return false
			}
		}
	}
	return true
}

// StringSlice wraps []string
type StringSlice struct {
	T []string
}

// Filter StringSlice
func (A StringSlice) Filter(fn func(string) bool) StringSlice {
	var r []string
	for _, v := range A.T {
		if fn(v) {
			r = append(r, v)
		}
	}
	return StringSlice{r}
}

// ForEach StringSlice
func (A StringSlice) ForEach(fn func(string)) {
	for _, v := range A.T {
		fn(v)
	}
}

// ForAll StringSlice
func (A StringSlice) ForAll(fn func(string) bool) bool {
	for _, v := range A.T {
		if !fn(v) {
			return false
		}
	}
	return true
}

// ForAny StringSlice
func (A StringSlice) ForAny(fn func(string) bool) bool {
	for _, v := range A.T {
		if fn(v) {
			return true
		}
	}
	return false
}

// Eq StringSlice
func (A StringSlice) Eq(other []string) bool {
	if len(A.T) == len(other) {
		for i, v := range A.T {
			if other[i] != v {
				return false
			}
		}
	}
	return true
}

// ByteSlice wraps []byte
type ByteSlice struct {
	T []byte
}

// Filter ByteSlice
func (A ByteSlice) Filter(fn func(byte) bool) ByteSlice {
	var r []byte
	for _, v := range A.T {
		if fn(v) {
			r = append(r, v)
		}
	}
	return ByteSlice{r}
}

// ForEach ByteSlice
func (A ByteSlice) ForEach(fn func(byte)) {
	for _, v := range A.T {
		fn(v)
	}
}

// ForAll ByteSlice
func (A ByteSlice) ForAll(fn func(byte) bool) bool {
	for _, v := range A.T {
		if !fn(v) {
			return false
		}
	}
	return true
}

// ForAny ByteSlice
func (A ByteSlice) ForAny(fn func(byte) bool) bool {
	for _, v := range A.T {
		if fn(v) {
			return true
		}
	}
	return false
}

// Eq ByteSlice
func (A ByteSlice) Eq(other []byte) bool {
	if len(A.T) == len(other) {
		for i, v := range A.T {
			if other[i] != v {
				return false
			}
		}
	}
	return true
}

// RuneSlice wraps []rune
type RuneSlice struct {
	T []rune
}

// Filter RuneSlice
func (A RuneSlice) Filter(fn func(rune) bool) RuneSlice {
	var r []rune
	for _, v := range A.T {
		if fn(v) {
			r = append(r, v)
		}
	}
	return RuneSlice{r}
}

// ForEach RuneSlice
func (A RuneSlice) ForEach(fn func(rune)) {
	for _, v := range A.T {
		fn(v)
	}
}

// ForAll RuneSlice
func (A RuneSlice) ForAll(fn func(rune) bool) bool {
	for _, v := range A.T {
		if !fn(v) {
			return false
		}
	}
	return true
}

// ForAny RuneSlice
func (A RuneSlice) ForAny(fn func(rune) bool) bool {
	for _, v := range A.T {
		if fn(v) {
			return true
		}
	}
	return false
}

// Eq RuneSlice
func (A RuneSlice) Eq(other []rune) bool {
	if len(A.T) == len(other) {
		for i, v := range A.T {
			if other[i] != v {
				return false
			}
		}
	}
	return true
}

// IntSlice wraps []int
type IntSlice struct {
	T []int
}

// Filter IntSlice
func (A IntSlice) Filter(fn func(int) bool) IntSlice {
	var r []int
	for _, v := range A.T {
		if fn(v) {
			r = append(r, v)
		}
	}
	return IntSlice{r}
}

// ForEach IntSlice
func (A IntSlice) ForEach(fn func(int)) {
	for _, v := range A.T {
		fn(v)
	}
}

// ForAll IntSlice
func (A IntSlice) ForAll(fn func(int) bool) bool {
	for _, v := range A.T {
		if !fn(v) {
			return false
		}
	}
	return true
}

// ForAny IntSlice
func (A IntSlice) ForAny(fn func(int) bool) bool {
	for _, v := range A.T {
		if fn(v) {
			return true
		}
	}
	return false
}

// Eq IntSlice
func (A IntSlice) Eq(other []int) bool {
	if len(A.T) == len(other) {
		for i, v := range A.T {
			if other[i] != v {
				return false
			}
		}
	}
	return true
}

// Float64Slice wraps []float64
type Float64Slice struct {
	T []float64
}

// Filter Float64Slice
func (A Float64Slice) Filter(fn func(float64) bool) Float64Slice {
	var r []float64
	for _, v := range A.T {
		if fn(v) {
			r = append(r, v)
		}
	}
	return Float64Slice{r}
}

// ForEach Float64Slice
func (A Float64Slice) ForEach(fn func(float64)) {
	for _, v := range A.T {
		fn(v)
	}
}

// ForAll Float64Slice
func (A Float64Slice) ForAll(fn func(float64) bool) bool {
	for _, v := range A.T {
		if !fn(v) {
			return false
		}
	}
	return true
}

// ForAny Float64Slice
func (A Float64Slice) ForAny(fn func(float64) bool) bool {
	for _, v := range A.T {
		if fn(v) {
			return true
		}
	}
	return false
}

// Eq Float64Slice
func (A Float64Slice) Eq(other []float64) bool {
	if len(A.T) == len(other) {
		for i, v := range A.T {
			if other[i] != v {
				return false
			}
		}
	}
	return true
}

// Complex128Slice wraps []complex128
type Complex128Slice struct {
	T []complex128
}

// Filter Complex128Slice
func (A Complex128Slice) Filter(fn func(complex128) bool) Complex128Slice {
	var r []complex128
	for _, v := range A.T {
		if fn(v) {
			r = append(r, v)
		}
	}
	return Complex128Slice{r}
}

// ForEach Complex128Slice
func (A Complex128Slice) ForEach(fn func(complex128)) {
	for _, v := range A.T {
		fn(v)
	}
}

// ForAll Complex128Slice
func (A Complex128Slice) ForAll(fn func(complex128) bool) bool {
	for _, v := range A.T {
		if !fn(v) {
			return false
		}
	}
	return true
}

// ForAny Complex128Slice
func (A Complex128Slice) ForAny(fn func(complex128) bool) bool {
	for _, v := range A.T {
		if fn(v) {
			return true
		}
	}
	return false
}

// Eq Complex128Slice
func (A Complex128Slice) Eq(other []complex128) bool {
	if len(A.T) == len(other) {
		for i, v := range A.T {
			if other[i] != v {
				return false
			}
		}
	}
	return true
}