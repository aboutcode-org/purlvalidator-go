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
	_ "embed"
	"fmt"
	"log"

	"github.com/blevesearch/vellum"
	"github.com/package-url/packageurl-go"
)

//go:embed purls.fst
var fstData []byte

var validator *vellum.FST

func init() {
	var err error
	validator, err = vellum.Load(fstData)
	if err != nil {
		log.Fatal(err)
	}
}

func validate_purl(packageURL string, fstMap *vellum.FST) (bool, error) {
	instance, err := packageurl.FromString(packageURL)
	if err != nil {
		return false, err
	}
	if instance.Version != "" || len(instance.Qualifiers) > 0 || instance.Subpath != "" {
		return false, fmt.Errorf("only base PURL is supported (no version, qualifiers, or subpath)")
	}

	result, err := fstMap.Contains([]byte(packageURL))
	return result, err
}

func Validate(packageURL string) (bool, error) {
	return validate_purl(packageURL, validator)
}
