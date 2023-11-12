package db

import (
	"context"
	"github.com/Wintersunner/xplor/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomFizzBuzz(t *testing.T) (int64, CreateFizzBuzzParams) {

	arg := CreateFizzBuzzParams{
		Useragent: "Test Browser",
		Message:   util.RandomMessage(),
		CreatedAt: time.Now().UTC(),
	}

	result, err := testQueries.CreateFizzBuzz(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	id, err := result.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, id)

	return id, arg
}

func TestQueries_CreateFizzBuzz(t *testing.T) {
	id, fizzBuzz := createRandomFizzBuzz(t)

	dbFizzBuzz, err := testQueries.GetFizzBuzz(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, dbFizzBuzz)
	require.Equal(t, dbFizzBuzz.ID, id)
	require.Equal(t, dbFizzBuzz.Message, fizzBuzz.Message)
}

func TestQueries_GetFizzBuzz(t *testing.T) {
	id, fizzBuzz := createRandomFizzBuzz(t)

	dbFizzBuzz, err := testQueries.GetFizzBuzz(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, dbFizzBuzz)
	require.Equal(t, dbFizzBuzz.ID, id)
	require.Equal(t, dbFizzBuzz.Message, fizzBuzz.Message)
}

func TestQueries_ListFizzBuzzes(t *testing.T) {
	_, _ = createRandomFizzBuzz(t)
	_, _ = createRandomFizzBuzz(t)

	dbFizzBuzzesList, err := testQueries.ListFizzBuzzes(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(dbFizzBuzzesList), 2)
}
