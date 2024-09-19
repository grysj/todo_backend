package db

import (
	"context"
	"testing"

	"github.com/grysj/todo_backend/util"

	"github.com/stretchr/testify/require"
)

func generateRandomList() (List, error) {
	randUser, _ := generateRandomUser()
	randTitle := util.RandomTitle()
	listParams := CreateListParams{
		CreatedBy: randUser.UserID,
		Title:     randTitle,
	}
	list, err := testQueries.CreateList(context.Background(), listParams)
	return list, err
}

func TestCreateList(t *testing.T) {
	randUser, _ := generateRandomUser()
	randTitle := util.RandomTitle()
	listParams := CreateListParams{
		CreatedBy: randUser.UserID,
		Title:     randTitle,
	}
	list, err := testQueries.CreateList(context.Background(), listParams)
	require.NoError(t, err)
	require.NotEmpty(t, list)

	require.Equal(t, randUser.UserID, list.CreatedBy)
	require.Equal(t, randTitle, list.Title)

	require.NotEmpty(t, list.ListID)
	require.NotEmpty(t, list.CreatedAt)
}

func TestEditTitle(t *testing.T) {
	list, _ := generateRandomList()

	newTitle := util.RandomTitle()

	editParams := EditTileParams{
		ListID: list.ListID,
		Title:  newTitle,
	}

	err := testQueries.EditTile(context.Background(), editParams)
	require.NoError(t, err)

	newList, _ := testQueries.GetList(context.Background(), list.ListID)

	require.Equal(t, list.CreatedAt, newList.CreatedAt)
	require.Equal(t, list.CreatedBy, newList.CreatedBy)
	require.Equal(t, newTitle, newList.Title)

}

func TestGetList(t *testing.T) {
	list, _ := generateRandomList()

	foundList, err := testQueries.GetList(context.Background(), list.ListID)
	require.NoError(t, err)
	require.NotEmpty(t, foundList)

	require.Equal(t, list.CreatedAt, foundList.CreatedAt)
	require.Equal(t, list.CreatedBy, foundList.CreatedBy)
	require.Equal(t, list.ListID, foundList.ListID)
	require.Equal(t, list.Title, foundList.Title)
}
