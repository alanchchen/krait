// MIT License
//
// Copyright (c) 2016 Alan Chen
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package krait

// errLevel is the severity of an error.
type errLevel int

const (
	warnLevel errLevel = iota + 1
	stopLevel
	exitLevel
	panicLevel
)

var (
	noHandlerErr = WarnErr("No handler registered for input.")
)

// shellError is an interractive shell error
type shellError struct {
	err   string
	level errLevel
}

func (s shellError) Error() string {
	return s.err
}

// WarnErr creates a Warn level error
func WarnErr(err string) error {
	return shellError{
		err:   err,
		level: warnLevel,
	}
}

// StopErr creates a Stop level error. Shell stops if encountered.
func StopErr(err string) error {
	return shellError{
		err:   err,
		level: stopLevel,
	}
}

// ExitErr creates a Exit level error. Program terminates if encountered.
func ExitErr(err string) error {
	return shellError{
		err:   err,
		level: exitLevel,
	}
}

// PanicErr creates a Panic level error. Program panics if encountered.
func PanicErr(err string) error {
	return shellError{
		err:   err,
		level: panicLevel,
	}
}
