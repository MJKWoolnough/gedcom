# gedcom

[![CI](https://github.com/MJKWoolnough/gedcom/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/gedcom/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/gedcom.svg)](https://pkg.go.dev/vimagination.zapto.org/gedcom)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/gedcom)](https://goreportcard.com/report/vimagination.zapto.org/gedcom)

--
    import "vimagination.zapto.org/gedcom"

Package gedcom implements a parser to read genealogical data in a standard format.

## Highlights

 - Parse GEDCOM into structured objects.
 - Configurable parsing options.

## Usage

```go
package main

import (
	"fmt"
	"io"
	"strings"

	"vimagination.zapto.org/gedcom"
)

func main() {
	const gedcomFile = `0 HEAD
1 SOUR Text
2 VERS 1.0
1 FILE TestFile
1 SUBM Unknown
1 GEDC
2 VERS 5.5
2 FORM LINEAGE-LINKED
1 CHAR UNICODE
0 @SUBM@ SUBM
0 @I1@ INDI
1 SEX F
1 NAME Jo /Wilby/
1 FAMS @F1@
0 @I2@ INDI
1 SEX M
1 NAME Fred /Bloggs/
1 FAMS @F1@
0 @I3@ INDI
1 SEX M
1 NAME Jack /Bloggs/
1 FAMC @F1@
1 FAMS @F2@
0 @I4@ INDI
1 SEX F
1 NAME Linda /Burrows/
1 FAMS @F2@
0 @I5@ INDI
1 SEX F
1 NAME Katie /Bloggs/
1 FAMC @F3@
0 @F1@ FAM
1 HUSB @I2@
1 WIFE @I1@
1 CHIL @I3@
0 @F2@ FAM
1 HUSB @I3@
1 WIFE @I4@
1 CHIL @I5@
0 TRLR
`

	g := gedcom.NewReader(strings.NewReader(gedcomFile))

	for {
		r, err := g.Record()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			fmt.Println("Error:", err)

			break
		}

		switch r := r.(type) {
		case *gedcom.Header:
			fmt.Println("Filename:", r.FileName)
		case *gedcom.Individual:
			fmt.Printf("Person (%s):\n", r.ID)

			for i, name := range r.PersonalNameStructure {
				fmt.Printf("	Name %d: %s\n", i+1, name.NamePersonal)
			}
		case *gedcom.Family:
			fmt.Printf("Family (%s):\n", r.ID)
			fmt.Println("	Father/Husband:", r.Husband)
			fmt.Println("	Mother/Wife:", r.Wife)

			for i, child := range r.Children {
				fmt.Printf("	Child %d: %s\n", i+1, child)
			}
		}
	}

	// Output:
	// Filename: TestFile
	// Person (I1):
	// 	Name 1: Jo /Wilby/
	// Person (I2):
	// 	Name 1: Fred /Bloggs/
	// Person (I3):
	// 	Name 1: Jack /Bloggs/
	// Person (I4):
	// 	Name 1: Linda /Burrows/
	// Person (I5):
	// 	Name 1: Katie /Bloggs/
	// Family (F1):
	// 	Father/Husband: I2
	// 	Mother/Wife: I1
	// 	Child 1: I3
	// Family (F2):
	// 	Father/Husband: I3
	// 	Mother/Wife: I4
	// 	Child 1: I5
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/gedcom
