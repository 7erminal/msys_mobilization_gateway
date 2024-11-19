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
