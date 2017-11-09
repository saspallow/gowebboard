package api_test

import (
	"github.com/saspallow/gowebboard/api"
	"testing"
	"time"
)

func TestValidateForumCreateRequest(t *testing.T) {
	req := api.ForumCreateRequest{
		Title: "",
		CreateAt: time.Now(),
	}
	err := req.Validate()
	if err == nil {
		t.Errorf("Expected error not nil; got nil")
	}

	req = api.ForumCreateRequest{
		Title: "a",
		CreateAt: time.Now(),
	}
	err = req.Validate()
	if err == nil {
		t.Errorf("Expected error not nil; got nil")
	}

	req = api.ForumCreateRequest{
		Title: "test1234",
		CreateAt: time.Now(),
	}
	err = req.Validate()
	if err != nil {
		t.Errorf("Expected error to be nil; got %v", err)
	}
}