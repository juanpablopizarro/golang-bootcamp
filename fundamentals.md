# Golang fundamentals

### Basic types
Go has almost all the typical type values:
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
This types can be manipulated with different operators. Strings can be added together with `+` as can integers and floats.  
Booleans have all the boolean operators as expected.

This types are used when declaring variables. In Go you can declare variables in three different ways that will be shown in the code below.  
Go also lets you declare constants by preceding the keyword `const` before the name. The type of the constant gets inferred.
[Run it online](https://goplay.space/#cmQu-3Uf58J)
```golang
package main

func main() {
	// Declare with no initialization, go will give the default value
	// which in this case is 0.
	var someInt int
	// Giving it a specific initial value
	var someBool, hi = true, "hi"
	// Inferring the type from the right side of the expression
	hello := "hello world" // hello will be of type string

	// declaring constants. Type gets inferred
	const number = 2;
	const str = "some string";

	fmt.Printf("someInt: %v\n", someInt)
	fmt.Printf("someBool: %v\nhi: %v\n", someBool, hi)
	fmt.Printf("helloWorld: %v\n", hello)
	fmt.Printf("number: %v\n", number)
	fmt.Printf("str: %v\n", str)
}
```
When you declare a variable in Go with no initialization that variable will hold the zero-value of the specified type. For example, in the case of `someInt`, its zero-value will be `0`. For a `bool` the zero-value will be `false`, strings will be `""`.  
If you want to know more about why Go uses the format `<name> <type>` for declaring variables and parameters you can checkout [Go's declaration syntax](https://blog.golang.org/gos-declaration-syntax)

Type casting in Go is done by enclosing the variable or value you want to cast with parenthesis, preceded by the type you want to cast it to.  
For example([GoPlay](https://goplay.space/#rAUMn61amkI)):
```golang
package main

func main() {
	someInt := 2
	castedFloat := float64(someInt)
	fmt.Printf("type of someint: %T\ntype of castedFloat: %T\n", someInt, castedFloat)
	// will result in:
	// type of someint: int
	// type of castedFloat: float64
}
```

### Pointers
Go has pointers. A pointer in Go will hold the memory address of a value. The zero-value of pointers it's `nil`.  
To declare a pointer to a type we use a syntax similar to C:
```golang
var somePointer *int
```
In this case `somePointer` is a pointer to an int. To access the value a pointer points to or the memory address of a given variable when can do this([GoPlay](https://goplay.space/#5_eTXtJ1h43)):
```golang
package main

func main() {
	var someInt *int
	otherInt := 2
	pointer := &otherInt
	fmt.Printf("zero-value of someInt: %v\n", someInt)
	fmt.Printf("value pointer points to: %v\n", *pointer)
}
```

### Structs
To group fields together Go defines `structs`, this structs are quite similar to those defined in the C language. They are defined with `type name struct`, each field it's going to have its own type and they can be initilized in different ways. For example([GoPlay](https://goplay.space/#x4njww69awl)):
```golang
package main

import "fmt"

type someStruct struct {
	intField  int
	boolField bool
}

func main() {
	// this will have the zero-value for all its fields
	s := someStruct{}
	// fields can get initialized by order
	s = someStruct{2, true}
	// we can name the fields we are initializing
	l := someStruct{
		boolField: false,
		intField:  2,
	}

	// We can then access the fields of the struct using the
	// dot notation:
	fmt.Printf("l bool field: %v\n", l.boolField)
	fmt.Printf("s int field field: %v\n", s.intField)
}
```
If you declare a pointer to a struct one could think that in order to access the fields we would have to use the pointer notation. But Go is nice and we don't have to type those 3 extra characters, we can simply do the following([GoPlay](https://goplay.space/#wKurxTznpp5)):
```golang
package main

type someStruct struct {
	intField int
	boolField bool
}

func main() {
	p := &someStruct{2, false}
	fmt.Printf("doing p.intField works: %v\n", p.intField)
	fmt.Printf("but also we can do (*p).boolField: %v\n", (*p).boolField)
}
```
Anonymous structs are something that Go gives us that comes in quite handy when we are testing. Basically we can group together fields without specifying a type for it. For example([GoPlay](https://goplay.space/#H8CneBlRDfh)):
```golang
package main

func main() {
	// see that it has no name that can be used to reference
	// the "type" of this structure
	s := struct {
		x int
		y int
	}{
		x: 2,
		y: 3,
   	}
	fmt.Printf("x: %v y: %v", s.x, s.y)
}
 ```

### Arrays, slices and maps
In Go there is something that is often confusing to newcomers and that is the difference between arrays and slices.  
Arrays are like you would think, a fixed size list of indexed values that all share the same type.  
Slices are something a bit more interesting. The first difference would be that arrays have a fixed size, while slices are dynamically allocated. In practice you will probably always see slices instead of arrays.  
Something that is worth noting is that slices themselves don't hold the array of the data. They are simply `structs` that contain three fields: Len, Cap and Data. Len is going to be the length of the array, Cap is the maximum capacity of the array and Data is a pointer to the backing array. [Here](https://golang.org/pkg/reflect/#SliceHeader) you can see the struct of the slice.  
Slices and arrays are initializaed differently and they let us do different type of operations. Below is a code that will explain all this in more detail([GoPlay](https://goplay.space/#NPh97D1qgEY)):
```golang
package main

import "fmt"

func main() {
	// This is how you would declare an array
	// var someArray [3]int
	// Slices can be declared in many different ways.
	// This slice will have its zero value that is going to be nil, since
	// no array has been alocated yet.
	// var names []string
	// slice of strings with an initial size of 2 and unlimited capacity
	// otherNames := make([]string, 2)
	// slice of strings with an initial size of 2 and maximum capacity of 4
	// capacity := make([]string, 2, 4)
	// We can initialize with values
	numbers := []int{1, 2, 3, 4, 5}
	// Slices let us do operations using the indices
	oneToThree := numbers[0:2]
	fmt.Println(oneToThree)
	// We can omit one of the indices and it will go to the last or the first
	threeToFive := numbers[2:]
	fmt.Println(threeToFive)
	fmt.Println(numbers)

	// Incrementing the size of slice.
	// If we want to append values to already declared slice we can use the append function
	oneToThree = append(oneToThree, 4)
	fmt.Println(oneToThree)
	// append doesn't care about the receiver, so we could declare a new variable with append:
	fiveToSix := append(threeToFive[len(threeToFive)-1:], 6)
	fmt.Println(fiveToSix)
}
```
To learn more about the usage of slices and how they work internally refer to [this blog post](https://blog.golang.org/go-slices-usage-and-internals).

Maps are basically like dictonaries in python or HashMaps in java. They map a key to a given value and let us access those values using the specified keys. The zero value of a map is going to be nil as with slices. Both keys and values can be of any given type, from structs to basic types to interfaces(which we will cover later). Much like slices, we can use `make` to create maps. Examples([GoPlay](https://goplay.space/#xjeZ48dknCc)):
```golang
package main

import "fmt"

type someStruct struct {
	intField  int
	boolField bool
}

func main() {
	// this a nil map since it's the zero-value
	// var m map[string]someStruct
	// declaring and initializing maps
	m := map[string]someStruct{
		"key1": someStruct{2, false},
		"key2": someStruct{3, true},
	}
	// We can reference the fields of the struct by accessing the value of the map
	fmt.Printf("m[\"key1\"].intField=%v\n", m["key1"].intField)
	// We can allocate new key-value pairs
	m["key3"] = someStruct{4, true}
	fmt.Printf("m[\"key3\"].boolField=%v\n", m["key3"].boolField)
}
```

### Function types
In Go, functions are first class citizens. This means that you can use functions as if they were like any other types(ints, string, what have you). In fact, along the standard library and many other packages you will find that lots and lots of function types. For example, in the `net/http` library there is a function type that is commonly used named [http.HandlerFunc](https://godoc.org/net/http#HandlerFunc) with a signature of `func (http.ResponseWriter, *http.Request)`.  
To declare and use function types you can do something like this([GoPlay](https://goplay.space/#rZ4Z48yEykG)):
```golang
package main

import (
	"errors"
	"fmt"
)

func main() {
	var fun func() error
	fun = func() error {
		fmt.Println("hi from fun")
		return errors.New("error from fun")
	}
	someFunc(fun)
}

func someFunc(f func() error) error {
	fmt.Println("hi from someFunc")
	// return the result of calling f
	return f()
}
```

### Methods
Go, unlike objected oriented programming, does not have classes. But it does have methods that you can define on concrete types. This might seem a bit weird, but a method is just a function with a special *receiver*. This receiver can be of **any** type, ints, structs, strings and anything you can think of(they can be defined on function types too, cool inception right?).  
Since the receiver can be of any type, it might hold a certain state. For example if I have a struct with two int fields then those two fields will be accessible from the methods I defined on top of that struct. If you want to modify those states and you want it to be reflected outside of the method you are going to have to define the receiver as a pointer receiver, this means that the receiver will hold the address of the state being manipulated. If we instead use value receivers then the methods will be operating on copies of the original values.  
For convenience, Go will let us call the methods that have pointer receivers with a value, so that we don't have to type 3 extra characters(doing `v.Method()` instead of `(&v).Method()`).  
Stop talking and show me the code([GoPlay](https://goplay.space/#sHBFKyoY5wf)):
```golang
package main

import "fmt"

type mstr string

func (m *mstr) Hi() {
	// print the value of the receiver
	fmt.Printf("Hi %v\n", *m)
}

func main() {
	// we can call hi using a value, go will translate to
	// a memory address since that is what Hi expects as the
	// receiver
	name := mstr("Gopher")
	name.Hi()

	// we can call hi too when we have an actual memory address
	otherName := &name
	otherName.Hi()
}
```

### Interfaces
Interfaces allow us to group together functions with specific signatures specified by us. The value of interfaces can be any value that implements those methods.  
In Go, interfaces are implemented implicitly. This there is no `implemetns` or anything like that. If a given type implements the functions specified by a given interface then that type will implement that interface without us saying anything.  
Interfaces hold a value, which is a value of a specific underlying concrete type, and type which is the concrete type mentioned. One cool thing about interface values in Go is that they can be nil, so essentially you can call a method on a nil value that implements the interface and that will be fine, in other languages this will probably result in null pointer exceptions.  
You might be wondering: well what happens if the interface has no methods? wouldn't any concrete type implement that interface?? ... Well you are right. The empty interface usually refered to as `interface{}` can hold values of any type. Is sometimes comes in handy when you are handling values of an unknown type. For example, in the `json` package of the standard library you have the [Encoder.Encode](https://godoc.org/encoding/json#Encoder.Encode) method that receives the data you want it to encode through an `interface{}` so can basically send everything to that method.  
**BE CAREFUL** when using the empty interface, don't over use it, whenever you are defining the API for your library try to see if you can instead use a user-defined interface with methods that express the meaning of the types the API will handle.  
When defining interfaces Go recommends to reduce the number of methods you put in that interface, since **the bigger the interface, the weaker the abstraction**.  
That was a lot of text, is time for some code([GoPlay](https://goplay.space/#FNh0cTZqrTb)):
```golang
package main

import "fmt"

// Note the format. Typically one method interfaces have
// a the name of the interface be the name of the method
// + er. For example, the interface that has the method
// Write is called Writer, in this case the method name
// is greet so the interface will be called Greeter.
type Greeter interface {
	Greet(name string)
}

type str struct{}

// str type implements the Greeter interface implicitly.
func (str) Greet(name string) {
	fmt.Printf("Hi %v\n", name)
}

func main() {
	// s is of type str so it implements Greeter
	s := str{}
	s.Greet("gopher")

	// whatType receives an empty interface so we can
	// send whatever we want to the function
	whatType(s)
	whatType(3)
	whatType("hi")

	// greetAndBye expects a Greeter, so whatever type
	// that implements the Greet(name string) method
	// can be sent to this function
	greetAndBye(s, "gopher")
}

func whatType(v interface{}) {
	fmt.Printf("type of %v: %T\n", v, v)
}

func greetAndBye(g Greeter, name string) {
	g.Greet(name)
	fmt.Printf("Bye %v\n", name)
}
```

#### Type assertions
There is this thing that you can do with interfaces called type assertion. This provides access to an interface value's underlying concrete value. This basically means that given an interface you can extract the value of a specific type. For example([GoPlay](https://goplay.space/#jOqaZ9T9TlU)):
```golang
package main

import "fmt"

type Greeter interface {
	Greet(name string)
}

type str struct{}

func (str) Greet(name string) {
	fmt.Printf("Hi %v\n", name)
}

func main() {
	// s is of type str so it implements Greeter
	s := str{}
	s.Greet("gopher")
	extract(s)
}

func extract(v interface{}) {
	// if v holds a value of type str then greeter will
	// have that value. If it does not then this will triger
	// a panic
	greeter := v.(str)
	greeter.Greet("gopher1")

	// instead of triggering a panic we can use a second
	// variable that will hold a boolean specifying if v
	// is of type str
	gr, ok := v.(str)
	if !ok {
		fmt.Println("v is not of type str")
		return
	}
	gr.Greet("gopher2")
}
```
You might be wondering, what if we want to do type assertion with multiple types not just one? Well, go has you covered. You can use type switches([GoPlay](https://goplay.space/#pomXwCJ9XzA)):
```golang
package main

import "fmt"

type Greeter interface {
	Greet(name string)
}

type str struct{}

func (str) Greet(name string) {
	fmt.Printf("Hi %v\n", name)
}

func main() {
	// s is of type str so it implements Greeter
	s := str{}
	extract(s)
	extract(2)
	extract("hello")
	extract(2.4)
}

func extract(v interface{}) {
	switch t := v.(type) {
	case str:
		t.Greet("GOPHER")
	case int:
		fmt.Printf("sent an int: %v\n", t)
	case string:
		fmt.Printf("sent a string: %v\n", t)
	default:
		fmt.Printf("unsupported type: %T\n", t)
	}
}
```
