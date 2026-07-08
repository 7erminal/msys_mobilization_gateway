package responses

import "time"

type ClientData struct {
	Id           int64
	ClientName   string
	ClientCode   string
	ClientUrl    string
	ClientCorpId int64
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
	Active       int
	HasPOS       int `json:"Has_pos"`
	HasApp       int `json:"Has_app"`
}

type ClientResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *ClientData
}

type ClientsResponse struct {
	StatusCode    int
	StatusMessage string
	Result        *[]interface{} // Changed to interface{} to handle different types
}
