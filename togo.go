package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

const (
	VERSION = "0.2.1"
	// @Implement: Make LABEL_PREFIX mutable from the command-line.
	LABEL_PREFIX = "@"
	REPO         = "https://gitlab.com/nihilism/togo"
)

type TogoDoc struct {
	Label   string
	Context string
	Line    uint32
}

func Usage() {
	// @Cleanup: Use a more elegant way of handling the usage page.
	fmt.Printf("togo %s, sourcecode annotation reviewer.\nUsage: togo [SOURCEFILES]...\n\nCommand-line utility to review annotations from the SOURCEFILES provided.\n\nFor more information, checkout the repo at %s.\n", VERSION, REPO)
}

func GetDoc(Annotation string, Line uint32) TogoDoc {
	Label := strings.Trim(strings.Fields(Annotation)[0], ":")
	var Context strings.Builder

	for i := 1; i < len(strings.Fields(Annotation)); i++ {
		Context.WriteString(strings.Fields(Annotation)[i] + " ")
	}

	return TogoDoc{Label, Context.String(), Line}
}

var WaitGroup = sync.WaitGroup{}

func Parse(Sourcefile string) {
	var docs []TogoDoc
	var sline uint32
	file, err := os.OpenFile(Sourcefile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	source := bufio.NewScanner(file)
	for source.Scan() {
		sline++
		// @Idea: Mutli-line comment support.
		regex := regexp.MustCompile(LABEL_PREFIX + "[a-zA-Z]*: .*")
		doc := regex.FindString(source.Text())
		if doc != "" {
			docs = append(docs, GetDoc(doc, sline))
		}
	}

	if len(docs) >= 1 {
		fmt.Printf("\t%s\n", Sourcefile)
		for _, doc := range docs {
			fmt.Printf("(L%d) %s: %s\n", doc.Line, doc.Label, doc.Context)
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
		Usage()
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
