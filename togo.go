package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type TogoDoc struct {
	Label   string
	Context string
	Line    uint32
}

func GetContext(Context string, Line uint32) TogoDoc {
	return TogoDoc{strings.Split(Context, ":")[0], strings.Split(Context, ":")[1], Line}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Print("error: please specify a source file to parse.")
		os.Exit(1)
	}

	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("could not open file: %v", err)
	}
	defer file.Close()

	var docs []TogoDoc
	var sline uint32
	source := bufio.NewScanner(file)
	for source.Scan() {
		sline++
		regex := regexp.MustCompile("@[a-zA-Z]*: .*")
		doc := regex.FindString(source.Text())
		if doc != "" {
			docs = append(docs, GetContext(doc, sline))
		}
	}

	if len(docs) >= 1 {
		fmt.Printf("\t%s\n", os.Args[1])
		for _, doc := range docs {
			fmt.Printf("(%d) %s: %s\n", doc.Line, doc.Label, doc.Context)
		}
	}
}
