// Extract wildcard coordinates from a list of concise coordinates
//
// exit codes:
//	1: wrong usage

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Gav struct {
	Group, Artifact, Version, Classifier, Packaging string
}

// Concise converts a coordinate in GAV notation into concise notation.
func (a Gav) Concise() string {
	var sb strings.Builder
	if len(a.Group) > 0 {
		sb.WriteString(a.Group)
	}
	if len(a.Artifact) > 0 || len(a.Version) > 0 || len(a.Classifier) > 0 {
		sb.WriteString(":")
	}
	if len(a.Artifact) > 0 {
		sb.WriteString(a.Artifact)
	}
	if len(a.Version) > 0 || len(a.Classifier) > 0 {
		sb.WriteString(":")
	}
	if len(a.Version) > 0 {
		sb.WriteString(a.Version)
	}
	if len(a.Classifier) > 0 {
		sb.WriteString(":")
		sb.WriteString(a.Classifier)
	}

	if len(a.Packaging) > 0 {
		sb.WriteString("@")
		sb.WriteString(a.Packaging)
	}

	return sb.String()
}

// Includes simulates wildcards on coordinates: empty fields on source will
// match destination, e.g.
// Gav{"g", "a"}.Includes(Gav{"g", "a", "1.0.0"})
// Output: true
func (a Gav) Includes(b Gav) bool {
	if len(a.Group) > 0 && a.Group != b.Group {
		return false
	}
	if len(a.Artifact) > 0 && a.Artifact != b.Artifact {
		return false
	}
	if len(a.Version) > 0 && a.Version != b.Version {
		return false
	}
	if len(a.Classifier) > 0 && a.Classifier != b.Classifier {
		return false
	}
	if len(a.Packaging) > 0 && a.Packaging != b.Packaging {
		return false
	}
	return true
}

// expected format: ^group:artifact:version:classifier@packaging$
func readConciseCoordinates(r io.Reader) []string {
	var ss []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ss = append(ss, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("reading standard input: %v\n", err)
	}
	return ss
}

// Concise converts a Maven coordinate in concise notation into a GAV
func Concise(c string) Gav {
	var gav Gav
	cs := strings.Split(c, "@")
	if len(cs) > 1 {
		gav.Packaging = cs[1]
		c = cs[0]
	}
	cs = strings.Split(c, ":")
	switch len(cs) {
	case 1:
		gav.Group = cs[0]
	case 2:
		gav.Group = cs[0]
		gav.Artifact = cs[1]
	case 3:
		gav.Group = cs[0]
		gav.Artifact = cs[1]
		gav.Version = cs[2]
	case 4:
		gav.Group = cs[0]
		gav.Artifact = cs[1]
		gav.Version = cs[2]
		gav.Classifier = cs[3]
	}
	return gav
}

func main() {
	flag.Usage = func() {
		log.Printf("Usage: %s <concise notation wildcard>...\n",
			os.Args[0])
		flag.PrintDefaults()
	}
	filename := flag.String("universe", "",
		"filename of concise notation list")
	flag.Parse()

	// Read commandline
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Read universe from file or stdin?
	var r io.Reader
	if len(*filename) > 0 {
		f, err := os.Open(*filename)
		if err != nil {
			log.Fatal(err)
		}
		// File implements Reader
		r = f
	} else {
		r = os.Stdin
	}

	universe := readConciseCoordinates(r)
	var gavs []Gav
	for _, u := range universe {
		gavs = append(gavs, Concise(u))
	}
	log.Printf("universe contains %d artifacts\n", len(universe))

	for _, concise := range flag.Args() {
		log.Printf("processing argument %s\n", concise)
		gav := Concise(concise)
		log.Printf("%s -> %+v\n", concise, gav)
		for _, g := range gavs {
			if gav.Includes(g) {
				fmt.Println(g.Concise())
			}
		}
	}
}
