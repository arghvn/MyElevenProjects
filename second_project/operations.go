package main

// Comparison operators in Go :
// This story is dedicated to 6 operators: ==, !=, <, <=, > and >=.

// In order to use aforementioned operators at least one of the operands needs to be assignable to the second one:



import "fmt"

type T struct {
    name string
}
func main() {
    s := struct{ name string }{"foo"}
    t := T{"foo"}
    fmt.Println(s == t)  // true




var a int = 1
var b rune = '1'

fmt.Println(a == b)


// Equality operators :

// Operands need to be comparable to apply == or != operator.

// boolean values are comparable (it’s evaluated without surprises so if both values are either true or false then f.ex. == evaluates to true),
// integers and floats are comparable:
var a int = 1
var b int = 2
var c float32 = 3.3
var d float32 = 4.4
fmt.Println(a == b) // false
fmt.Println(c == d) // false
// a == d would throw an error while compilation (mismatched types int and float32) so it isn’t possible to compare float with int.




// complex numbers are equal if their real and imaginary parts are equal:
var a complex64 = 1 + 1i
var b complex64 = 1 + 2i
var c complex64 = 1 + 2i
fmt.Println(a == b)  // false
fmt.Println(b == c)  // true
// string values are comparable
// pointer values are equal if both are nil or both are pointing at the same variable:
type T struct {
    name string
}

    t1 := T{"foo"}
    t2 := T{"bar"}
    p1 := &t1
    p2 := &t1
    p3 := &t2
    fmt.Println(p1 == p2)   // true
    fmt.Println(p2 == p3)   // false
    fmt.Println(p3 == nil)  // false
}
// Distinct zero-size variables may have the same memory address so we cannot assume anything about equality of pointers to such variables:

a1 := [0]int{}
a2 := [0]int{}
p1 := &a1
p2 := &a2
fmt.Println(p1 == p2)  // might be true or false.

// interface types are comparable.
// As with channels and pointers two interface types values are equal if both are nil or dynamic types are identical with equal dynamic values:
type I interface {
    m()
}
type J interface {
    m()
}
type T struct {
    name string
}
func (T) m() {}
type U struct {
    name string
}
func (U) m() {}
func main() {
    var i1, i2, i3, i4 I
    var j1 J
    i1 = T{"foo"}
    i2 = T{"foo"}
    i3 = T{"bar"}
    i4 = U{"foo"}
    fmt.Println(i1 == i2)  // true
    fmt.Println(i1 == i3)  // false
    fmt.Println(i1 == i4)  // false
    fmt.Println(i1 == j1)  // false
}

// alue i of interface type I and value t of non-interface type T are comparable if T implements I and values of type T are comparable.
// Such values are equal if I’s dynamic type is identical to T and i’s dynamic value is equal to t:
type I interface {
    m()
}
type T struct{}
func (T) m() {}
type S struct{}
func (S) m() {}
func main() {
    t := T{}
    s := S{}
    var i I
    i = T{}
    fmt.Println(t == i)  // true
    fmt.Println(s == i)  // false
}
// For structs to be comparable all their fields needs to be comparable. 
// To be equal it’s enough that all non blank fields are equal:
a := struct {
    name string
    _ int32
}{name: "foo"}
b := struct {
    name string
    _ int32
}{name: "foo"}
fmt.Println(a == b)  // true
// Arrays in Go are homogenous — only values of single type (array element type) can be placed inside.
// For array values to be comparable their element types needs to be comparable. Arrays are equal if corresponding elements are equal.
// This is it.

// There’re 3 types which aren’t comparable — maps, slices and functions. Go compiler won’t allow to
// do that and will f.ex. fail to compile the program comparing maps with an error map can only be compared to nil. 
// Presented error also tells that we can at least compare maps, slices or functions against nil.

// We know so far that interface values are comparable and e.g. maps are not. 
// If dynamic types of interface values are identical but aren’t comparable (like maps) then it causes a runtime error:

type I interface {
    m()
}
type T struct {
    meta map[string]string
}
func (T) m() {}
func main() {
    var i1 I = T{}
    var i2 I = T{}
    fmt.Println(i1 == i2)
}
panic: runtime error: comparing uncomparable type main.T
goroutine 1 [running]:
panic(0x8f060, 0x4201a2030)
    /usr/local/go/src/runtime/panic.go:500 +0x1a1
main.main()
    ...
// Ordering operators
// These operators can be only applied too three types: integer, floats and strings.
// There is nothing unusual or Go-specific here. It’s worth to note though that strings are ordered lexically,
// byte-wise so one byte at a time and without any collation algorithm:

fmt.Println("aaa" < "b") // true
fmt.Println("ł" > "z")   // true
// The outcome
// The result of any comparison operator is untyped boolean constant (true or false). 
// Since it has no type then can be assigned to any boolean variable:

var t T = true
t = 3.3 < 5
fmt.Println(t)
// This snippet outputs true. On the other hand, an attempt to assign a value of type bool:

var t T = true
var b bool = true
t = b
fmt.Println(t)
// yields an error cannot use b (type bool) as type T in assignment. Great introduction to 
// constants (both typed and untyped) is available on official blog.

