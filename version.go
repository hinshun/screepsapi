package screepsapi

import (
	"fmt"
	"net/http"
)

type VersionResponse struct {
	Ok         int                `json:"ok"`
	Package    int                `json:"package"`
	Protocol   int                `json:"protocol"`
	ServerData ServerDataResponse `json:"serverData"`
}

type ServerDataResponse struct {
	HistoryChunkSize int      `json:"historyChunkSize"`
	Shards           []string `json:"shards"`
}

func (w *VersionResponse) IsOk() bool {
	return w.Ok == 1
}

func (c *Client) Version() (VersionResponse, error) {
	versionResp := VersionResponse{}
	err := c.get(versionPath, &versionResp, nil, http.StatusOK)
	if err != nil {
		return versionResp, fmt.Errorf("failed to get version: %s", err)
	}

	return versionResp, nil
}
