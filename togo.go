package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

type TogoDoc struct {
	Label   string
	Context string
	Line    uint32
}

func GetDoc(Context string, Line uint32) TogoDoc {
	return TogoDoc{strings.Split(Context, ":")[0], strings.Split(Context, ":")[1], Line}
}

var WaitGroup = sync.WaitGroup{}
func Parse(Sourcefile string) {
	var docs []TogoDoc
	var sline uint32
	file, err := os.OpenFile(Sourcefile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("could not open file: %v", err)
	}

	source := bufio.NewScanner(file)
	for source.Scan() {
		sline++
		regex := regexp.MustCompile("@[a-zA-Z]*: .*")
		doc := regex.FindString(source.Text())
		if doc != "" {
			docs = append(docs, GetDoc(doc, sline))
		}
	}

	if len(docs) >= 1 {
		fmt.Printf("\t%s\n", Sourcefile)
		for _, doc := range docs {
			fmt.Printf("(%d) %s: %s\n", doc.Line, doc.Label, doc.Context)
		}
	}

	defer func() {
		file.Close()
		if len(os.Args) > 2 {
			WaitGroup.Done()
		}
	}()
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Print("error: please specify a source file to parse.")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		WaitGroup.Add(len(os.Args) - 1)
		for i := 1; i < len(os.Args); i++ {
			go Parse(os.Args[i])
		}
		WaitGroup.Wait()
	} else {
		Parse(os.Args[1])
	}
}
