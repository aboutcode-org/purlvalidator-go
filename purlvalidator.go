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
	"log"
	"strings"

	"github.com/blevesearch/vellum"
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

func validate_purl(packageURL string, fstMap *vellum.FST) bool {
	packageURL = strings.TrimSuffix(packageURL, "/")
	result, _ := fstMap.Contains([]byte(packageURL))
	return result
}

func Validate(packageURL string) bool {
	return validate_purl(packageURL, validator)
}
