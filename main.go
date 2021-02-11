package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dbzer0/markdownparse/lesparse"
)

func main() {
	p := lesparse.NewParser()

	f, err := os.Open("example.md")
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	chapter := p.Chapter(string(content))

	b, _ := json.MarshalIndent(chapter, "  ", "  ")
	fmt.Println(string(b))
}
