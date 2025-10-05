package requests

type AddServiceRequest struct {
	ServiceName        string
	ServiceDescription string
	CreatedBy          int
	ModifiedBy         int
	Active             int
}

type ValidatePin struct {
	Number   string
	Password string
}

type ResetPin struct {
	Number      string
	OldPassword string
	NewPassword string
}

type Id struct {
	Id string
}

type Number struct {
	Number string
}

type CreditAccountRequest struct {
	AccountNumber string
	Amount        string
	Reference     string
}

type CreditAccountRequestV2 struct {
	AccountNumber string
	Amount        string
	Reference     string
	Channel       string
}

type DebitAccountRequestV2 struct {
	AccountNumber string
	Amount        string
	Reference     string
	Channel       string
}

type AccountNumber struct {
	AccountNumber string
}

type FieldDepositRequest struct {
	AccountNumber string
	Amount        string
	MobileNumber  string
}

type RegisterCustomer struct {
	FirstName    string
	LastName     string
	Gender       string
	MobileNumber string
}

type CreateFieldAccount struct {
	FirstName         string
	LastName          string
	Gender            string
	MobileNumber      string
	AgentMobileNumber string
}

type VerifyCustomer struct {
	Username     string
	FirstName    string
	LastName     string
	Email        string
	Dob          string
	Gender       string
	MobileNumber string
}

type ActivateCustomer struct {
	Username     string
	MobileNumber string
}

type AccountStatementRequestV2 struct {
	AccountNumber string
	FromDate      string
	ToDate        string
}
