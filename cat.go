package main

//Cat contains parts to build a cat
type Cat struct {
	Body      string `json:"body"`
	BodyASCII string `json:"ascii"`
	MinLen    int    `json:"minimumlength"`
}

func getCat() *Cat {
	return &Cat{
		Body:      "  (\\__/)||\n  (•ㅅ•)||\n  /    \\っ\n",
		BodyASCII: "  (\\__/) ||\n  (*/\\*) ||\n  /    \\ >",
		MinLen:    9,
	}
}