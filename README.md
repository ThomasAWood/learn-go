<div align="center">
  <h1>Practice Algorithms in Go</h1>
  <p>Templates and tests for you to practice writing algorithms in Go! </p>
</div>

## About

A tool for practicing implementing common algorithms in Go. Great for preparing for technical interviews. Inspired to create this while doing the [FrontendMasters Algorithms Course](#acknowledgements).

## Getting Started

### Prerequisites

Please refer to [Installing Go](https://go.dev/doc/install)

### Usage
Write your implementations of the algorithms in the `algorithms` folder. 

#### Testing
To test your implementations run:
```
go test ./tests/{algorithm}_test.go -v
```
Or to test all implementations run:
```
go test ./tests -v
```

#### Reset Implementations 
To clear your implementations so that you can practice again run:
```
go run main.go
```

## Acknowledgements

I built this to follow along to the ThePrimaegens [The Last Algorithms Code You'll Ever Need](https://frontendmasters.com/courses/algorithms/) in Go. This is inspired by the tool he made for the course for Typescript implementation practice.