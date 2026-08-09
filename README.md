# Situation

[![run](https://github.com/asiffer/situation/actions/workflows/run.yaml/badge.svg)](https://github.com/asiffer/situation/actions/workflows/run.yaml)
[![docs](https://github.com/asiffer/situation/actions/workflows/docs.yaml/badge.svg)](https://asiffer.github.io/situation/)
[![Go Report Card](https://goreportcard.com/badge/github.com/asiffer/situation)](https://goreportcard.com/report/github.com/asiffer/situation)
[![gosec](https://github.com/asiffer/situation/actions/workflows/gosec.yaml/badge.svg)](https://github.com/asiffer/situation/security/code-scanning)
[![CodeQL](https://github.com/asiffer/situation/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/asiffer/situation/security/quality)

Situation provides the core infrastructure to automatically collect and consolidate IT data (machines, device, apps, network, flows...), on its own. 
Providing then an up-to-date and reliable view of the current state of your infra (or your home LAN), namely the *graph*.

Now you are ready to build a context-rich IT tool above Situation.

![infra](./excalidraw/architecture.svg)

## Installation

The agent currently supports Linux and Windows. 

### Github releases

Pre-built binaries are available through [github releases](https://github.com/asiffer/situation/releases/latest/).

### From sources

```shell
go install github.com/asiffer/situation/agent@latest
```

## Quick start

You can run the agent directly (without data persistence) and explore what has been discovered (experimental terminal ui)

```bash
situation run --explore
```

![tui](docs/img/tui.svg)


## Data persistence

Once you gives a db (sqlite or postgres) to `situation` you enable data persistence. 

```bash
situation run --db=situation.sqlite
```

To go further, several agents can collaborate by sharing the same postgres db (you own). 
The **IT data collection starts here**!

```bash
situation run --db="postgres://user:password@example.org:5432/situation"
```

See the [docs](https://asiffer.github.io/situation/) for more details.