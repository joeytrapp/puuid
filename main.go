package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/satori/go.uuid"
)

var (
	n       = flag.Int("n", 1, "Number of uuids to create")
	x       = flag.Int("x", 4, "Which uuid version to use (1 or 4")
	v       = flag.Bool("v", false, "Show version")
	h       = flag.Bool("h", false, "Show help")
	version = "1.0.1"
)

func main() {
	flag.Parse()
	if *v {
		fmt.Println(version)
		return
	}
	if *h {
		fmt.Println(HelpText())
		return
	}
	l := make([]string, 0)
	if *n < 1 {
		fmt.Println("-n must be a positive integer")
		os.Exit(1)
	}
	for i := 0; i < *n; i++ {
		u, err := NewUUID(*x)
		if err != nil {
			fmt.Println("-x must be 1 or 4")
			os.Exit(1)
		}
		l = append(l, u)
	}
	fmt.Printf("%s", strings.TrimSpace(strings.Join(l, "\n")))
}

func NewUUID(v int) (string, error) {
	var u uuid.UUID
	if v == 1 {
		u = uuid.NewV1()
	} else if v == 4 {
		u = uuid.NewV4()
	} else {
		return "", fmt.Errorf("Invalid UUID version %d", v)
	}
	return fmt.Sprintf("%s", u), nil
}

func HelpText() string {
	return `usage: puuid [OPTIONS]

Creates a UUID and prints it. Can create multiple UUID's with -n <num>.

	-v Show version
	-h Show help
	-n Number of uuids to create [1]
	-x UUID generation version: 1 or 4 [4]
`
}
