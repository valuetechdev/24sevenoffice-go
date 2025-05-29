package rest24

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type authResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type BearerToken struct {
	Token     string
	ExpiresAt *time.Time
}

type TokenOpts struct {
	ClientId       string
	ClientSecret   string
	OrganizationId string
}

func newToken(opts *TokenOpts) (*BearerToken, error) {
	payload := url.Values{}
	payload.Set("client_id", opts.ClientId)
	payload.Set("client_secret", opts.ClientSecret)
	payload.Set("login_organization", opts.OrganizationId)
	payload.Set("grant_type", "client_credentials")
	payload.Set("audience", "https://api.24sevenoffice.com")

	req, err := http.NewRequest("POST", "https://login.24sevenoffice.com/oauth/token", strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status %d\nbody: %s", resp.StatusCode, string(body))
	}

	token := authResponse{}
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(token.ExpiresIn))
	bt := BearerToken{
		Token:     token.AccessToken,
		ExpiresAt: &expiresAt,
	}

	return &bt, nil
}

func (bt *BearerToken) HasExpired() bool {
	return time.Now().After(*bt.ExpiresAt)
}

func (bt *BearerToken) Intercept(ctx context.Context, req *http.Request) error {
	req.Header.Set("Authorization", "bearer "+bt.Token)

	return nil
}
