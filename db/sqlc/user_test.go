package db

import (
	"backend/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func generateRandomUser() (User, error) {
	arg := util.RandomUser()
	user, err := testQueries.CreateUser(context.Background(), arg)
	return user, err
}

func TestCreateUser(t *testing.T) {

	arg := util.RandomUser()
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg, user.Username)

	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CreatedAt)
}

func TestGetUser(t *testing.T) {
	randUser, err1 := generateRandomUser()
	require.NoError(t, err1)
	user, err2 := testQueries.GetUser(context.Background(), randUser.UserID)
	require.NoError(t, err2)
	require.Equal(t, randUser.UserID, user.UserID)
	require.Equal(t, randUser.Username, user.Username)
	require.Equal(t, randUser.CreatedAt, user.CreatedAt)
}

func TestCountUser(t *testing.T) {
	countBefore, err1 := testQueries.CountUsers(context.Background())
	require.NoError(t, err1)
	require.NotEmpty(t, countBefore)
	n := util.RandomInt(1, 64)

	var i int64
	for i = 0; i < n; i++ {
		generateRandomUser()
	}

	countAfter, err2 := testQueries.CountUsers(context.Background())
	require.NoError(t, err2)
	require.NotEmpty(t, countAfter)
	require.Equal(t, n, countAfter-countBefore)

}
