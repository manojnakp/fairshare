---
title: Project Setup
description: Initialise and setup Fairshare, first steps.
editUrl: https://github.com/manojnakp/fairshare/edit/main/docs/walkthrough/01-setup.md
sidebar:
    badge: New
---

## First Steps

Initialise a simple [Go] project. This step consist of the following project directory:

```
/
├─ .github/workflows/  ⟶ GitHub Actions
├─ .editorconfig       ⟶ EditorConfig
├─ .gitignore          ⟶ Git repo ignore files
├─ .node-version       ⟶ supported node version (CI/CD)
├─ LICENSE             ⟶ Open source license (MIT)
├─ README.md           ⟶ Landing page
├─ go.mod              ⟶ Go module
├─ main.go             ⟶ Go entrypoint
╰─ package.json        ⟶ Nodejs dependency management
```

## HTTP Server

Setup Fairshare as an HTTP server that listens to incoming requests.

```
github.com/manojnakp/fairshare
├─ api/       ⟶ RESTful API endpoints and middlewares
├─ cli/       ⟶ CLI and configuration
├─ internal/  ⟶ Shared but private
╰─ main.go    ⟶ Entrypoint
```

So far, the necessary configuration options like HTTP port are parsed from
command-line flags. Errors are logged through `log.Println` and sibling
functions. This does not do so well in case of tracing logs like HTTP server
logs. So we need structured logging as well.

## Logging and Config improvements

[Go] supports custom config parsing through command line flags via `flag.Value`.
Custom flag types can be created that conform to `flag.Value` and supports other
configuration options like parsing from environment variables.

Logging can be improvised using `slog` package which has an elegant JSON logger
which (nearly) perfectly suits our requirements of writing HTTP server log. The
advantage of this approach is that any well known tool (like [Prometheus],
[Grafana] etc.) can consume the structured JSON logs for later processing.

## Documentation Site

This documentation site itself is set up [withastro]. [Starlight] from [Astro]
with [OpenAPI plugin] is used to present content and reference material for this
site.

[Go]: https://go.dev
[Prometheus]: https://prometheus.io
[Grafana]: https://grafana.com
[withastro]: https://github.com/withastro/astro
[Starlight]: https://starlight.astro.build
[Astro]: https://astro.build
[OpenAPI plugin]: https://starlight-openapi.vercel.app
