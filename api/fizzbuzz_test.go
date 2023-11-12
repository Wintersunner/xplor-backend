package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	mockdb "github.com/Wintersunner/xplor/db/mock"
	db "github.com/Wintersunner/xplor/db/sqlc"
	"github.com/Wintersunner/xplor/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var testConfig = util.Config{AllowedOrigins: "*", GinMode: gin.TestMode}

func TestGetFizzBuzz(t *testing.T) {
	fizzbuzz := randomFizzBuzz()
	store := createMockStore(t)
	store.EXPECT().GetFizzBuzz(gomock.Any(), gomock.Eq(fizzbuzz.ID)).Times(1).Return(fizzbuzz, nil)

	server := NewServer(store, testConfig)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/fizzbuzz/%d", fizzbuzz.ID), nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)
	var receivedFizzBuzz db.Fizzbuzz
	err = json.Unmarshal(recorder.Body.Bytes(), &receivedFizzBuzz)
	require.NoError(t, err)
	require.Equal(t, receivedFizzBuzz.Message, fizzbuzz.Message)
	require.Equal(t, http.StatusOK, recorder.Code)
}

func Test404ForInvalidFizzBuzzID(t *testing.T) {
	store := createMockStore(t)
	store.EXPECT().GetFizzBuzz(gomock.Any(), gomock.Any()).Times(1).Return(db.Fizzbuzz{}, errors.New("not found"))
	server := NewServer(store, testConfig)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/fizzbuzz/500", nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestListFizzBuzz(t *testing.T) {
	fizzbuzzList := []db.Fizzbuzz{
		randomFizzBuzz(), randomFizzBuzz(),
	}
	store := createMockStore(t)
	store.EXPECT().ListFizzBuzzes(gomock.Any()).Times(1).Return(fizzbuzzList, nil)
	server := NewServer(store, testConfig)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/fizzbuzz", nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)
	var receivedList []db.Fizzbuzz
	err = json.Unmarshal(recorder.Body.Bytes(), &receivedList)
	require.NoError(t, err)
	require.Equal(t, len(receivedList), len(fizzbuzzList))
	require.Equal(t, receivedList[0].ID, fizzbuzzList[0].ID)
}

func TestCreateFizzBuzz(t *testing.T) {
	req := db.CreateFizzBuzzParams{
		Message:   "Hello",
		CreatedAt: time.Now().UTC(),
	}
	jsonReq, _ := json.Marshal(req)
	store := createMockStore(t)
	execResult := execMockResult{
		affectedRows: 1,
		insertId:     120,
	}
	store.EXPECT().CreateFizzBuzz(gomock.Any(), gomock.AssignableToTypeOf(db.CreateFizzBuzzParams{})).Times(1).Return(execResult, nil)
	server := NewServer(store, testConfig)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/fizzbuzz", bytes.NewBuffer(jsonReq))
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)
	var response FizzBuzzResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	require.Equal(t, response.ID, int64(120))
	require.Equal(t, response.Message, req.Message)
	require.WithinDuration(t, response.CreatedAt, req.CreatedAt, time.Second)
}

func randomFizzBuzz() db.Fizzbuzz {
	return db.Fizzbuzz{
		ID:        util.RandomInt(1, 1000),
		Useragent: "Mock Test UserAgent",
		Message:   util.RandomMessage(),
		CreatedAt: time.Now(),
	}
}

func createMockStore(t *testing.T) *mockdb.MockStore {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return mockdb.NewMockStore(ctrl)
}

type execMockResult struct {
	affectedRows int64
	insertId     int64
}

func (res execMockResult) LastInsertId() (int64, error) {
	return res.insertId, nil
}

func (res execMockResult) RowsAffected() (int64, error) {
	return res.affectedRows, nil
}
