package rest24

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientInitialization(t *testing.T) {
	require := require.New(t)

	c := New(&Credentials{
		ClientId:       os.Getenv("TFSO_REST_APP_ID"),
		ClientSecret:   os.Getenv("TFSO_REST_SECRET"),
		OrganizationId: "363976382792241",
	})

	require.False(c.IsTokenValid())

	require.NoError(c.Authenticate())

	res, err := c.GetMeWithResponse(context.Background(), nil)
	require.NoError(err)
	require.NotNil(res)
	require.Equal(http.StatusOK, res.StatusCode())
	require.NotEmpty(res.Body)
	require.NotNil(res.JSON200)
	require.Equal("NO", *res.JSON200.CountryCode)
}
