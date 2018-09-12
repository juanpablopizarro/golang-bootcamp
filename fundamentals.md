# Golang fundamentals

### Packages
In Go, programs are structured using packages. The main function of a package should be in `package main`. If you create a project that does not have a package main then it means you are creating a library. Go's convetion says that the name of the package should match the last element of the import path. An import path is basically the path to the project from within your Go workspace.
Now that we are talking about imports, go gives us the statement `import "<path>"` to import a specific package either from the standard library or some package in our go workspace. Example([GoPlay](https://goplay.space/#bB6DC_CV-bF)):
```golang
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Make the zero value useful.")

	fmt.Printf("Square root of 8: %v\n", math.Sqrt(8))
}
```

#### Exporting names
If you've used Java then you probably already had fun with those endless method definitions, I'm talking about you `public static void main`.
Go has no *specific* access modifiers like public or private. In Go, names are either package-level or public. And you do this by making the first letter of the name be a capital letter. When you import a package you can only refer to those names that start with a capital letter. For example([GoPlay](https://goplay.space/#CrFk7n7AWd6)):
```golang
package main

import (
	"fmt"
	"math"
)

func main() {
	// Run this and you will get an error. After that
	// change the pi to be Pi and see what happens.
	fmt.Println(math.pi)

}
```

### Functions
Functions in Go can take any number of arguments and return any number of results. The typical model in Go is to return the results you want plus an error, but we'll cover errors later.
The return values of a function may be named, if so then Go will treat them as local variables to the scope of the function.  Example([GoPlay](https://goplay.space/#xPmMJhUXExA)):
```golang
package main

import "fmt"

// declare x and y as named return results
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```
Something interesting that Go has are variadic functions. This functions can be called with any number of trailing arguments. An example from the standard library may be [fmt.Println](https://golang.org/src/fmt/print.go?s=7595:7644#L253).
Here is an example that will illustrate all this([GoPlay](https://goplay.space/#9xI5pni7y-g)):
```golang
package main

import "fmt"

// Here's a function that will take an arbitrary number
// of `int`s as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// Variadic functions can be called in the usual way
	// with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using
	// `func(slice...)` like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
```

### Flow control statements
#### For
Go gives us only one looping construct, the `for` loop. This means we don't have a while, or a repeat until or anything like that.
For loops have a basic structure similar to the one used at C, except we don't use parenthesis, they will actually be a compilation error:
```golang
for i := 0; i < 2; i++ {
	// code
}
```
The first and last part of the for can be optional, meaning that we can only use the evaluation part and we basically have a while like so:
```golang
for i != 4 {
	// do some stuff
}
```
But you can also omit everything and you'll have an endless loop.

**range** iterates over elements in a variety of data structures. Let’s see how to use range with some of the data structures. ([GoPlay](https://play.golang.org/p/ChWJFN-Zaoy))

```golang

package main

import "fmt"

func main() {

    // Here we use `range` to sum the numbers in a slice.
    // Arrays work like this too.
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // `range` on arrays and slices provides both the
    // index and value for each entry. Above we didn't
    // need the index, so we ignored it with the
    // blank identifier `_`. Sometimes we actually want
    // the indexes though.
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // `range` on map iterates over key/value pairs.
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // `range` can also iterate over just the keys of a map.
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // `range` on strings iterates over Unicode code
    // points. The first value is the starting byte index
    // of the `rune` and the second the `rune` itself.
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
```

#### If
If statements, like for, does not used parenthesis. We can start if statements with a statement to execute before the condition, for example:
```golang
if err := funcReturnsError(); err != nil {
	// very important stuff
}
```
We also have an `else` and can concatenate it with another if like `} else if <condition> {`.

#### Switch
A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression. Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. This means we don't need that ugly break statement at the end of each case like in the mentioned languages. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers. For example([GoPlay](https://goplay.space/#SIsdHJKlgxe))
```golang
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
```
You could also write a switch case with no condition, providing a cleaner way to write long if-else chains.

#### Defer
Defer is something that gets used a lot in Go. This statements defers the execution of the function until the surrounding function returns. The arguments that it receives are evaluated immediately but the function does not get executed until the surrounding function returns.
Deferred function calls get pushed into a call stack. When a function returns the go runtime will pop each of the deferred functions and execute them(it's a LIFO structure). Example([GoPlay](https://goplay.space/#6hzIT6lIo5F)):
```golang
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

### Variables and types
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

### Exercise
Before moving along with the rest of the contents head over to the [loops and functions exercise](https://tour.golang.org/flowcontrol/8) and implement what it is requested there. Try to run your code locally so that you check that the environment you setup previously is properly working.

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

__NOTE:__ notice that in the above example the struct attributes are private because the first letter is in lower case.

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

### Composition over inheritance
Like we stated previously in the introduction. Go is all about composition, in fact, there is no inheritance. If you want to share behaviour and data of something you compose something with that type, no need of subclassing and specifying a complex tree of class hierarchy and inheritance.
Coming from an OOP world this might be weird at first. But composition is actually a well understood concept of OOP and in Go you will use it quite a lot. To learn a bit more about how this works you can read [this section](https://golang.org/doc/effective_go.html#embedding) of Effective Go to see some real-live examples of how you can benefit this.
Lets take a look at a simple example of how you would use this([GoPlay](https://goplay.space/#AnaEldEzcBl)):
```go
package main

import "fmt"

// User holds information of a given user.
type User struct {
	ID             int
	Name, Location string
}

// Here you can see that player embeds the User type.
// This way we are saying that a Player holds is composed
// of a User and a GameID.
type Player struct {
	User
	GameID int
}

func main() {
	p := Player{}
	p.ID = 42
	p.Name = "Globant"
	p.Location = "La Plata"
	p.GameID = 90404
	fmt.Printf("%+v", p)
	// This will print the following:
	// {User:{ID:42 Name:Globant Location:La Plata} GameID:90404}
	// You can see that the User type is embedded in the Player
	// structure.
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
**Excercise:** Go to [Go Tour slices excercise](https://tour.golang.org/moretypes/18) and implement what it's requested there.

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
**Exercise:** go to [Go tour maps excercise](https://tour.golang.org/moretypes/23) and implement what it's requested.

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
**Exercise:** go to [Go tour closure exercise](https://tour.golang.org/moretypes/26) and implement what it's requested.

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
In Go, interfaces are implemented implicitly. This there is no `implements` or anything like that. If a given type implements the functions specified by a given interface then that type will implement that interface without us saying anything.
Interfaces hold a value and a type, the value is the value of the specific underlying concrete type mentioned. One cool thing about interface values in Go is that they can be nil, so essentially you can call a method on a nil value that implements the interface and that will be fine, in other languages this will probably result in null pointer exceptions.
You might be wondering: well what happens if the interface has no methods? wouldn't any concrete type implement that interface?? ... Well you are right. The empty interface usually refered to as `interface{}` can hold values of any type. This sometimes comes in handy when you are handling values of an unknown type. For example, in the `json` package of the standard library you have the [Encoder.Encode](https://godoc.org/encoding/json#Encoder.Encode) method that receives the data you want it to encode through an `interface{}` so you can basically send anything you want to that method.
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
You are going to implement a solution to two different exercises that will show you interfaces from the standard library that are commonly used.
**Exercise:** go to [Go tour stringers exercise](https://tour.golang.org/methods/18) and implement what it's requested.
**Exercise:** go to [Go tour readers exercise](https://tour.golang.org/methods/22) and implement what it's requested.

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

### Errors
Errors are something of great discussion in Golang, mainly because it is very different to what we typically use in languages like Java or Ruby. If you want to express an error in those languages you would create your own exception and give it a meaningful name and maybe a description. You would then need aditional control structures like `try..catch ` to handle those errors.
In Go we instead treat errors as simple values. For this we have the `error` type that is nothing more than an [interface with a single method](https://godoc.org/builtin#error):
```golang
type error interface {
	Error() string
}
```
In go, functions usually return errors so if something went wrong you simply check whether the error is nil or no. This means that whenever you build your own functions you should return whatever you want to return and an error, if nothing went wrong simply return nil, otherwise return an error that expresses the problem. You will see a lot of code that looks like this:
```golang
// Handling errors
err := callFunc()
if err != nil {
	// handle the error in here
}

// Returning errors.
// Return the data you want + an error
func importantFunc() (string, error) {
	s, err := importantOperation()
	if err != nil {
		return "", err
	}
	return s, nil
}
```
**Exercise:** go to [Go tour error excercise](https://tour.golang.org/methods/20) and implement what it's requested.

### Concurrency
#### Goroutines
If you heard of Go then you probably also heard that Go has a kick-ass native support for concurrency. It is extremely simple and very efficient.
For this Go uses the concept of *goroutines*, a goroutine is lightweight thread managed by the Go runtime. This means that in a single OS Thread we can have multiple goroutines, thousands if you want.
This goroutines run in the same address space so you will have to be careful when you are accessing data from multiple goroutines.
This shows a bit how you can use goroutines([GoPlay](https://goplay.space/#cRTTaMlqWj0)):
```golang
package main

import (
	"fmt"
	"time"
)

func importantFeature() {
	fmt.Println("doing important work")
}

func main() {
	importantFeature()

	// you can spawn new goroutines with a simple
	// go funcName
	go importantFeature() // this now created a brand new goroutine

	s := "inside goroutine"
	// you can also spawn goroutines using an
	// anonymous functions
	go func(someString string) {
		fmt.Println("printing from the goroutine")
	}(s)

	for i := 0; i < 5; i++ {
		// we can spawn as many as we like. Try changing
		// i<5 to something like i<1000
		go func(i int) {
			time.Sleep(5 * time.Millisecond)
			fmt.Printf("i=%v\n", i)
		}(i)
	}
	// wait a bit so that we don't exit immediately
	// we need to wait so that at least some of the
	// goroutines get executed. Go doesn't automatically
	// wait for all goroutines, if we don't do synchronization
	// then we will exit and goroutines won't finish executing.
	time.Sleep(25 * time.Millisecond)
}
```

#### Channels
For communicating and synchronizing between different goroutines Go introduced the concept of *channels*. Channels are pipes that connect concurrent goroutines. You can send values from one goroutine and receive them from another. For example([GoPlay](https://goplay.space/#vteRsUOcf1w)):
```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		// The information flows in the direction of the arrow
		// In this case we are sending data into the messages
		// channel
		messages <- "hi from the goroutine"
	}()

	// now we are receiving the information that is sent to
	// the channel.
	// Note that this will block for a second until the
	// channel has something we can receive.
	m := <-messages
	fmt.Println(m)
}
```
In the comments we specified that the instruction `m := <-messages` will block until the channel has something we can read, so until the goroutine sends something to the messages channel that instruction will block the program there.
Receiving from a channel will *always* block, but sending to the channel may or may not block depending on how we initialize the channel. In the previous example we declared that channel as a synchronous channel, how did we do that? By not specifying a size when we created with `make`(line 9).
If we want to make an asynchronous channel, meaning that it will not block when we try to send something, then we need to declare a **buffered channel**. To create them you have to specify a size when you initialize it with `make`. For example([GoPlay](https://goplay.space/#5V-qhnN3zAs)):
```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 2)

	// we send two messages that will make
	// the channel full.
	messages <- "hi"
	messages <- "bye"

	// spawn a new goroutine that will read
	// from the channel after three seconds
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println(<-messages)
	}()

	// if we try to add a message to the channel
	// before one or all of the messages are consumed
	// then the operation will block since there
	// is no more space for that message to fit.
	// So after the previous go routine reads, we
	// will be able to send "hi again"
	messages <- "hi again"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
```
We can do something similar to synchronize execution across goroutines. For example, we could use a blocking receive to wait until the goroutine is done executing([GoPlay](https://goplay.space/#-BGwwrmU1ve)):
```golang
package main

import "fmt"
import "time"

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}

func main() {

	// Start a worker goroutine, giving it the channel to
	// notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the
	// worker on the channel.
	<-done
}
```
In the previous example we showed that you can send channels as parameters. What you can also do is specify the *direction* of the channel, this will limit the operations we can do. So essentialy we can say "this is a send only channel", "this a receive only channel" or "you can whatever you want, be free". Lets rewrite the previous example, note that inside the worker function we are only sending to the channel, so we can instead do the following:
```golang
func worker(done chan<- bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}
```
If we wanted to send a receive only channel to a function we could do this:
```golang
// note the direction of the arrow
func someFunc(done <-chan bool) {
	// do work
}
```

Closing a channel indicates that no more values will be sent on it. This can sometimes be useful to indicate a completion to the ones that are receiving from it. Reading from a closed channel does not return error, it returns the zero-value of the type of the channel. So if you have declared a channel of strings and you read from it after closing it you will get all empty strings. If you are using structs then the zero-value will be nil.
This allows us to use the `for range` structure to iterate over the values of the channel until the channel is closed. The next example will cover this use cases([GoPlay](https://goplay.space/#WGVtmD3Ck5C)):
```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 2)
	done := make(chan bool, 1)

	go func() {
		for {
			// this double assignment allows us to check
			// if the channel was closed or not
			m, ok := <-messages
			if !ok {
				// if ok is false it means that we cannot
				// receive from the channel so it was closed
				fmt.Println("channel closed!")
				// close the channel so that it sends the nil
				// value which will make the last sentence
				// of this program ends and the program exits
				close(done)
				return
			}
			fmt.Printf("received: %v\n", m)
		}
	}()
	// wait before closing the channel
	time.Sleep(2 * time.Second)
	messages <- "hi"
	messages <- "bye"
	close(messages)
	<-done
}
```

Go provides us with a statemente called `select`. This statemente allows a goroutine to wait on multiple communications operations and perform an action whenever we receive something from any of the specified channels. Example([GoPlay](https://goplay.space/#xP74KERfHXl)):
```golang
package main

import "time"
import "fmt"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
```
**Exercise:** go to [Go tour binary tree exercise](https://tour.golang.org/concurrency/7) and implement what it's requested

#### Mutexes
You've seen that channels are great for communicating between different goroutines, we can even use it for synchronization. But what if we want to guard a variable so that only one goroutine can access it at the same time. It would be quite cumbersome to implement a solution using only channels. Well, go is nice again and provides us with a library called `sync` that has a bunch of utilities that are useful for doing mutual exclusion using a generally refered to as `mutex`. This example was takend from the GoTour([GoPlay](https://goplay.space/#7L8oJihNE1G)):
```golang
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```
**Exercise:** go to [Go tour web crawler exercise](https://tour.golang.org/concurrency/10) and implement what it's requested.

### Recap
You just read a bunch of text so lets do a set of questions, sort of like a recap, to test your knowledge. This list has not only questions but also exercises that you need to solve. We recommend you to put all this in a repository so that the instructors can see how are doing, do a bit of code review and help you with whatever difficulty you may be having.

* How would you write a `while` statement in Go?
* What does the keyword `defer` do?
* Does Go support pointers? How do arguments get passed around(by value or by reference)?
* Are arrays in Go fixed length? How about slices?
* Say you have a map: `map[string]int`, how would do a lookup and check to see if the map holds the value of the key you were looking for?
* How does Go structure programs? What is the difference between a library and a program that executes?
* How do make a function or a type public? And how do you make it private?
* You are going to be building a simple calculator with 4 basic operations(add, subtract, multiply and divide). First build a library that provides those 4 methods. After that implement a program that reads from the command line the operation to be done and prints the result(by calling the library you implemented previously). The operation should be read however you'd like, but for simplicity sake limit yourself to 2 operands and 1 operation character. Something like `./program 1 + 2`.
* Suppose you are building a web server that needs a DB that can do a set of simple operations. You know that the requirements of what DB to use will change. You also now that it will be easier for testing purposes to not have to setup something like MySQL. How would you solve this problem using the feature that Go provide?
* How would you build a simple function that can receive *any* type of argument and prints the if that argument is of a primitive type. Limit to just `int`, `string`, `float` and `bool`.
* How are errors defined in Go?
* Ok, you know how errors are defined in Go now. Time to build a simple `errors` package that allows you to build errors that specify what kind of error is it, limit yourself to 3 kinds: `Internal`, `ThirdParty` and `Other`. Then provide a function in that package for users to check if the error they have is of the kind they care about. **NOTE**: remember not to break with the way errors are defined in Go, take advantage of that.
* What do you use to make two functions concurrent?
* How would you synchronize two concurrent functions?
* Write a program with three functions. One will send stuff(whatever you'd like) over a channel every one second and one will receive it and print it. The third function will tell the other two functions to stop and return(it could be the main func) after 5 seconds. **NOTE**: the program can not end until the sender and receiver have returned.