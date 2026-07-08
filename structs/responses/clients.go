package responses

import (
	"msys_api_gateway/models"
	"time"
)

type ClientData struct {
	ClientId     int64
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
	Result        *models.Clients
}

type ClientsResponse struct {
	StatusCode    int
	StatusMessage string
	Result        *[]models.Clients
}
