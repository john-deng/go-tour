package primary

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"testing"
	"github.com/stretchr/testify/assert"
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
func TestConst(t *testing.T) {
	fmt.Println(Pi)
	assert.Equal(t, 3.14, Pi)
}
/*****************************************************************************/
func TestFoo(t *testing.T) {
	var count = 2017
	fmt.Println("\nHello, foo", count)
	assert.Equal(t, 2017 , count)
}
/*****************************************************************************/
func TestZeroValueVars(t *testing.T) {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("a: %d [%v]\n", a, a)
	fmt.Printf("b: %s [%v]\n", b, b)
	fmt.Printf("c: %f [%v]\n", c, c)
	fmt.Printf("d: %v [%v]\n", d, d)

	assert.Equal(t, 0, a)
	assert.Equal(t, "", b)
	assert.Equal(t, float64(0), c)
	assert.Equal(t,false,d)
}
/*****************************************************************************/
// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

func TestShortVars(t *testing.T) {
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	aaa := int32(10)

	fmt.Printf("aa: %d [%v]\n", aa, aa)
	fmt.Printf("bb: %s [%v]\n", bb, bb)
	fmt.Printf("cc: %v [%v]\n", cc, cc)
	fmt.Printf("dd: %v [%v]\n", dd, dd)

	fmt.Printf("aaa: %d [%v]\n", aaa, aaa)
	assert.Equal(t, 10,aa)
	assert.Equal(t, "hello", bb)
	assert.Equal(t, 3.14159, cc)
	assert.Equal(t, true , dd)
}
/*****************************************************************************/
// Greeting funciton
// export Greeting
func TestGreeting(t *testing.T) {
	fmt.Println("Hello, world")
	fmt.Println("Hello, go")
	fmt.Printf("Home: %s\n", os.Getenv("HOME"))
}
/*****************************************************************************/
func TestConvertType(t *testing.T) {
	var x, y int = 3, 4
	dd := float64(5)
	fmt.Printf("x: %x\n", x)
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("f: %f\n", f)
	var z uint = uint(f)
	fmt.Println(x, y, z)
	assert.Equal(t, 3, x)
	assert.Equal(t, 4 ,y)
	assert.Equal(t, dd , f)
}
/*****************************************************************************/
func TestForLoop(t *testing.T) {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Printf("sum: %sum\n",sum)
	}
	fmt.Println(sum)
	assert.Equal(t, 45 , sum)
}
/*****************************************************************************/
func TestForContinued(t *testing.T) {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	assert.Equal(t, 1024 , sum)
}
/*****************************************************************************/
func TestForWhile(t *testing.T) {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	assert.Equal(t,1024,sum)
}
/*****************************************************************************/
func TestSwitchCase(t *testing.T) {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS;
	os {
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
func TestDeferStatement(t *testing.T) {
	fmt.Println("defer statement: ")

	defer fmt.Println("world")

	fmt.Print("hello ")
}
/*****************************************************************************/
func TestStackingDefers(t *testing.T) {
	fmt.Println("Stacking defers: ")

	fmt.Println("counting")
	i := 0
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	assert.Equal(t, 0, i)
}
/*****************************************************************************/
func add( x int, y int) int {
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
func testPointers(t *testing.T) {
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
func TestArrays(t *testing.T) {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 3, 5, 7, 9, 11}
	fmt.Println(primes)
	assert.Equal(t,[6]int{1, 3, 5, 7, 9, 11}, primes)
}
/*****************************************************************************/
func TestSlices(t *testing.T) {
	primes := [6]int{2, 4, 6, 8, 10, 12}

	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println(primes)

	s[1] = 99
	fmt.Println(s)
	fmt.Println(primes)
}
/*****************************************************************************/
func TestSlicesReference(t *testing.T) {
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
	assert.Equal(t, [4]string{
		"John",
		"XXX",
		"George",
		"Ringo",
	}, names)
}
/*****************************************************************************/
func TestRanges(t *testing.T) {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	assert.Equal(t, []int{1, 2, 4, 8, 16, 32, 64, 128},pow)
}
/*****************************************************************************/
func TestRangeContinued(t *testing.T) {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	assert.Equal(t, []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}, pow)
}
/*****************************************************************************/
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
/*****************************************************************************/
func TestSliceWithMake(t *testing.T) {
	a := make([]int, 5)
	printSlice("a", a)
	assert.Equal(t, []int{0,0,0,0,0},a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	assert.Equal(t, []int{},b)
	c := b[:2]
	printSlice("c", c)
	assert.Equal(t, []int{0,0},c)
	d := c[2:5]
	printSlice("d", d)
	assert.Equal(t, []int{0,0,0},d)


}
/*****************************************************************************/
// Coord struct
type Coord struct {
	Latitude, Longitude float64
}
/*****************************************************************************/
var m map[string]Coord

func TestMap(t *testing.T) {
	m = make(map[string]Coord)
	m["Bell Labs"] = Coord{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	assert.Equal(t, map[string]Coord{"Bell Labs":Coord{
		40.68433, -74.39967,
	}},m)
}
/*****************************************************************************/
func TestMapLiterals(t *testing.T) {
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
func TestMapLiteralsContinued(t *testing.T) {
	var m = map[string]Coord{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}
/*****************************************************************************/
func TestMutatingMaps(t *testing.T) {
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

func TestFunctionClosures(t *testing.T) {
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

func TestInterface(t *testing.T) {
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
func TestMultiMap(t *testing.T)  {

	fmt.Println(">>> Test multimap")

	m := map[string][]int{}
	m["x"] = append(m["x"], 17)
	m["x"] = append(m["x"], 42)
	m["y"] = append(m["y"], -1)
	fmt.Println(m)
}

/*****************************************************************************/
