package responses

import "time"

type ClientData struct {
	Id          int64
	ClientName  string
	ClientCode  string
	ClientUrl   string
	DateCreated time.Time
	Active      int
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
