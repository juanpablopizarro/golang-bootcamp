#   GoLang Bootcamp

<<<<<<< HEAD
**Week 1**: Novice to Beginner
## **Golang Fundamentals**:
* Types Systems and Type Casting
=======
**Before to start**:
## **Environment Setup**:
* [Download golang!](https://www.youtube.com/watch?v=nbkjGixXnlI&index=1&list=PL2ANXDJvvFeJLRIL_8NZl5LGtrRE3sE_E)
* [Setting up the environment](https://www.youtube.com/watch?v=FTDOW8UbKjQ&list=PL2ANXDJvvFeJLRIL_8NZl5LGtrRE3sE_E&index=2)
* Chose an IDE.. [Microsoft Visual Studio Code](https://code.visualstudio.com/) or [Atom](https://atom.io/)
* [Go for impatiens](https://www.youtube.com/watch?v=CF9S4QZuV30&t=1585s)
* [First program!](https://gobyexample.com/hello-world)
* [GIT!](https://try.github.io/levels/1/challenges/1)

**Week 1**:
## **Golang Fundamentals**:
* Native Types and Type Casting
>>>>>>> 764938489fc8054076c96e100aebc7e47456a79f
* Array, Maps, Slice
* Structures, Methods and Interfaces
* Pointers
* Error Handling and Logging
* Concurrancy
* Mutex
* Testing
* Package Management
* Dependency Management

**Assignments**

1.1. Create blabla1 script (Demonstrate implementor to interface).

1.2. Create blabla2 scripts (Demonstrate File Operations)

1.3. Create blabla3 scripts (Demonstrate Concurrancy with Mutex, WaitGroup)

1.4 Create blaBla script (Implement threadsafe map operation with mutex)

1.5. Create blabla scripts (Demonstrate Testing of application)

---

<<<<<<< HEAD
**Week 2**: Intermediate
=======
**Week 2**:
>>>>>>> 764938489fc8054076c96e100aebc7e47456a79f

## **Web Application Development**:
* Important Packages
* HTTP and other important package internals
* Testing Web Application
* Templating and Asset Management
* Middleware (Routing, Authentication, Database)
* File Operations and 

**Assignments**

2.1.    Create an REST application communicate to DB and would serve content on GET, POST, PUT and DELETE oprations.
<<<<<<< HEAD
TIP: 
*   For efficient request routing instead of net/http package you can use [gorilla/mux](https://github.com/gorilla/mux).
*   Database accessing [model](http://www.alexedwards.net/blog/organising-database-access) structure.
*   https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

2.2 Add Middleware layer for authentication, Routing filtration and Error handling


---

**Week 3**: Expertise
=======

2.2 Add Middleware layer for authentication, Routing filtration and Error handling

---

**Week 3**:
>>>>>>> 764938489fc8054076c96e100aebc7e47456a79f
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

**Week 4**

## **Build Enterprise ready application**

## Application 1:
### Distributed Remote Command Executor:-

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
