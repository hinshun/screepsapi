package screepsapi

import (
	"fmt"
	"net/http"
)

func (c *Client) Version() (VersionResponse, error) {
	versionResp := VersionResponse{}
	err := c.get(versionPath, &versionResp, nil, http.StatusOK)
	if err != nil {
		return versionResp, fmt.Errorf("failed to get version: %s", err)
	}

	return versionResp, nil
}
