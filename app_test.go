/*
MIT License

Copyright (c) 2018 Muhammad Muzzammil

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"bufio"
	"strings"
	"testing"
)

var message = buildMessage("Hello", "\tnewline", "one more newline with more data", "   one more with less data", "this\tone\thas\ttabs!")

func TestReadLines(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("Hello"))
	s := readLines(r)
	if s[0] != "Hello" {
		t.Fatalf("Couldn't parse bufio.Reader")
	}
}
func TestGetWidth(t *testing.T) {
	setup()
	r := getWidth(message)
	if r < 9 {
		t.Fatalf("Expected to be greater or equal to 9 but got %d", r)
	}
}
func TestRemoveTabs(t *testing.T) {
	r := removeTabs(message)
	if r[4] != "this    one    has    tabs!" {
		t.Fatalf("Tabs were not replaced with spaces")
	}
}
func TestCreateMessage(t *testing.T) {
	setup()
	data := removeTabs(message)
	r := createMessage(data, getWidth(data))
	var e = ` _________________________________
/ Hello                           \
|     newline                     |
| one more newline with more data |
|    one more with less data      |
\ this    one    has    tabs!     /
 ---------------------------------`

	if r != e {
		t.Fatalf("Expected\n%s\nbut got\n%s", e, r)
	}
}
