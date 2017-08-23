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
	userMemoryDataFormat = "gz"
)

type UserMemoryResponse struct {
	Ok   int         `json:"ok"`
	Data interface{} `json:"data"`
}

func (u *UserMemoryResponse) IsOk() bool {
	return u.Ok == 1
}

type RawUserMemoryResponse struct {
	Ok   int    `json:"ok"`
	Data string `json:"data"`
}

// UnmarshalJSON unmarshals the data string, which is of the form `gz:`
// followed by base64-encoded gzipped JSON encoding of the requested
// memory path, into a more functional map[string]interface{}.
func (u *UserMemoryResponse) UnmarshalJSON(b []byte) error {
	rawUserMemoryResp := RawUserMemoryResponse{}
	err := json.Unmarshal(b, &rawUserMemoryResp)
	if err != nil {
		return err
	}

	u.Ok = rawUserMemoryResp.Ok

	dataParts := strings.Split(rawUserMemoryResp.Data, fmt.Sprintf("%s:", userMemoryDataFormat))
	if len(dataParts) != 2 && dataParts[0] != userMemoryDataFormat {
		return fmt.Errorf("data not in format %s: %s", userMemoryDataFormat, rawUserMemoryResp.Data)
	}
	rawData := dataParts[1]

	decodedData, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return fmt.Errorf("failed to base64 decode data: %s", rawUserMemoryResp.Data)
	}

	gzipReader, err := gzip.NewReader(bytes.NewReader(decodedData))
	if err != nil {
		return fmt.Errorf("failed to create gzip reader over data: %s", rawUserMemoryResp.Data)
	}
	defer gzipReader.Close()

	unzippedData, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		return fmt.Errorf("failed to read gzip data: %s", rawUserMemoryResp.Data)
	}

	err = json.Unmarshal(unzippedData, &u.Data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal unzipped data: %s", unzippedData)
	}

	return nil
}

func (c *Client) UserMemory(shard, path string) (UserMemoryResponse, error) {
	userMemoryResp := UserMemoryResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)
	if path != "" {
		values.Add(pathKey, path)
	}

	err := c.get(userMemoryPath, &userMemoryResp, values, http.StatusOK)
	if err != nil {
		return userMemoryResp, fmt.Errorf("failed to get user memory: %s", err)
	}

	return userMemoryResp, nil
}

type UserMemoryRequest struct {
	Shard string  `json:"shard"`
	Path  string  `json:"path"`
	Value *string `json:"value,omitempty"`
}

func (c *Client) UpdateUserMemory(shard, path, value string) (UserConsoleResponse, error) {
	userMemoryReq := UserMemoryRequest{
		Shard: shard,
		Path:  path,
		Value: &value,
	}
	userConsoleResp := UserConsoleResponse{}

	err := c.post(userMemoryPath, &userMemoryReq, &userConsoleResp, nil, http.StatusOK)
	if err != nil {
		return userConsoleResp, fmt.Errorf("failed to post user memory: %s", err)
	}

	return userConsoleResp, nil
}

func (c *Client) DeleteUserMemory(shard, path string) (UserConsoleResponse, error) {
	userMemoryReq := UserMemoryRequest{
		Shard: shard,
		Path:  path,
	}
	userConsoleResp := UserConsoleResponse{}

	err := c.post(userMemoryPath, &userMemoryReq, &userConsoleResp, nil, http.StatusOK)
	if err != nil {
		return userConsoleResp, fmt.Errorf("failed to delete user memory: %s", err)
	}

	return userConsoleResp, nil
}
