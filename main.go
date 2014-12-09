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
	b       = flag.String("b", "", "String applied before the list of uuids")
	a       = flag.String("a", "", "String applied after the list of uuids")
	f       = flag.String("f", "%s", "Format string used when adding uuid to list")
	s       = flag.String("s", "\n", "Separator used when joining multiple uuids together")
	t       = flag.Bool("t", false, "Trim whitespace before printing")
	v       = flag.Bool("v", false, "Show version")
	h       = flag.Bool("h", false, "Show help")
	version = "1.3.1"
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
		l = append(l, fmt.Sprintf(ReplaceSpecial(*f), u))
	}
	str := strings.Join(l, ReplaceSpecial(*s))
	str = ReplaceSpecial(*b) + str + ReplaceSpecial(*a) + "\n"
	if *t {
		str = strings.TrimSpace(str)
	}
	fmt.Print(str)
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

func ReplaceSpecial(str string) string {
	r := map[string]string{
		`\n`: "\n",
		`\t`: "\t",
	}
	for k, v := range r {
		str = strings.Replace(str, k, v, -1)
	}
	return str
}

func HelpText() string {
	return `usage: puuid [OPTIONS]

Creates a UUID and prints it. Can create multiple UUID's with -n <num>.

	-v         Show version
	-h         Show help
	-n <num>   Number of uuids to create [1]
	-x <num>   UUID generation version: 1 or 4 [4]
	-b "<str>" String applied to the beginning of the list of uuids
	-a "<str>" String applied to the end of the list of uuids
	-f "<str>" Format string used when adding uuid to list
	-s "<str>" Separator used when joining multiple uuids together
	-t         Trim whitespace before printing
`
}
