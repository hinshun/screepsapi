package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) Branches() (BranchesResponse, error) {
	branchesResp := BranchesResponse{}
	err := c.get(branchesPath, &branchesResp, nil, http.StatusOK)
	if err != nil {
		return branchesResp, fmt.Errorf("failed to get branches: %s", err)
	}

	return branchesResp, nil
}

func (c *Client) Console(shard, expression string) (InsertResponse, error) {
	consoleReq := ConsoleRequest{
		Expression: expression,
	}
	insertResp := InsertResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.post(consolePath, &consoleReq, &insertResp, values, http.StatusOK)
	if err != nil {
		return insertResp, fmt.Errorf("failed to send console: %s", err)
	}

	return insertResp, nil
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

func (c *Client) SetActiveBranch(branch string, activeName ActiveName) (Response, error) {
	setActiveBranchReq := SetActiveBranchRequest{
		Branch:     branch,
		ActiveName: activeName,
	}
	resp := Response{}

	err := c.post(setActiveBranchPath, &setActiveBranchReq, &resp, nil, http.StatusOK)
	if err != nil {
		return resp, fmt.Errorf("failed to set active branch: %s", err)
	}

	return resp, nil
}
