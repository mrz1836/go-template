# 🚀 go-template
> From Zero to Go Hero: Pre-wired Template for Modern Libraries

<table>
  <thead>
    <tr>
      <th>CI&nbsp;/&nbsp;CD</th>
      <th>Quality&nbsp;&amp;&nbsp;Security</th>
      <th>Docs&nbsp;&amp;&nbsp;Meta</th>
      <th>Community</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td valign="top" align="left">
        <a href="https://github.com/mrz1836/go-template/releases">
          <img src="https://img.shields.io/github/release-pre/mrz1836/go-template?logo=github&style=flat" alt="Latest Release">
        </a><br/>
        <a href="https://github.com/mrz1836/go-template/actions">
          <img src="https://img.shields.io/github/actions/workflow/status/mrz1836/go-template/run-tests.yml?branch=master&logo=github&style=flat" alt="Build Status">
        </a><br/>
		<a href="https://github.com/mrz1836/go-template/actions">
          <img src="https://github.com/mrz1836/go-template/actions/workflows/codeql-analysis.yml/badge.svg?style=flat" alt="CodeQL">
        </a><br/>
        <a href="https://github.com/mrz1836/go-template/commits/master">
		  <img src="https://img.shields.io/github/last-commit/mrz1836/go-template?style=flat&logo=clockify&logoColor=white" alt="Last commit">
		</a>
      </td>
      <td valign="top" align="left">
        <a href="https://goreportcard.com/report/github.com/mrz1836/go-template">
          <img src="https://goreportcard.com/badge/github.com/mrz1836/go-template?style=flat" alt="Go Report Card">
        </a><br/>
		<a href="https://codecov.io/gh/mrz1836/go-template">
          <img src="https://codecov.io/gh/mrz1836/go-template/branch/master/graph/badge.svg?style=flat&token=iBdJQpOej9" alt="Code Coverage">
        </a><br/>
		<a href="https://scorecard.dev/viewer/?uri=github.com/mrz1836/go-template">
          <img src="https://api.scorecard.dev/projects/github.com/mrz1836/go-template/badge?logo=springsecurity&logoColor=white" alt="OpenSSF Scorecard">
        </a><br/>
		<a href=".github/SECURITY.md">
          <img src="https://img.shields.io/badge/security-policy-blue?style=flat&logo=springsecurity&logoColor=white" alt="Security policy">
        </a><br/>
		<a href="https://www.bestpractices.dev/projects/10822">
		  <img src="https://www.bestpractices.dev/projects/10822/badge?style=flat&logo=springsecurity&logoColor=white" alt="OpenSSF Best Practices">
		</a>
      </td>
      <td valign="top" align="left">
        <a href="https://golang.org/">
          <img src="https://img.shields.io/github/go-mod/go-version/mrz1836/go-template?style=flat" alt="Go version">
        </a><br/>
        <a href="https://pkg.go.dev/github.com/mrz1836/go-template?tab=doc">
          <img src="https://pkg.go.dev/badge/github.com/mrz1836/go-template.svg?style=flat" alt="Go docs">
        </a><br/>
        <a href=".github/AGENTS.md">
          <img src="https://img.shields.io/badge/AGENTS.md-found-40b814?style=flat&logo=openai" alt="AGENTS.md rules">
        </a><br/>
        <a href="Makefile">
          <img src="https://img.shields.io/badge/Makefile-supported-brightgreen?style=flat&logo=probot&logoColor=white" alt="Makefile Supported">
        </a><br/>
		<a href=".github/dependabot.yml">
          <img src="https://img.shields.io/badge/dependencies-automatic-blue?logo=dependabot&style=flat" alt="Dependabot">
        </a>
      </td>
      <td valign="top" align="left">
        <a href="https://github.com/mrz1836/go-template/graphs/contributors">
          <img src="https://img.shields.io/github/contributors/mrz1836/go-template?style=flat&logo=contentful&logoColor=white" alt="Contributors">
        </a><br/>
        <a href="https://github.com/sponsors/mrz1836">
          <img src="https://img.shields.io/badge/sponsor-MrZ-181717.svg?logo=github&style=flat" alt="Sponsor">
        </a><br/>
        <a href="https://mrz1818.com/?tab=tips&utm_source=github&utm_medium=sponsor-link&utm_campaign=go-template&utm_term=go-template&utm_content=go-template">
          <img src="https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat" alt="Donate Bitcoin">
        </a>
      </td>
    </tr>
  </tbody>
</table>

<br/>

## 🗂️ Table of Contents
* [What's Inside](#-whats-inside)
* [Installation](#-installation)
* [Documentation](#-documentation)
* [Examples & Tests](#-examples--tests)
* [Benchmarks](#-benchmarks)
* [Code Standards](#-code-standards)
* [AI Compliance](#-ai-compliance)
* [Maintainers](#-maintainers)
* [Contributing](#-contributing)
* [License](#-license)

<br/>

<!-- remove-this-section:start -->

## 🧩 What's Inside

**go-template** is a plug-and-play scaffold that lets you skip the boilerplate and jump straight to building your Go library. Clone it, [rename a few placeholders](#-installation), and you instantly inherit a production-grade setup:

<br/>

- **📚 Go Best Practices & Examples**  
  _Includes idiomatic Go patterns, table-driven tests, benchmarks, example functions, and fuzz tests—demonstrating how to write robust, maintainable, and production-grade Go code._

- **⚡ Zero-config CI/CD**  
  _GitHub Actions run tests, upload coverage, and enforce linting on every push—so you never forget to run the checks._

- **🛠️ One-command Makefile**  
  _`make test`, `make lint`, `make bench`, and more—common tasks stay muscle-memory simple._

- **🚢 Automated Releases**  
  _GoReleaser cuts signed, versioned artifacts the moment you push a tag—shipping new versions becomes a 10-second ritual._

- **🛡️ Security & Supply-chain Guardrails**  
  _[Dependabot](https://dependabot.com), [Nancy](https://github.com/sonatype-nexus-community/nancy), [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck), [CodeQL](https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/about-code-scanning), [OpenSSF Scorecard](https://openssf.org), and [gitleaks](https://github.com/gitleaks/gitleaks) give early warnings before bad things reach production._

- **🎨 Style & Quality Enforcement**  
  _[golangci-lint](https://github.com/golangci/golangci-lint) + [gofumpt](https://github.com/mvdan/gofumpt) keeps the codebase clean and idiomatic—no bikeshedding required._

- **🤖 AI-Friendly Policies**  
  _AGENTS.md, CLAUDE.md, cursorrules, and sweep.yaml ensure ChatGPT, Claude, Cursor & Sweep follow the same house rules._

- **🌍 Community-Ready Meta**  
  _Issue/PR templates, CODEOWNERS, CITATION, label sync, and a welcome bot to show contributors exactly how to get involved._

<br/>

### 🚀 Quick Wins

* **Clone → Tag → Release**: Go from idea to a published version in under five minutes.
* **Works Everywhere**: macOS, Linux, Windows (maybe lol), ARM64 – fully reproducible builds.
* **Battery-Included Examples**: ready-to-run demos, benchmarks, fuzz and race tests.
* **Flexible, Not Fragile**: swap or remove any piece without breaking the whole.

> **Tip:** Run `make help` right after cloning to see every command the template unlocks.

<br/>

### 🛠️ Template Kick-Off Guide

*(delete this section once your project is initialized)*

<br/>

#### 1) Clone or "[Use this template](https://github.com/new?template_name=go-template&template_owner=mrz1836)"

```bash
git clone https://github.com/mrz1836/go-template.git my-lib
cd my-lib
```

... or click **[Use this template](https://github.com/new?template_name=go-template&template_owner=mrz1836)** on GitHub and create a new repo.

<br/>

#### 2) Personalize the template in one command

```bash
make install-template owner=my_org repo=my-lib
````

<br/>

#### What does that command do?

1. **Finds & replaces names**

	* `mrz1836/go-template` → `my_org/my-lib`
	* `go-template` → `my-lib`
	* `mrz1836` → `my_org`

2. **Cleans up after itself**

	* Deletes the temporary install file `.make/temp.mk`
	* Removes its own `include .make/temp.mk` line from the root `Makefile`

3. **House-keeping**

	* Remove the default social-share image so you can drop in your own

> That’s it—open a diff, make sure you're happy, commit, and push. 🎉

<br/>

#### 3) Drop in your own share graphic (optional)

Replace `.github/IMAGES/go-share-image.png` with something on-brand and then [add it to GitHub as the default share image](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/customizing-your-repositorys-social-media-preview) for your repo.

<br/>

#### 4) Touch up metadata

Edit the highlighted files so they match your project:

* [`AGENTS.md`](.github/AGENTS.md)
  * Update the project name, description, and any other relevant details in the beginning of the file
* [`LICENSE`](LICENSE)
  * Update the year and your name or organization
* [`README.md`](README.md)
  * Remove the "remove-this-section" block at the bottom
  * Modify the "About" section to describe your library
* [`.github/SECURITY.md`](.github/SECURITY.md)
  * Update the security policy to match your project's needs

<br/>

#### (Optional) Configure Release Announcements

| Channel                        | How to wire it up                                                                                                 |
|--------------------------------|-------------------------------------------------------------------------------------------------------------------|
| **Slack**                      | 1. Create an [Incoming Webhook](https://api.slack.com/messaging/webhooks)<br>2. Add a repo secret `SLACK_WEBHOOK` |
| **Discord / Twitter / Reddit** | Follow the steps in [GoReleaser docs](https://goreleaser.com/customization/announce/#discord)                     |

<br/>

#### Give it a spin!
Push your initial commit and run `make tag version=0.1.0` - the CI/CD pipeline will take it from there. 🚀

<br/>

<!-- remove-this-section:end -->

## 📦 Installation

**go-template** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-template
```

<br/>

## 📚 Documentation
View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-template?tab=doc)

> **Heads up!** `go-template` is intentionally light on dependencies. The only
external package it uses is the excellent `testify` suite—and that's just for
our tests. You can drop this library into your projects without dragging along
extra baggage.

- Explore additional [usage examples](examples) for practical integration patterns
- Review [benchmark results](#benchmark-results) to assess performance characteristics
- Examine the comprehensive [test suite](template_test.go) for validation and coverage
- Fuzz tests [are available](template_fuzz_test.go) to ensure robustness against unexpected inputs

<br/>

<details>
<summary><strong><code>Repository Features</code></strong></summary>
<br/>

This repository was created using [MrZ's `go-template`](https://github.com/mrz1836/go-template#about)

### 🚀 Built-in Superpowers

* **Continuous Integration on Autopilot** with [GitHub Actions](https://github.com/features/actions) – every push is built, tested, and reported in minutes.
* **Pull‑Request Flow That Merges Itself** thanks to [auto‑merge](.github/workflows/auto-merge-on-approval.yml) and hands‑free [Dependabot auto‑merge](.github/workflows/dependabot-auto-merge.yml).
* **One‑Command Builds** powered by battle‑tested [Make](https://www.gnu.org/software/make) targets for linting, testing, releases, and more.
* **First‑Class Dependency Management** using native [Go Modules](https://github.com/golang/go/wiki/Modules).
* **Uniform Code Style** via [gofumpt](https://github.com/mvdan/gofumpt) plus zero‑noise linting with [golangci‑lint](https://github.com/golangci/golangci-lint).
* **Confidence‑Boosting Tests** with [testify](https://github.com/stretchr/testify), the Go [race detector](https://blog.golang.org/race-detector), crystal‑clear [HTML coverage](https://blog.golang.org/cover) snapshots, and automatic uploads to [Codecov](https://codecov.io/).
* **Hands‑Free Releases** delivered by [GoReleaser](https://github.com/goreleaser/goreleaser) whenever you create a [new Tag](https://git-scm.com/book/en/v2/Git-Basics-Tagging).
* **Relentless Dependency & Vulnerability Scans** via [Dependabot](https://dependabot.com), [Nancy](https://github.com/sonatype-nexus-community/nancy), and [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck).
* **Security Posture by Default** with [CodeQL](https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/about-code-scanning), [OpenSSF Scorecard](https://openssf.org), and secret‑leak detection via [gitleaks](https://github.com/gitleaks/gitleaks).
* **Automatic Syndication** to [pkg.go.dev](https://pkg.go.dev/) on every release for instant godoc visibility.
* **Polished Community Experience** using rich templates for [Issues & PRs](https://docs.github.com/en/communities/using-templates-to-encourage-useful-issues-and-pull-requests/configuring-issue-templates-for-your-repository).
* **All the Right Meta Files** (`LICENSE`, `CONTRIBUTING.md`, `CODE_OF_CONDUCT.md`, `SUPPORT.md`, `SECURITY.md`) pre‑filled and ready.
* **Code Ownership** clarified through a [CODEOWNERS](.github/CODEOWNERS) file, keeping reviews fast and focused.
* **Zero‑Noise Dev Environments** with tuned editor settings (`.editorconfig`) plus curated *ignore* files for [VS Code](.editorconfig), [Docker](.dockerignore), and [Git](.gitignore).
* **Label Sync Magic**: your repo labels stay in lock‑step with [.github/labels.yml](.github/labels.yml).
* **Friendly First PR Workflow** – newcomers get a warm welcome thanks to a dedicated [workflow](.github/workflows/pull-request-management.yml).
* **Standards‑Compliant Docs** adhering to the [standard‑readme](https://github.com/RichardLitt/standard-readme/blob/master/spec.md) spec.
* **Out‑of‑the‑Box VS Code Happiness** with a preconfigured [Go](https://code.visualstudio.com/docs/languages/go) workspace.
* **Optional Release Broadcasts** to your community via [Slack](https://slack.com), [Discord](https://discord.com), or [Twitter](https://twitter.com) – plug in your webhook.
* **AI Compliance Playbook** – machine‑readable guidelines ([AGENTS.md](.github/AGENTS.md), [CLAUDE.md](.github/CLAUDE.md), [.cursorrules](.cursorrules), [sweep.yaml](.github/sweep.yaml)) keep ChatGPT, Claude, Cursor & Sweep aligned with your repo’s rules.

</details>

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

This project uses [goreleaser](https://github.com/goreleaser/goreleaser) for streamlined binary and library deployment to GitHub. To get started, install it via:

```bash
brew install goreleaser
```

The release process is defined in the [.goreleaser.yml](.goreleaser.yml) configuration file.

To generate a snapshot (non-versioned) release for testing purposes, run:

```bash
make release-snap
```

Before tagging a new version, update the release metadata in the `CITATION.cff` file:

```bash
make citation version=0.2.1
```

Then create and push a new Git tag using:

```bash
make tag version=x.y.z
```

This process ensures consistent, repeatable releases with properly versioned artifacts and citation metadata.

</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands

```bash script
make help
```

List of all current commands:

<!-- make-help-start -->
```text
bench                 ## Run all benchmarks in the Go application
build-go              ## Build the Go application (locally)
citation              ## Update version in CITATION.cff (use version=X.Y.Z)
clean-mods            ## Remove all the Go mod cache
coverage              ## Show test coverage
diff                  ## Show git diff and fail if uncommitted changes exist
generate              ## Run go generate in the base of the repo
godocs                ## Trigger GoDocs tag sync
govulncheck-install   ## Install govulncheck
help                  ## Display this help message
install-go            ## Install using go install with specific version
install-releaser      ## Install GoReleaser
install-template      ## Kick-start a fresh copy of go-template
install               ## Install the application binary
lint                  ## Run the golangci-lint application (install if not found)
release-snap          ## Build snapshot binaries
release-test          ## Run release dry-run (no publish)
release               ## Run production release (requires github_token)
run-fuzz-tests        ## Run fuzz tests for all packages
tag-remove            ## Remove local and remote tag (use version=X.Y.Z)
tag-update            ## Force-update tag to current commit (use version=X.Y.Z)
tag                   ## Create and push a new tag (use version=X.Y.Z)
test-ci-no-race       ## CI test suite without race detector
test-ci-short         ## CI unit-only short tests
test-ci               ## CI full test suite with coverage
test-no-lint          ## Run only tests (no lint)
test-short            ## Run tests excluding integration
test-unit             ## Runs tests and outputs coverage
test                  ## Run lint and all tests
uninstall             ## Uninstall the Go binary
update-linter         ## Upgrade golangci-lint (macOS only)
update-releaser       ## Reinstall GoReleaser
update                ## Update dependencies
vet                   ## Run go vet
```
<!-- make-help-end -->

</details>

<details>
<summary><strong><code>GitHub Workflows</code></strong></summary>
<br/>

| Workflow Name                                                                | Description                                                                                                            |
|------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| [auto-merge-on-approval.yml](.github/workflows/auto-merge-on-approval.yml)   | Automatically merges PRs after approval and all required checks, following strict rules.                               |
| [check-for-leaks.yml](.github/workflows/check-for-leaks.yml)                 | Runs [gitleaks](https://github.com/gitleaks/gitleaks) to detect secrets on a daily schedule.                           |
| [clean-runner-cache.yml](.github/workflows/clean-runner-cache.yml)           | Removes GitHub Actions caches tied to closed pull requests.                                                            |
| [codeql-analysis.yml](.github/workflows/codeql-analysis.yml)                 | Analyzes code for security vulnerabilities using [GitHub CodeQL](https://codeql.github.com/).                          |
| [delete-merged-branches.yml](.github/workflows/delete-merged-branches.yml)   | Deletes feature branches after their pull requests are merged.                                                         |
| [dependabot-auto-merge.yml](.github/workflows/dependabot-auto-merge.yml)     | Automatically merges [Dependabot](https://github.com/dependabot) PRs that meet all requirements.                       |
| [pull-request-management.yml](.github/workflows/pull-request-management.yml) | Labels PRs by branch prefix, assigns a default user if none is assigned, and welcomes new contributors with a comment. |
| [release.yml](.github/workflows/release.yml)                                 | Builds and publishes releases via [GoReleaser](https://goreleaser.com/intro/) when a semver tag is pushed.             |
| [run-tests.yml](.github/workflows/run-tests.yml)                             | Runs linter, Go tests and dependency checks on every push and pull request.                                            |
| [scorecard.yml](.github/workflows/scorecard.yml)                             | Runs [OpenSSF](https://openssf.org/) Scorecard to assess supply chain security.                                        |
| [stale.yml](.github/workflows/stale.yml)                                     | Warns about (and optionally closes) inactive issues and PRs on a schedule or manual trigger.                           |
| [sync-labels.yml](.github/workflows/sync-labels.yml)                         | Keeps GitHub labels in sync with the declarative manifest at [`.github/labels.yml`](./.github/labels.yml).             |

</details>

<details>
<summary><strong><code>Updating Dependencies</code></strong></summary>
<br/>

To update all dependencies (Go modules, linters, and related tools), run:

```bash
make update
```

This command ensures all dependencies are brought up to date in a single step, including Go modules and any tools managed by the Makefile. It is the recommended way to keep your development environment and CI in sync with the latest versions.

</details>

<br/>

## 🧪 Examples & Tests

All unit tests and [examples](examples) run via [GitHub Actions](https://github.com/mrz1836/go-template/actions) and use [Go version 1.24.x](https://go.dev/doc/go1.24). View the [configuration file](.github/workflows/run-tests.yml).

Run all tests:

```bash script
make test
```

<br/>

## ⚡ Benchmarks

Run the Go [benchmarks](template_benchmark_test.go):

```bash script
make bench
```

<br/>

Performance benchmarks for the core functions in this library, executed on an Apple M1 Max (ARM64):

### Benchmark Results

| Benchmark                           | Iterations | ns/op | B/op | allocs/op |
|-------------------------------------|------------|------:|-----:|----------:|
| [Greet](template_benchmark_test.go) | 21,179,739 | 56.59 |   40 |         2 |

> These benchmarks reflect fast, allocation-free lookups for most retrieval functions, ensuring optimal performance in production environments.

<br/>

## 🛠️ Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## 🤖 AI Compliance
This project documents expectations for AI assistants using a few dedicated files:

- [AGENTS.md](.github/AGENTS.md) — canonical rules for coding style, workflows, and pull requests used by [Codex](https://chatgpt.com/codex).
- [CLAUDE.md](.github/CLAUDE.md) — quick checklist for the [Claude](https://www.anthropic.com/product) agent.
- [.cursorrules](.cursorrules) — machine-readable subset of the policies for [Cursor](https://www.cursor.so/) and similar tools.
- [sweep.yaml](.github/sweep.yaml) — rules for [Sweep](https://github.com/sweepai/sweep), a tool for code review and pull request management.

Edit `AGENTS.md` first when adjusting these policies, and keep the other files in sync within the same pull request.

<br/>

## 👥 Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:------------------------------------------------------------------------------------------------:|
|                                [MrZ](https://github.com/mrz1836)                                 |

<br/>

## 🤝 Contributing
View the [contributing guidelines](.github/CONTRIBUTING.md) and please follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:!
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:.
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap:
or by making a [**bitcoin donation**](https://gobitcoinsv.com/#sponsor?utm_source=github&utm_medium=sponsor-link&utm_campaign=go-template&utm_term=go-template&utm_content=go-template) to ensure this journey continues indefinitely! :rocket:

[![Stars](https://img.shields.io/github/stars/mrz1836/go-template?label=Please%20like%20us&style=social&v=1)](https://github.com/mrz1836/go-template/stargazers)

<br/>

## 📝 License

[![License](https://img.shields.io/github/license/mrz1836/go-template.svg?style=flat&v=1)](LICENSE)
