//go:build e2e
// +build e2e

package test

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	"github.com/wfen/go-rest-api-course/internal/comment"
)

func TestGetComments(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment")
	if err != nil {
		t.Fail()
	}
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/PostComment", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var cmt comment.Comment
	if err := json.Unmarshal(resp.Body(), &cmt); err != nil {
		t.Fatalf("Failed to unmarshal bytes: %s", err)
	}

	resp, err = client.R().
		SetPathParams(map[string]string{
			"id": strconv.FormatUint(uint64(cmt.ID), 10),
		}).
		Delete(BASE_URL + "/api/comment/{id}")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestGetComment(t *testing.T) {
	client := resty.New()

	resp, err := client.R().
		SetBody(`{"slug": "/GetComment", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var cmt comment.Comment
	if err := json.Unmarshal(resp.Body(), &cmt); err != nil {
		t.Fatalf("Failed to unmarshal bytes: %s", err)
	}

	resp, err = client.R().
		SetPathParams(map[string]string{
			"id": strconv.FormatUint(uint64(cmt.ID), 10),
		}).
		Get(BASE_URL + "/api/comment/{id}")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	resp, err = client.R().
		SetPathParams(map[string]string{
			"id": strconv.FormatUint(uint64(cmt.ID), 10),
		}).
		Delete(BASE_URL + "/api/comment/{id}")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUpdateComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/UpdateComment", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var cmt comment.Comment
	if err := json.Unmarshal(resp.Body(), &cmt); err != nil {
		t.Fatalf("Failed to unmarshal bytes: %s", err)
	}

	resp, err = client.R().
		SetPathParams(map[string]string{
			"id": strconv.FormatUint(uint64(cmt.ID), 10),
		}).
		SetBody(`{"slug": "/UpdateComment", "author": "12345", "body": "stop the world"}`).
		Put(BASE_URL + "/api/comment/{id}")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	resp, err = client.R().
		SetPathParams(map[string]string{
			"id": strconv.FormatUint(uint64(cmt.ID), 10),
		}).
		Delete(BASE_URL + "/api/comment/{id}")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestDeleteComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/DeleteComment", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var cmt comment.Comment
	if err := json.Unmarshal(resp.Body(), &cmt); err != nil {
		t.Fatalf("Failed to unmarshal bytes: %s", err)
	}

	resp, err = client.R().
		SetPathParams(map[string]string{
			"id": strconv.FormatUint(uint64(cmt.ID), 10),
		}).
		Delete(BASE_URL + "/api/comment/{id}")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
