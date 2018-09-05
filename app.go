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
	"unicode/utf8"
)

func main() {
	var ascii, version bool
	var data []string
	flag.BoolVar(&ascii, "ascii", false, "Use this if you see blocks here \"•ㅅ•っ\"")
	flag.BoolVar(&version, "version", false, "Check version of catsay")
	flag.Parse()
	if version {
		data = buildMessage("CatSay", "\tby Muhammad Muzzammil", "", "Version 2.0", "", "http://bit.ly/CATSAY")
	} else if info, _ := os.Stdin.Stat(); info.Mode()&os.ModeCharDevice != 0 {
		data = buildMessage("oh hai kittehs!", "use pipez... i doan knoe how 2 werk otherwize.", "example: echo \"y halo thar, im kat\" | catsay", "", "or try \"catsay -help\"", "btw, i hatz dis sign. lulz")
	} else {
		data = readLines(bufio.NewReader(os.Stdin))
	}
	var message = createMessage(data, getWidth(data))
	fmt.Println(message)
	fmt.Println(getCat(ascii))
}

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

func getWidth(message []string) int {
	ret := 9
	for _, l := range message {
		len := utf8.RuneCountInString(l)
		if len > ret {
			ret = len
		}
	}
	return ret
}

func createMessage(lines []string, width int) string {
	lines = formatMessage(removeTabs(lines), width)
	count := len(lines)
	var message []string

	message = append(message, " "+strings.Repeat("_", width+2))
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

	message = append(message, " "+strings.Repeat("-", width+2))
	return strings.Join(message, "\n")
}

func removeTabs(message []string) []string {
	var ret []string
	for _, l := range message {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}
	return ret
}

func formatMessage(message []string, width int) []string {
	var ret []string
	for _, l := range message {
		ret = append(ret, l+strings.Repeat(" ", width-utf8.RuneCountInString(l)))
	}
	return ret
}

func getCat(ascii bool) string {
	if ascii {
		return "  (\\__/) ||\n  (*/\\*) ||\n  /    \\ >"
	}
	return "  (\\__/)||\n  (•ㅅ•)||\n  /    \\っ\n"
}

func buildMessage(s ...string) []string {
	return s
}
