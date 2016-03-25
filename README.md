# Geonames for Go

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/marvell/geonames)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

## Usage

```go
    import (
    	"log"

    	"github.com/marvell/geonames"
    )

    func main() {
    	geonames.EnableDebugMode()

    	data, err := geonames.FetchCountryFeatures("US")
    	if err != nil {
    		log.Fatal(err)
    	}

    	log.Printf("%d", len(data))
    	if len(data) > 0 {
    		log.Printf("%#v", data[0])
    	}
    }
```
