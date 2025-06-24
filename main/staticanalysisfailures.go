package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// var timeoutInSeconds is of type time.Duration; don't use unit-specific suffix "Seconds" (ST1011)
const timeoutInSeconds = time.Duration(1) * time.Second

type Foo struct {
	Name string
	bar  int
	Done chan struct{}
}

// error should be returned as the last argument (ST1008)
func (f *Foo) DoSomething() (error, *os.File) {
	if f == nil {
		return fmt.Errorf("Foo is nil"), nil
	}

	done := false
	for !done {
		select {
		case <-f.Done:
			// ineffective break statement. Did you mean to break out of the outer loop? (SA4011)
			break
		case <-time.After(timeoutInSeconds):
			done = true
		}
	}

	// "io/ioutil" has been deprecated since Go 1.19: As of Go 1.16, the same functionality is now provided by package [io] or package [os], and those implementations should be preferred in new code. See the specific function documentation for details.  (SA1019)
	file, _ := ioutil.TempFile(".", "tempfile")

	return nil, file
}
