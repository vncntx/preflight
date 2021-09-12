![](./icon.svg)

# preflight

[![](https://github.com/vncntx/preflight/workflows/Unit%20Tests/badge.svg)](https://github.com/vncntx/preflight/actions?query=workflow%3A%22Unit+Tests%22)
[![](https://github.com/vncntx/preflight/workflows/Static%20Checks/badge.svg)](https://github.com/vncntx/preflight/actions?query=workflow%3A%22Static+Checks%22)
[![Go Reference](https://img.shields.io/badge/reference-007d9c.svg?labelColor=16161b&logo=go&logoColor=white)](https://pkg.go.dev/vincent.click/pkg/preflight?tab=doc)
[![Conventional Commits](https://img.shields.io/badge/commits-conventional-0047ab.svg?labelColor=16161b)](https://conventionalcommits.org)
[![License: BSD-3](https://img.shields.io/github/license/vncntx/preflight.svg?labelColor=16161b&color=0047ab)](./LICENSE)

Expectations and assertions for testing in Go

## Installation

```
go get vincent.click/pkg/preflight
```

## Getting Started

To write unit tests, use the **preflight** package to wrap around the standard **testing** package.

An [**Expectation**](./docs/expectation.md) provides a common interface for making assertions about values and behaviors. Expectations can be created from any variable or expression, [files and data streams](./docs/stream.md), or [builtin function calls](./docs/captor.md).

```go
func TestMethod(test *testing.T) {
    t := preflight.Unit(test)

    t.Expect(2 * 5).Equals(25)
}
```

## Development

Please read the [Contribution Guide](./CONTRIBUTING.md) before you proceed.

## Copyright

Copyright 2021 [Vincent Fiestada](mailto:vincent@vincent.click). This project is released under a [BSD-3-Clause License](./LICENSE).

Icon made by [Freepik](http://www.freepik.com/).
