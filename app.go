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
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
)

//Defining a main cat
var cat *Cat

func main() {
	//Declaring flags and parsing them
	var ascii, version bool
	var data []string
	flag.BoolVar(&ascii, "ascii", false, "Use this if you see blocks here \"•ㅅ•っ\"")
	flag.BoolVar(&version, "version", false, "Check version of catsay")
	flag.Parse()

	//Checking flags
	if version {
		data = buildMessage("CatSay", "\tby Muhammad Muzzammil", "", "Version 2.1", "", "http://bit.ly/CATSAY")
	} else if info, _ := os.Stdin.Stat(); info.Mode()&os.ModeCharDevice != 0 {
		data = buildMessage("oh hai kittehs!", "use pipez... i doan knoe how 2 werk otherwize.", "example: echo \"y halo thar, im kat\" | catsay", "", "or try \"catsay -help\"", "btw, i hatz dis sign. lulz")
	} else {
		data = readLines(bufio.NewReader(os.Stdin))
	}

	setup()
	data = removeTabs(data)                        //Replace tabs with spaces
	message := createMessage(data, getWidth(data)) //Create Message Dialog
	fmt.Println(message)                           //Print message dialog
	if ascii {                                     //If ascii is true, show ascii cat
		fmt.Println(cat.BodyASCII)
	} else {
		fmt.Println(cat.Body)
	}
}

//Setup a new cat
func setup() {
	cat = getCat()
}

//Reads and returns data from piped information
func readLines(reader *bufio.Reader) []string {
	var ret []string
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		ret = append(ret, string(line))
	}
	return ret
}

//Gets maximum width of block
func getWidth(message []string) int {
	ret := cat.MinLen
	for _, l := range message {
		width := runewidth.StringWidth(l)
		if width > ret {
			ret = width
		}
	}
	return ret
}

//Creates and returns a new message dialog
func createMessage(lines []string, width int) string {
	lines = formatMessage(lines, width)
	count := len(lines)
	var message []string

	message = append(message, " "+strings.Repeat("_", width+2)) //Top
	if count > 1 {
		message = append(message, fmt.Sprintf(`%s %s %s`, "/", lines[0], "\\"))
		i := 1
		for ; i < count-1; i++ {
			message = append(message, fmt.Sprintf(`%s %s %s`, "|", lines[i], "|"))
		}
		message = append(message, fmt.Sprintf(`%s %s %s`, "\\", lines[i], "/"))
	} else {
		message = append(message, fmt.Sprintf("%s %s %s", "<", lines[0], ">"))
	}

	message = append(message, " "+strings.Repeat("-", width+2)) //Bottom
	return strings.Join(message, "\n")
}

//Replaces tabs with 4 spaces
func removeTabs(message []string) []string {
	var ret []string
	for _, l := range message {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}
	return ret
}

//Formats message by adding whitespaces
func formatMessage(message []string, width int) []string {
	var ret []string
	for _, l := range message {
		w := runewidth.StringWidth(l)
		ret = append(ret, l+strings.Repeat(" ", width-w))
	}
	return ret
}

//Builds a message
func buildMessage(s ...string) []string {
	return s
}
