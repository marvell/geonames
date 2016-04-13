package geonames

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/c4milo/unzipit"
	"github.com/marvell/csvutil"
	"github.com/marvell/downloader"
	"github.com/palantir/stacktrace"
)

var cacheDir = "./cache"

func downloadFile(fileUrl string, useCache bool) (string, error) {
	filename := getUrlFilename(fileUrl)

	if useCache == true {
		if _, err := os.Stat(filename); os.IsNotExist(err) == false {
			return filename, nil
		}
	}

	err := downloader.New(fileUrl).SaveToFile(filename)
	if err != nil {
		return "", err
	}

	return filename, nil
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

func getUrlFilename(fileUrl string) string {
	u, _ := url.Parse(fileUrl)

	year, week := time.Now().ISOWeek()
	filename := fmt.Sprintf("geonames_y%dw%d_%s", year, week, path.Base(u.Path))

	os.Mkdir(cacheDir, 0755)

	return path.Join(cacheDir, filename)
}
