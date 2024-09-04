package db

import (
	"backend/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := util.RandomString(7)

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg, user.Username)

	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CreatedAt)
}

func TestGetUser(t *testing.T) {
	randUser, err := testQueries.CreateUser(context.Background(), util.RandomUser())
	require.NoError(t, err)
	user, err := testQueries.GetUser(context.Background(), randUser.UserID)
	require.NoError(t, err)
	require.Equal(t, randUser.UserID, user.UserID)
	require.Equal(t, randUser.Username, user.Username)
	require.Equal(t, randUser.CreatedAt, user.CreatedAt)
}
