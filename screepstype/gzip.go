package screepstype

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

const (
	GzipPrefix = "gz"
)

func Unzip(data string, compressionType CompressionType) ([]byte, error) {
	dataParts := strings.Split(data, fmt.Sprintf("%s:", GzipPrefix))
	if len(dataParts) != 2 && dataParts[0] != GzipPrefix {
		return nil, fmt.Errorf("data not in format %s: %s", GzipPrefix, data)
	}
	rawData := dataParts[1]

	decodedData, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return nil, fmt.Errorf("failed to base64 decode data: %s", err)
	}

	var gzipReader io.ReadCloser
	switch compressionType {
	case CompressionTypeGzip:
		gzipReader, err = gzip.NewReader(bytes.NewReader(decodedData))
	case CompressionTypeZlib:
		gzipReader, err = zlib.NewReader(bytes.NewReader(decodedData))
	default:
		return nil, fmt.Errorf("unsupported compression type: %s", compressionType)
	}
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
