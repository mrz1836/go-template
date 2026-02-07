<div align="center">

# 🚀&nbsp;&nbsp;go-template

**From Zero to Go Hero: Pre-wired Template for Modern Libraries**

<br/>

<a href="https://github.com/mrz1836/go-template/releases"><img src="https://img.shields.io/github/release-pre/mrz1836/go-template?include_prereleases&style=flat-square&logo=github&color=black" alt="Release"></a>
<a href="https://golang.org/"><img src="https://img.shields.io/github/go-mod/go-version/mrz1836/go-template?style=flat-square&logo=go&color=00ADD8" alt="Go Version"></a>
<a href="https://github.com/mrz1836/go-template/blob/master/LICENSE"><img src="https://img.shields.io/github/license/mrz1836/go-template?style=flat-square&color=blue" alt="License"></a>

<br/>

<table align="center" border="0">
  <tr>
    <td align="right">
       <code>CI / CD</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://github.com/mrz1836/go-template/actions"><img src="https://img.shields.io/github/actions/workflow/status/mrz1836/go-template/fortress.yml?branch=master&label=build&logo=github&style=flat-square" alt="Build"></a>
       <a href="https://github.com/mrz1836/go-template/actions"><img src="https://img.shields.io/github/last-commit/mrz1836/go-template?style=flat-square&logo=git&logoColor=white&label=last%20update" alt="Last Commit"></a>
    </td>
    <td align="right">
       &nbsp;&nbsp;&nbsp;&nbsp; <code>Quality</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://goreportcard.com/report/github.com/mrz1836/go-template"><img src="https://goreportcard.com/badge/github.com/mrz1836/go-template?style=flat-square" alt="Go Report"></a>
       <a href="https://codecov.io/gh/mrz1836/go-template"><img src="https://codecov.io/gh/mrz1836/go-template/branch/master/graph/badge.svg?style=flat-square" alt="Coverage"></a>
    </td>
  </tr>

  <tr>
    <td align="right">
       <code>Security</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://scorecard.dev/viewer/?uri=github.com/mrz1836/go-template"><img src="https://api.scorecard.dev/projects/github.com/mrz1836/go-template/badge?style=flat-square" alt="Scorecard"></a>
       <a href=".github/SECURITY.md"><img src="https://img.shields.io/badge/policy-active-success?style=flat-square&logo=security&logoColor=white" alt="Security"></a>
    </td>
    <td align="right">
       &nbsp;&nbsp;&nbsp;&nbsp; <code>Community</code> &nbsp;&nbsp;
    </td>
    <td align="left">
       <a href="https://github.com/mrz1836/go-template/graphs/contributors"><img src="https://img.shields.io/github/contributors/mrz1836/go-template?style=flat-square&color=orange" alt="Contributors"></a>
       <a href="https://mrz1818.com/"><img src="https://img.shields.io/badge/donate-bitcoin-ff9900?style=flat-square&logo=bitcoin" alt="Bitcoin"></a>
    </td>
  </tr>
</table>

</div>

<br/>
<br/>

<div align="center">

### <code>Project Navigation</code>

</div>

<table align="center">
  <tr>
    <td align="center" width="33%">
       🚀&nbsp;<a href="#-installation"><code>Installation</code></a>
    </td>
    <td align="center" width="33%">
       🧪&nbsp;<a href="#-examples--tests"><code>Examples&nbsp;&&nbsp;Tests</code></a>
    </td>
    <td align="center" width="33%">
       📚&nbsp;<a href="#-documentation"><code>Documentation</code></a>
    </td>
  </tr>
  <tr>
    <td align="center">
       🤝&nbsp;<a href="#-contributing"><code>Contributing</code></a>
    </td>
    <td align="center">
      🛠️&nbsp;<a href="#-code-standards"><code>Code&nbsp;Standards</code></a>
    </td>
    <td align="center">
      ⚡&nbsp;<a href="#-benchmarks"><code>Benchmarks</code></a>
    </td>
  </tr>
  <tr>
    <td align="center">
      🤖&nbsp;<a href="#-ai-usage--assistant-guidelines"><code>AI&nbsp;Usage</code></a>
    </td>
    <td align="center">
       ⚖️&nbsp;<a href="#-license"><code>License</code></a>
    </td>
    <td align="center">
       👥&nbsp;<a href="#-maintainers"><code>Maintainers</code></a>
    </td>
  </tr>
</table>
<br/>

## 🧩 What's Inside
<!-- remove-this-section:start -->
**go-template** is a plug-and-play scaffold that lets you skip the boilerplate and jump straight to building your Go library.
Clone it, [rename a few placeholders](#-template-kick-off-guide-3-easy-steps), and you instantly inherit a production-grade setup:

<br/>

- **📚 Go Best Practices & Examples**
  _Includes idiomatic Go patterns, [table-driven tests](template_test.go), [benchmarks](template_benchmark_test.go), [example functions](template_example_test.go), and [fuzz tests](template_fuzz_test.go)—demonstrating how to write robust, maintainable, and production-grade Go code._

- **⚡ Zero-config CI/CD**
  _[GitHub Actions](#-documentation) run tests, upload coverage, and enforce linting on every push—so you never forget to run the checks._

- **🛠️ One-command via [MAGE-X](https://github.com/mrz1836/mage-x)**
  _`magex test`, `magex lint`, `magex bench`, and more—common tasks stay muscle-memory simple._

- **🚢 Automated Releases**
  _[GoReleaser](https://goreleaser.com/) cuts signed, versioned artifacts the moment you push a tag—shipping new versions becomes a 10-second ritual._

- **🛡️ Security & Supply-chain Guardrails**
  _[Dependabot](https://dependabot.com), [Nancy](https://github.com/sonatype-nexus-community/nancy), [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck), [CodeQL](https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/about-code-scanning), [OpenSSF Scorecard](https://openssf.org) and [gitleaks](https://github.com/gitleaks/gitleaks) give early warnings before bad things reach production._

- **🎨 Style & Quality Enforcement**
  _[golangci-lint](https://github.com/golangci/golangci-lint) with 50+ linters and [gofumpt](https://github.com/mvdan/gofumpt) keeps the codebase clean and idiomatic - no bikeshedding required._

- **🤖 AI-Friendly Policies**
  _[AGENTS.md](.github/AGENTS.md), [CLAUDE.md](.github/CLAUDE.md), ensure [ChatGPT](https://openai.com), [Claude](https://claude.ai/) & [Gemini](https://gemini.google.com/) follow the same house rules._

- **🌍 Community-Ready Meta**
  _[Issue/PR templates](.github/ISSUE_TEMPLATE), [CODEOWNERS](.github/CODEOWNERS), [label sync](.github/labels.yml), and a welcome bot to show contributors exactly how to get involved._

<br/>

### 🚀 Quick Wins

* **Clone → Tag → Release**: Go from idea to a published version in under five minutes.
* **Works Everywhere**: macOS, Linux, Windows (maybe lol), ARM64 – fully reproducible builds.
* **Battery-Included Examples**: ready-to-run demos, benchmarks, fuzz and race tests.
* **Flexible, Not Fragile**: swap or remove any piece without breaking the whole.

> **Tip:** Run `magex help` right after cloning to see every command the template unlocks.

<br/>
<br/>

___

<br/>
<br/>

### 🛠 Template Kick-Off Guide (3 Easy Steps)

*(delete this section once your project is initialized)*

<br/>

#### 1) Clone or "[Use this template](https://github.com/new?template_name=go-template&template_owner=mrz1836)"

```bash
git clone https://github.com/mrz1836/go-template.git my-lib && cd my-lib
```

... or click **[Use this template](https://github.com/new?template_name=go-template&template_owner=mrz1836)** on GitHub and create a new repo.

<br/>

#### 2) Install [MAGE-X](https://github.com/mrz1836/mage-x) build tool and run the installation script

```bash
go install github.com/mrz1836/mage-x/cmd/magex@latest

# Run the install script to customize the project for your organization
magex InstallTemplate owner=yourorg repo=yourproject
```

**Example:**
```bash
# For GitHub user "acme" creating a project called "awesome-api"
magex InstallTemplate owner=acme repo=awesome-api
```

<br/>

#### What does that command do?

1. **Finds & replaces names across 70+ files**

	* `mrz1836/go-template` → `yourorg/yourproject`
	* `go-template` → `yourproject`
	* `mrz1836` → `yourorg`

2. **Cleans up template artifacts**

	* Removes the default social-share image so you can add your own
	* Updates module paths in `go.mod`
	* Fixes all GitHub badges and links

3. **Provides helpful feedback**

	* Shows exactly which files were modified
	* Gives you next steps to review and commit changes

<br/>

#### 3) Touch up metadata

Edit the highlighted files so they match your project:

* [`LICENSE`](LICENSE)
  * Update the year and your name or organization
* [`README.md`](README.md)
  * Remove the "remove-this-section" block in this file
  * Modify the "About" section to describe your library
* [`.github/SECURITY.md`](.github/SECURITY.md)
  * Update the security policy to match your project's needs
* [`.github/FUNDING.yml`](.github/FUNDING.yml)
  * If you want to accept funding, add your funding links here
* [`.goreleaser.yml`](.goreleaser.yml)
  * Modify settings for Slack, Discord, Twitter, or Reddit if you want to announce releases
* [`CODEOWNERS`](.github/CODEOWNERS)
  * Adjust rules for code ownership if needed

<br/>

#### Give it a spin!
Push your initial commit and run `magex version:bump push=true bump=minor` and the CI/CD pipeline will take it from there. 🚀


<br/>
<br/>

```text
┌──────────────────────────────────────────────────────────────────────────────────────────┐
│                                                                                          │
│           MR. Z'S GO-TEMPLATE – YOUR README STARTS RIGHT AFTER THIS BANNER               │
│                                                                                          │
└──────────────────────────────────────────────────────────────────────────────────────────┘
                                         ⬇ ⬇ ⬇
```

<br/>
<br/>

<!-- remove-this-section:end -->

## 📦 Installation

**go-template** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-template
```

Get the [MAGE-X](https://github.com/mrz1836/mage-x) build tool for development:
```shell script
go install github.com/mrz1836/mage-x/cmd/magex@latest
```

<br/>

## 📚 Documentation

- **API Reference** – Dive into the godocs at [pkg.go.dev/github.com/mrz1836/go-template](https://pkg.go.dev/github.com/mrz1836/go-template)
- **Usage Examples** – Browse practical patterns either the [examples directory](examples) or view the [example functions](template_example_test.go)
- **Benchmarks** – Check the latest numbers in the [benchmark results](#benchmark-results)
- **Test Suite** – Review both the [unit tests](template_test.go) and [fuzz tests](template_fuzz_test.go) (powered by [`testify`](https://github.com/stretchr/testify))

> **Good to know:** `go-template` ships with *zero* runtime dependencies.
> The only external package we use is `testify` and `magefile` — and that's strictly for tests and dev.

<br/>

<details>
<summary><strong>Repository Features</strong></summary>
<br/>

This repository includes 25+ built-in features covering CI/CD, security, code quality, developer experience, and community tooling.

**[View the full Repository Features list →](.github/docs/repository-features.md)**

</details>

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

This project uses [goreleaser](https://github.com/goreleaser/goreleaser) for streamlined binary and library deployment to GitHub. To get started, install it via:

```bash
brew install goreleaser
```

The release process is defined in the [.goreleaser.yml](.goreleaser.yml) configuration file.


Then create and push a new Git tag using:

```bash
magex version:bump push=true bump=patch branch=master
```

This process ensures consistent, repeatable releases with properly versioned artifacts and metadata.

</details>

<details>
<summary><strong><code>Pre-commit Hooks</code></strong></summary>
<br/>

Set up the Go-Pre-commit System to run the same formatting, linting, and tests defined in [AGENTS.md](.github/AGENTS.md) before every commit:

```bash
go install github.com/mrz1836/go-pre-commit/cmd/go-pre-commit@latest
go-pre-commit install
```

The system is configured via modular env files in [`.github/env/`](.github/env/README.md) and provides 17x faster execution than traditional Python-based pre-commit hooks. See the [complete documentation](http://github.com/mrz1836/go-pre-commit) for details.

</details>

<details>
<summary><strong>GitHub Workflows</strong></summary>
<br/>

All workflows are driven by modular configuration in [`.github/env/`](.github/env/README.md) — no YAML editing required.

**[View all workflows and the control center →](.github/docs/workflows.md)**

</details>

<details>
<summary><strong><code>Updating Dependencies</code></strong></summary>
<br/>

To update all dependencies (Go modules, linters, and related tools), run:

```bash
magex deps:update
```

This command ensures all dependencies are brought up to date in a single step, including Go modules and any tools managed by [MAGE-X](https://github.com/mrz1836/mage-x). It is the recommended way to keep your development environment and CI in sync with the latest versions.

</details>

<details>
<summary><strong><code>Build Commands</code></strong></summary>
<br/>

View all build commands

```bash script
magex help
```

</details>

<br/>

## 🧪 Examples & Tests

All unit tests and [examples](examples) run via [GitHub Actions](https://github.com/mrz1836/go-template/actions) and use [Go version 1.24.x](https://go.dev/doc/go1.24). View the [configuration file](.github/workflows/fortress.yml).

Run all tests (fast):

```bash script
magex test
```

Run all tests with race detector (slower):
```bash script
magex test:race
```

<br/>

## ⚡ Benchmarks

Run the Go [benchmarks](template_benchmark_test.go):

```bash script
magex bench
```

<br/>

### Benchmark Results

| Benchmark                           | Iterations | ns/op | B/op | allocs/op |
|-------------------------------------|------------|------:|-----:|----------:|
| [Greet](template_benchmark_test.go) | 21,179,739 | 56.59 |   40 |         2 |

> These benchmarks reflect fast, allocation-free lookups for most retrieval functions, ensuring optimal performance in production environments.
> Performance benchmarks for the core functions in this library, executed on an Apple M1 Max (ARM64).

<br/>

## 🛠️ Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## 🤖 AI Usage & Assistant Guidelines
Read the [AI Usage & Assistant Guidelines](.github/tech-conventions/ai-compliance.md) for details on how AI is used in this project and how to interact with the AI assistants.

<br/>

## 👥 Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" width="50" alt="MrZ" />](https://github.com/mrz1836) |
|:-----------------------------------------------------------------------------------------------------------:|
|                                      [MrZ](https://github.com/mrz1836)                                      |

<br/>

## 🤝 Contributing
View the [contributing guidelines](.github/CONTRIBUTING.md) and please follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:!
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:.
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap:
or by making a [**bitcoin donation**](https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-template&utm_term=go-template&utm_content=go-template) to ensure this journey continues indefinitely! :rocket:

[![Stars](https://img.shields.io/github/stars/mrz1836/go-template?label=Please%20like%20us&style=social&v=1)](https://github.com/mrz1836/go-template/stargazers)

<br/>

## 📝 License

[![License](https://img.shields.io/github/license/mrz1836/go-template.svg?style=flat&v=1)](LICENSE)
