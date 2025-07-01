package rest24

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

const orgId = "214517124256245"

func TestClientInitialization(t *testing.T) {
	require := require.New(t)

	c := New(&Credentials{
		ClientId:       os.Getenv("TFSO_REST_APP_ID"),
		ClientSecret:   os.Getenv("TFSO_REST_SECRET"),
		OrganizationId: orgId,
	})

	require.False(c.IsTokenValid(), "token should be invalid")
	require.NoError(c.Authenticate(), "client should authenticate")

	// Verify token properties
	token := c.GetToken()
	require.NotNil(token, "token should not be nil")
	require.NotEmpty(token.AccessToken, "access token should not be empty")
	require.Equal("Bearer", token.TokenType, "token type should be Bearer")

	t.Logf("Authentication successful. Token expires: %v", token.Expiry)

	res, err := c.GetAccountsWithResponse(context.Background(), &GetAccountsParams{})
	require.NoError(err, "GetAccounts should not error")
	require.NotNil(res, "GetAccounts should not be nil")
	require.Equal(http.StatusOK, res.StatusCode(), "GetAccounts status should be OK")
	require.NotNil(res.JSON200, "GetAccounts JSON200 should not be nil")
}

func TestClientTokenManagement(t *testing.T) {
	require := require.New(t)

	c := New(&Credentials{
		ClientId:       "test-client",
		ClientSecret:   "test-secret",
		OrganizationId: "test-org",
	})

	testToken := &oauth2.Token{
		AccessToken: "test-access-token",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(time.Hour),
	}
	c.SetToken(testToken)
	require.Equal(testToken.AccessToken, c.GetToken().AccessToken, "access token should match")
}
