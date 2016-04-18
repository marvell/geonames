# Geonames data for Go

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/marvell/geonames)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

Handy interface for access to geonames.org data.

## Installation

    go get github.com/marvell/geonames

## Usage

See `examples` folder.

```
go run examples/country_alternate_names.go
```

## Caching

Downloaded files save to `./cache` directory with `geonames_y<year>w<number_of_week>_<name_file>` prefix.

# Roadmap

1. Data structures:

    * [x] [Countries](https://godoc.org/github.com/marvell/geonames#Country)
    * [x] [Time zones](https://godoc.org/github.com/marvell/geonames#TimeZone)
    * [x] [Languages](https://godoc.org/github.com/marvell/geonames#Language)
    * [x] [Features](https://godoc.org/github.com/marvell/geonames#Feature)
    * [x] [Alternate names](https://godoc.org/github.com/marvell/geonames#AlternateName)
    * [x] [Postal codes](https://godoc.org/github.com/marvell/geonames#PostalCode)
    * [ ] Admin1 codes
    * [ ] Feature codes
    * [ ] User tags
    * [ ] Hierarchy

2. Functionality:

    * [x] Caching
    * [x] More examples
    * [ ] CLI client
