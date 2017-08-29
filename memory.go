package screepsapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/hinshun/screepsapi/screepstype"
)

type RawMemoryResponse struct {
	Response
	Data string `json:"data"`
}

// UnmarshalJSON unmarshals the data string, which is of the form `gz:`
// followed by base64-encoded gzipped JSON encoding of the requested
// memory path, into a more functional map[string]interface{}.
func (m *MemoryResponse) UnmarshalJSON(b []byte) error {
	rawMemoryResp := RawMemoryResponse{}
	err := json.Unmarshal(b, &rawMemoryResp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal raw response: %s", err)
	}

	m.Ok = rawMemoryResp.Ok
	if rawMemoryResp.Data != "" {
		unzippedData, err := screepstype.Unzip(rawMemoryResp.Data)
		if err != nil {
			return fmt.Errorf("failed to unzip gzipped data: %s", err)
		}

		err = json.Unmarshal(unzippedData, &m.Data)
		if err != nil {
			return fmt.Errorf("failed to unmarshal unzipped data: %s", unzippedData)
		}
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

func (c *Client) MemorySegment(shard string, segment int) (MemorySegmentResponse, error) {
	memorySegmentResp := MemorySegmentResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)
	values.Add(segmentKey, strconv.Itoa(segment))

	err := c.get(memorySegmentPath, &memorySegmentResp, values, http.StatusOK)
	if err != nil {
		return memorySegmentResp, fmt.Errorf("failed to get memory segment: %s", err)
	}

	return memorySegmentResp, nil
}

func (c *Client) UpdateMemorySegment(shard, data string, segment int) (Response, error) {
	updateMemoryReq := UpdateMemoryRequest{
		Shard:   shard,
		Data:    data,
		Segment: segment,
	}
	resp := Response{}

	err := c.post(memorySegmentPath, &updateMemoryReq, &resp, nil, http.StatusOK)
	if err != nil {
		return resp, fmt.Errorf("failed to update memory segment: %s", err)
	}

	return resp, nil
}
