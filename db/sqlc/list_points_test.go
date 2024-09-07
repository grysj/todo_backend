package db

import (
	"backend/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePoint(t *testing.T) {
	list, _ := generateRandomList()
	randPoint := util.RandomPoint()
	pointParams := CreatePointParams{
		ListID:   list.ListID,
		Content:  randPoint,
		Position: 1,
		AddedBy:  list.CreatedBy,
	}

	point, err := testQueries.CreatePoint(context.Background(), pointParams)
	require.NoError(t, err)
	require.NotEmpty(t, point)

	require.Equal(t, list.ListID, point.ListID)
	require.Equal(t, randPoint, point.Content)
	require.False(t, point.Checked)
	require.NotEmpty(t, point.PointID)
	require.NotEmpty(t, point.CreatedAt)

}
