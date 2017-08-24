package screepsapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	memoryDataFormat = "gz"
)

type MemoryResponse struct {
	Ok   int         `json:"ok"`
	Data interface{} `json:"data"`
}

func (m *MemoryResponse) IsOk() bool {
	return m.Ok == 1
}

type RawMemoryResponse struct {
	Ok   int    `json:"ok"`
	Data string `json:"data"`
}

// UnmarshalJSON unmarshals the data string, which is of the form `gz:`
// followed by base64-encoded gzipped JSON encoding of the requested
// memory path, into a more functional map[string]interface{}.
func (m *MemoryResponse) UnmarshalJSON(b []byte) error {
	rawMemoryResp := RawMemoryResponse{}
	err := json.Unmarshal(b, &rawMemoryResp)
	if err != nil {
		return err
	}

	m.Ok = rawMemoryResp.Ok

	dataParts := strings.Split(rawMemoryResp.Data, fmt.Sprintf("%s:", memoryDataFormat))
	if len(dataParts) != 2 && dataParts[0] != memoryDataFormat {
		return fmt.Errorf("data not in format %s: %s", memoryDataFormat, rawMemoryResp.Data)
	}
	rawData := dataParts[1]

	decodedData, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return fmt.Errorf("failed to base64 decode data: %s", rawMemoryResp.Data)
	}

	gzipReader, err := gzip.NewReader(bytes.NewReader(decodedData))
	if err != nil {
		return fmt.Errorf("failed to create gzip reader over data: %s", rawMemoryResp.Data)
	}
	defer gzipReader.Close()

	unzippedData, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		return fmt.Errorf("failed to read gzip data: %s", rawMemoryResp.Data)
	}

	err = json.Unmarshal(unzippedData, &m.Data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal unzipped data: %s", unzippedData)
	}

	return nil
}

func (c *Client) Memory(shard, path string) (MemoryResponse, error) {
	memoryResp := MemoryResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)
	if path != "" {
		values.Add(pathKey, path)
	}

	err := c.get(memoryPath, &memoryResp, values, http.StatusOK)
	if err != nil {
		return memoryResp, fmt.Errorf("failed to get memory: %s", err)
	}

	return memoryResp, nil
}

type MemoryRequest struct {
	Shard string  `json:"shard"`
	Path  string  `json:"path"`
	Value *string `json:"value,omitempty"`
}

func (c *Client) UpdateMemory(shard, path, value string) (InsertResponse, error) {
	memoryReq := MemoryRequest{
		Shard: shard,
		Path:  path,
		Value: &value,
	}
	insertResp := InsertResponse{}

	err := c.post(memoryPath, &memoryReq, &insertResp, nil, http.StatusOK)
	if err != nil {
		return insertResp, fmt.Errorf("failed to update memory: %s", err)
	}

	return insertResp, nil
}

func (c *Client) DeleteMemory(shard, path string) (InsertResponse, error) {
	memoryReq := MemoryRequest{
		Shard: shard,
		Path:  path,
	}
	insertResp := InsertResponse{}

	err := c.post(memoryPath, &memoryReq, &insertResp, nil, http.StatusOK)
	if err != nil {
		return insertResp, fmt.Errorf("failed to delete memory: %s", err)
	}

	return insertResp, nil
}
