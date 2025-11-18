/*

Copyright (c) nexB Inc. and others. All rights reserved.
ScanCode is a trademark of nexB Inc.
SPDX-License-Identifier: Apache-2.0
See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
See https://github.com/aboutcode-org/purl-validator-go for support or download.
See https://aboutcode.org for more information about nexB OSS projects.

*/

package purlvalidator

import (
	"log"
	"os"
	"testing"

	"github.com/blevesearch/vellum"
)

var testValidator *vellum.FST

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	fstData, err := os.ReadFile("testdata/testpurls.fst")
	if err != nil {
		panic(err)
	}

	testValidator, err = vellum.Load(fstData)
	if err != nil {
		log.Fatal(err)
	}
}

// `testpurls.fst` only contains following purls:
// - pkg:nuget/FluentUtils.EnumExtensions
// - pkg:nuget/FluentUtils.FromCompositeAttribute
// - pkg:nuget/FluentUtils.MediatR.Pagination
// - pkg:nuget/FluentUtils.MediatR.Pagination.AspNetCore

func TestNonexistentPurl(t *testing.T) {
	purl := "pkg:nuget/nonexistent"
	result := validate_purl(purl, testValidator)
	expected := false

	if result != expected {
		t.Errorf("Validate(\"%s\") = %t; expected %t", purl, result, expected)
	}
}

func TestValidPurl(t *testing.T) {
	purl := "pkg:nuget/FluentUtils.FromCompositeAttribute"
	result := validate_purl(purl, testValidator)
	expected := true

	if result != expected {
		t.Errorf("validate_purl(\"%s\") = %t; expected %t", purl, result, expected)
	}
}

func TestPurlWithTrailingSlash(t *testing.T) {
	purl := "pkg:nuget/FluentUtils.FromCompositeAttribute/"
	result := validate_purl(purl, testValidator)
	expected := true

	if result != expected {
		t.Errorf("validate_purl(\"%s\") = %t; expected %t", purl, result, expected)
	}
}
