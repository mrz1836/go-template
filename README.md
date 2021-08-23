# go-template
> Robust starter template for a new Go library

[![Release](https://img.shields.io/github/release-pre/mrz1836/go-template.svg?logo=github&style=flat&v=1)](https://github.com/mrz1836/go-template/releases)
[![Build Status](https://img.shields.io/github/workflow/status/mrz1836/go-template/run-go-tests?logo=github&v=1)](https://github.com/mrz1836/go-template/actions)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-template?style=flat&v=1)](https://goreportcard.com/report/github.com/mrz1836/go-template)
[![codecov](https://codecov.io/gh/mrz1836/go-template/branch/master/graph/badge.svg?v=1)](https://codecov.io/gh/mrz1836/go-template)
[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-template?v=1)](https://golang.org/)
<br>
[![Mergify Status](https://img.shields.io/endpoint.svg?url=https://gh.mergify.io/badges/mrz1836/go-template&style=flat&v=1)](https://mergify.io)
[![Sponsor](https://img.shields.io/badge/sponsor-mrz1836-181717.svg?logo=github&style=flat&v=1)](https://github.com/sponsors/mrz1836)
[![Donate](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat&v=1)](https://gobitcoinsv.com/#sponsor?utm_source=github&utm_medium=sponsor-link&utm_campaign=go-template&utm_term=go-template&utm_content=go-template)

<br/>

## Table of Contents
- [About](#about)
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

<br/>

## About
This is a GitHub repository starter template for a Go library. It has been created for ease-of-use for anyone who wants to:

- Quickly get into Go without losing too much time on environment setup
- Create a new repository with basic Continuous Integration
- Cheat codes for all the popular Go patterns

<br/>

**The `go-template` includes:**

- Continuous integration via [GitHub Actions](https://github.com/features/actions)
- Build automation via [Make](https://www.gnu.org/software/make)
- Dependency management using [Go Modules](https://github.com/golang/go/wiki/Modules)
- Code formatting using [gofumpt](https://github.com/mvdan/gofumpt) and linting with [golangci-lint](https://github.com/golangci/golangci-lint)
- Unit testing with [testify](https://github.com/stretchr/testify), [race detector](https://blog.golang.org/race-detector), code coverage [HTML report](https://blog.golang.org/cover) and [Codecov report](https://codecov.io/)
- Releasing using [GoReleaser](https://github.com/goreleaser/goreleaser)
- Dependency scanning and updating thanks to [Dependabot](https://dependabot.com)
- Security code analysis using [CodeQL Action](https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/about-code-scanning)
- Automatic syndication to [pkg.go.dev](https://pkg.go.dev/) on every release
- Basic templates for [Issues and Pull Requests](https://docs.github.com/en/communities/using-templates-to-encourage-useful-issues-and-pull-requests/configuring-issue-templates-for-your-repository) in Github
- Automatic sync for [labels](.github/labels.yml) into Github using a pre-defined [configuration](.github/labels.yml)
- Built-in powerful merging rules using [Mergify](https://mergify.io/)
- [Visual Studio Code](https://code.visualstudio.com) configuration with [Go](https://code.visualstudio.com/docs/languages/go)

<br/>

## Installation

**go-template** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-template
```

1. Sign up on [Codecov](https://codecov.io/) and configure [Codecov GitHub Application](https://github.com/apps/codecov) for all repositories
2. Click the `Use this template` button (alt. clone or download this repository)
3. Replace all occurrences of `mrz1836/go-template` to `your_org/repo_name` in all files
4. Replace all occurrences of `go-template` to `repo_name` in all files
5. Replace all occurrences of `mrz1836` to `repo_owner` in all files
6. Remove or update the [`go-share-image.png`](.github/IMAGES/go-share-image.png) and add to your Github repo (in settings)
7. Update the following files:
    - [LICENSE](LICENSE)
    - [README.md](README.md)
    - [SECURITY.md](SECURITY.md)

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-template)

[![GoDoc](https://godoc.org/github.com/mrz1836/go-template?status.svg&style=flat&v=1)](https://pkg.go.dev/github.com/mrz1836/go-template)

<br/>

<details>
<summary><strong><code>Package Dependencies</code></strong></summary>
<br/>

- [owner/repo](https://github.com/mrz1836/go-template)
- [owner/repo](https://github.com/mrz1836/go-template)
</details>

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed:
- **using make:** `make install-tools`
- **using brew:** `brew install goreleaser`

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
all                  Runs multiple commands
clean                Remove previous builds and any test cache data
clean-mods           Remove all the Go mod cache
coverage             Shows the test coverage
diff                 Show the git diff
generate             Runs the go generate command in the base of the repo
godocs               Sync the latest tag with GoDocs
help                 Show this help message
install              Install the application
install-go           Install the application (Using Native Go)
install-tools        Install go tools
lint                 Run the golangci-lint application (install if not found)
release              Full production release (creates release in Github)
release              Runs common.release then runs godocs
release-snap         Test the full release (build binaries)
release-test         Full production test release (everything except deploy)
replace-version      Replaces the version in HTML/JS (pre-deploy)
tag                  Generate a new tag and push (tag version=0.0.0)
tag-remove           Remove a tag if found (tag-remove version=0.0.0)
tag-update           Update an existing tag to current commit (tag-update version=0.0.0)
test                 Runs lint and ALL tests
test-ci              Runs all tests via CI (exports coverage)
test-ci-no-race      Runs all tests via CI (no race) (exports coverage)
test-ci-short        Runs unit tests via CI (exports coverage)
test-no-lint         Runs just tests
test-short           Runs vet, lint and tests (excludes integration tests)
test-unit            Runs tests and outputs coverage
uninstall            Uninstall the application (and remove files)
update-linter        Update the golangci-lint package (macOS only)
vet                  Run the Go vet application
```
</details>

<br/>

## Examples & Tests
All unit tests and [examples](examples) run via [Github Actions](https://github.com/mrz1836/go-template/actions) and
uses [Go version 1.16.x](https://golang.org/doc/go1.16). View the [configuration file](.github/workflows/run-tests.yml).

Run all tests (including integration tests)
```shell script
make test
```

Run tests (excluding integration tests)
```shell script
make test-short
```

<br/>

## Benchmarks
Run the Go benchmarks:
```shell script
make bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

<br/>

## Usage
Checkout all the [examples](examples)!

<br/>

## Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:---:|
| [MrZ](https://github.com/mrz1836) |

<br/>

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:!
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:.
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap:
or by making a [**bitcoin donation**](https://gobitcoinsv.com/#sponsor?utm_source=github&utm_medium=sponsor-link&utm_campaign=go-template&utm_term=go-template&utm_content=go-template) to ensure this journey continues indefinitely! :rocket:

[![Stars](https://img.shields.io/github/stars/mrz1836/go-template?label=Please%20like%20us&style=social)](https://github.com/mrz1836/go-template/stargazers)

<br/>

## License

![License](https://img.shields.io/github/license/mrz1836/go-template.svg?style=flat&v=1)