package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"github.com/john-deng/go-tour/reflection"
	"github.com/john-deng/go-tour/exporting"
	"github.com/john-deng/go-tour/concurrency"
	"github.com/john-deng/go-tour/slice"
)
/*****************************************************************************/
// Pi const
const Pi = 3.14

// Vertex struct
type Vertex struct {
	X int
	Y int
}
/*****************************************************************************/
func testConst() {
	fmt.Println(Pi)
}
/*****************************************************************************/
func foo() {
	var count = 2017
	fmt.Println("\nHello, foo", count)
}
/*****************************************************************************/
func testZeroValueVars() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("a: %d [%v]\n", a, a)
	fmt.Printf("b: %s [%v]\n", b, b)
	fmt.Printf("c: %f [%v]\n", c, c)
	fmt.Printf("d: %v [%v]\n", d, d)
}
/*****************************************************************************/
// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

func testShortVars() {
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	aaa := int32(10)

	fmt.Printf("aa: %d [%v]\n", aa, aa)
	fmt.Printf("bb: %s [%v]\n", bb, bb)
	fmt.Printf("cc: %f [%v]\n", cc, cc)
	fmt.Printf("dd: %v [%v]\n", dd, dd)

	fmt.Printf("aaa: %d [%v]\n", aaa, aaa)
}
/*****************************************************************************/
// Greeting funciton
// export Greeting
func Greeting() {
	fmt.Println("Hello, world")
	fmt.Println("Hello, go")
	fmt.Printf("Home: %s\n", os.Getenv("HOME"))
}
/*****************************************************************************/
func testConvertType() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("f: %f\n", f)
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
/*****************************************************************************/
func testForLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
/*****************************************************************************/
func testForContinued() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
/*****************************************************************************/
func testForWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
/*****************************************************************************/
func testSwitchCase() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
/*****************************************************************************/
func testDeferStatement() {
	fmt.Println("defer statement: ")

	defer fmt.Println("world")

	fmt.Print("hello ")
}
/*****************************************************************************/
func testStackingDefers() {
	fmt.Println("Stacking defers: ")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
/*****************************************************************************/
func add(x int, y int) int {
	return x + y
}
/*****************************************************************************/
// Multiple results
// A function can return any number of results.
// The swap function returns two strings.
func swap(x, y string) (string, string) {
	return y, x
}
/*****************************************************************************/
// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
/*****************************************************************************/
func testPointers() {
	fmt.Println("Pointers: ")

	i, j := 12, 36

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)

	p = &j
	*p = *p / 4
	fmt.Println(*p)

}
/*****************************************************************************/
func testArrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 3, 5, 7, 9, 11}
	fmt.Println(primes)
}
/*****************************************************************************/
func testSlices() {
	primes := [6]int{2, 4, 6, 8, 10, 12}

	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println(primes)

	s[1] = 99
	fmt.Println(s)
	fmt.Println(primes)
}
/*****************************************************************************/
func testSlicesReference() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
/*****************************************************************************/
func testRanges() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
/*****************************************************************************/
func testRangeContinued() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
/*****************************************************************************/
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
/*****************************************************************************/
func testSliceWithMake() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}
/*****************************************************************************/
// Coord struct
type Coord struct {
	Latitude, Longitude float64
}
/*****************************************************************************/
var m map[string]Coord

func testMap() {
	m = make(map[string]Coord)
	m["Bell Labs"] = Coord{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
/*****************************************************************************/
func testMapLiterals() {
	var m = map[string]Coord{
		"Bell Labs": Coord{
			40.68433, -74.39967,
		},
		"Google": Coord{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)
}
/*****************************************************************************/
func testMapLiteralsContinued() {
	var m = map[string]Coord{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}
/*****************************************************************************/
func testMutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
/*****************************************************************************/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func testFunctionClosures() {
	fmt.Println("Function closures: ")

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

/*****************************************************************************/
// https://play.golang.org/p/b-dWZ18WvKQ
// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

func testInterface() {
	fmt.Println(">>> Test interface")
	// Values of type user can be used to call methods
	// declared with a value receiver.
	john := user{"John", "john@email.com"}
	john.notify()

	// Values of type user can be used to call methods
	// declared with a pointer receiver.
	john.changeEmail("john@outlook.com")
	john.notify()

	// Values of type user can be used to call methods
	// declared with a value receiver.
	sophie := &user{"Sophie", "sophie@email.com"}
	sophie.notify()

	// Values of type user can be used to call methods
	// declared with a pointer receiver.
	sophie.changeEmail("sophie@outlook.com")
	sophie.notify()

	// Create a slice of users with two users.
	users := []user{
		{"John", "bill@gmail.com"},
		{"Sophie", "sophie@gmail.com"},
	}

	// Iterate over the slice of users
	// calling notify.
	for _, u := range users {
		u.notify()
	}
}
/*****************************************************************************/
func testMultiMap()  {

	fmt.Println(">>> Test multimap")

	m := map[string][]int{}
	m["x"] = append(m["x"], 17)
	m["x"] = append(m["x"], 42)
	m["y"] = append(m["y"], -1)
	fmt.Println(m)
}



/*****************************************************************************/
func main() {
	Greeting()

	fmt.Printf("add: %d\n", add(3, 5))

	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 3
	v.Y = 4
	fmt.Println(v)

	fmt.Println(split(17))

	testZeroValueVars()

	testShortVars()

	foo()

	testConvertType()

	testForLoop()

	testForContinued()

	testForWhile()

	testSwitchCase()

	testDeferStatement()

	testStackingDefers()

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	testConst()

	testPointers()

	testArrays()

	testSlices()

	testSlicesReference()

	testRanges()

	testSliceWithMake()

	testRangeContinued()

	testMap()

	testMapLiterals()

	testMapLiteralsContinued()

	testMutatingMaps()

	testFunctionClosures()

	testInterface()

	testMultiMap()

	exporting.Run()

	reflection.Run()

	concurrency.Run()

	slice.Run()

}
/*****************************************************************************/
