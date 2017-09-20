package myError

import "fmt"

// MyError
type MyError struct{}

// New
func New() *MyError {
	return &MyError{}
}

// Ok throw exception if !ok
func (m *MyError) Ok(ok bool, message interface{}) {
	if !ok {
		panic(message)
	}
}

// Err throw exception if err!=nil
func (m *MyError) Err(err error, message interface{}) {
	if err != nil {
		fmt.Println(err, message)
		panic(message)
	}
}
