package payroll24

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientInitialization(t *testing.T) {
	require := require.New(t)

	_, err := New(os.Getenv("TWENTYFOURSEVEN_API_PAYROLL"))
	require.NoError(err)
}

func TestServices(t *testing.T) {
	require := require.New(t)
	c, err := New(os.Getenv("TWENTYFOURSEVEN_API_PAYROLL"))
	require.NoError(err)

	a, err := c.GetAbsenceV2WithResponse(context.TODO())
	require.NoError(err, "GetAbsenceV2EmpIdWithResponse")
	require.Empty(a.JSON200)
}
