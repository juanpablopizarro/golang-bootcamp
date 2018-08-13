# Welcome to the Golang bootcamp!
Get ready to read the word Go a lot!

## Introduction
Go is a programming language built by Rob Pike, Robert Griesemer and Ken Thompson. It's a statically typed language with a gargabe collector and a kickass native concurrent-style programming model.

Go is mainly about three things:
* **Go is about composition**. It is *kind* of objected oriented just not in the usual way. There are no classes(but methods can be defined on any type), no subtype inheritance and interfaces are satisfied implicitly(we have structural typing). This results in simple pieces connected by small interfaces.
* **Go is about concurrency**. Go provides [CSP-like](https://en.wikipedia.org/wiki/Communicating_sequential_processes) concurrency primitives, it has lightweight threads called goroutines and typed thread-safe communication and synchronization with channels. This results in comprehensible concurrent code.
* **Go is about gophers!!**
![gopher](https://github.com/juanpablopizarro/golang-bootcamp/blob/master/docs/img/gophers.png)

In fact, we even have our own website where we can [create our own Gophers](http://gopherize.me/). Thanks to Mat Ryer and Ashley McNamara for making this.

## Setting up the environment
First download and install Go, you can follow [this video tutorial](https://www.youtube.com/watch?v=nbkjGixXnlI&index=1&list=PL2ANXDJvvFeJLRIL_8NZl5LGtrRE3sE_E) or you can head over to Go's website and follow the [installation section](https://golang.org/doc/install).  
Now you need to setup the development environment. Go organizes the code in what it's called a `workspace`. [Here](https://golang.org/doc/code.html) is a tutorial from Go's website on how to set it up.

As for what to use you have lots of options. All the mostly used editors have very good support for Go, [VisualStudioCode](https://code.visualstudio.com/) has a very good and highly maintained Go plugin. If are a vim user you are in luck, we have a kickass [vim-go](https://github.com/fatih/vim-go) plugin as well. [Atom](https://atom.io) has very good support too.

If you don't already know git then head over to [this tutorials](https://try.github.io/levels/1/challenges/1) to learn a bit. It will be necessary for this course but also along your career as a developer.

[Here](https://vimeo.com/53221560) is a very good introductory talk given by Andrew Gerrand, one of the core Go team members.

**Week 1**:  Novice to Beginner
## Golang fundamentals
In this section you will learn almost all the features that the Go language gives us. Along with a few of the sections you will find exercises to implement what was just shown.  
Before each code example you will find a link to a website called [GoPlay space](https://goplay.space) were you can run the code and look at the documentation of packages all in the same place.

Ok, head over to [Golang fundamentals](https://github.com/juanpablopizarro/golang-bootcamp/blob/master/fundamentals.md) to get started!

### Assignment

You are going to be implementing an in-memory database with persistence on files. You'll be working on this database through out the week, the excercise is divided into a few sections that will require you to implement the requested functionality before going to the next section.  
Inside your `GOPATH` you will have the `src` directory, this directory may or may not have the `github.com` folder, if it doesn't go ahead and create it. Inside the `github.com` folder create a folder that matches your github username. Finally inside that folder create a new one called `

#### Interface definition
When developing in Go often times you will start by defining an interface that will expose the methods of whatever it is you are implementing. In this case you are implementing a database, so think throughly of what methods you will need to implement, think what operations are tipically done on a database, do you need to open it or close it? Once you have the interface well defined you can start to code.

---

**Week 2**: Intermediate

## **Web Application Development**:
* Important Packages
* HTTP and other important package internals
* Testing Web Application
* Templating and Asset Management
* Middleware (Routing, Authentication, Database)
* File Operations

**Assignments**

2.1.    Create an REST application communicate to DB and would serve content on GET, POST, PUT and DELETE oprations.

2.2 Add Middleware layer for authentication, Routing filtration and Error handling

---

**Week 3**: Expertise
## **Microservice oriented architectural components**:

1. [Service Discovery](###Service-Discovery)
1. [Service Monitoring](###Service-Monitoring)
1. [Communication Protocols](###Communication-Protocols) (gRPC, RPC, TCP)
1. [Error Handling](###Error-Handling)
1. [Authentication](###Authentication)
1. [Caching](###Caching)
1. [Circuit Breaker](###Circuit-Breaker)
1. [Rate Limiting](###Rate-Limiting)
1. [Event Notification](###Event-Notification)
1. [Reverse Proxy](###Reverse-Proxy)
1. [Job Scheduler](###Job-Scheduler)


#### Service Discovery
    [Description and important links]
#### Communication Protocols (gRPC, RPC, TCP)
    [Description and important links]
#### Service Monitoring
    [Description and important links]
#### Error Handling
    [Description and important links]
#### Service Monitoring
    [Description and important links]
#### Authentication
    [Description and important links]
#### Caching
    [Description and important links]
#### Circuit Breaker
    [Description and important links]
#### Rate Limiting
    [Description and important links]
#### Event Notification
    [Description and important links]
#### Reverse Proxy
    [Description and important links]
#### Job Scheduler
    [Description and important links]

**Assignments**: 

3.1: Containerise load balanced application build in assignment 2.2

3.2: -

3.3: -

---

**Week 4** Application Development

## **Build Enterprise ready Cloud Native Application**

## Application 1:
### Distributed Remote Command Executor:-

**Purpose:** Create an distributed client-server application to facilitate system data (per second) retrival for monitoring and Provide UI interface to invoke defined set of Linux commands on target machines for on demand data retrival.

> We will be creating REST API interface, UI Dashboard, clients, server and data pipeline. Every component is independent and replaceable with respective alternate solution. Client can publish monitoring matrices/data to pipeline. Pipeline would retain data untill consumed by consumers. Single or Clustered Servers would subscribe to clients's data on pipeline and guarentees that data is consumed atleast ones by group of servers. It is the responsibility of server to analyse, process, tag consumed data and send it for persistent storage. REST API interface would allow us to retrieve data from database or retrive on-demand data from target machines and visualize on dashboard/ UI.

![alt text][RemoteCommander]

Knowledge Coverage:

-   Dashboard Development & Operations
    -   REST API Development
    -   Middleware for authentication, logging, distributed tracing, application metrices
-   Creating Server Development and Operation
    -   NATS Server
    -   Load Balanced Microservices
-   Creating Clients
    -   Golang Fundamentals and Language Intermediate experience
    -   NATS Client creation involving microservice development knowledge

## Application 2:
### Go-Kit oriented application: (need to decide application context)
    - [Application Overview]
---

[RemoteCommander]: docs/img/RemoteCommander.png "Remote Commander"
