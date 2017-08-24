package screepsapi

import (
	"fmt"
	"net/http"
)

type SetActiveBranchRequest struct {
	ActiveName ActiveNameType `json:"activeName"`
	Branch     string         `json:"branch"`
}

type ActiveNameType string

const (
	ActiveNameWorld ActiveNameType = "activeWorld"
)

type SetActiveBranchResponse struct {
	Ok int `json:"ok"`
}

func (s *SetActiveBranchResponse) IsOk() bool {
	return s.Ok == 1
}

func (c *Client) SetActiveBranch(branch string, activeName ActiveNameType) (SetActiveBranchResponse, error) {
	setActiveBranchReq := SetActiveBranchRequest{
		Branch:     branch,
		ActiveName: activeName,
	}
	setActiveBranchResp := SetActiveBranchResponse{}

	err := c.post(setActiveBranchPath, &setActiveBranchReq, &setActiveBranchResp, nil, http.StatusOK)
	if err != nil {
		return setActiveBranchResp, fmt.Errorf("failed to set active branch: %s", err)
	}

	return setActiveBranchResp, nil
}
