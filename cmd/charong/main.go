package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/piotrkowalczuk/pqt/pqtgo"
	"github.com/piotrkowalczuk/pqt/pqtsql"
)

var (
	schema, output string
	acronyms       = map[string]string{
		"id":   "ID",
		"http": "HTTP",
		"ip":   "IP",
		"net":  "NET",
		"irc":  "IRC",
		"io":   "IO",
		"kpi":  "KPI",
		"ssp":  "SSP",
		"uuid": "UUID",
	}
)

func init() {
	flag.StringVar(&schema, "schema", "charon", "")
	flag.StringVar(&output, "output", "schema.pqt", "")
}

func main() {
	file, err := openFile(output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sch := databaseSchema()
	if err := pqtgo.NewGenerator().
		SetAcronyms(acronyms).
		SetVisibility(pqtgo.Public).
		SetPackage("model").
		GenerateTo(sch, file); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(file, "const SQL = `\n")
	if err := pqtsql.NewGenerator().GenerateTo(sch, file); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(file, "`")
}

func openFile(output string) (io.WriteCloser, error) {
	return os.OpenFile(output+".go", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0660)
}
