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

import (
	"bytes"
	"sync"

	readline "gopkg.in/readline.v1"
)

type (
	lineString struct {
		line string
		err  error
	}

	shellReader struct {
		scanner      *readline.Instance
		consumers    chan lineString
		reading      bool
		readingMulti bool
		buf          *bytes.Buffer
		prompt       string
		multiPrompt  string
		showPrompt   bool
		completer    readline.AutoCompleter
		sync.Mutex
	}
)

// rlPrompt returns the proper prompt for readline based on showPrompt and
// prompt members.
func (s *shellReader) rlPrompt() string {
	if s.showPrompt {
		if s.readingMulti {
			return s.multiPrompt
		}
		return s.prompt
	}
	return ""
}

func (s *shellReader) readPassword() string {
	prompt := ""
	if s.buf.Len() > 0 {
		prompt = s.buf.String()
		s.buf.Truncate(0)
	}
	password, _ := s.scanner.ReadPassword(prompt)
	return string(password)
}

func (s *shellReader) setMultiMode(use bool) {
	s.readingMulti = use
}

func (s *shellReader) readLine(consumer chan lineString) {
	s.Lock()
	defer s.Unlock()

	// already reading
	if s.reading {
		return
	}
	s.reading = true
	// start reading

	// detect if print is called to
	// prevent readline lib from clearing line.
	// TODO find better way.
	shellPrompt := s.prompt
	prompt := s.rlPrompt()
	if s.buf.Len() > 0 {
		prompt += s.buf.String()
		s.buf.Truncate(0)
	}

	// use printed statement as prompt
	s.scanner.SetPrompt(prompt)

	line, err := s.scanner.Readline()

	// reset prompt
	s.scanner.SetPrompt(shellPrompt)

	ls := lineString{string(line), err}
	consumer <- ls
	s.reading = false
}
