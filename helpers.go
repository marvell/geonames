package geonames

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/c4milo/unzipit"
	"github.com/marvell/csvutil"
	"github.com/palantir/stacktrace"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
}

func downloadFile(url string) (string, error) {
	resp, err := httpClient.Get(url)

	if resp != nil {
		defer func() { resp.Body.Close() }()
	}

	if err != nil {
		return "", stacktrace.Propagate(err, "try to download file")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", stacktrace.Propagate(err, "try to read body of response")
	}

	file, err := ioutil.TempFile(os.TempDir(), "geonames")
	if err != nil {
		return "", stacktrace.Propagate(err, "try to create temporary file")
	}

	err = ioutil.WriteFile(file.Name(), body, 0644)
	if err != nil {
		return "", stacktrace.Propagate(err, "try to write to temporary file")
	}

	return file.Name(), nil
}

func parseCsvFile(name string, skipLines int, separateChar rune, commentsChar rune, handler func(raw []string) bool) error {
	file, err := os.Open(name)
	if err != nil {
		return stacktrace.Propagate(err, "try to open csv file")
	}

	reader := csvutil.NewReader(file, &csvutil.Config{
		Sep:           separateChar,
		Trim:          false,
		CommentPrefix: string(commentsChar),
		Comments:      true,
	})

	for i := 0; i < skipLines; i++ {
		reader.ReadRow()
	}

	reader.Do(func(row csvutil.Row) bool {
		if row.HasError() {
			log.Printf("[ERR] %s", row.Error)
			return false
		}

		return handler(row.Fields)
	})

	return nil
}

func unZip(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", stacktrace.Propagate(err, "try to open zip file")
	}

	defer file.Close()

	tempDir, err := ioutil.TempDir(os.TempDir(), "geonames")
	if err != nil {
		return "", stacktrace.Propagate(err, "try to create temporary directory")
	}

	destPath, err := unzipit.Unpack(file, tempDir)
	if err != nil {
		return "", stacktrace.Propagate(err, "try to unpack zip archive")
	}

	return destPath, nil
}
