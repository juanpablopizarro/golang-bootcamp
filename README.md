# Welcome to the Golang bootcamp!
Get ready to read the word Go a lot!

## Introduction
Go is a new programming language built by Rob Pike, Robert Griesemer and Ken Thompson. It's a statically typed language with a gargabe collector and a kickass native concurrent-style programming model.

Go is mainly about three things:
* **Go is about composition**. It is *kind* of objected oriented just not in the usual way. There are no classes(but methods can be defined on any type), no subtype inheritance and interfaces are satisfied implicitly(we have structural typing). This results in simple pieces connected by small interfaces.
* **Go is about concurrency**. Go provides [CSP-like](https://en.wikipedia.org/wiki/Communicating_sequential_processes) concurrency primitives, it has lightweight threads called goroutines and typed thread-safe communication and synchronization with channels. This results in comprehensible concurrent code.
* **Go is about gophers!!**
![gopher](https://github.com/juanpablopizarro/golang-bootcamp/blob/master/docs/img/gophers.png)

In fact, we even have our own website where we can [create our own Gophers](http://gopherize.me/). Thanks to Mat Ryer and Ashley McNamara for making this.

## Part 1  (3 days)
### Setting up the environment
First download and install Go, you can follow [this video tutorial](https://www.youtube.com/watch?v=nbkjGixXnlI&index=1&list=PL2ANXDJvvFeJLRIL_8NZl5LGtrRE3sE_E) or you can head over to Go's website and follow the [installation section](https://golang.org/doc/install).  
Now you need to setup the development environment. Go organizes the code in what it's called a `workspace`. [Here](https://golang.org/doc/code.html) is a tutorial from Go's website on how to set it up.

As for what to use you have lots of options. All the mostly used editors have very good support for Go, [VisualStudioCode](https://code.visualstudio.com/) has a very good and highly maintained Go plugin. If are a vim user you are in luck, we have a kickass [vim-go](https://github.com/fatih/vim-go) plugin as well. [Atom](https://atom.io) has very good support too.

If you don't already know git then head over to [this tutorials](https://try.github.io/levels/1/challenges/1) to learn a bit. It will be necessary for this course but also along your career as a developer.

[Here](https://vimeo.com/53221560) is a very good introductory talk given by Andrew Gerrand, one of the core Go team members.

### Golang fundamentals 
Before each code example you will find a link to a website called [GoPlay space](https://goplay.space) were you can run the code and look at the documentation of packages all in the same place.

Ok, head over to [Golang fundamentals](https://github.com/juanpablopizarro/golang-bootcamp/blob/master/fundamentals.md) to get started!


## Part 2 (2 days)
**INFO**: Along the course you will encounter a few challenges that will ask you to implement a solution that it's "idiomatic go". It might be hard to pin-point what idiomatic go looks like, to help you with this here is a guide from the Go team called [Effective Go](https://golang.org/doc/effective_go.html). This guide highlight most of Go's good practices.

*Implement an in-memory DB.*

The requested in-memory DB is a persistence tool that allow us to store data in memory by responding to basic CRUD operations: 
* Create
* Retrieve
* Update
* Delete

You are going to be implementing this in-memory database with persistence on files too, later. You'll be working on this database through out the week, the excercise is divided into a few sections that will require you to implement the requested functionality before going to the next section.
Inside your GOPATH you will have the src directory, this directory may or may not have the github.com folder, if it doesn't go ahead and create it. Inside the github.com folder create a folder that matches your github username. Finally inside that folder create a new one called `db`

### Interface definition
When developing in Go often times you will start by defining an interface that will expose the methods of whatever it is you are implementing. In this case you are implementing a database, so think throughly of what methods you will need to implement, think what operations are tipically done on a database, do you need to open it or close it? Once you have the interface well defined you can start to code.

Here is a very good talk by Francesc Campoy explaining how interfaces work in Go and how they worked: [Understanding Go interfaces](https://www.youtube.com/watch?v=F4wUrj6pmSI)

### Maps
This is an in-memory database which means that all the data will be kept in some data structure during program execution. Since what you are storing is `key`-`value` pairs you could use golang's builtin `map` structure. If you don't remember how they worked and what operations you could do head back to the [Fundaments section](https://github.com/juanpablopizarro/golang-bootcamp/blob/master/fundamentals.md#arrays-slices-and-maps) to do a little recap.  

Here are a few articles that delve into how maps are implemented in Go and explains just a bit about [Go's memory model](https://golang.org/ref/mem)
* [How the Go runtime implements maps efficiently](https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics)
* [If a map isn't a reference variable then what is it?](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it)
* [There is no pass by reference in Go](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go)

### Testing
Go has a builting library called (testing)[https://godoc.org/testing] this library provides us with primitives to perform tests. Although this library does not give us any type of assert or mocks or anything like that.  
To perform unit tests in Go you will tipically create a file called `ggg_test.go` where `ggg` is the name of the file you are testing. Inside that file you will create all your test functions, the convetion for the signature of a test function is like this: `func TestNameOfFuncBeingTested(t *testing.T)`. Once you have that done you can run `go test` in the command line and see the results of your tests. Here is a simple example of testing a `Sum` function using what it's called table testing:
```go
sum.go
-----
package sum
func Sum(a, b int) int {
    return a + b
}
-----
sum_test.go
-----
package sum
func TestSum(t *testing.T) {
    cases := []struct{
        name string
        a, b int
        expected int
    }{
        {name: "equals zero", a: 1, b: -1, expected: 0},
        {name: "equals 2", a: 1, b: 1, expected: 2},
    }
    for c, _ := range cases {
        t.Run(c.name, func(tt *testing.T){
            if res := Sum(c.a, c.b); res != c.expected {
                tt.Errorf("expected %d but got %d", c.expected, res)
            }

        })
    }
}
```
Try to run the tests of that code locally and see what happens.  

Here are a couple of useful articles on how to perform testing in Go:
* [Writing unit tests in Go](https://blog.alexellis.io/golang-writing-unit-tests/)
* [Tips and tricks for writing unit tests in Go](https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742)
* [Advanced testing techniques in Go](https://segment.com/blog/5-advanced-testing-techniques-in-go/)

## Part 3 (3 days)
### File I/O
What happens when you stop the program you just wrote? Does the data persist on disk?  
Right know we are storing everything in-memory so as soon as the program exits we loose all that data. How would you fix this? A simple approach would be to simply use files and dump all the data into a particular file as soon as the program finishes. You would also have to load all the data that a given file has when you are initializing your DB right?

Go packages that *might* be useful for this:
* [bufio](https://godoc.org/bufio)
* [strings](https://godoc.org/strings)
* [io](https://godoc.org/io) and [io/ioutil](https://godoc.org/io/ioutil)
* [os](https://godoc.org/os)
* [bytes](https://godoc.org/bytes)

Implement a new feature to the in-memory DB that:
* Loads initial data from a local file
* Dumps memory content to a local file when connection is closed

*Note: it's not a need to dump data to file while DB session is open*

### Concurrency
Say you want to use this database in an HTTP server that you have running in production, this server is probably serving more than one request at a time. What happens if both requests try to use the DB to write some data into the same section of the map? You'll have a race condition which will cause your http server to panic. Panics are not good, that's why they are called panic. To avoid this you need to guard the operations that are done on the map using whatever Go provides for this. We talked about something that does this in the [Fundamentals section](https://github.com/juanpablopizarro/golang-bootcamp/blob/master/fundamentals.md#mutexes).  
Once you've implemented that you should write a bit of code that would perform multiple operations on the DB at the same time(look into how the `testing.T.Run()` method works). Remember to run your tests with the `race` flag, this will allow the go runtime

Useful reading and libraries:
* [sync package](https://godoc.org/sync)
* [Synchronizing state in golang](https://kylewbanks.com/blog/tutorial-synchronizing-state-with-mutexes-golang)

## Part 4 - RESTFul
### RESTFul Architectures
If you are not familiar with REST, it would be necessary to get introduced into it:
* [Intro to REST](https://www.youtube.com/watch?v=YCcAE2SCQ6k)
* [REST Tutorial](https://www.restapitutorial.com/)

In order to proceed you should have a good understanding of HTTP, and what RESTFul architectures are. 


### REST API design (2 days)
The exercise in this section it design a RESTFul API which exposes web services for managing a *Shopping Cart*. 
The features needed are:
* Create a new cart
* Adding items to a cart
* List all items of a specific cart 
* Changing the quantity of an existent item in a cart 
* Removing an item from a cart
* Clear a specific cart (remove all items).


**Note:** Available articles to be added to the shopping cart are provided by the following third party web service.

| Method   |      URL    | Desc |
|----------|-------------|---   |
| GET | http://challenge.getsandbox.com/articles | To get all available articles |
| GET | http://challenge.getsandbox.com/articles/{articleId} | To get an specific artible by id. It returns `404` if the _articleId_ is not found |
  
**Important**: In order to consider an API well defined, you should consider a way to describe all the endpoints as much as possible, i.e.:
* URIs with methods allowed 
* headers, if needed
* parameters (path and query params)
* request & response bodys required with format type and expected values (payloads)
* http statuses responses

### REST Impl (3 days)
Once the API is defined, we are ready to implement it. 
There are several way to implement it, you can choose 
* Go Native: https://godoc.org/net/http
* Gorilla mux: https://godoc.org/github.com/gorilla/mux

You should also consider adding some integration testing for he implemented API, using https://godoc.org/net/http/httptest.
Consider also taking a look at:
* [unit testing HTTP servers](https://www.youtube.com/watch?v=hVFEV-ieeew) _(video)_
* [Go Testing Technique: Testing JSON HTTP Requests](https://medium.com/@xoen/go-testing-technique-testing-json-http-requests-76d9ce0e11f)


### DB (2 days)
At this point the RESTFul API is already implemented and working fine, so, the next challenge is to change the persistence engine.
So, you need to change the in-memory DB used with a SQL and NoSQL local DBs. You have to connect you application to a local **MongoDB** and **MySQL**

1. MongoDB: https://godoc.org/github.com/mongodb/mongo-go-driver/mongo
2. MySQL: https://godoc.org/github.com/go-sql-driver/mysql


### Bonus track 
If you have enough time, it would be nice to containerise your application using Docker
