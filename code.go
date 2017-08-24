package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type PullCodeResponse struct {
	Ok      int               `json:"ok"`
	Branch  string            `json:"branch"`
	Modules map[string]string `json:"modules"`
}

func (l *PullCodeResponse) IsOk() bool {
	return l.Ok == 1
}

func (c *Client) PullCode(branch string) (PullCodeResponse, error) {
	pullCodeResp := PullCodeResponse{}

	values := make(url.Values)
	values.Add(branchKey, branch)

	err := c.get(codePath, &pullCodeResp, values, http.StatusOK)
	if err != nil {
		return pullCodeResp, fmt.Errorf("failed to pull code: %s", err)
	}

	return pullCodeResp, nil
}

type PushCodeRequest struct {
	Branch  string            `json:"branch"`
	Modules map[string]string `json:"modules"`
}

type PushCodeResponse struct {
	Ok            int                    `json:"ok"`
	Result        PushCodeResultResponse `json:"result"`
	Ops           []PushCodeOpsResponse  `json:"ops"`
	InsertedCount int                    `json:"insertedCount"`
	InsertedIDs   []string               `json:"insertedIDs"`
}

type PushCodeResultResponse struct {
	Ok int `json:"ok"`
	N  int `json:"n"`
}

type PushCodeOpsResponse struct {
	ID         string `json:"_id"`
	User       string `json:"user"`
	Expression string `json:"expression"`
	Hidden     bool   `json:"hidden"`
}

func (c *PushCodeResponse) IsOk() bool {
	return c.Ok == 1
}

func (c *Client) PushCode(shard, expression string) (PushCodeResponse, error) {
	pushCodeReq := PushCodeRequest{
		Expression: expression,
	}
	pushCodeResp := PushCodeResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.post(codePath, &pushCodeReq, &pushCodeResp, values, http.StatusOK)
	if err != nil {
		return pushCodeResp, fmt.Errorf("failed to push code: %s", err)
	}

	return pushCodeResp, nil
}
