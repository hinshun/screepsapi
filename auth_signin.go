package screepsapi

import (
	"fmt"
	"net/http"
)

type AuthSigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthSigninResponse struct {
	Ok    int    `json:"ok"`
	Token string `json:"token"`
}

func (a *AuthSigninResponse) IsOk() bool {
	return a.Ok == 1
}

func (c *Client) Signin(email, password string) error {
	authSigninReq := AuthSigninRequest{
		Email:    email,
		Password: password,
	}
	authSigninResp := AuthSigninResponse{}

	err := c.post(authSigninPath, &authSigninReq, &authSigninResp, nil, http.StatusOK)
	if err != nil {
		return fmt.Errorf("failed to post auth signin: %s", err)
	}

	c.token = authSigninResp.Token
	return nil
}
