package api_test

import (
	"testing"
	
	"github.com/saspallow/gowebboard/api"
)

func TestValidateForumCreateRequest(t *testing.T) {

	cases := []struct {
		req api.ForumCreateRequest
		hasError bool
	}{
		{api.ForumCreateRequest{}, true},
		{api.ForumCreateRequest{Title: "q"}, true},
		{api.ForumCreateRequest{Title: "123"}, true},
		{api.ForumCreateRequest{Title: "test1234"}, false},
	}

	for _, c := range cases {
		err := c.req.Validate()
		if c.hasError && err == nil {
			t.Errorf("Expected has error; got nul")
		}
		if !c.hasError && err != nil {
			t.Errorf("Expected not error; got error")
		}
	}
}