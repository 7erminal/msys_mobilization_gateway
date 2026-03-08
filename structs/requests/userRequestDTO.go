package requests

type LoginApiRequest struct {
	PhoneNumber  string
	Password     string
	SourceSystem string
	ClientId     string
}

type RefreshTokenRequest struct {
	RefreshToken string
}

type RefreshTokenApiRequest struct {
	RefreshToken string
}

type LoginRequest struct {
	PhoneNumber string
	Password    string
	ClientId    string `validate:"omitempty"`
}

type RegisterRequest struct {
	ClientId     int64  `validate:"required"`
	FirstName    string `validate:"required"`
	LastName     string `validate:"required"`
	Email        string
	Gender       string `validate:"required"`
	MobileNumber string `validate:"required"`
	Password     string `validate:"required"`
	Dob          string `validate:"required"`
	Username     string `validate:"required"`
}

type RegisterApiRequest struct {
	FirstName    string
	LastName     string
	Gender       string
	MobileNumber string
	ClientId     string
	Password     string
	Dob          string
	Username     string
}

type OpenAccountRequest struct {
	ClientId int64 `validate:"required"`
}

type ResetPassword struct {
	Username    string `validate:"required"`
	NewPassword string `validate:"required"`
	// ClientId    string `validate:"required"`
}

type ChangePasswordRequest struct {
	Username    string `validate:"required"`
	OldPassword string `validate:"required"`
	NewPassword string `validate:"required"`
	ClientId    string `validate:"required"`
}
