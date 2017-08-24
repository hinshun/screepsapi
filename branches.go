package screepsapi

import (
	"fmt"
	"net/http"
)

type BranchesResponse struct {
	Ok       int              `json:"ok"`
	Branches []BranchResponse `json:"list"`
}

type BranchResponse struct {
	ID               string `json:"_id"`
	Branch           string `json:"branch"`
	ActiveSimulation bool   `json:"activeSim"`
	ActiveWorld      bool   `json:"activeWorld"`
}

func (b *BranchesResponse) IsOk() bool {
	return b.Ok == 1
}

func (c *Client) Branches() (BranchesResponse, error) {
	branchesResp := BranchesResponse{}
	err := c.get(branchesPath, &branchesResp, nil, http.StatusOK)
	if err != nil {
		return branchesResp, fmt.Errorf("failed to get branches: %s", err)
	}

	return branchesResp, nil
}
