/*

Copyright (c) nexB Inc. and others. All rights reserved.
ScanCode is a trademark of nexB Inc.
SPDX-License-Identifier: Apache-2.0
See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
See https://github.com/aboutcode-org/purl-validator-go for support or download.
See https://aboutcode.org for more information about nexB OSS projects.

*/

package main

import (
	"log"
	"os"
	"sort"
	"strings"

	"github.com/blevesearch/vellum"
)

func main() {
	data, _ := os.ReadFile("cmd/data/purls.txt")
	lines := strings.FieldsFunc(string(data), func(r rune) bool {
		return r == '\n'
	})
	sort.Strings(lines)
	output := strings.Join(lines, "\n")
	err := os.WriteFile("cmd/data/purls.txt", []byte(output), 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("purls.fst")
	if err != nil {
		log.Fatal(err)
	}
	builder, err := vellum.New(f, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		err = builder.Insert([]byte(line), 0)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = builder.Close()
	if err != nil {
		log.Fatal(err)
	}
}
