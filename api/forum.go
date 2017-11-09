package api

import (
	"time"
)

type ForumController interface {
	Create(*ForumCreateRequest) (*ForumCreateResponse, error)
	Update(*ForumUpdateRequest) (*ForumUpdateResponse, error)
	List() (*ForumListReponse, error)
	Delete(int) (*ForumDeleteResponse, error)
	GET(int) (*Forum, error)
}

// Create
// ForumCreateRequest is create request for ForumController
type ForumCreateRequest struct {
	Title string
	CreateAt time.Time
}

// Validate reuest
func (req *ForumCreateRequest) Validate() error {
	return nil
}

// ForumCreateResponse is create response for ForumController
type ForumCreateResponse struct {
	ID int
}

// Update
// ForumUpdateRequest is update request for ForumController
type ForumUpdateRequest struct {
	ID int
	Title string
}

// ForumUpdateResponse is update response for ForumController
type ForumUpdateResponse struct {

}

// List
// ForumListReponse is list response for ForumController
type ForumListReponse struct {
	Forum []*Forum
}

// ForumDeleteResponse is delete response for ForumController
type ForumDeleteResponse struct {

}

// Forum type
type Forum struct {
	ID int
	Title string
}