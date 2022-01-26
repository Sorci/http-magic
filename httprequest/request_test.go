package httprequest

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	SUCCESS = "SUCCESS"
)

func TestRequestGetSuccess(t *testing.T) {
	serverHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(SUCCESS))
	}
	server := httptest.NewServer(http.HandlerFunc(serverHandler))
	defer server.Close()
	r := NewRequest(server.URL)
	content, resp, err := r.Get().ResponseString()
	require.NoError(t, err, "should not have failed to make a GET request")

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, SUCCESS, content)
}

func TestRequestPostSuccess(t *testing.T) {
	serverHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(SUCCESS))
	}
	server := httptest.NewServer(http.HandlerFunc(serverHandler))
	defer server.Close()
	r := NewRequest(server.URL)
	content, resp, err := r.Post().ResponseString()
	require.NoError(t, err, "should not have failed to make a POST request")

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, SUCCESS, content)
}

func TestRequestPutSuccess(t *testing.T) {
	serverHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(SUCCESS))
	}
	server := httptest.NewServer(http.HandlerFunc(serverHandler))
	defer server.Close()
	r := NewRequest(server.URL)
	content, resp, err := r.Put().ResponseString()
	require.NoError(t, err, "should not have failed to make a PUT request")

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, SUCCESS, content)
}

func TestRequestDeleteSuccess(t *testing.T) {
	serverHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(SUCCESS))
	}
	server := httptest.NewServer(http.HandlerFunc(serverHandler))
	defer server.Close()
	r := NewRequest(server.URL)
	content, resp, err := r.Delete().ResponseString()
	require.NoError(t, err, "should not have failed to make a DELETE request")

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, SUCCESS, content)
}

func TestRequestPatchSuccess(t *testing.T) {
	serverHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(SUCCESS))
	}
	server := httptest.NewServer(http.HandlerFunc(serverHandler))
	defer server.Close()
	r := NewRequest(server.URL)
	content, resp, err := r.Patch().ResponseString()
	require.NoError(t, err, "should not have failed to make a PATCH request")

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, SUCCESS, content)
}

func TestRequestFailureTimeout(t *testing.T) {
	r := NewRequest("https://www.baidu.com", Options{Timeout: 1 * time.Microsecond})
	resp, err := r.Get().Response()
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Response code:", resp.StatusCode)
	}
}

func TestRequestFailureRetry(t *testing.T) {
	r := NewRequest("https://www.baidu.com", Options{
		Timeout: 1 * time.Microsecond,
		RetryCount: 3,
	})
	resp, err := r.Get().Response()
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %+v", err))
	} else {
		fmt.Println("Response code:", resp.StatusCode)
	}
}

