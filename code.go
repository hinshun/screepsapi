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
	if branch != "" {
		values.Add(branchKey, branch)
	}

	err := c.get(codePath, &pullCodeResp, values, http.StatusOK)
	if err != nil {
		return pullCodeResp, fmt.Errorf("failed to pull code: %s", err)
	}

	return pullCodeResp, nil
}

type PushCodeRequest struct {
	Branch  string            `json:"branch,omitempty"`
	Modules map[string]string `json:"modules"`
}

type PushCodeResponse struct {
	Ok        int `json:"ok"`
	Timestamp int `json:"timestamp"`
	Hash      int `json:"_hash"`
}

func (c *PushCodeResponse) IsOk() bool {
	return c.Ok == 1
}

func (c *Client) PushCode(branch string, modules map[string]string) (PushCodeResponse, error) {
	pushCodeReq := PushCodeRequest{
		Branch:  branch,
		Modules: modules,
	}
	pushCodeResp := PushCodeResponse{}

	err := c.post(codePath, &pushCodeReq, &pushCodeResp, nil, http.StatusOK)
	if err != nil {
		return pushCodeResp, fmt.Errorf("failed to push code: %s", err)
	}

	return pushCodeResp, nil
}
