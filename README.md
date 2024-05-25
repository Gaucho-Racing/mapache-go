# mapache-go

[![Build Status](https://github.com/gin-gonic/gin/workflows/Run%20Tests/badge.svg?branch=master)](https://github.com/gin-gonic/gin/actions?query=branch%3Amaster)
[![codecov](https://codecov.io/gh/gin-gonic/gin/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-gonic/gin)
[![GoDoc](https://pkg.go.dev/badge/github.com/gaucho-racing/mapache-go?status.svg)](https://pkg.go.dev/github.com/gaucho-racing/mapache-go?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/gaucho-racing/mapache-go/-/badge.svg)](https://sourcegraph.com/github.com/gaucho-racing/mapache-go?badge)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/release/gaucho-racing/mapache-go.svg?style=flat-square)](https://github.com/gaucho-racing/mapache-go/releases)

mapache-go is a library that provides access to various abstractions used by [Mapache](https://github.com/gaucho-racing/mapache) services.

## Getting started

### Prerequisites

mapache-go requires [Go](https://go.dev/) version [1.21](https://go.dev/doc/devel/release#go1.21.0) or above.

### Installing

With [Go's module support](https://go.dev/wiki/Modules#how-to-use-modules), `go [build|run|test]` automatically fetches the necessary dependencies when you add the import in your code:

```sh
import "github.com/gaucho-racing/mapache-go"
```

Alternatively, use `go get`:

```sh
go get -u github.com/gaucho-racing/mapache-go
```

## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b gh-username/my-amazing-feature`)
3. Commit your Changes (`git commit -m 'Add my amazing feature'`)
4. Push to the Branch (`git push origin gh-username/my-amazing-feature`)
5. Open a Pull Request