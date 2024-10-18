package bearer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	ValidFromUtc string `json:"validFromUtc"`
	ValidToUtc   string `json:"validToUtc"`
}

type BearerToken struct {
	Token     string
	ExpiresAt *time.Time
}

func New() (*BearerToken, error) {
	apiToken, ok := os.LookupEnv("TWENTYFOURSEVEN_API_PAYROLL")
	if !ok {
		// Safe-return
		return nil, nil
	}

	url := fmt.Sprintf("https://payroll.24sevenoffice.com/api/auth?token=%s", apiToken)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad status %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type accessToken struct {
		AccessToken string `json:"accessToken"`
		ExpiresOn   string `json:"expiresOn"`
	}
	token := accessToken{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	expiresOnSplitted := strings.Split(token.ExpiresOn, "+")
	expiresAt, err := time.Parse("2006-01-02T15:04:05", expiresOnSplitted[0])
	if err != nil {
		return nil, err
	}
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
	req.Header.Set("Authorization", "Bearer "+bt.Token)

	return nil
}
