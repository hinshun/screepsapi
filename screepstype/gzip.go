package screepstype

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	GzipPrefix = "gz"
)

func Unzip(data string) ([]byte, error) {
	dataParts := strings.Split(data, fmt.Sprintf("%s:", GzipPrefix))
	if len(dataParts) != 2 && dataParts[0] != GzipPrefix {
		return nil, fmt.Errorf("data not in format %s: %s", GzipPrefix, data)
	}
	rawData := dataParts[1]

	decodedData, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return nil, fmt.Errorf("failed to base64 decode data: %s", err)
	}

	gzipReader, err := gzip.NewReader(bytes.NewReader(decodedData))
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader over data: %s", err)
	}
	defer gzipReader.Close()

	unzippedData, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		return nil, fmt.Errorf("failed to read gzip data: %s", err)
	}

	return unzippedData, nil
}
