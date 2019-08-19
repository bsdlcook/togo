package main

import (
	"bufio"
	"fmt"
	"github.com/mgutz/ansi"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	Line    string
}

func Usage() {
	// @Cleanup: Use a more elegant way of handling the usage page.
	fmt.Printf("togo %s, sourcecode annotation reviewer.\nUsage: togo [SOURCEFILES]...\n\nCommand-line utility to review annotations from the SOURCEFILES provided.\n\nFor more information, checkout the repo at %s.\n", VERSION, REPO)
}

func GetDoc(Annotation string, Line string) TogoDoc {
	Label := strings.Trim(strings.Fields(Annotation)[0], ":")
	var Context strings.Builder

	for i := 1; i < len(strings.Fields(Annotation)); i++ {
		Context.WriteString(strings.Fields(Annotation)[i] + " ")
	}

	return TogoDoc{Label, Context.String(), Line}
}

func Parse(Sourcefile string) {
	var docs []TogoDoc
	var sline uint64
	file, err := os.OpenFile(Sourcefile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	source := bufio.NewScanner(file)
	for source.Scan() {
		sline++
		// @Idea: Mutli-line comment support.
		regex := regexp.MustCompile(LABEL_PREFIX + "[a-zA-Z]*: .*")
		doc := regex.FindString(source.Text())
		if doc != "" {
			docs = append(docs, GetDoc(doc, strconv.FormatUint(sline, 10)))
		}
	}

	if len(docs) >= 1 {
		fmt.Printf("\t%s\n", ansi.Color(Sourcefile, "white+b"))
		for _, doc := range docs {
			fmt.Printf("(%s) %s: %s\n", ansi.Color("L"+doc.Line, "blue+b"), ansi.Color(doc.Label, "green+b"), ansi.Color(doc.Context, "white"))
		}
	}
}

func main() {
	if len(os.Args) <= 1 {
		Usage()
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		for i := 1; i < len(os.Args); i++ {
			Parse(os.Args[i])
		}
	} else {
		Parse(os.Args[1])
	}
}
