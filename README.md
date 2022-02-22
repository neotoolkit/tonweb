# TonWeb

[![CI-img]][CI-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Unofficial tonweb golang sdk

## Installation
```shell
go get github.com/neotoolkit/tonweb
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/neotoolkit/tonweb"
)

func main() {
	c := tonweb.NewClient()

	req, err := c.Blocks().GetConsensusBlock()

	fmt.Println(req, err)
}
```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[CI-img]: https://github.com/neotoolkit/tonweb/workflows/CI/badge.svg
[CI-url]: https://github.com/neotoolkit/tonweb/actions
[pkg-img]: https://pkg.go.dev/badge/neotoolkit/tonweb
[pkg-url]: https://pkg.go.dev/github.com/neotoolkit/tonweb
[reportcard-img]: https://goreportcard.com/badge/neotoolkit/tonweb
[reportcard-url]: https://goreportcard.com/report/neotoolkit/tonweb
[coverage-img]: https://codecov.io/gh/neotoolkit/tonweb/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/neotoolkit/tonweb
[version-img]: https://img.shields.io/github/v/release/neotoolkit/tonweb
[version-url]: https://github.com/neotoolkit/tonweb/releases

## Sponsors
<p>
  <a href="https://evrone.com/?utm_source=github&utm_campaign=neotoolkit">
    <img src="https://raw.githubusercontent.com/neotoolkit/.github/main/assets/sponsored_by_evrone.svg"
      alt="Sponsored by Evrone">
  </a>
</p>
