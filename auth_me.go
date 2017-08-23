package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type AuthMeResponse struct {
	Ok                 int            `json:"ok"`
	ID                 string         `json:"_id"`
	Email              string         `json:"email"`
	Username           string         `json:"username"`
	CPU                int            `json:"cpu"`
	Badge              BadgeResponse  `json:"badge"`
	Password           bool           `json:"password"`
	Credits            int            `json:"credits"`
	PromoPeriodUntil   int            `json:"promoPeriodUntil"`
	Money              int            `json:"money"`
	SubscriptionTokens int            `json:"subscriptionTokens"`
	GitHub             GitHubResponse `json:"github"`
	Steam              SteamResponse  `json:"steam"`
}

type GitHubResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type SteamResponse struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Ownership   []int  `json:"ownership"`
}

func (a *AuthMeResponse) IsOk() bool {
	return a.Ok == 1
}

func (c *Client) AuthMe(shard string) (AuthMeResponse, error) {
	authMeResp := AuthMeResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.Get(authMePath, &authMeResp, values, http.StatusOK)
	if err != nil {
		return authMeResp, fmt.Errorf("failed to get auth me: %s", err)
	}

	return authMeResp, nil
}
