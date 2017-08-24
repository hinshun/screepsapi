package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type GenUniqueObjectNameRequest struct {
	Type GenType `json:"expression"`
}

type GenType string

const (
	GenTypeFlag  GenType = "flag"
	GenTypeSpawn         = "spawn"
)

type GenUniqueObjectNameResponse struct {
	Ok   int    `json:"ok"`
	Name string `json:"name"`
}

func (g *GenUniqueObjectNameResponse) IsOk() bool {
	return g.Ok == 1
}

func (c *Client) GenUniqueObjectName(shard string, genType GenType) (GenUniqueObjectNameResponse, error) {
	genUniqueObjectNameReq := GenUniqueObjectNameRequest{
		Type: genType,
	}
	genUniqueObjectNameResp := GenUniqueObjectNameResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.post(genUniqueObjectNamePath, &genUniqueObjectNameReq, &genUniqueObjectNameResp, values, http.StatusOK)
	if err != nil {
		return genUniqueObjectNameResp, fmt.Errorf("failed to generate unique object name: %s", err)
	}

	return genUniqueObjectNameResp, nil
}
