# po

Vendor your Go dependencies, nothing more.

## Installation

    go get github.com/kyleconroy/po

## Usage

Po sets `$GOPATH` to a local `_vendor` directory. That's it.

To use, prefix all your `go` commands with `po`.

    po go command [arguments]

When you're done, make sure to save `_vendor`.
