package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"runtime"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

var allErr = make([]error, 0)
func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error{
  return oops2()
}

func oops2() error{
   return oops3()
}



func returnStack() []byte{
  stackSlice := make([]byte, 512)
  s := runtime.Stack(stackSlice, false)

  return stackSlice[0:s]

}

func oops3() error {
  err := MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
    if len(err.Error()) > 0{
     allErr = append(allErr, errors.New(string(returnStack()))) 
  } 
    
    return err
}

func ProcessFile(){

  	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
      cErr := errors.New(fmt.Errorf("Failed Error Path : %w ---- Origin : %v",pathError, string(returnStack())).Error())
      allErr = append(allErr, cErr)
		} else {
			fmt.Println(err)
		}
	}
}

func main() {
	if err := oops(); err != nil {
		//fmt.Println(err)
	}


  go ProcessFile()

  time.Sleep(1000 * time.Millisecond)

  fmt.Println("Length of all errors", len(allErr))

  err1 := fmt.Errorf("Unwrap Error 1 : [%w] ", allErr[0])
  err2 := fmt.Errorf("Two errors are happening: %w", err1) 
  err3 := fmt.Errorf("Unwrap Error 3 - %w", allErr[1])
  

  fmt.Println(err2)
	fmt.Println(errors.Unwrap(err2))
  fmt.Println(errors.Unwrap(err3))
}
