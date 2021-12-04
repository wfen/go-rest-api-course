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

const (
	// can be prepared at jwt.io with the appropriate 256-bit-secret
	jwt string = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI" +
		"6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.yNO7b5Ai_C_6A-WedtDaCcb2eP7odajno5lmBvrRwpo"
)

func TestGetComments(t *testing.T) {
	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().Get(BASE_URL + "/api/comment")
	if err != nil {
		t.Fail()
	}
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().
		//SetBasicAuth("admin", "password").
		SetHeader("Authorization", "Bearer "+jwt).
		SetBody(`{"slug": "/PostComment", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var cmt comment.Comment
	if err := json.Unmarshal(resp.Body(), &cmt); err != nil {
		t.Fatalf("Failed to unmarshal bytes: %s", err)
	}

	resp, err = client.R().
		//SetBasicAuth("admin", "password").
		SetHeader("Authorization", "Bearer "+jwt).
		SetPathParams(map[string]string{
			"id": strconv.FormatUint(uint64(cmt.ID), 10),
		}).
		Delete(BASE_URL + "/api/comment/{id}")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestGetComment(t *testing.T) {
	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().
		//SetBasicAuth("admin", "password").
		SetHeader("Authorization", "Bearer "+jwt).
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
		//SetBasicAuth("admin", "password").
		SetHeader("Authorization", "Bearer "+jwt).
		SetPathParams(map[string]string{
			"id": strconv.FormatUint(uint64(cmt.ID), 10),
		}).
		Delete(BASE_URL + "/api/comment/{id}")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUpdateComment(t *testing.T) {
	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().
		//SetBasicAuth("admin", "password").
		SetHeader("Authorization", "Bearer "+jwt).
		SetBody(`{"slug": "/UpdateComment", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var cmt comment.Comment
	if err := json.Unmarshal(resp.Body(), &cmt); err != nil {
		t.Fatalf("Failed to unmarshal bytes: %s", err)
	}

	resp, err = client.R().
		//SetBasicAuth("admin", "password").
		SetHeader("Authorization", "Bearer "+jwt).
		SetPathParams(map[string]string{
			"id": strconv.FormatUint(uint64(cmt.ID), 10),
		}).
		SetBody(`{"slug": "/UpdateComment", "author": "12345", "body": "stop the world"}`).
		Put(BASE_URL + "/api/comment/{id}")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	resp, err = client.R().
		//SetBasicAuth("admin", "password").
		SetHeader("Authorization", "Bearer "+jwt).
		SetPathParams(map[string]string{
			"id": strconv.FormatUint(uint64(cmt.ID), 10),
		}).
		Delete(BASE_URL + "/api/comment/{id}")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestDeleteComment(t *testing.T) {
	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().
		//SetBasicAuth("admin", "password").
		SetHeader("Authorization", "Bearer "+jwt).
		SetBody(`{"slug": "/DeleteComment", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var cmt comment.Comment
	if err := json.Unmarshal(resp.Body(), &cmt); err != nil {
		t.Fatalf("Failed to unmarshal bytes: %s", err)
	}

	resp, err = client.R().
		//SetBasicAuth("admin", "password").
		SetHeader("Authorization", "Bearer "+jwt).
		SetPathParams(map[string]string{
			"id": strconv.FormatUint(uint64(cmt.ID), 10),
		}).
		Delete(BASE_URL + "/api/comment/{id}")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
