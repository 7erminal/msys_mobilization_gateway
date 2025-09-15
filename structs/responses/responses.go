package responses

type VerifyCustomerResponse struct {
	StatusCode int    `json:"StatusCode"`
	StatusDesc string `json:"StatusDesc"`
	Result     string `json:"Result,omitempty"`
	Client     string `json:"Client,omitempty"`
}

type VerifyCustomerResponseData struct {
	Result     string `json:"Result,omitempty"`
	Client     string `json:"Client,omitempty"`
	StatusCode int    `json:"StatusCode,omitempty"`
	StatusDesc string `json:"StatusDesc,omitempty"`
}

type VerifyCustomerApiResponse struct {
	Data VerifyCustomerResponseData `json:"data"`
}

type CustomerAccountsApiData struct {
	Username     string `json:"username"`
	MobileNumber string `json:"mobileNumber"`
}

type CustomerAccountsData struct {
	Username     string
	MobileNumber string
}

type ApprovedCustomersApiData struct {
	StatusCode int                        `json:"StatusCode"`
	StatusDesc string                     `json:"StatusDesc"`
	Result     *[]CustomerAccountsApiData `json:"Result,omitempty"`
	Client     string                     `json:"Client,omitempty"`
}

type ApprovedCustomersApiResponse struct {
	Data ApprovedCustomersApiData `json:"data"`
}

type ApprovedCustomersResponse struct {
	StatusCode int                    `json:"StatusCode"`
	StatusDesc string                 `json:"StatusDesc"`
	Result     []CustomerAccountsData `json:"Result,omitempty"`
}

type ActivateCustomerResponse struct {
	StatusCode int    `json:"StatusCode"`
	StatusDesc string `json:"StatusDesc"`
	Result     string `json:"Result,omitempty"`
}

type AccountsApiData struct {
	AccountNumber string `json:"accountNumber"`
	BankCode      string `json:"bankCode"`
	BankName      string `json:"bankName"`
	AccountName   string `json:"accountName"`
}

type AccountsResponse struct {
	StatusCode int                `json:"StatusCode"`
	StatusDesc string             `json:"StatusDesc"`
	Result     *[]AccountsApiData `json:"Result,omitempty"`
}
