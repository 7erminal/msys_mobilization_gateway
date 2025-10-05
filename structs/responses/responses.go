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
	AccountNumber string
	Product       string
}

type AccountsResponse struct {
	StatusCode int                `json:"StatusCode"`
	StatusDesc string             `json:"StatusDesc"`
	Result     *[]AccountsApiData `json:"Result,omitempty"`
}

type AccountsApiResponse struct {
	StatusCode int                `json:"StatusCode"`
	StatusDesc string             `json:"StatusDesc"`
	Result     *[]AccountsApiData `json:"Result,omitempty"`
	Client     string             `json:"Client,omitempty"`
}

type AccountsApiDataResponse struct {
	Data AccountsApiResponse `json:"data"`
}

type CreditAccountApiResponse struct {
	StatusCode int    `json:"StatusCode"`
	StatusDesc string `json:"StatusDesc"`
	Result     string `json:"Result,omitempty"`
	Client     string `json:"Client,omitempty"`
}

type CreditAccountApiDataResponse struct {
	Data CreditAccountApiResponse `json:"data"`
}

type CreditAccountResponse struct {
	StatusCode int    `json:"StatusCode"`
	StatusDesc string `json:"StatusDesc"`
	Result     string `json:"Result,omitempty"`
	Client     string `json:"Client,omitempty"`
}

type DebitAccountApiResponse struct {
	StatusCode int    `json:"StatusCode"`
	StatusDesc string `json:"StatusDesc"`
	Result     string `json:"Result,omitempty"`
	Client     string `json:"Client,omitempty"`
}

type DebitAccountApiDataResponse struct {
	Data CreditAccountApiResponse `json:"data"`
}

type DebitAccountResponse struct {
	StatusCode int    `json:"StatusCode"`
	StatusDesc string `json:"StatusDesc"`
	Result     string `json:"Result,omitempty"`
	Client     string `json:"Client,omitempty"`
}

type AccountBalanceData struct {
	AvailableBalance string
	ClearBalance     string
	LoanBalance      string
	SharesBalance    string
	AccountStatus    string
}

type AccountBalanceApiData struct {
	StatusCode int                   `json:"StatusCode"`
	StatusDesc string                `json:"StatusDesc"`
	Result     *[]AccountBalanceData `json:"Result,omitempty"`
	Client     string                `json:"Client,omitempty"`
}

type AccountBalanceApiResponse struct {
	Data AccountBalanceApiData `json:"data"`
}

type AccountBalanceResponse struct {
	StatusCode int                 `json:"StatusCode"`
	StatusDesc string              `json:"StatusDesc"`
	Result     *AccountBalanceData `json:"Result,omitempty"`
	Client     string              `json:"Client,omitempty"`
}

type AccountStatementData struct {
	TransactionDate   string
	TransactionType   string
	TransactionAmount string
	Reference         string
	Narration         string
	Description       string
	Amount            string
	Balance           string
}

type AccountStatementApiResponse struct {
	StatusCode int                     `json:"StatusCode"`
	StatusDesc string                  `json:"StatusDesc"`
	Result     *[]AccountStatementData `json:"Result,omitempty"`
	Client     string                  `json:"Client,omitempty"`
}

// type AccountStatementApiResponse struct {
// 	Data AccountStatementApiData `json:"data"`
// }

type AccountStatementResponse struct {
	StatusCode int                     `json:"StatusCode"`
	StatusDesc string                  `json:"StatusDesc"`
	Result     *[]AccountStatementData `json:"Result,omitempty"`
	Client     string                  `json:"Client,omitempty"`
}
