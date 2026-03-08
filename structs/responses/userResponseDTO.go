package responses

type LoginDataResponse struct {
	PhoneNumber  string `json:"phone_number"`
	SourceSystem string `json:"source_system"`
	ClientId     string `json:"client_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type LoginDataResponseResult struct {
	StatusCode    int
	StatusMessage string
	Result        bool
	Client        string
}

type LoginApiResponse struct {
	StatusCode int
	Value      string
	StatusDesc string
}

type LoginAccountApiResponse struct {
	Data LoginDataResponseResult `json:"data"`
}

type LoginResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        string
}

type TokenResponseDTO struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresIn    int64
}

type LoginDataResponseDTO struct {
	UserType string
	Token    *TokenResponseDTO
}

type LoginResponseDTO struct {
	StatusCode int
	Result     string
	StatusDesc string
}

type CustomerLoginApiResponseDTO struct {
	StatusCode int
	Result     *LoginDataResponseDTO
	StatusDesc string
}

type CustomerLoginDataResponseDTO struct {
	UserType string
	Token    *TokenResponseDTO
}

type CustomerLoginResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *CustomerLoginDataResponseDTO
}
