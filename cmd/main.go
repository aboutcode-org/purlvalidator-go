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
	"path/filepath"
	"sort"
	"strings"

	"github.com/blevesearch/vellum"
)

func main() {
	f, err := os.Create("purls.fst")
	var purlCount int
	if err != nil {
		log.Fatal(err)
	}

	dirname := "cmd/data/"
	entries, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	builder, err := vellum.New(f, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".txt") {
			fullPath := filepath.Join(dirname, entry.Name())
			purlCount += insert_purls(builder, fullPath)
		}
	}

	err = builder.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("FST generated with %d base PackageURLs", purlCount)
	log.Printf("FST generated at %s", f.Name())
}

func insert_purls(builder *vellum.Builder, file string) int {
	var err error
	// #nosec G304
	data, _ := os.ReadFile(file)
	lines := strings.FieldsFunc(string(data), func(r rune) bool {
		return r == '\n'
	})
	sort.Strings(lines)

	log.Printf("Insert PURLs from %s in FST", file)
	for _, line := range lines {
		err = builder.Insert([]byte(line), 0)
		if err != nil {
			log.Fatal(err)
		}
	}
	return len(lines)
}
