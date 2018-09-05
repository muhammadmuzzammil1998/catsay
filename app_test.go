package main

import (
	"testing"
)

var message = buildMessage("Hello", "\tnewline", "one more newline with more data", "   one more with less data", "this\tone\thas\ttabs!")

func TestGetWidth(t *testing.T) {
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
	r := createMessage(message, getWidth(message))
	var e = ` _________________________________
/ Hello                           \
|     newline                     |
| one more newline with more data |
|    one more with less data      |
\ this    one    has    tabs!     /
 ---------------------------------`

	if r != e {
		t.Fatalf("Expected %s but got %s", e, r)
	}
}
