# purlvalidator

[![License](https://img.shields.io/badge/License-Apache--2.0-blue.svg?style=for-the-badge)](https://opensource.org/licenses/Apache-2.0)
[![Version](https://img.shields.io/github/v/release/aboutcode-org/purlvalidator-go?style=for-the-badge)](https://github.com/aboutcode-org/purlvalidator-go/releases)
[![Test](https://img.shields.io/github/actions/workflow/status/aboutcode-org/purlvalidator-go/ci.yml?style=for-the-badge&logo=github)](https://github.com/aboutcode-org/purlvalidator-go/actions)

**purlvalidator** is a Go library for validating [Package-URLs (PURLs)](https://github.com/package-url/purl-spec). It works fully offline, including in **air-gapped** or **restricted environments**, and answers one key question: **Does the package this PURL represents actually exist?**

## How It Works

**purlvalidator** is shipped with a pre-built FST (Finite State Transducer), a set of compact automata containing latest Package-URLs mined by the MineCode[^1]. Library uses this FST to perform lookups and confirm whether the **base PURL**[^2] exists.

## Currently Supported Ecosystems

- apk
- cargo
- composer
- conan
- cpan
- cran
- debian
- maven
- npm
- nuget
- pypi
- swift

## Usage

Add `purlvalidator` as a dependency:

```bash
go get github.com/aboutcode-org/purlvalidator-go
```

Or add it to `go.mod`:

```text
require github.com/aboutcode-org/purlvalidator-go v1.0.0
```

Use it in your code like this:

```go
package main

import (
	"fmt"
	"log"

	purlvalidator "github.com/aboutcode-org/purlvalidator-go"
)

func main() {
	exists, err := purlvalidator.Validate("pkg:nuget/FluentValidation")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(exists)
}
```

Examples and errors:

```go
exists, err := purlvalidator.Validate("pkg:nuget/FluentValidation")
// exists == true, err == nil

exists, err = purlvalidator.Validate("pkg:nuget/non-existent-foo-bar")
// exists == false, err == nil

exists, err = purlvalidator.Validate("pkg:nuget/FluentValidation@10.2.3")
// err reports that only base PURLs are supported.

exists, err = purlvalidator.Validate("test:nuget/FluentValidation")
// err reports that the PURL scheme is invalid.
```

`Validate` returns:

- `true, nil` when the base PURL exists in the packaged data.
- `false, nil` when the base PURL is syntactically valid but unknown.
- `false, err` when the input is not a valid PURL or contains a version,
  qualifiers, or subpath.

Use the released module version when you need reproducible validation results.
Use a newer patch release when you need newer packaged PURL data.


## How to get latest Package-URL data?

A patch release is published daily with the latest FST generated from newly mined package-urls.


## Contribution

We welcome contributions from the community! If you find a bug or have an idea for a new feature, please open an issue on the GitHub repository. If you want to contribute code, you can fork the repository, make your changes, and submit a pull request.

* Please try to write a good commit message, see [good commit message wiki](https://aboutcode.readthedocs.io/en/latest/contributing/writing_good_commit_messages.html).
* Add DCO `Sign Off` to your commits.

## Development Setup

Run these commands, starting from a git clone of [https://github.com/aboutcode-org/purlvalidator-go.git](https://github.com/aboutcode-org/purlvalidator-go.git)

Generate FST:

```bash
make build-fst
```

Run tests:

```bash
make test
```

Fix formatting and linting:

```bash
make valid
```

## License

SPDX-License-Identifier: Apache-2.0

purl-validator is licensed under Apache License version 2.0.

```text
You may not use this software except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[^1]: MineCode continuously collects package metadata from various package ecosystems to maintain an up-to-date catalog of known packages.
[^2]: A Base Package-URL is a Package-URL without a version, qualifiers, or subpath.
