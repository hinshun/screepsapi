package screepsapi

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) Me(shard string) (MeResponse, error) {
	meResp := MeResponse{}

	values := make(url.Values)
	values.Add(shardKey, shard)

	err := c.get(mePath, &meResp, values, http.StatusOK)
	if err != nil {
		return meResp, fmt.Errorf("failed to get auth me: %s", err)
	}

	return meResp, nil
}

func (c *Client) SignIn(email, password string) error {
	signInReq := SignInRequest{
		Email:    email,
		Password: password,
	}
	signInResp := SignInResponse{}

	err := c.post(signInPath, &signInReq, &signInResp, nil, http.StatusOK)
	if err != nil {
		return fmt.Errorf("failed to post auth signin: %s", err)
	}

	c.token = signInResp.Token
	return nil
}
