package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

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
