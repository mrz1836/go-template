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
          <img src="https://img.shields.io/github/release-pre/mrz1836/go-template?logo=github&style=flat&v=1" alt="Latest Release">
        </a><br/>
        <a href="https://github.com/mrz1836/go-template/actions">
          <img src="https://img.shields.io/github/actions/workflow/status/mrz1836/go-template/fortress.yml?branch=master&logo=github&style=flat" alt="Build Status">
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
		<a href="https://codecov.io/gh/mrz1836/go-template/tree/master">
          <img src="https://codecov.io/gh/mrz1836/go-template/branch/master/graph/badge.svg?style=flat" alt="Code Coverage">
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
		<a href="https://magefile.org/">
          <img src="https://img.shields.io/badge/mage-powered-brightgreen?style=flat&logo=probot&logoColor=white" alt="Mage Powered">
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
  _[AGENTS.md](.github/AGENTS.md), [CLAUDE.md](.github/CLAUDE.md), [cursorrules](.cursorrules), and [sweep.yaml](.github/sweep.yaml) ensure [ChatGPT](https://openai.com), [Claude](https://claude.ai/), [Cursor](https://www.cursor.com/) & [Sweep](https://sweep.dev/) follow the same house rules._

- **🌍 Community-Ready Meta**
  _[Issue/PR templates](.github/ISSUE_TEMPLATE), [CODEOWNERS](.github/CODEOWNERS), [CITATION](CITATION.cff), [label sync](.github/labels.yml), and a welcome bot to show contributors exactly how to get involved._

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
<summary><strong><code>Repository Features</code></strong></summary>
<br/>

* **Continuous Integration on Autopilot** with [GitHub Actions](https://github.com/features/actions) – every push is built, tested, and reported in minutes.
* **Pull‑Request Flow That Merges Itself** thanks to [auto‑merge](.github/workflows/auto-merge-on-approval.yml) and hands‑free [Dependabot auto‑merge](.github/workflows/dependabot-auto-merge.yml).
* **One‑Command Builds** powered by battle‑tested [MAGE-X](https://github.com/mrz1836/mage-x) targets for linting, testing, releases, and more.
* **First‑Class Dependency Management** using native [Go Modules](https://github.com/golang/go/wiki/Modules).
* **Uniform Code Style** via [gofumpt](https://github.com/mvdan/gofumpt) plus zero‑noise linting with [golangci‑lint](https://github.com/golangci/golangci-lint).
* **Confidence‑Boosting Tests** with [testify](https://github.com/stretchr/testify), the Go [race detector](https://blog.golang.org/race-detector), crystal‑clear [HTML coverage](https://blog.golang.org/cover) snapshots, and automatic uploads to [Codecov](https://codecov.io/).
* **Hands‑Free Releases** delivered by [GoReleaser](https://github.com/goreleaser/goreleaser) whenever you create a [new Tag](https://git-scm.com/book/en/v2/Git-Basics-Tagging).
* **Relentless Dependency & Vulnerability Scans** via [Dependabot](https://dependabot.com), [Nancy](https://github.com/sonatype-nexus-community/nancy) and [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck).
* **Security Posture by Default** with [CodeQL](https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/about-code-scanning), [OpenSSF Scorecard](https://openssf.org) and secret‑leak detection via [gitleaks](https://github.com/gitleaks/gitleaks).
* **Automatic Syndication** to [pkg.go.dev](https://pkg.go.dev/) on every release for instant godoc visibility.
* **Polished Community Experience** using rich templates for [Issues & PRs](https://docs.github.com/en/communities/using-templates-to-encourage-useful-issues-and-pull-requests/configuring-issue-templates-for-your-repository).
* **All the Right Meta Files** (`LICENSE`, `CITATION.cff`, `CONTRIBUTING.md`, `CODE_OF_CONDUCT.md`, `SUPPORT.md`, `SECURITY.md`) pre‑filled and ready.
* **Code Ownership** clarified through a [CODEOWNERS](.github/CODEOWNERS) file, keeping reviews fast and focused.
* **Zero‑Noise Dev Environments** with tuned editor settings (`.editorconfig`) plus curated *ignore* files for [VS Code](.editorconfig), [Docker](.dockerignore), and [Git](.gitignore).
* **Label Sync Magic**: your repo labels stay in lock‑step with [.github/labels.yml](.github/labels.yml).
* **Friendly First PR Workflow** – newcomers get a warm welcome thanks to a dedicated [workflow](.github/workflows/pull-request-management.yml).
* **Standards‑Compliant Docs** adhering to the [standard‑readme](https://github.com/RichardLitt/standard-readme/blob/master/spec.md) spec.
* **Instant Cloud Workspaces** via [Gitpod](https://gitpod.io/) – spin up a fully configured dev environment with automatic linting and tests.
* **Out‑of‑the‑Box VS Code Happiness** with a preconfigured [Go](https://code.visualstudio.com/docs/languages/go) workspace and [`.vscode`](.vscode) folder with all the right settings.
* **Optional Release Broadcasts** to your community via [Slack](https://slack.com), [Discord](https://discord.com), or [Twitter](https://twitter.com) – plug in your webhook.
* **AI Compliance Playbook** – machine‑readable guidelines ([AGENTS.md](.github/AGENTS.md), [CLAUDE.md](.github/CLAUDE.md), [.cursorrules](.cursorrules), [sweep.yaml](.github/sweep.yaml)) keep ChatGPT, Claude, Cursor & Sweep aligned with your repo's rules.
* **Go-Pre-commit System** - [High-performance Go-native pre-commit hooks](https://github.com/mrz1836/go-pre-commit) with 17x faster execution—run the same formatting, linting, and tests before every commit, just like CI.
* **Zero Python Dependencies** - Pure Go implementation with environment-based configuration via [.env.base](.github/.env.base).
* **DevContainers for Instant Onboarding** – Launch a ready-to-code environment in seconds with [VS Code DevContainers](https://containers.dev/) and the included [.devcontainer.json](.devcontainer.json) config.

</details>

<details>
<summary><strong><code>Repository File Glossary</code></strong></summary>
<br/>

This glossary describes each tracked file in the repository and notes if it is required for GitHub or another external service.

| File Path                                                                                      | Description                                     | Service          |
|------------------------------------------------------------------------------------------------|-------------------------------------------------|------------------|
| [.cursorrules](.cursorrules)                                                                   | Rules for Cursor AI integrations                | Cursor           |
| [.devcontainer.json](.devcontainer.json)                                                       | VS Code dev or GitHub container configuration   | VS Code & GitHub |
| [.dockerignore](.dockerignore)                                                                 | Paths ignored by Docker builds                  | Docker           |
| [.editorconfig](.editorconfig)                                                                 | Editor configuration defaults                   | Editor           |
| [.gitattributes](.gitattributes)                                                               | Git attributes and export settings              | Git              |
| [.github/.env.base](.github/.env.base)                                                         | Shared environment variables for GitHub Actions | GitHub Actions   |
| [.github/.env.custom](.github/.env.custom)                                                     | Custom environment variables for GitHub Actions | GitHub Actions   |
| [.github/AGENTS.md](.github/AGENTS.md)                                                         | Contribution rules and guidelines               | GitHub           |
| [.github/CLAUDE.md](.github/CLAUDE.md)                                                         | Claude agent instructions                       | Claude           |
| [.github/CODEOWNERS](.github/CODEOWNERS)                                                       | Code ownership declarations for GitHub          | GitHub           |
| [.github/CODE_OF_CONDUCT.md](.github/CODE_OF_CONDUCT.md)                                       | Community behavior standards                    | GitHub           |
| [.github/CODE_STANDARDS.md](.github/CODE_STANDARDS.md)                                         | Coding style guide                              | GitHub           |
| [.github/CONTRIBUTING.md](.github/CONTRIBUTING.md)                                             | How to contribute to the project                | GitHub           |
| [.github/FUNDING.yml](.github/FUNDING.yml)                                                     | Funding links displayed by GitHub               | GitHub           |
| [.github/IMAGES/go-share-image.png](.github/IMAGES/go-share-image.png)                         | Social sharing image                            | GitHub           |
| [.github/ISSUE_TEMPLATE/bug_report.yml](.github/ISSUE_TEMPLATE/bug_report.yml)                 | Issue template for bug reports                  | GitHub           |
| [.github/ISSUE_TEMPLATE/feature_request.yml](.github/ISSUE_TEMPLATE/feature_request.yml)       | Issue template for feature requests             | GitHub           |
| [.github/ISSUE_TEMPLATE/question.yml](.github/ISSUE_TEMPLATE/question.yml)                     | Issue template for questions                    | GitHub           |
| [.github/SECURITY.md](.github/SECURITY.md)                                                     | Security policy                                 | GitHub           |
| [.github/SUPPORT.md](.github/SUPPORT.md)                                                       | Support guidelines                              | GitHub           |
| [.github/dependabot.yml](.github/dependabot.yml)                                               | Dependabot configuration                        | GitHub           |
| [.github/labels.yml](.github/labels.yml)                                                       | Repository label definitions                    | GitHub           |
| [.github/pull_request_template.md](.github/pull_request_template.md)                           | Pull request description template               | GitHub           |
| [.github/sweep.yaml](.github/sweep.yaml)                                                       | Sweep AI configuration                          | Sweep AI         |
| [.github/workflows/auto-merge-on-approval.yml](.github/workflows/auto-merge-on-approval.yml)   | Workflow for automatic merges                   | GitHub Actions   |
| [.github/workflows/codeql-analysis.yml](.github/workflows/codeql-analysis.yml)                 | CodeQL security analysis workflow               | GitHub Actions   |
| [.github/workflows/dependabot-auto-merge.yml](.github/workflows/dependabot-auto-merge.yml)     | Auto merge Dependabot PRs                       | GitHub Actions   |
| [.github/workflows/fortress.yml](.github/workflows/fortress.yml)                               | Fortress security & testing workflow            | GitHub Actions   |
| [.github/workflows/pull-request-management.yml](.github/workflows/pull-request-management.yml) | Pull request triage workflow                    | GitHub Actions   |
| [.github/workflows/scorecard.yml](.github/workflows/scorecard.yml)                             | OpenSSF Scorecard workflow                      | GitHub Actions   |
| [.github/workflows/stale.yml](.github/workflows/stale-check.yml)                               | Close stale issues and PRs                      | GitHub Actions   |
| [.github/workflows/sync-labels.yml](.github/workflows/sync-labels.yml)                         | Sync repository labels                          | GitHub Actions   |
| [.gitignore](.gitignore)                                                                       | Files and directories Git should ignore         | Git              |
| [.gitpod.yml](.gitpod.yml)                                                                     | Gitpod workspace configuration                  | Gitpod           |
| [.golangci.json](.golangci.json)                                                               | GolangCI-Lint configuration                     | GolangCI-Lint    |
| [.goreleaser.yml](.goreleaser.yml)                                                             | GoReleaser configuration for release automation | GoReleaser       |
| [.prettierignore](.github/.prettierignore)                                                     | Paths ignored by Prettier formatting            | Prettier         |
| [.prettierrc.yml](.github/.prettierrc.yml)                                                     | Prettier configuration file                     | Prettier         |
| [.vscode/extensions.json](.vscode/extensions.json)                                             | Recommended VS Code extensions                  | VS Code          |
| [.vscode/launch.json](.vscode/launch.json)                                                     | VS Code debugging configuration                 | VS Code          |
| [.vscode/settings.json](.vscode/settings.json)                                                 | VS Code workspace settings                      | VS Code          |
| [.vscode/tasks.json](.vscode/tasks.json)                                                       | VS Code tasks configuration                     | VS Code          |
| [CITATION.cff](CITATION.cff)                                                                   | Citation metadata recognized by GitHub          | GitHub           |
| [Dockerfile](Dockerfile)                                                                       | Docker image build instructions                 | Docker           |
| [LICENSE](LICENSE)                                                                             | Project license                                 | Yours!           |
| [README.md](README.md)                                                                         | Project overview and usage                      | Yours!           |
| [codecov.yml](codecov.yml)                                                                     | Codecov upload configuration                    | Codecov          |
| [examples/example.go](examples/example.go)                                                     | Example usage of the library                    | None             |
| [go.mod](go.mod)                                                                               | Go module definition                            | Go               |
| [go.sum](go.sum)                                                                               | Dependency checksums generated by Go            | Go               |
| [magefile.go](magefile.go)                                                                     | Magefile with MAGE-X targets                    | MAGE-X           |
| [template.go](template.go)                                                                     | Main package source code                        | Yours!           |
| [template_benchmark_test.go](template_benchmark_test.go)                                       | Go benchmark tests                              | Go test          |
| [template_example_test.go](template_example_test.go)                                           | Example tests for documentation                 | Go test          |
| [template_fuzz_test.go](template_fuzz_test.go)                                                 | Go fuzz tests                                   | Go test          |
| [template_test.go](template_test.go)                                                           | Unit tests                                      | Go test          |
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
magex version:bump push=true bump=patch
```

This process ensures consistent, repeatable releases with properly versioned artifacts and citation metadata.

</details>

<details>
<summary><strong><code>Pre-commit Hooks</code></strong></summary>
<br/>

Set up the Go-Pre-commit System to run the same formatting, linting, and tests defined in [AGENTS.md](.github/AGENTS.md) before every commit:

```bash
go install github.com/mrz1836/go-pre-commit/cmd/go-pre-commit@latest
go-pre-commit install
```

The system is configured via [.env.base](.github/.env.base) and can be customized using also using [.env.custom](.github/.env.custom) and provides 17x faster execution than traditional Python-based pre-commit hooks. See the [complete documentation](http://github.com/mrz1836/go-pre-commit) for details.

</details>

<details>
<summary><strong><code>GitHub Workflows</code></strong></summary>
<br/>


### 🎛️ The Workflow Control Center

All GitHub Actions workflows in this repository are powered by a single configuration files – your one-stop shop for tweaking CI/CD behavior without touching a single YAML file! 🎯

**Configuration Files:**
- **[.env.base](.github/.env.base)** – Default configuration that works for most Go projects
- **[.env.custom](.github/.env.custom)** – Optional project-specific overrides

This magical file controls everything from:
- **⚙️ Go version matrix** (test on multiple versions or just one)
- **🏃 Runner selection** (Ubuntu or macOS, your wallet decides)
- **🔬 Feature toggles** (coverage, fuzzing, linting, race detection, benchmarks)
- **🛡️ Security tool versions** (gitleaks, nancy, govulncheck)
- **🤖 Auto-merge behaviors** (how aggressive should the bots be?)
- **🏷️ PR management rules** (size labels, auto-assignment, welcome messages)

<br/>

| Workflow Name                                                                      | Description                                                                                                            |
|------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| [auto-merge-on-approval.yml](.github/workflows/auto-merge-on-approval.yml)         | Automatically merges PRs after approval and all required checks, following strict rules.                               |
| [codeql-analysis.yml](.github/workflows/codeql-analysis.yml)                       | Analyzes code for security vulnerabilities using [GitHub CodeQL](https://codeql.github.com/).                          |
| [dependabot-auto-merge.yml](.github/workflows/dependabot-auto-merge.yml)           | Automatically merges [Dependabot](https://github.com/dependabot) PRs that meet all requirements.                       |
| [fortress.yml](.github/workflows/fortress.yml)                                     | Runs the GoFortress security and testing workflow, including linting, testing, releasing, and vulnerability checks.    |
| [pull-request-management.yml](.github/workflows/pull-request-management.yml)       | Labels PRs by branch prefix, assigns a default user if none is assigned, and welcomes new contributors with a comment. |
| [scorecard.yml](.github/workflows/scorecard.yml)                                   | Runs [OpenSSF](https://openssf.org/) Scorecard to assess supply chain security.                                        |
| [stale.yml](.github/workflows/stale-check.yml)                                     | Warns about (and optionally closes) inactive issues and PRs on a schedule or manual trigger.                           |
| [sync-labels.yml](.github/workflows/sync-labels.yml)                               | Keeps GitHub labels in sync with the declarative manifest at [`.github/labels.yml`](./.github/labels.yml).             |

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

## 🤖 AI Compliance
This project documents expectations for AI assistants using a few dedicated files:

- [AGENTS.md](.github/AGENTS.md) — canonical rules for coding style, workflows, and pull requests used by [Codex](https://chatgpt.com/codex).
- [CLAUDE.md](.github/CLAUDE.md) — quick checklist for the [Claude](https://www.anthropic.com/product) agent.
- [.cursorrules](.cursorrules) — machine-readable subset of the policies for [Cursor](https://www.cursor.so/) and similar tools.
- [sweep.yaml](.github/sweep.yaml) — rules for [Sweep](https://github.com/sweepai/sweep), a tool for code review and pull request management.

Edit `AGENTS.md` first when adjusting these policies, and keep the other files in sync within the same pull request.

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
